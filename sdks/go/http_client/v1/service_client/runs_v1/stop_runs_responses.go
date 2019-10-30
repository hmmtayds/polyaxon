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

package runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// StopRunsReader is a Reader for the StopRuns structure.
type StopRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewStopRunsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStopRunsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopRunsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStopRunsOK creates a StopRunsOK with default headers values
func NewStopRunsOK() *StopRunsOK {
	return &StopRunsOK{}
}

/*StopRunsOK handles this case with default header values.

A successful response.
*/
type StopRunsOK struct {
}

func (o *StopRunsOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/stop][%d] stopRunsOK ", 200)
}

func (o *StopRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewStopRunsNoContent creates a StopRunsNoContent with default headers values
func NewStopRunsNoContent() *StopRunsNoContent {
	return &StopRunsNoContent{}
}

/*StopRunsNoContent handles this case with default header values.

No content.
*/
type StopRunsNoContent struct {
	Payload interface{}
}

func (o *StopRunsNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/stop][%d] stopRunsNoContent  %+v", 204, o.Payload)
}

func (o *StopRunsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *StopRunsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopRunsForbidden creates a StopRunsForbidden with default headers values
func NewStopRunsForbidden() *StopRunsForbidden {
	return &StopRunsForbidden{}
}

/*StopRunsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type StopRunsForbidden struct {
	Payload interface{}
}

func (o *StopRunsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/stop][%d] stopRunsForbidden  %+v", 403, o.Payload)
}

func (o *StopRunsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *StopRunsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopRunsNotFound creates a StopRunsNotFound with default headers values
func NewStopRunsNotFound() *StopRunsNotFound {
	return &StopRunsNotFound{}
}

/*StopRunsNotFound handles this case with default header values.

Resource does not exist.
*/
type StopRunsNotFound struct {
	Payload interface{}
}

func (o *StopRunsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/runs/stop][%d] stopRunsNotFound  %+v", 404, o.Payload)
}

func (o *StopRunsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *StopRunsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
