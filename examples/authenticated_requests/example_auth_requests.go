package main

import (
	"context"
	"fmt"
	gocod "github.com/carlocayos/go-cod"
	"github.com/segmentio/ksuid"
	"log"
)

// ==================  IMPORTANT!!!!  ====================
// You MUST change these values before running this code
// This is configured to run on my account
//
// =======================================================
var (
	myEmail      = "<< CHANGE ME >>"
	myPassword   = "<< CHANGE ME >>"
	myGamerTag   = "MrExcitement#6438"
	myGameType   = gocod.Multiplayer
	myLookupType = gocod.BattlenetLookup
	myPlatform   = gocod.Battlenet
	myTitleCw    = gocod.ColdWar
	myTitleMw    = gocod.ModernWarfare
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
	registerDeviceRes, err := c.RegisterDevice(context.Background(), deviceId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nStatus = %v\n", registerDeviceRes.Status)
	fmt.Printf("AuthHeader = %v\n\n", *registerDeviceRes.Data.AuthHeader)

	// =======================================================
	// 2) Next is sending a Login request. Change to your email and password
	// =======================================================
	loginRes, err := c.Login(context.Background(), deviceId, myEmail, myPassword, *registerDeviceRes.Data.AuthHeader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(loginRes)

	// =======================================================
	// 3) Get Gamer Match Details
	//    Example request which requires authentication.
	// =======================================================
	// create Gamer struct
	// NOTE: Change this to your own account
	gamer := &gocod.Gamer{
		Platform:   myPlatform,
		LookupType: myLookupType,
		GamerTag:   myGamerTag,
	}

	gamerMatchDetailsResp, err := c.GamerMatchDetails(context.Background(), myTitleCw, gamer, myGameType,
		int64(1602815242000), int64(1608407321000), 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nGamerMatchDetails Status = %v\n", gamerMatchDetailsResp.Status)
	for _, v := range gamerMatchDetailsResp.Data.Matches {
		fmt.Printf("MatchId = %v\n", v.MatchID)
		fmt.Printf("Map = %v\n", v.Map)
		fmt.Printf("Mode = %v\n\n", v.Mode)
	}

	// =======================================================
	// 4) Gamer Match List
	// =======================================================
	gamerMatchListResp, err := c.GamerMatchList(context.Background(), myTitleMw, gamer, myGameType, 1592845714000, 1594381028000, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nGamerMatchList Status = %v\n", gamerMatchListResp.Status)

	for _, v := range gamerMatchListResp.Data {
		fmt.Printf("\tMap = %v\n", v.Map)
		fmt.Printf("\tMatchID = %v\n", v.MatchID)
		fmt.Printf("\tPlatform = %v\n", v.Platform)
		fmt.Printf("\tTimestamp = %v\n", v.Timestamp)
		fmt.Printf("\tTitle = %v\n", v.Title)
		fmt.Printf("\tType = %v\n\n", v.Type)
	}

	// =======================================================
	// 5) Gamer Stats Profile
	// =======================================================
	gamerStatsResp, err := c.GamerStats(context.Background(), myTitleMw, gamer, myGameType)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nGamerStats Status = %v\n", gamerStatsResp.Status)
	fmt.Printf("Username = %v\n", gamerStatsResp.Data.Username)
	fmt.Printf("Platform = %v\n", gamerStatsResp.Data.Platform)
	fmt.Printf("Title = %v\n", gamerStatsResp.Data.Title)
	fmt.Printf("Message = %v\n", gamerStatsResp.Data.Message)
	fmt.Printf("Level = %v\n", gamerStatsResp.Data.Level)
	fmt.Printf("Prestige = %v\n", gamerStatsResp.Data.Prestige)
	fmt.Printf("PrestigeID = %v\n", gamerStatsResp.Data.PrestigeID)
	fmt.Printf("ParagonID = %v\n", gamerStatsResp.Data.ParagonID)
	fmt.Printf("ParagonRank = %v\n", gamerStatsResp.Data.ParagonRank)
	fmt.Printf("LevelXpGained = %v\n", gamerStatsResp.Data.LevelXpGained)
	fmt.Printf("LevelXpRemainder = %v\n", gamerStatsResp.Data.LevelXpRemainder)

	// =======================================================
	// 6) Friend Stats
	// =======================================================
	friendStatsResp, err := c.FriendsStats(context.Background(), myTitleMw, gamer, myGameType)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFriendsStats Status = %v\n", friendStatsResp.Status)

	for _, v := range friendStatsResp.Data {
		fmt.Printf("Friend Username: = %v\n", v.Username)
		fmt.Printf("\t Level : = %v\n", v.Level)
		fmt.Printf("\t TotalXp : = %v\n", v.TotalXp)
		fmt.Printf("\t Platform : = %v\n", v.Platform)
		fmt.Printf("\t Title : = %v\n", v.Title)

		fmt.Printf("\t ALL Properties:\n")
		props := v.Lifetime.All.Properties
		fmt.Printf("\t\t KdRatio : = %v\n", props.KdRatio)
		fmt.Printf("\t\t RecordKillsInAMatch : = %v\n", props.RecordKillsInAMatch)
		fmt.Printf("\t\t RecordLongestWinStreak : = %v\n", props.RecordLongestWinStreak)
		fmt.Printf("\t\t WinLossRatio : = %v\n", props.WinLossRatio)
		fmt.Printf("\t\t Deaths : = %v\n", props.Deaths)
		fmt.Printf("\t\t BestAssists : = %v\n", props.BestAssists)
		fmt.Printf("\t\t RecordDeathsInAMatch : = %v\n", props.RecordDeathsInAMatch)
	}

	// =======================================================
	// 6) Gamer Loot
	// =======================================================
	gamerLootResp, err := c.GamerLoot(context.Background(), myTitleMw, gamer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nGamerLoot Status = %v\n", gamerLootResp.Status)
	// TODO: Review swagger spec and change the object definitions. Streams must be type map[string]SomeStruct
	// 	See https://github.com/carlocayos/go-cod/issues/3
	fmt.Printf("Loot Season 0 Name = %v\n", gamerLootResp.Data.Streams.LootSeason0.Name)
	fmt.Printf("\tRarity = %v\n", gamerLootResp.Data.Streams.LootSeason0.Rarity)

	// =======================================================
	// 7) Match Analysis
	// =======================================================
	matchAnalysisResp, err := c.MatchAnalysis(context.Background(), myTitleMw, gamer, 1594381028000)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nMatchAnalysis Status = %v\n", matchAnalysisResp.Status)

	for _, v := range matchAnalysisResp.Data {
		fmt.Printf("Match ID = %v\n", v.PayLoad.MatchID)
		fmt.Printf("\tTeam1Score = %v\n", v.PayLoad.Team1Score)
		fmt.Printf("\tMode = %v\n", v.PayLoad.Mode)
		fmt.Printf("\tKdRatio = %v\n", v.PayLoad.KdRatio)
		fmt.Printf("\tWinningTeam = %v\n", v.PayLoad.WinningTeam)
		fmt.Printf("\tPlayer = %v\n", v.PayLoad.Player)
	}

	// =======================================================
	// 7) COD Points
	// =======================================================
	codPointsResp, err := c.CODPoints(context.Background(), myTitleMw, gamer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nCODPoints Status = %v\n", codPointsResp.Status)
	fmt.Printf("COD Points = %v\n", codPointsResp.Data.CodPoints)
}
