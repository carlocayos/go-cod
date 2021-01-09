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

// MatchAnalysisResponse match analysis response
//
// swagger:model matchAnalysisResponse
type MatchAnalysisResponse struct {

	// data
	Data []*MatchAnalysisResponseDataItems0 `json:"data"`

	// status
	Status string `json:"status,omitempty"`

	// match analysis response additional properties
	MatchAnalysisResponseAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *MatchAnalysisResponse) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// data
		Data []*MatchAnalysisResponseDataItems0 `json:"data"`

		// status
		Status string `json:"status,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv MatchAnalysisResponse

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
		m.MatchAnalysisResponseAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m MatchAnalysisResponse) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// data
		Data []*MatchAnalysisResponseDataItems0 `json:"data"`

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

	if len(m.MatchAnalysisResponseAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.MatchAnalysisResponseAdditionalProperties)
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

// Validate validates this match analysis response
func (m *MatchAnalysisResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MatchAnalysisResponse) validateData(formats strfmt.Registry) error {

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
func (m *MatchAnalysisResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchAnalysisResponse) UnmarshalBinary(b []byte) error {
	var res MatchAnalysisResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MatchAnalysisResponseDataItems0 match analysis response data items0
//
// swagger:model MatchAnalysisResponseDataItems0
type MatchAnalysisResponseDataItems0 struct {

	// pay load
	PayLoad *MatchAnalysisResponseDataItems0PayLoad `json:"payLoad,omitempty"`

	// ws response
	WsResponse *MatchAnalysisResponseDataItems0WsResponse `json:"wsResponse,omitempty"`

	// match analysis response data items0 additional properties
	MatchAnalysisResponseDataItems0AdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *MatchAnalysisResponseDataItems0) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// pay load
		PayLoad *MatchAnalysisResponseDataItems0PayLoad `json:"payLoad,omitempty"`

		// ws response
		WsResponse *MatchAnalysisResponseDataItems0WsResponse `json:"wsResponse,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv MatchAnalysisResponseDataItems0

	rcv.PayLoad = stage1.PayLoad
	rcv.WsResponse = stage1.WsResponse
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "payLoad")
	delete(stage2, "wsResponse")
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
		m.MatchAnalysisResponseDataItems0AdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m MatchAnalysisResponseDataItems0) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// pay load
		PayLoad *MatchAnalysisResponseDataItems0PayLoad `json:"payLoad,omitempty"`

		// ws response
		WsResponse *MatchAnalysisResponseDataItems0WsResponse `json:"wsResponse,omitempty"`
	}

	stage1.PayLoad = m.PayLoad
	stage1.WsResponse = m.WsResponse

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.MatchAnalysisResponseDataItems0AdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.MatchAnalysisResponseDataItems0AdditionalProperties)
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

// Validate validates this match analysis response data items0
func (m *MatchAnalysisResponseDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayLoad(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWsResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MatchAnalysisResponseDataItems0) validatePayLoad(formats strfmt.Registry) error {

	if swag.IsZero(m.PayLoad) { // not required
		return nil
	}

	if m.PayLoad != nil {
		if err := m.PayLoad.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payLoad")
			}
			return err
		}
	}

	return nil
}

