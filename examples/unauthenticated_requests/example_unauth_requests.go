package main

import (
	"context"
	"fmt"
	gocod "github.com/carlocayos/go-cod/v2"
	"log"
)

// ==================  IMPORTANT!!!!  ====================
// You MUST change these values before running this code
// This is configured to run on my account
//
// =======================================================
var (
	myGamerTag    = "MrExcitement#6438"
	myGameTypeMp  = gocod.Multiplayer
	myGameTypeWz  = gocod.Warzone
	myPlatform    = gocod.Battlenet
	myGameTitleMw = gocod.ModernWarfare
)

func main() {
	// =======================================================
	// 1) Setup the default client
	//	  https://my.callofduty.com/api/papi-client
	// =======================================================
	c := gocod.NewClient(nil)

	// =======================================================
	// 2) Leaderboard request
	// =======================================================
	leaderBoardResp, err := c.LeaderBoard(context.Background(), myGameTitleMw, myPlatform, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nLeaderBoard Status = %v\n", leaderBoardResp.Status)
	data := leaderBoardResp.Data

	fmt.Printf("\ntitle = %v\n", data.Title)
	fmt.Printf("platform = %v\n", data.Platform)
	fmt.Printf("leaderboardType = %v\n", data.LeaderboardType)
	fmt.Printf("totalPages = %v\n", data.TotalPages)
	fmt.Printf("resultsRequested = %v\n", data.ResultsRequested)

	fmt.Println("\nPrinting Column Values...")

	for _, v := range data.Columns {
		fmt.Println(v)
	}
	// =======================================================
	// 2) Map List request
	// =======================================================
	mapListResp, err := c.MapList(context.Background(), myGameTitleMw, myPlatform, myGameTypeMp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nMapList Status = %v\n", mapListResp.Status)
	fmt.Printf("Map Rust = %v\n", mapListResp.Data.MpRust)

	// =======================================================
	// 3) Full Match Info request
	// =======================================================
	fullMatchInfoResp, err := c.FullMatchInfo(context.Background(), myGameTitleMw, myPlatform, myGameTypeWz, "4568842382991718691")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFullMatchInfo Status = %v\n", fullMatchInfoResp.Status)
	fmt.Printf("GameType = %v\n", fullMatchInfoResp.Data.AllPlayers[0].GameType)
	fmt.Printf("PlayerCount = %v\n", fullMatchInfoResp.Data.AllPlayers[0].PlayerCount)
	fmt.Printf("Uno ID = %v\n", fullMatchInfoResp.Data.AllPlayers[0].Player.Uno)

	// =======================================================
	// 4) Battlepass Loot request
	// =======================================================
	battlePassLootResp, err := c.BattlePassLoot(context.Background(), myGameTitleMw, myPlatform, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nBattlePassLoot Status = %v\n", battlePassLootResp.Status)
	fmt.Printf("CategoryName = %v\n", battlePassLootResp.Data.CategoryName)
	fmt.Printf("CategoryTitle = %v\n", battlePassLootResp.Data.CategoryTitle)

	for k, v := range battlePassLootResp.Data.Tiers {
		fmt.Printf("Tier#%v:\n", k)
		fmt.Printf("\tName = %v\n", v.Name)
		fmt.Printf("\tRarity = %v\n", v.Rarity)
		fmt.Printf("\tLabel = %v\n", v.Label)
		fmt.Printf("\tTier = %v\n", v.Tier)
		fmt.Printf("\tType = %v\n", v.Type)
		fmt.Printf("\tDescriptionBody = %v\n", v.DescriptionBody)
		fmt.Printf("\tDescriptionTitle = %v\n", v.DescriptionTitle)
		fmt.Printf("\tFree = %v\n", v.Free)
		fmt.Printf("\tExclusiveTitle = %v\n", v.ExclusiveTitle)
		fmt.Printf("\tImage = %v\n", v.Image)
		fmt.Printf("\tImageDetail = %v\n", v.ImageDetail)
	}

	// =======================================================
	// 5) Purchasable request
	// =======================================================
	purchasableResp, err := c.Purchasable(context.Background(), myGameTitleMw, myPlatform)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nPurchasable Status = %v\n", purchasableResp.Status)

	// LootStream is of type map[string]map[string]PurchasableItem
	lootStream := purchasableResp.Data.LootStream

	for k, v := range lootStream {
		fmt.Printf("Loot Stream Name (Key) %v:\n", k)

		for k2, v2 := range v {
			fmt.Printf("\tPurchasable Item Name (Key) = %v:\n", k2)

			fmt.Printf("\t\tName = %v\n", v2.Name)
			fmt.Printf("\t\tLabel = %v\n", v2.Label)
			fmt.Printf("\t\tDescription = %v\n", v2.Description)
			fmt.Printf("\t\tBackgroundImage = %v\n", v2.BackgroundImage)
			fmt.Printf("\t\tLogoImage = %v\n", v2.LogoImage)
			fmt.Printf("\t\tItems = %v\n", v2.Items)
			fmt.Printf("\t\tCosts.CODPoints = %v\n", v2.Costs.CODPoints)
		}
	}

	// =======================================================
	// 6) Loadout request
	//    See example_loadout.go
	// =======================================================
}
