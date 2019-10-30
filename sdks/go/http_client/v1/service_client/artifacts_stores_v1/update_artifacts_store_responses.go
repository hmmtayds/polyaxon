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

package artifacts_stores_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateArtifactsStoreReader is a Reader for the UpdateArtifactsStore structure.
type UpdateArtifactsStoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateArtifactsStoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateArtifactsStoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateArtifactsStoreNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateArtifactsStoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateArtifactsStoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateArtifactsStoreOK creates a UpdateArtifactsStoreOK with default headers values
func NewUpdateArtifactsStoreOK() *UpdateArtifactsStoreOK {
	return &UpdateArtifactsStoreOK{}
}

/*UpdateArtifactsStoreOK handles this case with default header values.

A successful response.
*/
type UpdateArtifactsStoreOK struct {
	Payload *service_model.V1ArtifactsStore
}

func (o *UpdateArtifactsStoreOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/artifacts_stores/{artifact_store.uuid}][%d] updateArtifactsStoreOK  %+v", 200, o.Payload)
}

func (o *UpdateArtifactsStoreOK) GetPayload() *service_model.V1ArtifactsStore {
	return o.Payload
}

func (o *UpdateArtifactsStoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ArtifactsStore)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateArtifactsStoreNoContent creates a UpdateArtifactsStoreNoContent with default headers values
func NewUpdateArtifactsStoreNoContent() *UpdateArtifactsStoreNoContent {
	return &UpdateArtifactsStoreNoContent{}
}

/*UpdateArtifactsStoreNoContent handles this case with default header values.

No content.
*/
type UpdateArtifactsStoreNoContent struct {
	Payload interface{}
}

func (o *UpdateArtifactsStoreNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/artifacts_stores/{artifact_store.uuid}][%d] updateArtifactsStoreNoContent  %+v", 204, o.Payload)
}

func (o *UpdateArtifactsStoreNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateArtifactsStoreNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateArtifactsStoreForbidden creates a UpdateArtifactsStoreForbidden with default headers values
func NewUpdateArtifactsStoreForbidden() *UpdateArtifactsStoreForbidden {
	return &UpdateArtifactsStoreForbidden{}
}

/*UpdateArtifactsStoreForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateArtifactsStoreForbidden struct {
	Payload interface{}
}

func (o *UpdateArtifactsStoreForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/artifacts_stores/{artifact_store.uuid}][%d] updateArtifactsStoreForbidden  %+v", 403, o.Payload)
}

func (o *UpdateArtifactsStoreForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateArtifactsStoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateArtifactsStoreNotFound creates a UpdateArtifactsStoreNotFound with default headers values
func NewUpdateArtifactsStoreNotFound() *UpdateArtifactsStoreNotFound {
	return &UpdateArtifactsStoreNotFound{}
}

/*UpdateArtifactsStoreNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateArtifactsStoreNotFound struct {
	Payload interface{}
}

func (o *UpdateArtifactsStoreNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/artifacts_stores/{artifact_store.uuid}][%d] updateArtifactsStoreNotFound  %+v", 404, o.Payload)
}

func (o *UpdateArtifactsStoreNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateArtifactsStoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
