package main

import (
	"context"
	"fmt"
	gocod "github.com/carlocayos/go-cod"
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
	fmt.Printf("Tier No: = %v\n", battlePassLootResp.Data.Tiers.Nr103.Tier)
	fmt.Printf("\tName = %v\n", battlePassLootResp.Data.Tiers.Nr103.Name)
	fmt.Printf("\tRarity = %v\n", battlePassLootResp.Data.Tiers.Nr103.Rarity)

	// =======================================================
	// 5) Purchasable request
	// =======================================================
	purchasableResp, err := c.Purchasable(context.Background(), myGameTitleMw, myPlatform)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nPurchasable Status = %v\n", purchasableResp.Status)
	fmt.Printf("Battle Pass Bundle Description = %v\n", purchasableResp.Data.LootStream.LootSeason6.BattlePassUpgradeBundle6.Description)
	fmt.Printf("\tLabel = %v\n", purchasableResp.Data.LootStream.LootSeason6.BattlePassUpgrade6.Label)
	fmt.Printf("\tCOD Points Cost = %v\n", purchasableResp.Data.LootStream.LootSeason6.BattlePassUpgrade6.Costs.CODPoints)

	// =======================================================
	// 6) Loadout request
	// =======================================================
	loadoutResp, err := c.Loadout(context.Background(), myGameTitleMw, myGameTypeWz)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nLoadout Status = %v\n", loadoutResp.Status)

	// Example Loadout Response Payload
	// {
	//  "status": "success",
	//  "data": {
	//    "weapons": {
	//      "iw8_sn_romeo700_variant_0": {
	//        "dwid": "985",
	//        "name": "iw8_sn_romeo700_variant_0",
	//        "variantId": 0,
	//        "rootName": "iw8_sn_romeo700",
	//        "category": "weapon_dmr",
	//        "imageLoot": "ui_loot_weapon_sn_mike14",
	//        "imageIcon": "icon_weapon_sn_romeo700",
	//        "attachmentsEquipped": "",
	//        "attachments": "",
	//        "perk": "",
	//        "stringKey": "WEAPON/SN_ROMEO700",
	//        "label": "SP-R 208",
	//        "rarity": "base"
	//      },
	//		...
	//	}
	// }

	// This is an example how to retrieve unmapped fields using *additionalProperties
	weapons := loadoutResp.Data.LoadoutResponseDataAdditionalProperties["weapons"].(map[string]interface{})
	romeo := weapons["iw8_sn_romeo700_variant_0"].(map[string]interface{})
	fmt.Printf("iw8_sn_romeo700_variant_0\n")
	fmt.Printf("\tdwid = %v\n", romeo["dwid"])
	fmt.Printf("\tname = %v\n", romeo["name"])
	fmt.Printf("\tvariantId = %v\n", romeo["variantId"])
	fmt.Printf("\trootName = %v\n", romeo["rootName"])
	fmt.Printf("\tcategory = %v\n", romeo["category"])
	fmt.Printf("\timageLoot = %v\n", romeo["imageLoot"])
	fmt.Printf("\timageIcon = %v\n", romeo["imageIcon"])
	fmt.Printf("\tattachmentsEquipped = %v\n", romeo["attachmentsEquipped"])
	fmt.Printf("\tattachments = %v\n", romeo["attachments"])
	fmt.Printf("\tperk = %v\n", romeo["perk"])
	fmt.Printf("\tstringKey = %v\n", romeo["stringKey"])
	fmt.Printf("\tlabel = %v\n", romeo["label"])
	fmt.Printf("\trarity = %v\n", romeo["rarity"])
}
