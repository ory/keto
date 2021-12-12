// Code generated by go-swagger; DO NOT EDIT.

package read

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetCheckParams creates a new GetCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCheckParams() *GetCheckParams {
	return &GetCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCheckParamsWithTimeout creates a new GetCheckParams object
// with the ability to set a timeout on a request.
func NewGetCheckParamsWithTimeout(timeout time.Duration) *GetCheckParams {
	return &GetCheckParams{
		timeout: timeout,
	}
}

// NewGetCheckParamsWithContext creates a new GetCheckParams object
// with the ability to set a context for a request.
func NewGetCheckParamsWithContext(ctx context.Context) *GetCheckParams {
	return &GetCheckParams{
		Context: ctx,
	}
}

// NewGetCheckParamsWithHTTPClient creates a new GetCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCheckParamsWithHTTPClient(client *http.Client) *GetCheckParams {
	return &GetCheckParams{
		HTTPClient: client,
	}
}

/* GetCheckParams contains all the parameters to send to the API endpoint
   for the get check operation.

   Typically these are written to a http.Request.
*/
type GetCheckParams struct {

	// MaxDepth.
	//
	// Format: int64
	MaxDepth *int64

	/* Namespace.

	   Namespace of the Relation Tuple
	*/
	Namespace string

	/* Object.

	   Object of the Relation Tuple
	*/
	Object string

	/* Relation.

	   Relation of the Relation Tuple
	*/
	Relation string

	/* SubjectID.

	   SubjectID of the Relation Tuple
	*/
	SubjectID *string

	/* SubjectSetNamespace.

	   Namespace of the Subject Set
	*/
	SubjectSetNamespace *string

	/* SubjectSetObject.

	   Object of the Subject Set
	*/
	SubjectSetObject *string

	/* SubjectSetRelation.

	   Relation of the Subject Set
	*/
	SubjectSetRelation *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCheckParams) WithDefaults() *GetCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get check params
