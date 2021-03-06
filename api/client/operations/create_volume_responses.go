package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/opensds/opensds/api/models"
)

// CreateVolumeReader is a Reader for the CreateVolume structure.
type CreateVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateVolumeOK creates a CreateVolumeOK with default headers values
func NewCreateVolumeOK() *CreateVolumeOK {
	return &CreateVolumeOK{}
}

/*CreateVolumeOK handles this case with default header values.

OK
*/
type CreateVolumeOK struct {
	Payload *models.VolumeResponse
}

func (o *CreateVolumeOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/volumes/{resourceType}][%d] createVolumeOK  %+v", 200, o.Payload)
}

func (o *CreateVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
