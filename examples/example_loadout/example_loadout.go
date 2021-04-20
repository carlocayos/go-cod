package main

import (
	"context"
	"fmt"
	gocod "github.com/carlocayos/go-cod/v2"
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
	// Base Weapons
	baseWeapons := loadoutResp.Data.BaseWeapons
	for k, v := range baseWeapons {
		fmt.Printf("BaseWeapon = %v:\n", k)
		fmt.Printf("\tName = %v\n", v.Name)
		fmt.Printf("\tAttachmentCategories = %v\n", v.AttachmentCategories)
		fmt.Printf("\tCategory = %v\n", v.Category)
		fmt.Printf("\tAccuracy = %v\n", v.Accuracy)
		fmt.Printf("\tDamage = %v\n", v.Damage)
		fmt.Printf("\tRange = %v\n", v.Range)
		fmt.Printf("\tFirerate = %v\n", v.Firerate)
		fmt.Printf("\tMobility = %v\n", v.Mobility)
		fmt.Printf("\tControl = %v\n", v.Control)
		fmt.Printf("\tCustomizationDisabled = %v\n", v.CustomizationDisabled)
		fmt.Printf("\tGunsmithDisabled = %v\n", v.GunsmithDisabled)
		fmt.Printf("\tUIDisplayOrder = %v\n", v.UIDisplayOrder)
		fmt.Printf("\tDecalSlots = %v\n", v.DecalSlots)
		fmt.Printf("\tRankUnlocked = %v\n", v.RankUnlocked)
		fmt.Printf("\tWeaponSlot = %v\n", v.WeaponSlot)
		fmt.Printf("\tDescriptionKey = %v\n", v.DescriptionKey)
		fmt.Printf("\tDescriptionLabel = %v\n", v.DescriptionLabel)
		fmt.Printf("\tProgressionImage = %v\n", v.ProgressionImage)
	}

	// Weapons
	weapons := loadoutResp.Data.Weapons
	for k, v := range weapons {
		fmt.Printf("Weapon = %v:\n", k)
		fmt.Printf("\tdwid = %v\n", v.Dwid)
		fmt.Printf("\tname = %v\n", v.Name)
		fmt.Printf("\tvariantId = %v\n", v.VariantID)
		fmt.Printf("\trootName = %v\n", v.RootName)
		fmt.Printf("\tcategory = %v\n", v.Category)
		fmt.Printf("\timageLoot = %v\n", v.ImageLoot)
		fmt.Printf("\timageIcon = %v\n", v.ImageIcon)
		fmt.Printf("\tattachmentsEquipped = %v\n", v.AttachmentsEquipped)
		fmt.Printf("\tattachments = %v\n", v.Attachments)
		fmt.Printf("\tperk = %v\n", v.Perk)
		fmt.Printf("\tstringKey = %v\n", v.StringKey)
		fmt.Printf("\tlabel = %v\n", v.Label)
		fmt.Printf("\trarity = %v\n", v.Rarity)
	}

	// WeaponLevels is of type `map[string]map[string]Level`
	weaponLevels := loadoutResp.Data.WeaponLevels
	for k, v := range weaponLevels {
		fmt.Printf("Weapon Name %v:\n", k)

		for k2, v2 := range v {
			fmt.Printf("\tLevel %v:\n", k2)
			fmt.Printf("\t\tRank = %v\n", v2.Rank)
			fmt.Printf("\t\tUnlockName = %v\n", v2.UnlockName)
			fmt.Printf("\t\tUnlockCategory = %v\n", v2.UnlockCategory)
			fmt.Printf("\t\tUnlockOverrideName = %v\n", v2.UnlockOverrideName)
		}
	}

	// Attachments
	attachments := loadoutResp.Data.Attachments
	for k, v := range attachments {
		fmt.Printf("Attachment %v:\n", k)
		fmt.Printf("\tName = %v\n", v.Name)
		fmt.Printf("\tLabel = %v\n", v.Label)
		fmt.Printf("\tBaseRef = %v\n", v.BaseRef)
		fmt.Printf("\tCategory = %v\n", v.Category)
		fmt.Printf("\tLabel = %v\n", v.Label)
		fmt.Printf("\tLabelKey = %v\n", v.LabelKey)
		fmt.Printf("\tDescriptionLabel = %v\n", v.DescriptionLabel)
		fmt.Printf("\tDescriptionKey = %v\n", v.DescriptionKey)
		fmt.Printf("\tPerk = %v\n", v.Perk)
		fmt.Printf("\tBlocksCategory = %v\n", v.BlocksCategory)
		fmt.Printf("\tAccuracy = %v\n", v.Accuracy)
		fmt.Printf("\tDamage = %v\n", v.Damage)
		fmt.Printf("\tRange = %v\n", v.Range)
		fmt.Printf("\tFirerate = %v\n", v.Firerate)
		fmt.Printf("\tMobility = %v\n", v.Mobility)
		fmt.Printf("\tControl = %v\n", v.Control)
		fmt.Printf("\tMod1 = %v\n", v.Mod1)
		fmt.Printf("\tMod2 = %v\n", v.Mod2)
		fmt.Printf("\tMod2 = %v\n", v.Mod2) // up to mod8
	}

	// Perks
	perks := loadoutResp.Data.Perks
	for k, v := range perks {
		fmt.Printf("Perk Slot No. %v:\n", k)

		for k2, v2 := range v {
			fmt.Printf("\tPerk Name %v:\n", k2)
			fmt.Printf("\t\tName = %v\n", v2.Name)
			fmt.Printf("\t\tLabelKey = %v\n", v2.LabelKey)
			fmt.Printf("\t\tLabel = %v\n", v2.Label)
			fmt.Printf("\t\tDescriptionKey = %v\n", v2.DescriptionKey)
			fmt.Printf("\t\tDescriptionLabel = %v\n", v2.DescriptionLabel)
			fmt.Printf("\t\tExtraInfoKey = %v\n", v2.ExtraInfoKey)
			fmt.Printf("\t\tExtraInfoLabel = %v\n", v2.ExtraInfoLabel)
			fmt.Printf("\t\tPatchNotesKey = %v\n", v2.PatchNotesKey)
			fmt.Printf("\t\tPatchNotes = %v\n", v2.PatchNotes)
			fmt.Printf("\t\tImageMainUI = %v\n", v2.ImageMainUI)
			fmt.Printf("\t\tImageProgression = %v\n", v2.ImageProgression)
			fmt.Printf("\t\tCost = %v\n", v2.Cost)
			fmt.Printf("\t\tIndexID = %v\n", v2.IndexID)
			fmt.Printf("\t\tSlot = %v\n", v2.Slot)
			fmt.Printf("\t\tBrDescriptionKey = %v\n", v2.BrDescriptionKey)
			fmt.Printf("\t\tBrDescriptionLabel = %v\n", v2.BrDescriptionLabel)
			fmt.Printf("\t\tRankUnlocked = %v\n", v2.RankUnlocked)
		}
	}

	// Equipment is of type `map[string][]EquipmentItem`
	equipment := loadoutResp.Data.Equipment
	for k, v := range equipment {
		fmt.Printf("Equipment Type %v:\n", k)

		for k2, v2 := range v {
			fmt.Printf("\tPerk Name %v:\n", k2)

			fmt.Printf("\t\tName = %v\n", v2.Name)
			fmt.Printf("\t\tLabel = %v\n", v2.Label)
			fmt.Printf("\t\tDescription = %v\n", v2.Description)
			fmt.Printf("\t\tDescriptionLabel = %v\n", v2.DescriptionLabel)
			fmt.Printf("\t\tImage = %v\n", v2.Image)
			fmt.Printf("\t\tImageLarge = %v\n", v2.ImageLarge)
			fmt.Printf("\t\tUiOrder = %v\n", v2.UIOrder)
			fmt.Printf("\t\tType = %v\n", v2.Type)
			fmt.Printf("\t\tProgressionImage = %v\n", v2.ProgressionImage)
			fmt.Printf("\t\tRankUnlocked = %v\n", v2.RankUnlocked)
		}
	}

	// Killstreaks
	killstreaks := loadoutResp.Data.Killstreaks
	for k, v := range killstreaks {
		fmt.Printf("Kill Streak Name(Key) %v:\n", k)

		fmt.Printf("\tDwid = %v\n", v.Dwid)
		fmt.Printf("\tName = %v\n", v.Name)
		fmt.Printf("\tDpadIcon = %v\n", v.DpadIcon)
		fmt.Printf("\tFullImage = %v\n", v.FullImage)
		fmt.Printf("\tIcon = %v\n", v.Icon)
		fmt.Printf("\tKills = %v\n", v.Kills)
		fmt.Printf("\tOverheadIcon = %v\n", v.OverheadIcon)
		fmt.Printf("\tProgressionImage = %v\n", v.ProgressionImage)
		fmt.Printf("\tScoreCost = %v\n", v.ScoreCost)
		fmt.Printf("\tSmallImage = %v\n", v.SmallImage)
		fmt.Printf("\tStreakType = %v\n", v.StreakType)
		fmt.Printf("\tLabelKey = %v\n", v.LabelKey)
		fmt.Printf("\tLabel = %v\n", v.Label)
		fmt.Printf("\tSupportCost = %v\n", v.SupportCost)
		fmt.Printf("\tUnearnedIcon = %v\n", v.UnearnedIcon)
		fmt.Printf("\tDescription = %v\n", v.Description)
		fmt.Printf("\tDescriptionLabel = %v\n", v.DescriptionLabel)
	}

	// WeaponCategoriesBySlot is of type map[string][]string
	weaponCategoriesBySlot := loadoutResp.Data.WeaponCategoriesBySlot
	for k, v := range weaponCategoriesBySlot {
		fmt.Printf("Weapon Categories By Slot (Key) %v:\n", k)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}

	// WeaponsByCategory is of type map[string][]string
	weaponsByCategory := loadoutResp.Data.WeaponsByCategory
	for k, v := range weaponsByCategory {
		fmt.Printf("Weapons By Category (Key) %v:\n", k)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}

	// WeaponsByBaseWeapon is of type map[string][]string
	weaponsByBaseWeapon := loadoutResp.Data.WeaponsByBaseWeapon
	for k, v := range weaponsByBaseWeapon {
		fmt.Printf("Weapons By Base Weapon (Key) %v:\n", k)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}

	// BaseWeaponsByCategory is of type map[string][]string
	baseWeaponsByCategory := loadoutResp.Data.BaseWeaponsByCategory
	for k, v := range baseWeaponsByCategory {
		fmt.Printf("Base Weapons By Category (Key) %v:\n", k)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}

	// AttachmentsByBaseRef is of type map[string][]string
	attachmentsByBaseRef := loadoutResp.Data.AttachmentsByBaseRef
	for k, v := range attachmentsByBaseRef {
		fmt.Printf("Attachments By Base Ref (Key) %v:\n", k)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}

	// AttachmentsByCategory is of type map[string][]string
	attachmentsByCategory := loadoutResp.Data.AttachmentsByCategory
	for k, v := range attachmentsByCategory {
		fmt.Printf("Attachments By Category (Key) %v:\n", k)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}
}
