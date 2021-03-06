// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetNeighboringAccessPointsResponse get neighboring access points response
// swagger:model getNeighboringAccessPointsResponse
type GetNeighboringAccessPointsResponse struct {

	// data
	Data []*GetNeighboringAccessPointsResponseDataItems0 `json:"data"`
}

// Validate validates this get neighboring access points response
func (m *GetNeighboringAccessPointsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetNeighboringAccessPointsResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetNeighboringAccessPointsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNeighboringAccessPointsResponse) UnmarshalBinary(b []byte) error {
	var res GetNeighboringAccessPointsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetNeighboringAccessPointsResponseDataItems0 get neighboring access points response data items0
// swagger:model GetNeighboringAccessPointsResponseDataItems0
type GetNeighboringAccessPointsResponseDataItems0 struct {

	// id
	ID string `json:"_id,omitempty"`

	// age
	Age int64 `json:"age,omitempty"`

	// ap mac
	ApMac string `json:"ap_mac,omitempty"`

	// band
	Band string `json:"band,omitempty"`

	// bssid
	Bssid string `json:"bssid,omitempty"`

	// bw
	Bw int64 `json:"bw,omitempty"`

	// center freq
	CenterFreq int64 `json:"center_freq,omitempty"`

	// channel
	Channel int64 `json:"channel,omitempty"`

	// essid
	Essid string `json:"essid,omitempty"`

	// freq
	Freq int64 `json:"freq,omitempty"`

	// is adhoc
	IsAdhoc bool `json:"is_adhoc,omitempty"`

	// is rogue
	IsRogue bool `json:"is_rogue,omitempty"`

	// is ubnt
	IsUbnt bool `json:"is_ubnt,omitempty"`

	// last seen
	LastSeen int64 `json:"last_seen,omitempty"`

	// noise
	Noise int64 `json:"noise,omitempty"`

	// oui
	Oui string `json:"oui,omitempty"`

	// radio
	Radio string `json:"radio,omitempty"`

	// radio name
	RadioName string `json:"radio_name,omitempty"`

	// report time
	ReportTime int64 `json:"report_time,omitempty"`

	// rssi
	Rssi int64 `json:"rssi,omitempty"`

	// rssi age
	RssiAge int64 `json:"rssi_age,omitempty"`

	// security
	Security string `json:"security,omitempty"`

	// signal
	Signal int64 `json:"signal,omitempty"`

	// site id
	SiteID string `json:"site_id,omitempty"`
}

// Validate validates this get neighboring access points response data items0
func (m *GetNeighboringAccessPointsResponseDataItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetNeighboringAccessPointsResponseDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetNeighboringAccessPointsResponseDataItems0) UnmarshalBinary(b []byte) error {
	var res GetNeighboringAccessPointsResponseDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
