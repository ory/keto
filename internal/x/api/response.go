package api

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	v1alpha2 "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type response struct {
	w               http.ResponseWriter
	shouldIntercept bool
	buf             *bytes.Buffer
	code            int
}

func (r *response) Header() http.Header {
	return r.w.Header()
}

func (r *response) Write(bytes []byte) (int, error) {
	if !r.shouldIntercept {
		return r.w.Write(bytes)
	}

	// If we are here, it means that WriteHeader was called with a non-200 status code.
	return r.buf.Write(bytes)
}

func (r *response) WriteHeader(statusCode int) {
	if statusCode == http.StatusOK {
		if codeFromHeader := r.w.Header().Get("x-http-code"); codeFromHeader != "" {
			r.w.Header().Del("x-http-code")
			statusCodeFromHeader, err := strconv.Atoi(codeFromHeader)
			if err != nil {
				log.Println("error:", err)
			}
			r.w.WriteHeader(statusCodeFromHeader)
		}
		return
	}

	r.shouldIntercept = true
	r.buf = &bytes.Buffer{}
	r.code = statusCode
}

func (r *response) Flush() {
	if !r.shouldIntercept {
		if f, ok := r.w.(http.Flusher); ok {
			f.Flush()
		}
	}
}

func (r *response) flush() {
	if !r.shouldIntercept {
		return
	}

	var spb statuspb.Status
	if err := protojson.Unmarshal(r.buf.Bytes(), &spb); err != nil {
		r.w.WriteHeader(r.code)
		_, _ = r.w.Write(r.buf.Bytes())
		return
	}

	// Proto errors in the REST request path are always bad requests.
	if strings.HasPrefix(spb.Message, "proto:") ||
		spb.Message == "empty field path" ||
		invalidFieldPathRe.MatchString(spb.Message) ||
		unexpectedBodyRe.MatchString(spb.Message) ||
		strings.HasPrefix(spb.Message, "grpc: error unmarshalling request:") {
		// Special case: error deserializing request into a protobuf is a bad request.
		r.code = http.StatusBadRequest
	}
	r.w.WriteHeader(r.code)

	if spb.Code == 0 && spb.Message == "" {
		_, _ = r.w.Write(r.buf.Bytes())
		return
	}

	s := status.FromProto(&spb)
	errResponse := v1alpha2.ErrorResponse{
		Error: &v1alpha2.ErrorResponse_Error{
			Code:    int64(r.code),
			Status:  http.StatusText(r.code),
			Message: s.Message(),
		},
	}
	for _, detail := range s.Details() {
		switch t := detail.(type) {
		case *errdetails.ErrorInfo:
			errResponse.Error.Reason = t.Reason
		case *errdetails.DebugInfo:
			errResponse.Error.Debug = t.Detail
		case *errdetails.RequestInfo:
			errResponse.Error.Request = t.RequestId
		case *errdetails.BadRequest:
			errResponse.Error.Details = make(map[string]string, len(t.FieldViolations))
			for _, v := range t.FieldViolations {
				errResponse.Error.Details[v.Field] = v.Description
			}
		}
	}

	out, err := protojson.Marshal(&errResponse)
	if err != nil {
		log.Println("error:", err)
	}
	_, _ = r.w.Write(out)
}