func (m *MatchAnalysisResponseDataItems0) validateWsResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.WsResponse) { // not required
		return nil
	}

	if m.WsResponse != nil {
		if err := m.WsResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wsResponse")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MatchAnalysisResponseDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchAnalysisResponseDataItems0) UnmarshalBinary(b []byte) error {
	var res MatchAnalysisResponseDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MatchAnalysisResponseDataItems0PayLoad match analysis response data items0 pay load
//
// swagger:model MatchAnalysisResponseDataItems0PayLoad
type MatchAnalysisResponseDataItems0PayLoad struct {

	// accuracy
	Accuracy string `json:"accuracy,omitempty"`

	// all players
	AllPlayers string `json:"allPlayers,omitempty"`

	// arena
	Arena string `json:"arena,omitempty"`

	// assists
	Assists string `json:"assists,omitempty"`

	// average speed during match
	AverageSpeedDuringMatch string `json:"averageSpeedDuringMatch,omitempty"`

	// damage done
	DamageDone string `json:"damageDone,omitempty"`

	// damage taken
	DamageTaken string `json:"damageTaken,omitempty"`

	// deaths
	Deaths string `json:"deaths,omitempty"`

	// distance traveled
	DistanceTraveled string `json:"distanceTraveled,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// executions
	Executions string `json:"executions,omitempty"`

	// game battle
	GameBattle string `json:"gameBattle,omitempty"`

	// game type
	GameType string `json:"gameType,omitempty"`

	// game type long
	GameTypeLong string `json:"game_type,omitempty"`

	// game name
	GameName string `json:"game_name,omitempty"`

	// headshots
	Headshots string `json:"headshots,omitempty"`

	// is present at end
	IsPresentAtEnd string `json:"isPresentAtEnd,omitempty"`

	// kd ratio
	KdRatio string `json:"kdRatio,omitempty"`

	// kills
	Kills string `json:"kills,omitempty"`

	// longest streak
	LongestStreak string `json:"longestStreak,omitempty"`

	// map
	Map string `json:"map,omitempty"`

	// map name
	MapName string `json:"map_name,omitempty"`

	// match ID
	MatchID string `json:"matchID,omitempty"`

	// match xp
	MatchXp string `json:"matchXp,omitempty"`

	// medal xp
	MedalXp string `json:"medalXp,omitempty"`

	// misc xp
	MiscXp string `json:"miscXp,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// mode name
	ModeName string `json:"mode_name,omitempty"`

	// nearmisses
	Nearmisses string `json:"nearmisses,omitempty"`

	// objective destroyed equipment
	ObjectiveDestroyedEquipment string `json:"objectiveDestroyedEquipment,omitempty"`

	// percent time moving
	PercentTimeMoving string `json:"percentTimeMoving,omitempty"`

	// place
	Place string `json:"place,omitempty"`

	// player
	Player string `json:"player,omitempty"`

	// playlist name
	PlaylistName string `json:"playlistName,omitempty"`

	// private match
	PrivateMatch string `json:"privateMatch,omitempty"`

	// rank
	Rank string `json:"rank,omitempty"`

	// result
	Result string `json:"result,omitempty"`

	// score
	Score string `json:"score,omitempty"`

	// score per minute
	ScorePerMinute string `json:"scorePerMinute,omitempty"`

	// score xp
	ScoreXp string `json:"scoreXp,omitempty"`

	// season rank
	SeasonRank string `json:"seasonRank,omitempty"`

	// shots fired
	ShotsFired string `json:"shotsFired,omitempty"`

	// shots landed
	ShotsLanded string `json:"shotsLanded,omitempty"`

	// shots missed
	ShotsMissed string `json:"shotsMissed,omitempty"`

	// suicides
	Suicides string `json:"suicides,omitempty"`

	// team1 score
	Team1Score string `json:"team1Score,omitempty"`

	// team2 score
	Team2Score string `json:"team2Score,omitempty"`

	// time played
	TimePlayed string `json:"timePlayed,omitempty"`

	// total xp
	TotalXp string `json:"totalXp,omitempty"`

	// utc end seconds
	UtcEndSeconds string `json:"utcEndSeconds,omitempty"`

	// utc start seconds
	UtcStartSeconds string `json:"utcStartSeconds,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// wall bangs
	WallBangs string `json:"wallBangs,omitempty"`

	// weapon stats
	WeaponStats string `json:"weaponStats,omitempty"`

	// winning team
	WinningTeam string `json:"winningTeam,omitempty"`

	// match analysis response data items0 pay load additional properties
	MatchAnalysisResponseDataItems0PayLoadAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *MatchAnalysisResponseDataItems0PayLoad) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// accuracy
		Accuracy string `json:"accuracy,omitempty"`

		// all players
		AllPlayers string `json:"allPlayers,omitempty"`

		// arena
		Arena string `json:"arena,omitempty"`

		// assists
		Assists string `json:"assists,omitempty"`

		// average speed during match
		AverageSpeedDuringMatch string `json:"averageSpeedDuringMatch,omitempty"`

		// damage done
		DamageDone string `json:"damageDone,omitempty"`

		// damage taken
		DamageTaken string `json:"damageTaken,omitempty"`

		// deaths
		Deaths string `json:"deaths,omitempty"`

		// distance traveled
		DistanceTraveled string `json:"distanceTraveled,omitempty"`

		// duration
		Duration string `json:"duration,omitempty"`

		// executions
		Executions string `json:"executions,omitempty"`

		// game battle
		GameBattle string `json:"gameBattle,omitempty"`

		// game type
		GameType string `json:"gameType,omitempty"`

		// game type long
		GameTypeLong string `json:"game_type,omitempty"`

		// game name
		GameName string `json:"game_name,omitempty"`

		// headshots
		Headshots string `json:"headshots,omitempty"`

		// is present at end
		IsPresentAtEnd string `json:"isPresentAtEnd,omitempty"`

		// kd ratio
		KdRatio string `json:"kdRatio,omitempty"`

		// kills
		Kills string `json:"kills,omitempty"`

		// longest streak
		LongestStreak string `json:"longestStreak,omitempty"`

		// map
		Map string `json:"map,omitempty"`

		// map name
		MapName string `json:"map_name,omitempty"`

		// match ID
		MatchID string `json:"matchID,omitempty"`

		// match xp
		MatchXp string `json:"matchXp,omitempty"`

		// medal xp
		MedalXp string `json:"medalXp,omitempty"`

		// misc xp
		MiscXp string `json:"miscXp,omitempty"`

		// mode
		Mode string `json:"mode,omitempty"`

		// mode name
		ModeName string `json:"mode_name,omitempty"`

		// nearmisses
		Nearmisses string `json:"nearmisses,omitempty"`

		// objective destroyed equipment
		ObjectiveDestroyedEquipment string `json:"objectiveDestroyedEquipment,omitempty"`

		// percent time moving
		PercentTimeMoving string `json:"percentTimeMoving,omitempty"`

		// place
		Place string `json:"place,omitempty"`

		// player
		Player string `json:"player,omitempty"`

		// playlist name
		PlaylistName string `json:"playlistName,omitempty"`

		// private match
		PrivateMatch string `json:"privateMatch,omitempty"`

		// rank
		Rank string `json:"rank,omitempty"`

		// result
		Result string `json:"result,omitempty"`

		// score
		Score string `json:"score,omitempty"`

		// score per minute
		ScorePerMinute string `json:"scorePerMinute,omitempty"`

		// score xp
		ScoreXp string `json:"scoreXp,omitempty"`

		// season rank
		SeasonRank string `json:"seasonRank,omitempty"`

		// shots fired
		ShotsFired string `json:"shotsFired,omitempty"`

		// shots landed
		ShotsLanded string `json:"shotsLanded,omitempty"`

		// shots missed
		ShotsMissed string `json:"shotsMissed,omitempty"`

		// suicides
		Suicides string `json:"suicides,omitempty"`

		// team1 score
		Team1Score string `json:"team1Score,omitempty"`

		// team2 score
		Team2Score string `json:"team2Score,omitempty"`

		// time played
		TimePlayed string `json:"timePlayed,omitempty"`

		// total xp
		TotalXp string `json:"totalXp,omitempty"`

		// utc end seconds
		UtcEndSeconds string `json:"utcEndSeconds,omitempty"`

		// utc start seconds
		UtcStartSeconds string `json:"utcStartSeconds,omitempty"`

		// version
		Version string `json:"version,omitempty"`

		// wall bangs
		WallBangs string `json:"wallBangs,omitempty"`

		// weapon stats
		WeaponStats string `json:"weaponStats,omitempty"`

		// winning team
		WinningTeam string `json:"winningTeam,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv MatchAnalysisResponseDataItems0PayLoad

	rcv.Accuracy = stage1.Accuracy
	rcv.AllPlayers = stage1.AllPlayers
	rcv.Arena = stage1.Arena
	rcv.Assists = stage1.Assists
	rcv.AverageSpeedDuringMatch = stage1.AverageSpeedDuringMatch
	rcv.DamageDone = stage1.DamageDone
	rcv.DamageTaken = stage1.DamageTaken
	rcv.Deaths = stage1.Deaths
	rcv.DistanceTraveled = stage1.DistanceTraveled
	rcv.Duration = stage1.Duration
	rcv.Executions = stage1.Executions
	rcv.GameBattle = stage1.GameBattle
	rcv.GameType = stage1.GameType
	rcv.GameTypeLong = stage1.GameTypeLong
	rcv.GameName = stage1.GameName
	rcv.Headshots = stage1.Headshots
	rcv.IsPresentAtEnd = stage1.IsPresentAtEnd
	rcv.KdRatio = stage1.KdRatio
	rcv.Kills = stage1.Kills
	rcv.LongestStreak = stage1.LongestStreak
	rcv.Map = stage1.Map
	rcv.MapName = stage1.MapName
	rcv.MatchID = stage1.MatchID
	rcv.MatchXp = stage1.MatchXp
	rcv.MedalXp = stage1.MedalXp
	rcv.MiscXp = stage1.MiscXp
	rcv.Mode = stage1.Mode
	rcv.ModeName = stage1.ModeName
	rcv.Nearmisses = stage1.Nearmisses
	rcv.ObjectiveDestroyedEquipment = stage1.ObjectiveDestroyedEquipment
	rcv.PercentTimeMoving = stage1.PercentTimeMoving
	rcv.Place = stage1.Place
	rcv.Player = stage1.Player
	rcv.PlaylistName = stage1.PlaylistName
	rcv.PrivateMatch = stage1.PrivateMatch
	rcv.Rank = stage1.Rank
	rcv.Result = stage1.Result
	rcv.Score = stage1.Score
	rcv.ScorePerMinute = stage1.ScorePerMinute
	rcv.ScoreXp = stage1.ScoreXp
	rcv.SeasonRank = stage1.SeasonRank
	rcv.ShotsFired = stage1.ShotsFired
	rcv.ShotsLanded = stage1.ShotsLanded
	rcv.ShotsMissed = stage1.ShotsMissed
	rcv.Suicides = stage1.Suicides
	rcv.Team1Score = stage1.Team1Score
	rcv.Team2Score = stage1.Team2Score
	rcv.TimePlayed = stage1.TimePlayed
	rcv.TotalXp = stage1.TotalXp
	rcv.UtcEndSeconds = stage1.UtcEndSeconds
	rcv.UtcStartSeconds = stage1.UtcStartSeconds
	rcv.Version = stage1.Version
	rcv.WallBangs = stage1.WallBangs
	rcv.WeaponStats = stage1.WeaponStats
	rcv.WinningTeam = stage1.WinningTeam
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "accuracy")
	delete(stage2, "allPlayers")
	delete(stage2, "arena")
	delete(stage2, "assists")
	delete(stage2, "averageSpeedDuringMatch")
	delete(stage2, "damageDone")
	delete(stage2, "damageTaken")
	delete(stage2, "deaths")
	delete(stage2, "distanceTraveled")
	delete(stage2, "duration")
	delete(stage2, "executions")
	delete(stage2, "gameBattle")
	delete(stage2, "gameType")
	delete(stage2, "gameTypeLong")
	delete(stage2, "game_name")
	delete(stage2, "headshots")
	delete(stage2, "isPresentAtEnd")
	delete(stage2, "kdRatio")
	delete(stage2, "kills")
	delete(stage2, "longestStreak")
	delete(stage2, "map")
	delete(stage2, "map_name")
	delete(stage2, "matchID")
	delete(stage2, "matchXp")
	delete(stage2, "medalXp")
	delete(stage2, "miscXp")
	delete(stage2, "mode")
	delete(stage2, "mode_name")
	delete(stage2, "nearmisses")
	delete(stage2, "objectiveDestroyedEquipment")
	delete(stage2, "percentTimeMoving")
	delete(stage2, "place")
	delete(stage2, "player")
	delete(stage2, "playlistName")
	delete(stage2, "privateMatch")
	delete(stage2, "rank")
	delete(stage2, "result")
	delete(stage2, "score")
	delete(stage2, "scorePerMinute")
	delete(stage2, "scoreXp")
	delete(stage2, "seasonRank")
	delete(stage2, "shotsFired")
	delete(stage2, "shotsLanded")
	delete(stage2, "shotsMissed")
	delete(stage2, "suicides")
	delete(stage2, "team1Score")
	delete(stage2, "team2Score")
	delete(stage2, "timePlayed")
	delete(stage2, "totalXp")
	delete(stage2, "utcEndSeconds")
	delete(stage2, "utcStartSeconds")
	delete(stage2, "version")
	delete(stage2, "wallBangs")
	delete(stage2, "weaponStats")
	delete(stage2, "winningTeam")
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
		m.MatchAnalysisResponseDataItems0PayLoadAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m MatchAnalysisResponseDataItems0PayLoad) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// accuracy
		Accuracy string `json:"accuracy,omitempty"`

		// all players
		AllPlayers string `json:"allPlayers,omitempty"`

		// arena
		Arena string `json:"arena,omitempty"`

		// assists
		Assists string `json:"assists,omitempty"`

		// average speed during match
		AverageSpeedDuringMatch string `json:"averageSpeedDuringMatch,omitempty"`

		// damage done
		DamageDone string `json:"damageDone,omitempty"`

		// damage taken
		DamageTaken string `json:"damageTaken,omitempty"`

		// deaths
		Deaths string `json:"deaths,omitempty"`

		// distance traveled
		DistanceTraveled string `json:"distanceTraveled,omitempty"`

		// duration
		Duration string `json:"duration,omitempty"`

		// executions
		Executions string `json:"executions,omitempty"`

		// game battle
		GameBattle string `json:"gameBattle,omitempty"`

		// game type
		GameType string `json:"gameType,omitempty"`

		// game type long
		GameTypeLong string `json:"game_type,omitempty"`

		// game name
		GameName string `json:"game_name,omitempty"`

		// headshots
		Headshots string `json:"headshots,omitempty"`

		// is present at end
		IsPresentAtEnd string `json:"isPresentAtEnd,omitempty"`

		// kd ratio
		KdRatio string `json:"kdRatio,omitempty"`

		// kills
		Kills string `json:"kills,omitempty"`

		// longest streak
		LongestStreak string `json:"longestStreak,omitempty"`

		// map
		Map string `json:"map,omitempty"`

		// map name
		MapName string `json:"map_name,omitempty"`

		// match ID
		MatchID string `json:"matchID,omitempty"`

		// match xp
		MatchXp string `json:"matchXp,omitempty"`

		// medal xp
		MedalXp string `json:"medalXp,omitempty"`

		// misc xp
		MiscXp string `json:"miscXp,omitempty"`

		// mode
		Mode string `json:"mode,omitempty"`

		// mode name
		ModeName string `json:"mode_name,omitempty"`

		// nearmisses
		Nearmisses string `json:"nearmisses,omitempty"`

		// objective destroyed equipment
		ObjectiveDestroyedEquipment string `json:"objectiveDestroyedEquipment,omitempty"`

		// percent time moving
		PercentTimeMoving string `json:"percentTimeMoving,omitempty"`

		// place
		Place string `json:"place,omitempty"`

		// player
		Player string `json:"player,omitempty"`

		// playlist name
		PlaylistName string `json:"playlistName,omitempty"`

		// private match
		PrivateMatch string `json:"privateMatch,omitempty"`

		// rank
		Rank string `json:"rank,omitempty"`

		// result
		Result string `json:"result,omitempty"`

		// score
		Score string `json:"score,omitempty"`

		// score per minute
		ScorePerMinute string `json:"scorePerMinute,omitempty"`

		// score xp
		ScoreXp string `json:"scoreXp,omitempty"`

		// season rank
		SeasonRank string `json:"seasonRank,omitempty"`

		// shots fired
		ShotsFired string `json:"shotsFired,omitempty"`

		// shots landed
		ShotsLanded string `json:"shotsLanded,omitempty"`

		// shots missed
		ShotsMissed string `json:"shotsMissed,omitempty"`

		// suicides
		Suicides string `json:"suicides,omitempty"`

		// team1 score
		Team1Score string `json:"team1Score,omitempty"`

		// team2 score
		Team2Score string `json:"team2Score,omitempty"`

		// time played
		TimePlayed string `json:"timePlayed,omitempty"`

		// total xp
		TotalXp string `json:"totalXp,omitempty"`

		// utc end seconds
		UtcEndSeconds string `json:"utcEndSeconds,omitempty"`

		// utc start seconds
		UtcStartSeconds string `json:"utcStartSeconds,omitempty"`

		// version
		Version string `json:"version,omitempty"`

		// wall bangs
		WallBangs string `json:"wallBangs,omitempty"`

		// weapon stats
		WeaponStats string `json:"weaponStats,omitempty"`

		// winning team
		WinningTeam string `json:"winningTeam,omitempty"`
	}

	stage1.Accuracy = m.Accuracy
	stage1.AllPlayers = m.AllPlayers
	stage1.Arena = m.Arena
	stage1.Assists = m.Assists
	stage1.AverageSpeedDuringMatch = m.AverageSpeedDuringMatch
	stage1.DamageDone = m.DamageDone
	stage1.DamageTaken = m.DamageTaken
	stage1.Deaths = m.Deaths
	stage1.DistanceTraveled = m.DistanceTraveled
	stage1.Duration = m.Duration
	stage1.Executions = m.Executions
	stage1.GameBattle = m.GameBattle
	stage1.GameType = m.GameType
	stage1.GameTypeLong = m.GameTypeLong
	stage1.GameName = m.GameName
	stage1.Headshots = m.Headshots
	stage1.IsPresentAtEnd = m.IsPresentAtEnd
	stage1.KdRatio = m.KdRatio
	stage1.Kills = m.Kills
	stage1.LongestStreak = m.LongestStreak
	stage1.Map = m.Map
	stage1.MapName = m.MapName
	stage1.MatchID = m.MatchID
	stage1.MatchXp = m.MatchXp
	stage1.MedalXp = m.MedalXp
	stage1.MiscXp = m.MiscXp
	stage1.Mode = m.Mode
	stage1.ModeName = m.ModeName
	stage1.Nearmisses = m.Nearmisses
	stage1.ObjectiveDestroyedEquipment = m.ObjectiveDestroyedEquipment
	stage1.PercentTimeMoving = m.PercentTimeMoving
	stage1.Place = m.Place
	stage1.Player = m.Player
	stage1.PlaylistName = m.PlaylistName
	stage1.PrivateMatch = m.PrivateMatch
	stage1.Rank = m.Rank
	stage1.Result = m.Result
	stage1.Score = m.Score
	stage1.ScorePerMinute = m.ScorePerMinute
	stage1.ScoreXp = m.ScoreXp
	stage1.SeasonRank = m.SeasonRank
	stage1.ShotsFired = m.ShotsFired
	stage1.ShotsLanded = m.ShotsLanded
	stage1.ShotsMissed = m.ShotsMissed
	stage1.Suicides = m.Suicides
	stage1.Team1Score = m.Team1Score
	stage1.Team2Score = m.Team2Score
	stage1.TimePlayed = m.TimePlayed
	stage1.TotalXp = m.TotalXp
	stage1.UtcEndSeconds = m.UtcEndSeconds
	stage1.UtcStartSeconds = m.UtcStartSeconds
	stage1.Version = m.Version
	stage1.WallBangs = m.WallBangs
	stage1.WeaponStats = m.WeaponStats
	stage1.WinningTeam = m.WinningTeam

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.MatchAnalysisResponseDataItems0PayLoadAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.MatchAnalysisResponseDataItems0PayLoadAdditionalProperties)
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

