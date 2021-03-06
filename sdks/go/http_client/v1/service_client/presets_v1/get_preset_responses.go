// Copyright 2018-2020 Polyaxon, Inc.
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

package presets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetPresetReader is a Reader for the GetPreset structure.
type GetPresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetPresetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetPresetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPresetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPresetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPresetOK creates a GetPresetOK with default headers values
func NewGetPresetOK() *GetPresetOK {
	return &GetPresetOK{}
}

/*GetPresetOK handles this case with default header values.

A successful response.
*/
type GetPresetOK struct {
	Payload *service_model.V1Preset
}

func (o *GetPresetOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets/{uuid}][%d] getPresetOK  %+v", 200, o.Payload)
}

func (o *GetPresetOK) GetPayload() *service_model.V1Preset {
	return o.Payload
}

func (o *GetPresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Preset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresetNoContent creates a GetPresetNoContent with default headers values
func NewGetPresetNoContent() *GetPresetNoContent {
	return &GetPresetNoContent{}
}

/*GetPresetNoContent handles this case with default header values.

No content.
*/
type GetPresetNoContent struct {
	Payload interface{}
}

func (o *GetPresetNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets/{uuid}][%d] getPresetNoContent  %+v", 204, o.Payload)
}

func (o *GetPresetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetPresetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresetForbidden creates a GetPresetForbidden with default headers values
func NewGetPresetForbidden() *GetPresetForbidden {
	return &GetPresetForbidden{}
}

/*GetPresetForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetPresetForbidden struct {
	Payload interface{}
}

func (o *GetPresetForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets/{uuid}][%d] getPresetForbidden  %+v", 403, o.Payload)
}

func (o *GetPresetForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetPresetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresetNotFound creates a GetPresetNotFound with default headers values
func NewGetPresetNotFound() *GetPresetNotFound {
	return &GetPresetNotFound{}
}

/*GetPresetNotFound handles this case with default header values.

Resource does not exist.
*/
type GetPresetNotFound struct {
	Payload interface{}
}

func (o *GetPresetNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets/{uuid}][%d] getPresetNotFound  %+v", 404, o.Payload)
}

func (o *GetPresetNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetPresetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPresetDefault creates a GetPresetDefault with default headers values
func NewGetPresetDefault(code int) *GetPresetDefault {
	return &GetPresetDefault{
		_statusCode: code,
	}
}

/*GetPresetDefault handles this case with default header values.

An unexpected error response
*/
type GetPresetDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get preset default response
func (o *GetPresetDefault) Code() int {
	return o._statusCode
}

func (o *GetPresetDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/presets/{uuid}][%d] GetPreset default  %+v", o._statusCode, o.Payload)
}

func (o *GetPresetDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetPresetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
