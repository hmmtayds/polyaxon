// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package k8_s_secrets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteK8SSecretReader is a Reader for the DeleteK8SSecret structure.
type DeleteK8SSecretReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteK8SSecretReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteK8SSecretOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteK8SSecretNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteK8SSecretForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteK8SSecretNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteK8SSecretOK creates a DeleteK8SSecretOK with default headers values
func NewDeleteK8SSecretOK() *DeleteK8SSecretOK {
	return &DeleteK8SSecretOK{}
}

/*DeleteK8SSecretOK handles this case with default header values.

A successful response.
*/
type DeleteK8SSecretOK struct {
}

func (o *DeleteK8SSecretOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/k8s_secrets/{uuid}][%d] deleteK8SSecretOK ", 200)
}

func (o *DeleteK8SSecretOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteK8SSecretNoContent creates a DeleteK8SSecretNoContent with default headers values
func NewDeleteK8SSecretNoContent() *DeleteK8SSecretNoContent {
	return &DeleteK8SSecretNoContent{}
}

/*DeleteK8SSecretNoContent handles this case with default header values.

No content.
*/
type DeleteK8SSecretNoContent struct {
	Payload interface{}
}

func (o *DeleteK8SSecretNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/k8s_secrets/{uuid}][%d] deleteK8SSecretNoContent  %+v", 204, o.Payload)
}

func (o *DeleteK8SSecretNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteK8SSecretNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteK8SSecretForbidden creates a DeleteK8SSecretForbidden with default headers values
func NewDeleteK8SSecretForbidden() *DeleteK8SSecretForbidden {
	return &DeleteK8SSecretForbidden{}
}

/*DeleteK8SSecretForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteK8SSecretForbidden struct {
	Payload interface{}
}

func (o *DeleteK8SSecretForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/k8s_secrets/{uuid}][%d] deleteK8SSecretForbidden  %+v", 403, o.Payload)
}

func (o *DeleteK8SSecretForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteK8SSecretForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteK8SSecretNotFound creates a DeleteK8SSecretNotFound with default headers values
func NewDeleteK8SSecretNotFound() *DeleteK8SSecretNotFound {
	return &DeleteK8SSecretNotFound{}
}

/*DeleteK8SSecretNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteK8SSecretNotFound struct {
	Payload interface{}
}

func (o *DeleteK8SSecretNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/{owner}/k8s_secrets/{uuid}][%d] deleteK8SSecretNotFound  %+v", 404, o.Payload)
}

func (o *DeleteK8SSecretNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteK8SSecretNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