func (o *GetCheckParams) WithTimeout(timeout time.Duration) *GetCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get check params
func (o *GetCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get check params
func (o *GetCheckParams) WithContext(ctx context.Context) *GetCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get check params
func (o *GetCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get check params
func (o *GetCheckParams) WithHTTPClient(client *http.Client) *GetCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get check params
func (o *GetCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMaxDepth adds the maxDepth to the get check params
func (o *GetCheckParams) WithMaxDepth(maxDepth *int64) *GetCheckParams {
	o.SetMaxDepth(maxDepth)
	return o
}

// SetMaxDepth adds the maxDepth to the get check params
func (o *GetCheckParams) SetMaxDepth(maxDepth *int64) {
	o.MaxDepth = maxDepth
}

// WithNamespace adds the namespace to the get check params
func (o *GetCheckParams) WithNamespace(namespace string) *GetCheckParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the get check params
func (o *GetCheckParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithObject adds the object to the get check params
func (o *GetCheckParams) WithObject(object string) *GetCheckParams {
	o.SetObject(object)
	return o
}

// SetObject adds the object to the get check params
func (o *GetCheckParams) SetObject(object string) {
	o.Object = object
}

// WithRelation adds the relation to the get check params
func (o *GetCheckParams) WithRelation(relation string) *GetCheckParams {
	o.SetRelation(relation)
	return o
}

// SetRelation adds the relation to the get check params
func (o *GetCheckParams) SetRelation(relation string) {
	o.Relation = relation
}

// WithSubjectID adds the subjectID to the get check params
func (o *GetCheckParams) WithSubjectID(subjectID *string) *GetCheckParams {
	o.SetSubjectID(subjectID)
	return o
}

// SetSubjectID adds the subjectId to the get check params
func (o *GetCheckParams) SetSubjectID(subjectID *string) {
	o.SubjectID = subjectID
}

// WithSubjectSetNamespace adds the subjectSetNamespace to the get check params
func (o *GetCheckParams) WithSubjectSetNamespace(subjectSetNamespace *string) *GetCheckParams {
	o.SetSubjectSetNamespace(subjectSetNamespace)
	return o
}

// SetSubjectSetNamespace adds the subjectSetNamespace to the get check params
func (o *GetCheckParams) SetSubjectSetNamespace(subjectSetNamespace *string) {
	o.SubjectSetNamespace = subjectSetNamespace
}

// WithSubjectSetObject adds the subjectSetObject to the get check params
func (o *GetCheckParams) WithSubjectSetObject(subjectSetObject *string) *GetCheckParams {
	o.SetSubjectSetObject(subjectSetObject)
	return o
}

// SetSubjectSetObject adds the subjectSetObject to the get check params
func (o *GetCheckParams) SetSubjectSetObject(subjectSetObject *string) {
	o.SubjectSetObject = subjectSetObject
}

// WithSubjectSetRelation adds the subjectSetRelation to the get check params
func (o *GetCheckParams) WithSubjectSetRelation(subjectSetRelation *string) *GetCheckParams {
	o.SetSubjectSetRelation(subjectSetRelation)
	return o
}

// SetSubjectSetRelation adds the subjectSetRelation to the get check params
func (o *GetCheckParams) SetSubjectSetRelation(subjectSetRelation *string) {
	o.SubjectSetRelation = subjectSetRelation
}

// WriteToRequest writes these params to a swagger request
func (o *GetCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MaxDepth != nil {

		// query param max-depth
		var qrMaxDepth int64

		if o.MaxDepth != nil {
			qrMaxDepth = *o.MaxDepth
		}
		qMaxDepth := swag.FormatInt64(qrMaxDepth)
		if qMaxDepth != "" {

			if err := r.SetQueryParam("max-depth", qMaxDepth); err != nil {
				return err
			}
		}
	}

	// query param namespace
	qrNamespace := o.Namespace
	qNamespace := qrNamespace
	if qNamespace != "" {

		if err := r.SetQueryParam("namespace", qNamespace); err != nil {
			return err
		}
	}

	// query param object
	qrObject := o.Object
	qObject := qrObject
	if qObject != "" {

		if err := r.SetQueryParam("object", qObject); err != nil {
			return err
		}
	}

	// query param relation
	qrRelation := o.Relation
	qRelation := qrRelation
	if qRelation != "" {

		if err := r.SetQueryParam("relation", qRelation); err != nil {
			return err
		}
	}

	if o.SubjectID != nil {

		// query param subject_id
		var qrSubjectID string

		if o.SubjectID != nil {
			qrSubjectID = *o.SubjectID
		}
		qSubjectID := qrSubjectID
		if qSubjectID != "" {

			if err := r.SetQueryParam("subject_id", qSubjectID); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetNamespace != nil {

		// query param subject_set.namespace
		var qrSubjectSetNamespace string

		if o.SubjectSetNamespace != nil {
			qrSubjectSetNamespace = *o.SubjectSetNamespace
		}
		qSubjectSetNamespace := qrSubjectSetNamespace
		if qSubjectSetNamespace != "" {

			if err := r.SetQueryParam("subject_set.namespace", qSubjectSetNamespace); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetObject != nil {

		// query param subject_set.object
		var qrSubjectSetObject string

		if o.SubjectSetObject != nil {
			qrSubjectSetObject = *o.SubjectSetObject
		}
		qSubjectSetObject := qrSubjectSetObject
		if qSubjectSetObject != "" {

			if err := r.SetQueryParam("subject_set.object", qSubjectSetObject); err != nil {
				return err
			}
		}
	}

	if o.SubjectSetRelation != nil {

		// query param subject_set.relation
		var qrSubjectSetRelation string

		if o.SubjectSetRelation != nil {
			qrSubjectSetRelation = *o.SubjectSetRelation
		}
		qSubjectSetRelation := qrSubjectSetRelation
		if qSubjectSetRelation != "" {

			if err := r.SetQueryParam("subject_set.relation", qSubjectSetRelation); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
