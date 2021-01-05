package main

import (
	"context"
	"fmt"
	gocod "github.com/carlocayos/go-cod"
	"github.com/carlocayos/go-cod/api/client/service"
	"log"
)

func main() {
	// =======================================================
	// 1) Setup the default client
	//	  https://my.callofduty.com/api/papi-client
	// =======================================================
	c := gocod.NewClient(nil)
	serviceOperations := c.ServiceClient.Operations

	// =======================================================
	// 2) Leaderboard request
	// =======================================================
	leaderBoardParams := service.LeaderBoardParams{
		Page:     1,
		Platform: "battle",
		Title:    "mw",
		Context:  context.Background(),
	}

	leaderBoardResponse, err := serviceOperations.LeaderBoard(&leaderBoardParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nStatus = %v\n", leaderBoardResponse.Payload.Status)
	data := leaderBoardResponse.Payload.Data

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
	// 3) Battle Pass Loot request
	// =======================================================
	battlePassLootParams := service.BattlePassLootParams{
		Title:    "mw",
		Platform: "uno",
		Season:   1,
		Context:  context.Background(),
	}

	bpLootResp, err := serviceOperations.BattlePassLoot(&battlePassLootParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nStatus = %v\n", bpLootResp.Payload.Status)
	bpLootdata := bpLootResp.Payload.Data

	fmt.Printf("\ncategoryName = %v\n", bpLootdata.CategoryName)
	fmt.Printf("categoryTitle = %v\n", bpLootdata.CategoryTitle)

	fmt.Println("\nPrinting Tier Values...")
	fmt.Printf("Tier 0: %v %v\n", bpLootdata.Tiers.Nr0.Name, bpLootdata.Tiers.Nr0.Rarity)
	fmt.Printf("Tier 1: %v %v\n", bpLootdata.Tiers.Nr1.Name, bpLootdata.Tiers.Nr1.Rarity)
	fmt.Printf("Tier 2: %v %v\n", bpLootdata.Tiers.Nr2.Name, bpLootdata.Tiers.Nr2.Rarity)
}
