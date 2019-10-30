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

package k8_s_config_maps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateK8SConfigMapReader is a Reader for the UpdateK8SConfigMap structure.
type UpdateK8SConfigMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateK8SConfigMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateK8SConfigMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateK8SConfigMapNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateK8SConfigMapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateK8SConfigMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateK8SConfigMapOK creates a UpdateK8SConfigMapOK with default headers values
func NewUpdateK8SConfigMapOK() *UpdateK8SConfigMapOK {
	return &UpdateK8SConfigMapOK{}
}

/*UpdateK8SConfigMapOK handles this case with default header values.

A successful response.
*/
type UpdateK8SConfigMapOK struct {
	Payload *service_model.V1K8SResource
}

func (o *UpdateK8SConfigMapOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] updateK8SConfigMapOK  %+v", 200, o.Payload)
}

func (o *UpdateK8SConfigMapOK) GetPayload() *service_model.V1K8SResource {
	return o.Payload
}

func (o *UpdateK8SConfigMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1K8SResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateK8SConfigMapNoContent creates a UpdateK8SConfigMapNoContent with default headers values
func NewUpdateK8SConfigMapNoContent() *UpdateK8SConfigMapNoContent {
	return &UpdateK8SConfigMapNoContent{}
}

/*UpdateK8SConfigMapNoContent handles this case with default header values.

No content.
*/
type UpdateK8SConfigMapNoContent struct {
	Payload interface{}
}

func (o *UpdateK8SConfigMapNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] updateK8SConfigMapNoContent  %+v", 204, o.Payload)
}

func (o *UpdateK8SConfigMapNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateK8SConfigMapNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateK8SConfigMapForbidden creates a UpdateK8SConfigMapForbidden with default headers values
func NewUpdateK8SConfigMapForbidden() *UpdateK8SConfigMapForbidden {
	return &UpdateK8SConfigMapForbidden{}
}

/*UpdateK8SConfigMapForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateK8SConfigMapForbidden struct {
	Payload interface{}
}

func (o *UpdateK8SConfigMapForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] updateK8SConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *UpdateK8SConfigMapForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateK8SConfigMapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateK8SConfigMapNotFound creates a UpdateK8SConfigMapNotFound with default headers values
func NewUpdateK8SConfigMapNotFound() *UpdateK8SConfigMapNotFound {
	return &UpdateK8SConfigMapNotFound{}
}

/*UpdateK8SConfigMapNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateK8SConfigMapNotFound struct {
	Payload interface{}
}

func (o *UpdateK8SConfigMapNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/k8s_config_maps/{k8s_resource.uuid}][%d] updateK8SConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *UpdateK8SConfigMapNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateK8SConfigMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
