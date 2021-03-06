// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LeaderBoardResponse leader board response
//
// swagger:model leaderBoardResponse
type LeaderBoardResponse struct {

	// data
	Data *LeaderBoardResponseData `json:"data,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// leader board response additional properties
	LeaderBoardResponseAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *LeaderBoardResponse) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// data
		Data *LeaderBoardResponseData `json:"data,omitempty"`

		// status
		Status string `json:"status,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv LeaderBoardResponse

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
		m.LeaderBoardResponseAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m LeaderBoardResponse) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// data
		Data *LeaderBoardResponseData `json:"data,omitempty"`

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

	if len(m.LeaderBoardResponseAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.LeaderBoardResponseAdditionalProperties)
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

// Validate validates this leader board response
func (m *LeaderBoardResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LeaderBoardResponse) validateData(formats strfmt.Registry) error {

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
func (m *LeaderBoardResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaderBoardResponse) UnmarshalBinary(b []byte) error {
	var res LeaderBoardResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LeaderBoardResponseData leader board response data
//
// swagger:model LeaderBoardResponseData
type LeaderBoardResponseData struct {

	// columns
	Columns []string `json:"columns"`

	// entries
	Entries []*LeaderBoardResponseDataEntriesItems0 `json:"entries"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// leaderboard type
	LeaderboardType string `json:"leaderboardType,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// page
	Page int64 `json:"page,omitempty"`

	// platform
	Platform string `json:"platform,omitempty"`

	// results requested
	ResultsRequested int64 `json:"resultsRequested,omitempty"`

	// sort
	Sort string `json:"sort,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// total pages
	TotalPages int64 `json:"totalPages,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// leader board response data additional properties
	LeaderBoardResponseDataAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *LeaderBoardResponseData) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// columns
		Columns []string `json:"columns"`

		// entries
		Entries []*LeaderBoardResponseDataEntriesItems0 `json:"entries"`

		// game mode
		GameMode string `json:"gameMode,omitempty"`

		// leaderboard type
		LeaderboardType string `json:"leaderboardType,omitempty"`

		// message
		Message string `json:"message,omitempty"`

		// page
		Page int64 `json:"page,omitempty"`

		// platform
		Platform string `json:"platform,omitempty"`

		// results requested
		ResultsRequested int64 `json:"resultsRequested,omitempty"`

		// sort
		Sort string `json:"sort,omitempty"`

		// title
		Title string `json:"title,omitempty"`

		// total pages
		TotalPages int64 `json:"totalPages,omitempty"`

		// type
		Type string `json:"type,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv LeaderBoardResponseData

	rcv.Columns = stage1.Columns
	rcv.Entries = stage1.Entries
	rcv.GameMode = stage1.GameMode
	rcv.LeaderboardType = stage1.LeaderboardType
	rcv.Message = stage1.Message
	rcv.Page = stage1.Page
	rcv.Platform = stage1.Platform
	rcv.ResultsRequested = stage1.ResultsRequested
	rcv.Sort = stage1.Sort
	rcv.Title = stage1.Title
	rcv.TotalPages = stage1.TotalPages
	rcv.Type = stage1.Type
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "columns")
	delete(stage2, "entries")
	delete(stage2, "gameMode")
	delete(stage2, "leaderboardType")
	delete(stage2, "message")
	delete(stage2, "page")
	delete(stage2, "platform")
	delete(stage2, "resultsRequested")
	delete(stage2, "sort")
	delete(stage2, "title")
	delete(stage2, "totalPages")
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
		m.LeaderBoardResponseDataAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m LeaderBoardResponseData) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// columns
		Columns []string `json:"columns"`

		// entries
		Entries []*LeaderBoardResponseDataEntriesItems0 `json:"entries"`

		// game mode
		GameMode string `json:"gameMode,omitempty"`

		// leaderboard type
		LeaderboardType string `json:"leaderboardType,omitempty"`

		// message
		Message string `json:"message,omitempty"`

		// page
		Page int64 `json:"page,omitempty"`

		// platform
		Platform string `json:"platform,omitempty"`

		// results requested
		ResultsRequested int64 `json:"resultsRequested,omitempty"`

		// sort
		Sort string `json:"sort,omitempty"`

		// title
		Title string `json:"title,omitempty"`

		// total pages
		TotalPages int64 `json:"totalPages,omitempty"`

		// type
		Type string `json:"type,omitempty"`
	}

	stage1.Columns = m.Columns
	stage1.Entries = m.Entries
	stage1.GameMode = m.GameMode
	stage1.LeaderboardType = m.LeaderboardType
	stage1.Message = m.Message
	stage1.Page = m.Page
	stage1.Platform = m.Platform
	stage1.ResultsRequested = m.ResultsRequested
	stage1.Sort = m.Sort
	stage1.Title = m.Title
	stage1.TotalPages = m.TotalPages
	stage1.Type = m.Type

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.LeaderBoardResponseDataAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.LeaderBoardResponseDataAdditionalProperties)
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

// Validate validates this leader board response data
func (m *LeaderBoardResponseData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LeaderBoardResponseData) validateEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.Entries) { // not required
		return nil
	}

	for i := 0; i < len(m.Entries); i++ {
		if swag.IsZero(m.Entries[i]) { // not required
			continue
		}

		if m.Entries[i] != nil {
			if err := m.Entries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + "entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LeaderBoardResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaderBoardResponseData) UnmarshalBinary(b []byte) error {
	var res LeaderBoardResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LeaderBoardResponseDataEntriesItems0 leader board response data entries items0
//
// swagger:model LeaderBoardResponseDataEntriesItems0
type LeaderBoardResponseDataEntriesItems0 struct {

	// rank
	Rank float64 `json:"rank,omitempty"`

	// rating
	Rating float64 `json:"rating,omitempty"`

	// update time
	UpdateTime float64 `json:"updateTime,omitempty"`

	// username
	Username string `json:"username,omitempty"`

	// values
	Values *LeaderBoardResponseDataEntriesItems0Values `json:"values,omitempty"`

	// leader board response data entries items0 additional properties
	LeaderBoardResponseDataEntriesItems0AdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *LeaderBoardResponseDataEntriesItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// rank
		Rank float64 `json:"rank,omitempty"`

		// rating
		Rating float64 `json:"rating,omitempty"`

		// update time
		UpdateTime float64 `json:"updateTime,omitempty"`

		// username
		Username string `json:"username,omitempty"`

		// values
		Values *LeaderBoardResponseDataEntriesItems0Values `json:"values,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv LeaderBoardResponseDataEntriesItems0

	rcv.Rank = stage1.Rank
	rcv.Rating = stage1.Rating
	rcv.UpdateTime = stage1.UpdateTime
	rcv.Username = stage1.Username
	rcv.Values = stage1.Values
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "rank")
	delete(stage2, "rating")
	delete(stage2, "updateTime")
	delete(stage2, "username")
	delete(stage2, "values")
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
		m.LeaderBoardResponseDataEntriesItems0AdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m LeaderBoardResponseDataEntriesItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// rank
		Rank float64 `json:"rank,omitempty"`

		// rating
		Rating float64 `json:"rating,omitempty"`

		// update time
		UpdateTime float64 `json:"updateTime,omitempty"`

		// username
		Username string `json:"username,omitempty"`

		// values
		Values *LeaderBoardResponseDataEntriesItems0Values `json:"values,omitempty"`
	}

	stage1.Rank = m.Rank
	stage1.Rating = m.Rating
	stage1.UpdateTime = m.UpdateTime
	stage1.Username = m.Username
	stage1.Values = m.Values

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.LeaderBoardResponseDataEntriesItems0AdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.LeaderBoardResponseDataEntriesItems0AdditionalProperties)
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

