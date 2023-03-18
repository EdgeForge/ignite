// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Balloon Balloon device descriptor.
// swagger:model Balloon
type Balloon struct {

	// Target balloon size in MiB.
	// Required: true
	AmountMib *int64 `json:"amount_mib"`

	// Whether the balloon should deflate when the guest has memory pressure.
	// Required: true
	DeflateOnOom *bool `json:"deflate_on_oom"`

	// Interval in seconds between refreshing statistics. A non-zero value will enable the statistics. Defaults to 0.
	StatsPollingIntervals int64 `json:"stats_polling_interval_s,omitempty"`
}

// Validate validates this balloon
func (m *Balloon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmountMib(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeflateOnOom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Balloon) validateAmountMib(formats strfmt.Registry) error {

	if err := validate.Required("amount_mib", "body", m.AmountMib); err != nil {
		return err
	}

	return nil
}

func (m *Balloon) validateDeflateOnOom(formats strfmt.Registry) error {

	if err := validate.Required("deflate_on_oom", "body", m.DeflateOnOom); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Balloon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Balloon) UnmarshalBinary(b []byte) error {
	var res Balloon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}