package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit"
)

/*GetInventoryOK successful operation

swagger:response getInventoryOK
*/
type GetInventoryOK struct {

	// In: body
	Payload GetInventoryOKBodyBody `json:"body,omitempty"`
}

// NewGetInventoryOK creates GetInventoryOK with default headers values
func NewGetInventoryOK() *GetInventoryOK {
	return &GetInventoryOK{}
}

// WithPayload adds the payload to the get inventory o k response
func (o *GetInventoryOK) WithPayload(payload GetInventoryOKBodyBody) *GetInventoryOK {
	o.Payload = payload
	return o
}

// WriteResponse to the client
func (o *GetInventoryOK) WriteResponse(rw http.ResponseWriter, producer httpkit.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
