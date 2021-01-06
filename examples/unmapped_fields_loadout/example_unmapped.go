package main

import (
	"context"
	"fmt"
	gocod "github.com/carlocayos/go-cod"
	"log"
)

func main() {
	// =======================================================
	// 1) Setup the default client
	//	  https://my.callofduty.com/api/papi-client
	// =======================================================
	c := gocod.NewClient(nil)

	// =======================================================
	// 2) Loadout request
	// =======================================================
	loadoutResp, err := c.Loadout(context.Background(), gocod.ModernWarfare, gocod.Warzone)
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
	// TODO: Review swagger spec and change the object definitions. Weapons must be of type map[string]SomeStruct
	// 	See https://github.com/carlocayos/go-cod/issues/3 for a similar issue
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