// Validate validates this match analysis response data items0 pay load
func (m *MatchAnalysisResponseDataItems0PayLoad) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MatchAnalysisResponseDataItems0PayLoad) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchAnalysisResponseDataItems0PayLoad) UnmarshalBinary(b []byte) error {
	var res MatchAnalysisResponseDataItems0PayLoad
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MatchAnalysisResponseDataItems0WsResponse match analysis response data items0 ws response
//
// swagger:model MatchAnalysisResponseDataItems0WsResponse
type MatchAnalysisResponseDataItems0WsResponse struct {

	// data
	Data *MatchAnalysisResponseDataItems0WsResponseData `json:"data,omitempty"`

	// match analysis response data items0 ws response additional properties
	MatchAnalysisResponseDataItems0WsResponseAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *MatchAnalysisResponseDataItems0WsResponse) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// data
		Data *MatchAnalysisResponseDataItems0WsResponseData `json:"data,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv MatchAnalysisResponseDataItems0WsResponse

	rcv.Data = stage1.Data
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "data")
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
		m.MatchAnalysisResponseDataItems0WsResponseAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m MatchAnalysisResponseDataItems0WsResponse) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// data
		Data *MatchAnalysisResponseDataItems0WsResponseData `json:"data,omitempty"`
	}

	stage1.Data = m.Data

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.MatchAnalysisResponseDataItems0WsResponseAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.MatchAnalysisResponseDataItems0WsResponseAdditionalProperties)
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

// Validate validates this match analysis response data items0 ws response
func (m *MatchAnalysisResponseDataItems0WsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MatchAnalysisResponseDataItems0WsResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wsResponse" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MatchAnalysisResponseDataItems0WsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchAnalysisResponseDataItems0WsResponse) UnmarshalBinary(b []byte) error {
	var res MatchAnalysisResponseDataItems0WsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MatchAnalysisResponseDataItems0WsResponseData match analysis response data items0 ws response data
//
// swagger:model MatchAnalysisResponseDataItems0WsResponseData
type MatchAnalysisResponseDataItems0WsResponseData struct {

	// content
	Content string `json:"content,omitempty"`

	// errors
	Errors string `json:"errors,omitempty"`

	// match analysis response data items0 ws response data additional properties
	MatchAnalysisResponseDataItems0WsResponseDataAdditionalProperties map[string]interface{} `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *MatchAnalysisResponseDataItems0WsResponseData) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// content
		Content string `json:"content,omitempty"`

		// errors
		Errors string `json:"errors,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv MatchAnalysisResponseDataItems0WsResponseData

	rcv.Content = stage1.Content
	rcv.Errors = stage1.Errors
	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "content")
	delete(stage2, "errors")
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
		m.MatchAnalysisResponseDataItems0WsResponseDataAdditionalProperties = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m MatchAnalysisResponseDataItems0WsResponseData) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// content
		Content string `json:"content,omitempty"`

		// errors
		Errors string `json:"errors,omitempty"`
	}

	stage1.Content = m.Content
	stage1.Errors = m.Errors

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.MatchAnalysisResponseDataItems0WsResponseDataAdditionalProperties) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.MatchAnalysisResponseDataItems0WsResponseDataAdditionalProperties)
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

// Validate validates this match analysis response data items0 ws response data
func (m *MatchAnalysisResponseDataItems0WsResponseData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MatchAnalysisResponseDataItems0WsResponseData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchAnalysisResponseDataItems0WsResponseData) UnmarshalBinary(b []byte) error {
	var res MatchAnalysisResponseDataItems0WsResponseData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
