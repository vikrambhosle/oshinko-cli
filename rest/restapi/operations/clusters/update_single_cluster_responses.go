package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/radanalyticsio/oshinko-cli/rest/models"
)

/*UpdateSingleClusterAccepted Cluster update response

swagger:response updateSingleClusterAccepted
*/
type UpdateSingleClusterAccepted struct {

	// In: body
	Payload *models.SingleCluster `json:"body,omitempty"`
}

// NewUpdateSingleClusterAccepted creates UpdateSingleClusterAccepted with default headers values
func NewUpdateSingleClusterAccepted() *UpdateSingleClusterAccepted {
	return &UpdateSingleClusterAccepted{}
}

// WithPayload adds the payload to the update single cluster accepted response
func (o *UpdateSingleClusterAccepted) WithPayload(payload *models.SingleCluster) *UpdateSingleClusterAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update single cluster accepted response
func (o *UpdateSingleClusterAccepted) SetPayload(payload *models.SingleCluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSingleClusterAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateSingleClusterDefault Unexpected error

swagger:response updateSingleClusterDefault
*/
type UpdateSingleClusterDefault struct {
	_statusCode int

	// In: body
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewUpdateSingleClusterDefault creates UpdateSingleClusterDefault with default headers values
func NewUpdateSingleClusterDefault(code int) *UpdateSingleClusterDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateSingleClusterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update single cluster default response
func (o *UpdateSingleClusterDefault) WithStatusCode(code int) *UpdateSingleClusterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update single cluster default response
func (o *UpdateSingleClusterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update single cluster default response
func (o *UpdateSingleClusterDefault) WithPayload(payload *models.ErrorResponse) *UpdateSingleClusterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update single cluster default response
func (o *UpdateSingleClusterDefault) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateSingleClusterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}