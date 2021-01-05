package main

import (
	"context"
	"fmt"
	gocod "github.com/carlocayos/go-cod"
	"github.com/carlocayos/go-cod/api/client/authentication"
	"github.com/carlocayos/go-cod/api/client/service"
	"github.com/carlocayos/go-cod/api/models"
	"github.com/segmentio/ksuid"
	"log"
	"net/http"
)

// NOTE: You must change these values before running this code
var (
	email      = "<< CHANGE ME >>"
	password   = "<< CHANGE ME >>"
	gameType   = "mp"
	gamerTag   = "MrExcitement#6438"
	lookupType = "gamer"
	platform   = "battle"
	title      = "cw"
)

func main() {
	// =======================================================
	// 1) First step is to send a Register Device POST request to https://profile.callofduty.com/cod/mapp/registerDevice
	// with a unique device id in the payload. It would return a security token value in the authHeader field which
	// will be used to complete teh login process
	// =======================================================
	// create the COD API client
	c := gocod.NewClient(nil)

	// send a register device request
	deviceId := ksuid.New().String()

	param := authentication.RegisterDeviceParams{
		RegisterDeviceRequest: &models.RegisterDeviceRequest{DeviceID: &deviceId},
		Context:               context.Background(),
	}
	fmt.Printf("\nDevice ID = %v\n", deviceId)

	resp, err := c.AuthenticationClient.Operations.RegisterDevice(&param)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nStatus = %v\n", resp.Payload.Status)
	fmt.Printf("AuthHeader = %v\n\n", *resp.Payload.Data.AuthHeader)

	// =======================================================
	// 2) Next is sending a Login request. Change to your email and password
	// =======================================================
	loginParams := authentication.LoginParams{
		Authorization: "Bearer " + *resp.Payload.Data.AuthHeader,
		XCodDeviceID:  deviceId,
		LoginRequest: &models.LoginRequest{
			Email:    &email,
			Password: &password,
		},
		Context: context.Background(),
	}

	ok, err := c.AuthenticationClient.Operations.Login(&loginParams)
	if err != nil {
		log.Fatal(err)
	}
	if ok.Payload.Status == http.StatusUnauthorized {
		log.Fatal(ok.Payload.LoginResponseAdditionalProperties["message"])
	}

	fmt.Println(ok)

	// =======================================================
	// 3) Example request which requires authentication.
	//    Get Match Details
	// =======================================================
	limit := int32(10)

	// send a match detail request
	matchDetailsParams := service.MatchDetailsParams{
		Start:      1602815242000,
		End:        1608407321000,
		GameType:   gameType,
		Gamertag:   gamerTag,
		LookupType: lookupType,
		Platform:   platform,
		Title:      title,
		Context:    context.Background(),
		Limit:      &limit,
		Cookie:     "ACT_SSO_COOKIE=" + ok.Payload.ACTSSOCOOKIE,
	}

	matchDetailsResponse, err := c.ServiceClient.Operations.MatchDetails(&matchDetailsParams)
	if err != nil {
		log.Fatal(err)
	}
	matchDetailsStatus := matchDetailsResponse.Payload.Status
	fmt.Printf("\nStatus = %v\n\n", matchDetailsStatus)
	data := matchDetailsResponse.Payload.Data

	if matchDetailsStatus == "error" {
		log.Fatalf("error in match details request. %v : %v", data.Type, data.Message)
	}
	for k, v := range data.Matches {
		fmt.Printf("%v = %v\n", k, v)
	}

	// =======================================================
	// 4) Another example request which requires authentication.
	//    Get Gamer Match List
	// =======================================================
	// send a gamer match list request
	matchListParams := service.MatchListParams{
		Start:      0,
		End:        0,
		Cookie:     "ACT_SSO_COOKIE=" + ok.Payload.ACTSSOCOOKIE,
		GameType:   gameType,
		Gamertag:   gamerTag,
		Limit:      &limit,
		LookupType: lookupType,
		Platform:   platform,
		Title:      title,
		Context:    context.Background(),
	}

	matchListResponse, err := c.ServiceClient.Operations.MatchList(&matchListParams)
	if err != nil {
		log.Fatal(err)
	}
	matchListStatus := matchListResponse.Payload.Status
	fmt.Printf("\nStatus = %v\n\n", matchListStatus)

	if matchListStatus == "error" {
		// data field is an object type for error cases
		//
		//	"data": {
		//    "type": "com.activision.mt.common.stdtools.exceptions.NoStackTraceException",
		//    "message": "Not permitted: not authenticated"
		//  }
		additionalProps := matchListResponse.Payload.GamerAllMatchListResponseAdditionalProperties
		log.Fatalf("error in match details request. %v : %v", additionalProps["type"], additionalProps["message"])
	}

	// Match list response contains a "data" field array type
	//
	//     "data": [
	//        {
	//            "platform": "battle",
	//            "title": "mw",
	//            "timestamp": 1594381028000,
	//            "type": "war",
	//            "matchId": "15402031512165119432",
	//            "map": "mp_runner"
	//        },
	//       ...
	d := matchListResponse.Payload.Data

	for _, v := range d {
		fmt.Printf("Map = %v\n", v.Map)
		fmt.Printf("MatchID = %v\n", v.MatchID)
		fmt.Printf("Platform = %v\n", v.Platform)
		fmt.Printf("Timestamp = %v\n", v.Timestamp)
		fmt.Printf("Title = %v\n", v.Title)
		fmt.Printf("Type = %v\n", v.Type)
	}
}
