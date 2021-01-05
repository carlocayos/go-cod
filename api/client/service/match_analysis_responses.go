// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/carlocayos/go-cod/api/models"
)

// MatchAnalysisReader is a Reader for the MatchAnalysis structure.
type MatchAnalysisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MatchAnalysisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMatchAnalysisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMatchAnalysisOK creates a MatchAnalysisOK with default headers values
func NewMatchAnalysisOK() *MatchAnalysisOK {
	return &MatchAnalysisOK{}
}

/*MatchAnalysisOK handles this case with default header values.

Match Analysis Response
*/
type MatchAnalysisOK struct {
	Payload *models.MatchAnalysisResponse
}

func (o *MatchAnalysisOK) Error() string {
	return fmt.Sprintf("[GET /ce/v2/title/{title}/platform/{platform}/gametype/all/{lookupType}/{gamertag}/summary/match_analysis/contentType/full/end/{end}/matchAnalysis/mobile/en][%d] matchAnalysisOK  %+v", 200, o.Payload)
}

func (o *MatchAnalysisOK) GetPayload() *models.MatchAnalysisResponse {
	return o.Payload
}

func (o *MatchAnalysisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MatchAnalysisResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
