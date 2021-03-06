package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ListShareResourcesReader is a Reader for the ListShareResources structure.
type ListShareResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListShareResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListShareResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListShareResourcesOK creates a ListShareResourcesOK with default headers values
func NewListShareResourcesOK() *ListShareResourcesOK {
	return &ListShareResourcesOK{}
}

/*ListShareResourcesOK handles this case with default header values.

OK
*/
type ListShareResourcesOK struct {
	Payload ListShareResourcesOKBody
}

func (o *ListShareResourcesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/shares][%d] listShareResourcesOK  %+v", 200, o.Payload)
}

func (o *ListShareResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListShareResourcesOKBody list share resources o k body
swagger:model ListShareResourcesOKBody
*/
type ListShareResourcesOKBody interface{}
