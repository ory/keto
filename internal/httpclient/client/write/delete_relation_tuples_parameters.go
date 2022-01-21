// Code generated by go-swagger; DO NOT EDIT.

package write

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
)

// NewDeleteRelationTuplesParams creates a new DeleteRelationTuplesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRelationTuplesParams() *DeleteRelationTuplesParams {
	return &DeleteRelationTuplesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRelationTuplesParamsWithTimeout creates a new DeleteRelationTuplesParams object
// with the ability to set a timeout on a request.
func NewDeleteRelationTuplesParamsWithTimeout(timeout time.Duration) *DeleteRelationTuplesParams {
	return &DeleteRelationTuplesParams{
		timeout: timeout,
	}
}

// NewDeleteRelationTuplesParamsWithContext creates a new DeleteRelationTuplesParams object
// with the ability to set a context for a request.
func NewDeleteRelationTuplesParamsWithContext(ctx context.Context) *DeleteRelationTuplesParams {
	return &DeleteRelationTuplesParams{
		Context: ctx,
	}
}

// NewDeleteRelationTuplesParamsWithHTTPClient creates a new DeleteRelationTuplesParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRelationTuplesParamsWithHTTPClient(client *http.Client) *DeleteRelationTuplesParams {
	return &DeleteRelationTuplesParams{
		HTTPClient: client,
	}
}

/* DeleteRelationTuplesParams contains all the parameters to send to the API endpoint
   for the delete relation tuples operation.

   Typically these are written to a http.Request.
*/
type DeleteRelationTuplesParams struct {

	/* Namespace.

	   Namespace of the Relation Tuple
	*/
	Namespace *string

	/* Object.

	   Object of the Relation Tuple
	*/
	Object *string

	/* Relation.

	   Relation of the Relation Tuple
	*/
	Relation *string

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

// WithDefaults hydrates default values in the delete relation tuples params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRelationTuplesParams) WithDefaults() *DeleteRelationTuplesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete relation tuples params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRelationTuplesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithTimeout(timeout time.Duration) *DeleteRelationTuplesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithContext(ctx context.Context) *DeleteRelationTuplesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithHTTPClient(client *http.Client) *DeleteRelationTuplesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithNamespace(namespace *string) *DeleteRelationTuplesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithObject adds the object to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithObject(object *string) *DeleteRelationTuplesParams {
	o.SetObject(object)
	return o
}

// SetObject adds the object to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetObject(object *string) {
	o.Object = object
}

// WithRelation adds the relation to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithRelation(relation *string) *DeleteRelationTuplesParams {
	o.SetRelation(relation)
	return o
}

// SetRelation adds the relation to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetRelation(relation *string) {
	o.Relation = relation
}

// WithSubjectID adds the subjectID to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithSubjectID(subjectID *string) *DeleteRelationTuplesParams {
	o.SetSubjectID(subjectID)
	return o
}

// SetSubjectID adds the subjectId to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetSubjectID(subjectID *string) {
	o.SubjectID = subjectID
}

// WithSubjectSetNamespace adds the subjectSetNamespace to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithSubjectSetNamespace(subjectSetNamespace *string) *DeleteRelationTuplesParams {
	o.SetSubjectSetNamespace(subjectSetNamespace)
	return o
}

// SetSubjectSetNamespace adds the subjectSetNamespace to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetSubjectSetNamespace(subjectSetNamespace *string) {
	o.SubjectSetNamespace = subjectSetNamespace
}

// WithSubjectSetObject adds the subjectSetObject to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithSubjectSetObject(subjectSetObject *string) *DeleteRelationTuplesParams {
	o.SetSubjectSetObject(subjectSetObject)
	return o
}

// SetSubjectSetObject adds the subjectSetObject to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetSubjectSetObject(subjectSetObject *string) {
	o.SubjectSetObject = subjectSetObject
}

// WithSubjectSetRelation adds the subjectSetRelation to the delete relation tuples params
func (o *DeleteRelationTuplesParams) WithSubjectSetRelation(subjectSetRelation *string) *DeleteRelationTuplesParams {
	o.SetSubjectSetRelation(subjectSetRelation)
	return o
}

// SetSubjectSetRelation adds the subjectSetRelation to the delete relation tuples params
func (o *DeleteRelationTuplesParams) SetSubjectSetRelation(subjectSetRelation *string) {
	o.SubjectSetRelation = subjectSetRelation
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRelationTuplesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string

		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {

			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}
	}

	if o.Object != nil {

		// query param object
		var qrObject string

		if o.Object != nil {
			qrObject = *o.Object
		}
		qObject := qrObject
		if qObject != "" {

			if err := r.SetQueryParam("object", qObject); err != nil {
				return err
			}
		}
	}

	if o.Relation != nil {

		// query param relation
		var qrRelation string

		if o.Relation != nil {
			qrRelation = *o.Relation
		}
		qRelation := qrRelation
		if qRelation != "" {

			if err := r.SetQueryParam("relation", qRelation); err != nil {
				return err
			}
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