// Validate validates this leader board response data entries items0
func (m *LeaderBoardResponseDataEntriesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LeaderBoardResponseDataEntriesItems0) validateValues(formats strfmt.Registry) error {

	if swag.IsZero(m.Values) { // not required
		return nil
	}

	if m.Values != nil {
		if err := m.Values.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("values")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LeaderBoardResponseDataEntriesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaderBoardResponseDataEntriesItems0) UnmarshalBinary(b []byte) error {
	var res LeaderBoardResponseDataEntriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LeaderBoardResponseDataEntriesItems0Values leader board response data entries items0 values
//
// swagger:model LeaderBoardResponseDataEntriesItems0Values
type LeaderBoardResponseDataEntriesItems0Values struct {

	// accuracy
	Accuracy float64 `json:"accuracy,omitempty"`

	// assists
	Assists float64 `json:"assists,omitempty"`

	// average time
	AverageTime float64 `json:"averageTime,omitempty"`

	// deaths
	Deaths float64 `json:"deaths,omitempty"`

	// games played
	GamesPlayed float64 `json:"gamesPlayed,omitempty"`

	// headshots
	Headshots float64 `json:"headshots,omitempty"`

	// hits
	Hits float64 `json:"hits,omitempty"`

	// kd ratio
	KdRatio float64 `json:"kdRatio,omitempty"`

	// kills
	Kills float64 `json:"kills,omitempty"`

	// killstreak
	Killstreak float64 `json:"killstreak,omitempty"`

	// level
	Level float64 `json:"level,omitempty"`

	// losses
	Losses float64 `json:"losses,omitempty"`

	// misses
	Misses float64 `json:"misses,omitempty"`

	// prestige
	Prestige float64 `json:"prestige,omitempty"`

	// score
	Score float64 `json:"score,omitempty"`

	// score per minute
	ScorePerMinute float64 `json:"scorePerMinute,omitempty"`

	// shots
	Shots float64 `json:"shots,omitempty"`

	// time played
	TimePlayed float64 `json:"timePlayed,omitempty"`

	// wins
	Wins float64 `json:"wins,omitempty"`

	// xp
	Xp float64 `json:"xp,omitempty"`

	// leader board response data entries items0 values additional properties
	LeaderBoardResponseDataEntriesItems0ValuesAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *LeaderBoardResponseDataEntriesItems0Values) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// accuracy
		Accuracy float64 `json:"accuracy,omitempty"`

		// assists
		Assists float64 `json:"assists,omitempty"`

		// average time
		AverageTime float64 `json:"averageTime,omitempty"`

		// deaths
		Deaths float64 `json:"deaths,omitempty"`

		// games played
		GamesPlayed float64 `json:"gamesPlayed,omitempty"`

		// headshots
		Headshots float64 `json:"headshots,omitempty"`

		// hits
		Hits float64 `json:"hits,omitempty"`

		// kd ratio
		KdRatio float64 `json:"kdRatio,omitempty"`

		// kills
		Kills float64 `json:"kills,omitempty"`

		// killstreak
		Killstreak float64 `json:"killstreak,omitempty"`

		// level
		Level float64 `json:"level,omitempty"`

		// losses
		Losses float64 `json:"losses,omitempty"`

		// misses
		Misses float64 `json:"misses,omitempty"`

		// prestige
		Prestige float64 `json:"prestige,omitempty"`

		// score
		Score float64 `json:"score,omitempty"`

		// score per minute
		ScorePerMinute float64 `json:"scorePerMinute,omitempty"`

		// shots
		Shots float64 `json:"shots,omitempty"`

		// time played
		TimePlayed float64 `json:"timePlayed,omitempty"`

		// wins
		Wins float64 `json:"wins,omitempty"`

		// xp
		Xp float64 `json:"xp,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv LeaderBoardResponseDataEntriesItems0Values

	rcv.Accuracy = stage1.Accuracy
	rcv.Assists = stage1.Assists
	rcv.AverageTime = stage1.AverageTime
	rcv.Deaths = stage1.Deaths
	rcv.GamesPlayed = stage1.GamesPlayed
	rcv.Headshots = stage1.Headshots
	rcv.Hits = stage1.Hits
	rcv.KdRatio = stage1.KdRatio
	rcv.Kills = stage1.Kills
	rcv.Killstreak = stage1.Killstreak
	rcv.Level = stage1.Level
	rcv.Losses = stage1.Losses
	rcv.Misses = stage1.Misses
	rcv.Prestige = stage1.Prestige
	rcv.Score = stage1.Score
	rcv.ScorePerMinute = stage1.ScorePerMinute
	rcv.Shots = stage1.Shots
	rcv.TimePlayed = stage1.TimePlayed
	rcv.Wins = stage1.Wins
	rcv.Xp = stage1.Xp
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "accuracy")
	delete(stage2, "assists")
	delete(stage2, "averageTime")
	delete(stage2, "deaths")
	delete(stage2, "gamesPlayed")
	delete(stage2, "headshots")
	delete(stage2, "hits")
	delete(stage2, "kdRatio")
	delete(stage2, "kills")
	delete(stage2, "killstreak")
	delete(stage2, "level")
	delete(stage2, "losses")
	delete(stage2, "misses")
	delete(stage2, "prestige")
	delete(stage2, "score")
	delete(stage2, "scorePerMinute")
	delete(stage2, "shots")
	delete(stage2, "timePlayed")
	delete(stage2, "wins")
	delete(stage2, "xp")
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
		m.LeaderBoardResponseDataEntriesItems0ValuesAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m LeaderBoardResponseDataEntriesItems0Values) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// accuracy
		Accuracy float64 `json:"accuracy,omitempty"`

		// assists
		Assists float64 `json:"assists,omitempty"`

		// average time
		AverageTime float64 `json:"averageTime,omitempty"`

		// deaths
		Deaths float64 `json:"deaths,omitempty"`

		// games played
		GamesPlayed float64 `json:"gamesPlayed,omitempty"`

		// headshots
		Headshots float64 `json:"headshots,omitempty"`

		// hits
		Hits float64 `json:"hits,omitempty"`

		// kd ratio
		KdRatio float64 `json:"kdRatio,omitempty"`

		// kills
		Kills float64 `json:"kills,omitempty"`

		// killstreak
		Killstreak float64 `json:"killstreak,omitempty"`

		// level
		Level float64 `json:"level,omitempty"`

		// losses
		Losses float64 `json:"losses,omitempty"`

		// misses
		Misses float64 `json:"misses,omitempty"`

		// prestige
		Prestige float64 `json:"prestige,omitempty"`

		// score
		Score float64 `json:"score,omitempty"`

		// score per minute
		ScorePerMinute float64 `json:"scorePerMinute,omitempty"`

		// shots
		Shots float64 `json:"shots,omitempty"`

		// time played
		TimePlayed float64 `json:"timePlayed,omitempty"`

		// wins
		Wins float64 `json:"wins,omitempty"`

		// xp
		Xp float64 `json:"xp,omitempty"`
	}

	stage1.Accuracy = m.Accuracy
	stage1.Assists = m.Assists
	stage1.AverageTime = m.AverageTime
	stage1.Deaths = m.Deaths
	stage1.GamesPlayed = m.GamesPlayed
	stage1.Headshots = m.Headshots
	stage1.Hits = m.Hits
	stage1.KdRatio = m.KdRatio
	stage1.Kills = m.Kills
	stage1.Killstreak = m.Killstreak
	stage1.Level = m.Level
	stage1.Losses = m.Losses
	stage1.Misses = m.Misses
	stage1.Prestige = m.Prestige
	stage1.Score = m.Score
	stage1.ScorePerMinute = m.ScorePerMinute
	stage1.Shots = m.Shots
	stage1.TimePlayed = m.TimePlayed
	stage1.Wins = m.Wins
	stage1.Xp = m.Xp

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.LeaderBoardResponseDataEntriesItems0ValuesAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.LeaderBoardResponseDataEntriesItems0ValuesAdditionalProperties)
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

// Validate validates this leader board response data entries items0 values
func (m *LeaderBoardResponseDataEntriesItems0Values) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LeaderBoardResponseDataEntriesItems0Values) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LeaderBoardResponseDataEntriesItems0Values) UnmarshalBinary(b []byte) error {
	var res LeaderBoardResponseDataEntriesItems0Values
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
