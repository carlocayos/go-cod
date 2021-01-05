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
	// 2) Loadout request
	// =======================================================
	loadoutParams := service.LoadoutParams{
		GameType: "wz",
		Title:    "mw",
		Context:  context.Background(),
	}

	loadoutResponse, err := serviceOperations.Loadout(&loadoutParams)
	if err != nil {
		log.Fatal(err)
	}

	//	{
	//	  "status": "success",
	//	  "data": {
	//	    "weapons": {
	//	  	   "iw8_knife_variant_119": {
	//			 	"dwid": "892",
	//				"name": "iw8_knife_variant_119",
	//				"variantId": 119,
	//				"rootName": "iw8_knife",
	//				"category": "weapon_melee2",
	//				"imageLoot": "ui_loot_weapon_me_hatchet",
	//				"imageIcon": "icon_weapon_me_soscar_knife",
	//				"attachmentsEquipped": "",
	//				"attachments": "",
	//				"perk": "",
	//				"stringKey": "WEAPON/KNIFE_V72",
	//				"label": "<WEAPON/KNIFE_V120>",
	//				"rarity": "common"
	//		},
	//		...
	//  }

	fmt.Printf("Status = %v\n", loadoutResponse.Payload.Status)

	// get the weapons of type map[string]string
	weapons := loadoutResponse.Payload.Data.LoadoutResponseDataAdditionalProperties["weapons"].(map[string]interface{})

	// iterate through the weapons
	for k, v := range weapons {
		fmt.Printf("Key = %v\n", k)

		// print all the fields
		for k2, v2 := range v.(map[string]interface{}) {
			fmt.Printf("\t%v = %v\n", k2, v2)
		}
		fmt.Printf("\n\n")
	}
}
