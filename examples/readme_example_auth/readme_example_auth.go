package main

import (
	"context"
	"fmt"
	go_cod "github.com/carlocayos/go-cod"
	"github.com/segmentio/ksuid"
)

func main() {
	// =======================================================
	// 1) First, is creating a client and sending a register device request
	// =======================================================
	c := go_cod.NewClient(nil)

	// send a register device request with a unique device id
	// ksuid is used here to generate a unique id, but any uid generator would be fine
	deviceId := ksuid.New().String()
	registerDeviceRes, _ := c.RegisterDevice(context.Background(), deviceId)

	// =======================================================
	// 2) Next is the Login request. Replace it with your email and password
	//      and pass the returned AuthHeader and Device ID
	// =======================================================
	email := "<< CHANGE ME >>"
	password := "<< CHANGE ME >>"
	loginRes, _ := c.Login(context.Background(), deviceId, email,
		password, *registerDeviceRes.Data.AuthHeader)
	fmt.Println(loginRes.ACTSSOCOOKIE)

	// =======================================================
	// 3) Final step is to send a Get Gamer Match List.
	//    You need to call Login() before using this authenticated request.
	//	  The token is stored in the client and will be implicitly sent along
	//    each Authenticated request.
	// =======================================================
	// create Gamer struct - MUST CHANGE this to your own account
	gamer := &go_cod.Gamer{
		Platform:   go_cod.Battlenet,
		LookupType: go_cod.BattlenetLookup,
		GamerTag:   "MrExcitement#6438",
	}

	gamerMatchListResp, _ := c.GamerMatchList(context.Background(),
		go_cod.ModernWarfare, gamer, go_cod.Multiplayer, 0, 0, 3)

	fmt.Println(gamerMatchListResp.Status)
	for _, v := range gamerMatchListResp.Data {
		fmt.Printf("\tMap = %v\n", v.Map)
		fmt.Printf("\tMatchID = %v\n", v.MatchID)
		fmt.Printf("\tPlatform = %v\n", v.Platform)
		fmt.Printf("\tTimestamp = %v\n", v.Timestamp)
		fmt.Printf("\tTitle = %v\n", v.Title)
		fmt.Printf("\tType = %v\n\n", v.Type)
	}
}
