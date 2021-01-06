package go_cod

import (
	"context"
	"fmt"
	"github.com/carlocayos/go-cod/api/client/service"
	"github.com/carlocayos/go-cod/api/models"
)

//	This file contains the authenticated and unauthenticated COD API operations
//
//	COD API Operations - Unauthenticated
//		- LeaderBoard
//		- Map List
//		- Full Match Info
//		- Battle Pass Loot
//		- Purchasable
//		- Loadout
//
//	COD API Operations - Authenticated
//		- Gamer Match Details
//		- Gamer Match List
//		- Gamer Stats Profile
//		- Gamer Friend Stats Profile
//		- Gamer Loot
//		- Match Analysis
//		- COD Points

// Gamer ID details
type Gamer struct {
	Platform   Platform
	LookupType LookupType
	GamerTag   string
}

//	------------------------------------------
//	COD API Operations - Unauthenticated
//	------------------------------------------

// LeaderBoard returns the leader board details
func (c *Client) LeaderBoard(ctx context.Context, title GameTitle, platform Platform, page int64) (*models.LeaderBoardResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.LeaderBoardParams{
		Context:  ctx,
		Title:    string(title),
		Platform: string(platform),
		Page:     page,
	}

	res, err := c.ServiceClient.Operations.LeaderBoard(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in leader board request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

// Map List returns a list of maps and its available game modes
func (c *Client) MapList(ctx context.Context, title GameTitle, platform Platform, gameType GameType) (*models.MapListResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.MapListParams{
		Context:  context.Background(),
		Title:    string(title),
		Platform: string(platform),
		GameType: string(gameType),
	}

	res, err := c.ServiceClient.Operations.MapList(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in map list request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

// FullMatchInfo returns the match information
func (c *Client) FullMatchInfo(ctx context.Context, title GameTitle, platform Platform, gameType GameType,
	matchId string) (*models.FullMatchInfoResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.FullMatchInfoParams{
		Context:  context.Background(),
		Title:    string(title),
		Platform: string(platform),
		GameType: string(gameType),
		MatchID:  matchId,
	}

	res, err := c.ServiceClient.Operations.FullMatchInfo(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in full match info request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

// BattlePassLoot returns the available loots for the specified battle pass season
func (c *Client) BattlePassLoot(ctx context.Context, title GameTitle, platform Platform, season int64) (*models.BattlePassLootResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.BattlePassLootParams{
		Context:  context.Background(),
		Title:    string(title),
		Platform: string(platform),
		Season:   season,
	}

	res, err := c.ServiceClient.Operations.BattlePassLoot(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in battlepass loot request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

// Purchasable returns all the purchasable items
func (c *Client) Purchasable(ctx context.Context, title GameTitle, platform Platform) (*models.PurchasableResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.PurchasableParams{
		Context:  context.Background(),
		Title:    string(title),
		Platform: string(platform),
	}

	res, err := c.ServiceClient.Operations.Purchasable(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in purchasbale request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

// Loadout returns all the loadout information
func (c *Client) Loadout(ctx context.Context, title GameTitle, gameType GameType) (*models.LoadoutResponse, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.LoadoutParams{
		Context:  context.Background(),
		Title:    string(title),
		GameType: string(gameType),
	}

	res, err := c.ServiceClient.Operations.Loadout(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in loadout request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

//	------------------------------------------
//	COD API Operations - Authenticated
//	------------------------------------------

// LoggedIn checks if the user has logged in and generated a valid token
func (c *Client) LoggedIn() error {
	if c.AuthToken == nil || c.AuthToken.ActSSOCookie == "" {
		return fmt.Errorf("must be logged in to use an authenticated request")
	}
	return nil
}

// GamerMatchDetails returns the gamer match details. Authentication required
func (c *Client) GamerMatchDetails(ctx context.Context, title GameTitle, gamer *Gamer, gameType GameType,
	start int64, end int64, limit int32) (*models.GamerMatchDetailsResponse, error) {

	if err := c.LoggedIn(); err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.GamerMatchDetailsParams{
		Context:    ctx,
		Cookie:     ActSSOCookie + "=" + c.AuthToken.ActSSOCookie,
		Title:      string(title),
		Platform:   string(gamer.Platform),
		LookupType: string(gamer.LookupType),
		Gamertag:   gamer.GamerTag,
		GameType:   string(gameType),
		Start:      start,
		End:        end,
	}

	// assign limit if set to a positive integer, otherwise omit
	if limit > 0 {
		param.Limit = &limit
	}
	res, err := c.ServiceClient.Operations.GamerMatchDetails(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in gamer match details request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

// GamerMatchList returns a list of gamer matches. Authentication required
func (c *Client) GamerMatchList(ctx context.Context, title GameTitle, gamer *Gamer, gameType GameType,
	start int64, end int64, limit int32) (*models.GamerMatchListResponse, error) {

	if err := c.LoggedIn(); err != nil {
		return nil, err
	}

	if ctx == nil {
		ctx = context.Background()
	}

	param := service.GamerMatchListParams{
		Context:    ctx,
		Cookie:     ActSSOCookie + "=" + c.AuthToken.ActSSOCookie,
		Title:      string(title),
		Platform:   string(gamer.Platform),
		LookupType: string(gamer.LookupType),
		Gamertag:   gamer.GamerTag,
		GameType:   string(gameType),
		Start:      start,
		End:        end,
	}

	// assign limit if set to a positive integer, otherwise omit
	if limit > 0 {
		param.Limit = &limit
	}
	res, err := c.ServiceClient.Operations.GamerMatchList(&param)
	if err != nil {
		return nil, err
	}

	if res.Payload.Status == "error" {
		// data field is an object type for error cases, these fields are unmapped for this API.
		// Extract unmapped fields - type and message
		//	"data": {
		//    "type": "com.activision.mt.common.stdtools.exceptions.NoStackTraceException",
		//    "message": "Not permitted: not authenticated"
		//  }
		additionalProps := res.Payload.GamerMatchListResponseAdditionalProperties
		return nil, fmt.Errorf("error in gamer match list request. %v - %v", additionalProps["type"], additionalProps["message"])
	}

	return res.Payload, nil
}

// GamerStats returns the gamer stats profile. Authentication required
func (c *Client) GamerStats(ctx context.Context, title GameTitle, gamer *Gamer, gameType GameType) (*models.GamerStatsResponse, error) {
	if err := c.LoggedIn(); err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.GamerStatsParams{
		Context:    ctx,
		Cookie:     ActSSOCookie + "=" + c.AuthToken.ActSSOCookie,
		Title:      string(title),
		Platform:   string(gamer.Platform),
		LookupType: string(gamer.LookupType),
		Gamertag:   gamer.GamerTag,
		GameType:   string(gameType),
	}

	res, err := c.ServiceClient.Operations.GamerStats(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in gamer stats profile request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

// FriendsStats returns the friend stats profile. Authentication required
func (c *Client) FriendsStats(ctx context.Context, title GameTitle, gamer *Gamer, gameType GameType) (*models.FriendStatsResponse, error) {
	if err := c.LoggedIn(); err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.FriendStatsParams{
		Context:    ctx,
		Cookie:     ActSSOCookie + "=" + c.AuthToken.ActSSOCookie,
		Title:      string(title),
		Platform:   string(gamer.Platform),
		LookupType: string(gamer.LookupType),
		Gamertag:   gamer.GamerTag,
		GameType:   string(gameType),
	}

	res, err := c.ServiceClient.Operations.FriendStats(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		// data field is an object type for error cases, these fields are unmapped for this API.
		// Extract unmapped fields - type and message
		//	"data": {
		//    "type": "com.activision.mt.common.stdtools.exceptions.NoStackTraceException",
		//    "message": "Not permitted: not authenticated"
		//  }
		additionalProps := res.Payload.FriendStatsResponseAdditionalProperties
		return nil, fmt.Errorf("error in friend stats request. %v - %v", additionalProps["type"], additionalProps["message"])
	}

	return res.Payload, nil
}

// GamerLoot returns the gamer loots. Authentication required
func (c *Client) GamerLoot(ctx context.Context, title GameTitle, gamer *Gamer) (*models.GamerLootResponse, error) {
	if err := c.LoggedIn(); err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.GamerLootParams{
		Context:    ctx,
		Cookie:     ActSSOCookie + "=" + c.AuthToken.ActSSOCookie,
		Title:      string(title),
		Platform:   string(gamer.Platform),
		LookupType: string(gamer.LookupType),
		Gamertag:   gamer.GamerTag,
	}

	res, err := c.ServiceClient.Operations.GamerLoot(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in gamer loot request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}

// MatchAnalysis returns the full match information. Authentication required
func (c *Client) MatchAnalysis(ctx context.Context, title GameTitle, gamer *Gamer, end int64) (*models.MatchAnalysisResponse, error) {
	if err := c.LoggedIn(); err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.MatchAnalysisParams{
		Context:    ctx,
		Cookie:     ActSSOCookie + "=" + c.AuthToken.ActSSOCookie,
		Title:      string(title),
		Platform:   string(gamer.Platform),
		LookupType: string(gamer.LookupType),
		Gamertag:   gamer.GamerTag,
		End:        end,
	}

	res, err := c.ServiceClient.Operations.MatchAnalysis(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		// data field is an object type for error cases, these fields are unmapped for this API.
		// Extract unmapped fields - type and message
		//	"data": {
		//    "type": "com.activision.mt.common.stdtools.exceptions.NoStackTraceException",
		//    "message": "Not permitted: not authenticated"
		//  }
		additionalProps := res.Payload.MatchAnalysisResponseAdditionalProperties
		return nil, fmt.Errorf("error in match analysis request. %v - %v", additionalProps["type"], additionalProps["message"])
	}

	return res.Payload, nil
}

// CODPoints returns the gamer COD Points. Authentication required
func (c *Client) CODPoints(ctx context.Context, title GameTitle, gamer *Gamer) (*models.CodPointsResponse, error) {

	if err := c.LoggedIn(); err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}

	param := service.CodPointsParams{
		Context:    ctx,
		Cookie:     ActSSOCookie + "=" + c.AuthToken.ActSSOCookie,
		Title:      string(title),
		Platform:   string(gamer.Platform),
		LookupType: string(gamer.LookupType),
		Gamertag:   gamer.GamerTag,
	}

	res, err := c.ServiceClient.Operations.CodPoints(&param)
	if err != nil {
		return nil, err
	}
	if res.Payload.Status == "error" {
		return nil, fmt.Errorf("error in COD Points request. %v - %v", res.Payload.Data.Type, res.Payload.Data.Message)
	}

	return res.Payload, nil
}
