// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GamerLootResponse gamer loot response
//
// swagger:model gamerLootResponse
type GamerLootResponse struct {

	// data
	Data *GamerLootResponseData `json:"data,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// gamer loot response additional properties
	GamerLootResponseAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *GamerLootResponse) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// data
		Data *GamerLootResponseData `json:"data,omitempty"`

		// status
		Status string `json:"status,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv GamerLootResponse

	rcv.Data = stage1.Data
	rcv.Status = stage1.Status
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "data")
	delete(stage2, "status")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.GamerLootResponseAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m GamerLootResponse) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// data
		Data *GamerLootResponseData `json:"data,omitempty"`

		// status
		Status string `json:"status,omitempty"`
	}

	stage1.Data = m.Data
	stage1.Status = m.Status

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.GamerLootResponseAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.GamerLootResponseAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this gamer loot response
func (m *GamerLootResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GamerLootResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GamerLootResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GamerLootResponse) UnmarshalBinary(b []byte) error {
	var res GamerLootResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GamerLootResponseData gamer loot response data
//
// swagger:model GamerLootResponseData
type GamerLootResponseData struct {

	// message
	Message string `json:"message,omitempty"`

	// streams
	Streams map[string]GamerLootResponseDataStreamsAnon `json:"streams,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// gamer loot response data additional properties
	GamerLootResponseDataAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *GamerLootResponseData) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// message
		Message string `json:"message,omitempty"`

		// streams
		Streams map[string]GamerLootResponseDataStreamsAnon `json:"streams,omitempty"`

		// type
		Type string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv GamerLootResponseData

	rcv.Message = stage1.Message
	rcv.Streams = stage1.Streams
	rcv.Type = stage1.Type
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "message")
	delete(stage2, "streams")
	delete(stage2, "type")
	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]interface{})
		for k, v := range stage2 {
			var toadd interface{}
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.GamerLootResponseDataAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m GamerLootResponseData) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// message
		Message string `json:"message,omitempty"`

		// streams
		Streams map[string]GamerLootResponseDataStreamsAnon `json:"streams,omitempty"`

		// type
		Type string `json:"type,omitempty"`
	}

	stage1.Message = m.Message
	stage1.Streams = m.Streams
	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.GamerLootResponseDataAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.GamerLootResponseDataAdditionalProperties)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this gamer loot response data
func (m *GamerLootResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStreams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GamerLootResponseData) validateStreams(formats strfmt.Registry) error {

	if swag.IsZero(m.Streams) { // not required
		return nil
	}

	for k := range m.Streams {

		if swag.IsZero(m.Streams[k]) { // not required
			continue
		}
		if val, ok := m.Streams[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GamerLootResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GamerLootResponseData) UnmarshalBinary(b []byte) error {
	var res GamerLootResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GamerLootResponseDataStreamsAnon gamer loot response data streams anon
//
// swagger:model GamerLootResponseDataStreamsAnon
type GamerLootResponseDataStreamsAnon struct {

	// category name label
	CategoryNameLabel string `json:"categoryNameLabel,omitempty"`

	// category title label
	CategoryTitleLabel string `json:"categoryTitleLabel,omitempty"`

	// items obtained
	ItemsObtained string `json:"itemsObtained,omitempty"`

	// loot type
	LootType string `json:"lootType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// premium
	Premium bool `json:"premium,omitempty"`

	// premium token owned not redeemed
	PremiumTokenOwnedNotRedeemed bool `json:"premiumTokenOwnedNotRedeemed,omitempty"`

	// rarity
	Rarity string `json:"rarity,omitempty"`

	// season info Url key
	SeasonInfoURLKey string `json:"seasonInfoUrlKey,omitempty"`

	// stream type
	StreamType string `json:"streamType,omitempty"`

	// tier
	Tier int64 `json:"tier,omitempty"`

	// tier skip tokens unredeemed
	TierSkipTokensUnredeemed string `json:"tierSkipTokensUnredeemed,omitempty"`
}

// Validate validates this gamer loot response data streams anon
func (m *GamerLootResponseDataStreamsAnon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GamerLootResponseDataStreamsAnon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GamerLootResponseDataStreamsAnon) UnmarshalBinary(b []byte) error {
	var res GamerLootResponseDataStreamsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
