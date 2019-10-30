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

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateK8SSecretsReader is a Reader for the CreateK8SSecrets structure.
type CreateK8SSecretsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateK8SSecretsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateK8SSecretsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateK8SSecretsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateK8SSecretsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateK8SSecretsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateK8SSecretsOK creates a CreateK8SSecretsOK with default headers values
func NewCreateK8SSecretsOK() *CreateK8SSecretsOK {
	return &CreateK8SSecretsOK{}
}

/*CreateK8SSecretsOK handles this case with default header values.

A successful response.
*/
type CreateK8SSecretsOK struct {
	Payload *service_model.V1K8SResource
}

func (o *CreateK8SSecretsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/k8s_secrets][%d] createK8SSecretsOK  %+v", 200, o.Payload)
}

func (o *CreateK8SSecretsOK) GetPayload() *service_model.V1K8SResource {
	return o.Payload
}

func (o *CreateK8SSecretsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1K8SResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateK8SSecretsNoContent creates a CreateK8SSecretsNoContent with default headers values
func NewCreateK8SSecretsNoContent() *CreateK8SSecretsNoContent {
	return &CreateK8SSecretsNoContent{}
}

/*CreateK8SSecretsNoContent handles this case with default header values.

No content.
*/
type CreateK8SSecretsNoContent struct {
	Payload interface{}
}

func (o *CreateK8SSecretsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/k8s_secrets][%d] createK8SSecretsNoContent  %+v", 204, o.Payload)
}

func (o *CreateK8SSecretsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateK8SSecretsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateK8SSecretsForbidden creates a CreateK8SSecretsForbidden with default headers values
func NewCreateK8SSecretsForbidden() *CreateK8SSecretsForbidden {
	return &CreateK8SSecretsForbidden{}
}

/*CreateK8SSecretsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateK8SSecretsForbidden struct {
	Payload interface{}
}

func (o *CreateK8SSecretsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/k8s_secrets][%d] createK8SSecretsForbidden  %+v", 403, o.Payload)
}

func (o *CreateK8SSecretsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateK8SSecretsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateK8SSecretsNotFound creates a CreateK8SSecretsNotFound with default headers values
func NewCreateK8SSecretsNotFound() *CreateK8SSecretsNotFound {
	return &CreateK8SSecretsNotFound{}
}

/*CreateK8SSecretsNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateK8SSecretsNotFound struct {
	Payload interface{}
}

func (o *CreateK8SSecretsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/k8s_secrets][%d] createK8SSecretsNotFound  %+v", 404, o.Payload)
}

func (o *CreateK8SSecretsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateK8SSecretsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
