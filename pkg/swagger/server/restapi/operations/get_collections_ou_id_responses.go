// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Donders-Institute/dr-gateway/pkg/swagger/server/models"
)

// GetCollectionsOuIDOKCode is the HTTP code returned for type GetCollectionsOuIDOK
const GetCollectionsOuIDOKCode int = 200

/*
GetCollectionsOuIDOK success

swagger:response getCollectionsOuIdOK
*/
type GetCollectionsOuIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBodyCollections `json:"body,omitempty"`
}

// NewGetCollectionsOuIDOK creates GetCollectionsOuIDOK with default headers values
func NewGetCollectionsOuIDOK() *GetCollectionsOuIDOK {

	return &GetCollectionsOuIDOK{}
}

// WithPayload adds the payload to the get collections ou Id o k response
func (o *GetCollectionsOuIDOK) WithPayload(payload *models.ResponseBodyCollections) *GetCollectionsOuIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get collections ou Id o k response
func (o *GetCollectionsOuIDOK) SetPayload(payload *models.ResponseBodyCollections) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCollectionsOuIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCollectionsOuIDInternalServerErrorCode is the HTTP code returned for type GetCollectionsOuIDInternalServerError
const GetCollectionsOuIDInternalServerErrorCode int = 500

/*
GetCollectionsOuIDInternalServerError failure

swagger:response getCollectionsOuIdInternalServerError
*/
type GetCollectionsOuIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseBody500 `json:"body,omitempty"`
}

// NewGetCollectionsOuIDInternalServerError creates GetCollectionsOuIDInternalServerError with default headers values
func NewGetCollectionsOuIDInternalServerError() *GetCollectionsOuIDInternalServerError {

	return &GetCollectionsOuIDInternalServerError{}
}

// WithPayload adds the payload to the get collections ou Id internal server error response
func (o *GetCollectionsOuIDInternalServerError) WithPayload(payload *models.ResponseBody500) *GetCollectionsOuIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get collections ou Id internal server error response
func (o *GetCollectionsOuIDInternalServerError) SetPayload(payload *models.ResponseBody500) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCollectionsOuIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}