// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/carlocayos/go-cod/v2/api/models"
)

// MapListReader is a Reader for the MapList structure.
type MapListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MapListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMapListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMapListOK creates a MapListOK with default headers values
func NewMapListOK() *MapListOK {
	return &MapListOK{}
}

/*MapListOK handles this case with default header values.

Map List Response
*/
type MapListOK struct {
	Payload *models.MapListResponse
}

func (o *MapListOK) Error() string {
	return fmt.Sprintf("[GET /ce/v1/title/{title}/platform/{platform}/gameType/{gameType}/communityMapData/availability][%d] mapListOK  %+v", 200, o.Payload)
}

func (o *MapListOK) GetPayload() *models.MapListResponse {
	return o.Payload
}

func (o *MapListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MapListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
