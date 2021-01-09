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

	// get match summary for each game mode
	for k, v := range gamerMatchDetailsResp.Data.Summary {
		fmt.Printf("Game Mode = %v:\n", k)
		fmt.Printf("\tKills = %v\n", v.Kills)
		fmt.Printf("\tAccuracy = %v\n", v.Accuracy)
		fmt.Printf("\tHeadshots = %v\n", v.Headshots)
		fmt.Printf("\tHeadshotPercentage = %v\n", v.HeadshotPercentage)
		fmt.Printf("\tAssists = %v\n", v.Assists)
		fmt.Printf("\tScore = %v\n", v.Score)
		fmt.Printf("\tKdRatio = %v\n", v.KdRatio)
		fmt.Printf("\tDamageDealt = %v\n", v.DamageDealt)
		fmt.Printf("\tHighestMultikill = %v\n", v.HighestMultikill)
		fmt.Printf("\tAssists = %v\n", v.Assists)
	}

	// list the match details
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
	err = sendGamerStats(c, gamer)
	if err != nil {
		log.Fatal(err)
	}

	// =======================================================
	// 6) Friend Stats
	// =======================================================
	err = sendFriendStats(c, gamer)
	if err != nil {
		log.Fatal(err)
	}

	// =======================================================
	// 6) Gamer Loot
	// =======================================================
	gamerLootResp, err := c.GamerLoot(context.Background(), myTitleMw, gamer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nGamerLoot Status = %v\n", gamerLootResp.Status)

	for k, v := range gamerLootResp.Data.Streams {
		fmt.Printf("Stream = %v:\n", k)
		fmt.Printf("\tName = %v:\n", v.Name)
		fmt.Printf("\tTier = %v:\n", v.Tier)
		fmt.Printf("\tRarity = %v:\n", v.Rarity)
		fmt.Printf("\tCategoryNameLabel = %v:\n", v.CategoryNameLabel)
		fmt.Printf("\tCategoryTitleLabel = %v:\n", v.CategoryTitleLabel)
		fmt.Printf("\tItemsObtained = %v:\n", v.ItemsObtained)
		fmt.Printf("\tLootType = %v:\n", v.LootType)
		fmt.Printf("\tPremium = %v:\n", v.Premium)
		fmt.Printf("\tPremiumTokenOwnedNotRedeemed = %v:\n", v.PremiumTokenOwnedNotRedeemed)
	}

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

// Separate function for sending FriendStats request
func sendFriendStats(c *gocod.Client, gamer *gocod.Gamer) error {
	friendStatsResp, err := c.FriendsStats(context.Background(), myTitleMw, gamer, myGameType)
	if err != nil {
		return err
	}
	fmt.Printf("\nFriendsStats Status = %v\n", friendStatsResp.Status)

	// Friend Stats Data is an array type
	//	"data": [
	//		{
	//			"title": "mw",
	//			"platform": "battle",
	//			"username": "myfriend#4172170",
	//			...
	//		}
	//	...
	//	]

	for _, v := range friendStatsResp.Data {
		fmt.Printf("Friend Username: = %v\n", v.Username)

		fmt.Printf("\tTitle = %v\n", v.Title)
		fmt.Printf("\tPlatform = %v\n", v.Platform)
		fmt.Printf("\tUsername = %v\n", v.Username)
		fmt.Printf("\tType = %v\n", v.Type)
		fmt.Printf("\tLevel = %v\n", v.Level)
		fmt.Printf("\tMaxLevel = %v\n", v.MaxLevel)
		fmt.Printf("\tLevelXpRemainder = %v\n", v.LevelXpRemainder)
		fmt.Printf("\tLevelXpGained = %v\n", v.LevelXpGained)
		fmt.Printf("\tPrestige = %v\n", v.Prestige)
		fmt.Printf("\tPrestigeID = %v\n", v.PrestigeID)
		fmt.Printf("\tMaxPrestige = %v\n", v.MaxPrestige)
		fmt.Printf("\tTotalXp = %v\n", v.TotalXp)
		fmt.Printf("\tParagonRank = %v\n", v.ParagonRank)
		fmt.Printf("\tParagonID = %v\n", v.ParagonID)
		fmt.Printf("\tS = %v\n", v.S)
		fmt.Printf("\tP = %v\n", v.P)

		// ALL - gamer stats summary
		allProperties := v.Lifetime.All.Properties
		fmt.Printf("ALL Gamer Stats Summary\n")

		fmt.Printf("\t\tRecordLongestWinStreak = %v\n", allProperties.RecordLongestWinStreak)
		fmt.Printf("\t\tRecordXpInAMatch = %v\n", allProperties.RecordXpInAMatch)
		fmt.Printf("\t\tAccuracy = %v\n", allProperties.Accuracy)
		fmt.Printf("\t\tLosses = %v\n", allProperties.Losses)
		fmt.Printf("\t\tTotalGamesPlayed = %v\n", allProperties.TotalGamesPlayed)
		fmt.Printf("\t\tScore = %v\n", allProperties.Score)
		fmt.Printf("\t\tWinLossRatio = %v\n", allProperties.WinLossRatio)
		fmt.Printf("\t\tTotalShots = %v\n", allProperties.TotalShots)
		fmt.Printf("\t\tBestScoreXp = %v\n", allProperties.BestScoreXp)
		fmt.Printf("\t\tGamesPlayed = %v\n", allProperties.GamesPlayed)
		fmt.Printf("\t\tBestSquardKills = %v\n", allProperties.BestSquardKills)
		fmt.Printf("\t\tBestSguardWave = %v\n", allProperties.BestSguardWave)
		fmt.Printf("\t\tBestConfirmed = %v\n", allProperties.BestConfirmed)
		fmt.Printf("\t\tDeaths = %v\n", allProperties.Deaths)
		fmt.Printf("\t\tWins = %v\n", allProperties.Wins)
		fmt.Printf("\t\tBestSquardCrates = %v\n", allProperties.BestSquardCrates)
		fmt.Printf("\t\tKdRatio = %v\n", allProperties.KdRatio)
		fmt.Printf("\t\tBestAssists = %v\n", allProperties.BestAssists)
		fmt.Printf("\t\tBestFieldgoals = %v\n", allProperties.BestFieldgoals)
		fmt.Printf("\t\tBestScore = %v\n", allProperties.BestScore)
		fmt.Printf("\t\tRecordDeathsInAMatch = %v\n", allProperties.RecordDeathsInAMatch)
		fmt.Printf("\t\tScorePerGame = %v\n", allProperties.ScorePerGame)
		fmt.Printf("\t\tBestSPM = %v\n", allProperties.BestSPM)
		fmt.Printf("\t\tBestKillChains = %v\n", allProperties.BestKillChains)
		fmt.Printf("\t\tRecordKillsInAMatch = %v\n", allProperties.RecordKillsInAMatch)
		fmt.Printf("\t\tSuicides = %v\n", allProperties.Suicides)
		fmt.Printf("\t\tWlRatio = %v\n", allProperties.WlRatio)
		fmt.Printf("\t\tCurrentWinStreak = %v\n", allProperties.CurrentWinStreak)
		fmt.Printf("\t\tBestMatchBonusXp = %v\n", allProperties.BestMatchBonusXp)
		fmt.Printf("\t\tBestMatchXp = %v\n", allProperties.BestMatchXp)
		fmt.Printf("\t\tBestSguardWeaponLevel = %v\n", allProperties.BestSguardWeaponLevel)
		fmt.Printf("\t\tBestKD = %v\n", allProperties.BestKD)
		fmt.Printf("\t\tKills = %v\n", allProperties.Kills)
		fmt.Printf("\t\tBestKillsAsInfected = %v\n", allProperties.BestKillsAsInfected)
		fmt.Printf("\t\tBestReturns = %v\n", allProperties.BestReturns)
		fmt.Printf("\t\tBestStabs = %v\n", allProperties.BestStabs)
		fmt.Printf("\t\tBestKillsAsSurvivor = %v\n", allProperties.BestKillsAsSurvivor)
		fmt.Printf("\t\tTimePlayedTotal = %v\n", allProperties.TimePlayedTotal)
		fmt.Printf("\t\tBestDestructions = %v\n", allProperties.BestDestructions)
		fmt.Printf("\t\tHeadshots = %v\n", allProperties.Headshots)
		fmt.Printf("\t\tBestRescues = %v\n", allProperties.BestRescues)
		fmt.Printf("\t\tAssists = %v\n", allProperties.Assists)
		fmt.Printf("\t\tTies = %v\n", allProperties.Ties)
		fmt.Printf("\t\tRecordKillStreak = %v\n", allProperties.RecordKillStreak)
		fmt.Printf("\t\tBestPlants = %v\n", allProperties.BestPlants)
		fmt.Printf("\t\tMisses = %v\n", allProperties.Misses)
		fmt.Printf("\t\tBestDamage = %v\n", allProperties.BestDamage)
		fmt.Printf("\t\tBestSetbacks = %v\n", allProperties.BestSetbacks)
		fmt.Printf("\t\tBestTouchdowns = %v\n", allProperties.BestTouchdowns)
		fmt.Printf("\t\tScorePerMinute = %v\n", allProperties.ScorePerMinute)
		fmt.Printf("\t\tBestDeaths = %v\n", allProperties.BestDeaths)
		fmt.Printf("\t\tBestMedalXp = %v\n", allProperties.BestMedalXp)
		fmt.Printf("\t\tBestDefends = %v\n", allProperties.BestDefends)
		fmt.Printf("\t\tBestSquardRevives = %v\n", allProperties.BestSquardRevives)
		fmt.Printf("\t\tBestKills = %v\n", allProperties.BestKills)
		fmt.Printf("\t\tBestDefuses = %v\n", allProperties.BestDefuses)
		fmt.Printf("\t\tBestCaptures = %v\n", allProperties.BestCaptures)
		fmt.Printf("\t\tHits = %v\n", allProperties.Hits)
		fmt.Printf("\t\tBestKillStreak = %v\n", allProperties.BestKillStreak)
		fmt.Printf("\t\tBestDenied = %v\n", allProperties.BestDenied)

		// Data.Lifetime.Mode is of type map[string]GamerStatsResponseDataLifetimeModeAnon
		mode := v.Lifetime.Mode
		for k, v := range mode {
			fmt.Printf("Mode (Key) = %v:\n", k)

			fmt.Printf("\t\tKills = %v\n", v.Kills)
			fmt.Printf("\t\tScore = %v\n", v.Score)
			fmt.Printf("\t\tTimePlayed = %v\n", v.TimePlayed)
			fmt.Printf("\t\tKdRatio = %v\n", v.KdRatio)
			fmt.Printf("\t\tSetBacks = %v\n", v.SetBacks)
			fmt.Printf("\t\tScorePerMinute = %v\n", v.ScorePerMinute)
			fmt.Printf("\t\tStabs = %v\n", v.Stabs)
			fmt.Printf("\t\tDeaths = %v\n", v.Deaths)
		}

		// Data.Lifetime.ItemData is of type map[string]map[string]GamerStatsResponseDataLifetimeItemDataAnon
		itemData := v.Lifetime.ItemData
		for k, v := range itemData {
			fmt.Printf("Item Category (Key) = %v:\n", k)

			for k2, v2 := range v {
				fmt.Printf("\t\tItem (Key) %v:\n", k2)

				fmt.Printf("\t\t\tHits = %v\n", v2.Hits)
				fmt.Printf("\t\t\tKills = %v\n", v2.Kills)
				fmt.Printf("\t\t\tKdRatio = %v\n", v2.KdRatio)
				fmt.Printf("\t\t\tHeadshots = %v\n", v2.Headshots)
				fmt.Printf("\t\t\tAccuracy = %v\n", v2.Accuracy)
				fmt.Printf("\t\t\tShots = %v\n", v2.Shots)
				fmt.Printf("\t\t\tDeaths = %v\n", v2.Deaths)
			}
		}

		// Data.Lifetime.ScorestreakData. is of type map[string]GamerStatsResponseDataLifetimeScorestreakDataLethalScorestreakDataAnon
		lethalScorestreakData := v.Lifetime.ScorestreakData.LethalScorestreakData
		for k, v := range lethalScorestreakData {
			fmt.Printf("Lethal Core Streak (Key) = %v:\n", k)

			fmt.Printf("\t\tExtraStat1 = %v\n", v.ExtraStat1)
			fmt.Printf("\t\tUses = %v\n", v.Uses)
			fmt.Printf("\t\tAwardedCount = %v\n", v.AwardedCount)
		}

		// Data.Lifetime.ScorestreakData.SupportScorestreakData. is of type map[string]GamerStatsResponseDataLifetimeScorestreakDataSupportScorestreakDataAnon
		supportScorestreakData := v.Lifetime.ScorestreakData.SupportScorestreakData
		for k, v := range supportScorestreakData {
			fmt.Printf("Support Core Streak (Key) = %v:\n", k)

			fmt.Printf("\t\tExtraStat1 = %v\n", v.ExtraStat1)
			fmt.Printf("\t\tUses = %v\n", v.Uses)
			fmt.Printf("\t\tAwardedCount = %v\n", v.AwardedCount)
		}

		// Accolade Data
		accoladeDataProperties := v.Lifetime.AccoladeData.Properties
		fmt.Printf("Accolade Data\n")

		fmt.Printf("\t\tClassChanges = %v\n", accoladeDataProperties.ClassChanges)
		fmt.Printf("\t\tHighestAvgAltitude = %v\n", accoladeDataProperties.HighestAvgAltitude)
		fmt.Printf("\t\tKillsFromBehind = %v\n", accoladeDataProperties.KillsFromBehind)
		fmt.Printf("\t\tLmgDeaths = %v\n", accoladeDataProperties.LmgDeaths)
		fmt.Printf("\t\tRiotShieldDamageAbsorbed = %v\n", accoladeDataProperties.RiotShieldDamageAbsorbed)
		fmt.Printf("\t\tFlashbangHits = %v\n", accoladeDataProperties.FlashbangHits)
		fmt.Printf("\t\tMeleeKills = %v\n", accoladeDataProperties.MeleeKills)
		fmt.Printf("\t\tTagsLargestBank = %v\n", accoladeDataProperties.TagsLargestBank)
		fmt.Printf("\t\tShotgunKills = %v\n", accoladeDataProperties.ShotgunKills)
		fmt.Printf("\t\tSniperDeaths = %v\n", accoladeDataProperties.SniperDeaths)
		fmt.Printf("\t\tTimeProne = %v\n", accoladeDataProperties.TimeProne)
		fmt.Printf("\t\tKillstreakWhitePhosphorousKillsAssists = %v\n", accoladeDataProperties.KillstreakWhitePhosphorousKillsAssists)
		fmt.Printf("\t\tShortestLife = %v\n", accoladeDataProperties.ShortestLife)
		fmt.Printf("\t\tDeathsFromBehind = %v\n", accoladeDataProperties.DeathsFromBehind)
		fmt.Printf("\t\tHigherRankedKills = %v\n", accoladeDataProperties.HigherRankedKills)
		fmt.Printf("\t\tMostAssists = %v\n", accoladeDataProperties.MostAssists)
		fmt.Printf("\t\tLeastKills = %v\n", accoladeDataProperties.LeastKills)
		fmt.Printf("\t\tTagsDenied = %v\n", accoladeDataProperties.TagsDenied)
		fmt.Printf("\t\tKillstreakWheelsonKills = %v\n", accoladeDataProperties.KillstreakWheelsonKills)
		fmt.Printf("\t\tSniperHeadshots = %v\n", accoladeDataProperties.SniperHeadshots)
		fmt.Printf("\t\tKillstreakJuggernautKills = %v\n", accoladeDataProperties.KillstreakJuggernautKills)
		fmt.Printf("\t\tSmokesUsed = %v\n", accoladeDataProperties.SmokesUsed)
		fmt.Printf("\t\tAvengerKills = %v\n", accoladeDataProperties.AvengerKills)
		fmt.Printf("\t\tDecoyHits = %v\n", accoladeDataProperties.DecoyHits)
		fmt.Printf("\t\tKillstreakCarePackageUsed = %v\n", accoladeDataProperties.KillstreakCarePackageUsed)
		fmt.Printf("\t\tMolotovKills = %v\n", accoladeDataProperties.MolotovKills)
		fmt.Printf("\t\tGasHits = %v\n", accoladeDataProperties.GasHits)
		fmt.Printf("\t\tComebackKills = %v\n", accoladeDataProperties.ComebackKills)
		fmt.Printf("\t\tLmgHeadshots = %v\n", accoladeDataProperties.LmgHeadshots)
		fmt.Printf("\t\tSmgDeaths = %v\n", accoladeDataProperties.SmgDeaths)
		fmt.Printf("\t\tCarrierKills = %v\n", accoladeDataProperties.CarrierKills)
		fmt.Printf("\t\tDeployableCoverUsed = %v\n", accoladeDataProperties.DeployableCoverUsed)
		fmt.Printf("\t\tThermiteKills = %v\n", accoladeDataProperties.ThermiteKills)
		fmt.Printf("\t\tArKills = %v\n", accoladeDataProperties.ArKills)
		fmt.Printf("\t\tC4Kills = %v\n", accoladeDataProperties.C4Kills)
		fmt.Printf("\t\tSuicides = %v\n", accoladeDataProperties.Suicides)
		fmt.Printf("\t\tClutch = %v\n", accoladeDataProperties.Clutch)
		fmt.Printf("\t\tSurvivorKills = %v\n", accoladeDataProperties.SurvivorKills)
		fmt.Printf("\t\tKillstreakGunshipKills = %v\n", accoladeDataProperties.KillstreakGunshipKills)
		fmt.Printf("\t\tTimeSpentAsPassenger = %v\n", accoladeDataProperties.TimeSpentAsPassenger)
		fmt.Printf("\t\tReturns = %v\n", accoladeDataProperties.Returns)
		fmt.Printf("\t\tSmgHeadshots = %v\n", accoladeDataProperties.SmgHeadshots)
		fmt.Printf("\t\tLauncherDeaths = %v\n", accoladeDataProperties.LauncherDeaths)
		fmt.Printf("\t\tOneShotOneKills = %v\n", accoladeDataProperties.OneShotOneKills)
		fmt.Printf("\t\tAmmoBoxUsed = %v\n", accoladeDataProperties.AmmoBoxUsed)
		fmt.Printf("\t\tSpawnSelectSquad = %v\n", accoladeDataProperties.SpawnSelectSquad)
		fmt.Printf("\t\tWeaponPickups = %v\n", accoladeDataProperties.WeaponPickups)
		fmt.Printf("\t\tPointBlankKills = %v\n", accoladeDataProperties.PointBlankKills)
		fmt.Printf("\t\tTagsCaptured = %v\n", accoladeDataProperties.TagsCaptured)
		fmt.Printf("\t\tKillstreakGroundKills = %v\n", accoladeDataProperties.KillstreakGroundKills)
		fmt.Printf("\t\tDistanceTraveledInVehicle = %v\n", accoladeDataProperties.DistanceTraveledInVehicle)
		fmt.Printf("\t\tLongestLife = %v\n", accoladeDataProperties.LongestLife)
		fmt.Printf("\t\tStunHits = %v\n", accoladeDataProperties.StunHits)
		fmt.Printf("\t\tSpawnSelectFlag = %v\n", accoladeDataProperties.SpawnSelectFlag)
		fmt.Printf("\t\tShotgunHeadshots = %v\n", accoladeDataProperties.ShotgunHeadshots)
		fmt.Printf("\t\tBombDefused = %v\n", accoladeDataProperties.BombDefused)
		fmt.Printf("\t\tSnapshotHits = %v\n", accoladeDataProperties.SnapshotHits)
		fmt.Printf("\t\tNoKillsWithDeath = %v\n", accoladeDataProperties.NoKillsWithDeath)
		fmt.Printf("\t\tKillstreakAUAVAssists = %v\n", accoladeDataProperties.KillstreakAUAVAssists)
		fmt.Printf("\t\tKillstreakPersonalUAVKills = %v\n", accoladeDataProperties.KillstreakPersonalUAVKills)
		fmt.Printf("\t\tTacticalInsertionSpawns = %v\n", accoladeDataProperties.TacticalInsertionSpawns)
		fmt.Printf("\t\tLauncherKills = %v\n", accoladeDataProperties.LauncherKills)
		fmt.Printf("\t\tSpawnSelectVehicle = %v\n", accoladeDataProperties.SpawnSelectVehicle)
		fmt.Printf("\t\tMostKillsLeastDeaths = %v\n", accoladeDataProperties.MostKillsLeastDeaths)
		fmt.Printf("\t\tMostKills = %v\n", accoladeDataProperties.MostKills)
		fmt.Printf("\t\tDefends = %v\n", accoladeDataProperties.Defends)
		fmt.Printf("\t\tTimeSpentAsDriver = %v\n", accoladeDataProperties.TimeSpentAsDriver)
		fmt.Printf("\t\tBombDetonated = %v\n", accoladeDataProperties.BombDetonated)
		fmt.Printf("\t\tArHeadshots = %v\n", accoladeDataProperties.ArHeadshots)
		fmt.Printf("\t\tTimeOnPoint = %v\n", accoladeDataProperties.TimeOnPoint)
		fmt.Printf("\t\tLmgKills = %v\n", accoladeDataProperties.LmgKills)
		fmt.Printf("\t\tKillstreakUAVAssists = %v\n", accoladeDataProperties.KillstreakUAVAssists)
		fmt.Printf("\t\tCarepackagesCaptured = %v\n", accoladeDataProperties.CarepackagesCaptured)
		fmt.Printf("\t\tMostKillsLongestStreak = %v\n", accoladeDataProperties.MostKillsLongestStreak)
		fmt.Printf("\t\tKillstreakCruiseMissileKills = %v\n", accoladeDataProperties.KillstreakCruiseMissileKills)
		fmt.Printf("\t\tLongestStreak = %v\n", accoladeDataProperties.LongestStreak)
		fmt.Printf("\t\tDestroyedKillstreaks = %v\n", accoladeDataProperties.DestroyedKillstreaks)
		fmt.Printf("\t\tHipfireKills = %v\n", accoladeDataProperties.HipfireKills)
		fmt.Printf("\t\tStimDamageHealed = %v\n", accoladeDataProperties.StimDamageHealed)
		fmt.Printf("\t\tSkippedKillcams = %v\n", accoladeDataProperties.SkippedKillcams)
		fmt.Printf("\t\tLeastAssists = %v\n", accoladeDataProperties.LeastAssists)
		fmt.Printf("\t\tMostMultikills = %v\n", accoladeDataProperties.MostMultikills)
		fmt.Printf("\t\tHighestRankedKills = %v\n", accoladeDataProperties.HighestRankedKills)
		fmt.Printf("\t\tKillstreakAirstrikeKills = %v\n", accoladeDataProperties.KillstreakAirstrikeKills)
		fmt.Printf("\t\tDistanceTravelled = %v\n", accoladeDataProperties.DistanceTravelled)
		fmt.Printf("\t\tKillstreakKills = %v\n", accoladeDataProperties.KillstreakKills)
		fmt.Printf("\t\tSemtexKills = %v\n", accoladeDataProperties.SemtexKills)
		fmt.Printf("\t\tPenetrationKills = %v\n", accoladeDataProperties.PenetrationKills)
		fmt.Printf("\t\tExplosionsSurvived = %v\n", accoladeDataProperties.ExplosionsSurvived)
		fmt.Printf("\t\tHighestMultikill = %v\n", accoladeDataProperties.HighestMultikill)
		fmt.Printf("\t\tArDeaths = %v\n", accoladeDataProperties.ArDeaths)
		fmt.Printf("\t\tLongshotKills = %v\n", accoladeDataProperties.LongshotKills)
		fmt.Printf("\t\tProximityMineKills = %v\n", accoladeDataProperties.ProximityMineKills)
		fmt.Printf("\t\tTagsMegaBanked = %v\n", accoladeDataProperties.TagsMegaBanked)
		fmt.Printf("\t\tMostKillsMostHeadshots = %v\n", accoladeDataProperties.MostKillsMostHeadshots)
		fmt.Printf("\t\tFirstInfected = %v\n", accoladeDataProperties.FirstInfected)
		fmt.Printf("\t\tKillstreakCUAVAssists = %v\n", accoladeDataProperties.KillstreakCUAVAssists)
		fmt.Printf("\t\tThrowingKnifeKills = %v\n", accoladeDataProperties.ThrowingKnifeKills)
		fmt.Printf("\t\tExecutionKills = %v\n", accoladeDataProperties.ExecutionKills)
		fmt.Printf("\t\tLastSurvivor = %v\n", accoladeDataProperties.LastSurvivor)
		fmt.Printf("\t\tReconDroneMarks = %v\n", accoladeDataProperties.ReconDroneMarks)
		fmt.Printf("\t\tDeadSilenceKills = %v\n", accoladeDataProperties.DeadSilenceKills)
		fmt.Printf("\t\tRevengeKills = %v\n", accoladeDataProperties.RevengeKills)
		fmt.Printf("\t\tInfectedKills = %v\n", accoladeDataProperties.InfectedKills)
		fmt.Printf("\t\tKillEnemyTeam = %v\n", accoladeDataProperties.KillEnemyTeam)
		fmt.Printf("\t\tSniperKills = %v\n", accoladeDataProperties.SniperKills)
		fmt.Printf("\t\tKillstreakCluserStrikeKills = %v\n", accoladeDataProperties.KillstreakCluserStrikeKills)
		fmt.Printf("\t\tMeleeDeaths = %v\n", accoladeDataProperties.MeleeDeaths)
		fmt.Printf("\t\tTimeWatchingKillcams = %v\n", accoladeDataProperties.TimeWatchingKillcams)
		fmt.Printf("\t\tKillstreakTankKills = %v\n", accoladeDataProperties.KillstreakTankKills)
		fmt.Printf("\t\tNoKillNoDeath = %v\n", accoladeDataProperties.NoKillNoDeath)
		fmt.Printf("\t\tShotgunDeaths = %v\n", accoladeDataProperties.ShotgunDeaths)
		fmt.Printf("\t\tKillstreakChopperGunnerKills = %v\n", accoladeDataProperties.KillstreakChopperGunnerKills)
		fmt.Printf("\t\tShotsFired = %v\n", accoladeDataProperties.ShotsFired)
		fmt.Printf("\t\tStoppingPowerKills = %v\n", accoladeDataProperties.StoppingPowerKills)
		fmt.Printf("\t\tPistolPeaths = %v\n", accoladeDataProperties.PistolPeaths)
		fmt.Printf("\t\tKillstreakShieldTurretKills = %v\n", accoladeDataProperties.KillstreakShieldTurretKills)
		fmt.Printf("\t\tTimeCrouched = %v\n", accoladeDataProperties.TimeCrouched)
		fmt.Printf("\t\tNoDeathsFromBehind = %v\n", accoladeDataProperties.NoDeathsFromBehind)
		fmt.Printf("\t\tBombPlanted = %v\n", accoladeDataProperties.BombPlanted)
		fmt.Printf("\t\tSetbacks = %v\n", accoladeDataProperties.Setbacks)
		fmt.Printf("\t\tSmgKills = %v\n", accoladeDataProperties.SmgKills)
		fmt.Printf("\t\tClaymoreKills = %v\n", accoladeDataProperties.ClaymoreKills)
		fmt.Printf("\t\tKills10NoDeaths = %v\n", accoladeDataProperties.Kills10NoDeaths)
		fmt.Printf("\t\tPistolHeadshots = %v\n", accoladeDataProperties.PistolHeadshots)
		fmt.Printf("\t\tKillstreakVTOLJetKills = %v\n", accoladeDataProperties.KillstreakVTOLJetKills)
		fmt.Printf("\t\tHeadshots = %v\n", accoladeDataProperties.Headshots)
		fmt.Printf("\t\tMostDeaths = %v\n", accoladeDataProperties.MostDeaths)
		fmt.Printf("\t\tAdsKills = %v\n", accoladeDataProperties.AdsKills)
		fmt.Printf("\t\tEmpDroneHits = %v\n", accoladeDataProperties.EmpDroneHits)
		fmt.Printf("\t\tDefenderKills = %v\n", accoladeDataProperties.DefenderKills)
		fmt.Printf("\t\tLauncherHeadshots = %v\n", accoladeDataProperties.LauncherHeadshots)
		fmt.Printf("\t\tTimesSelectedAsSquadLeader = %v\n", accoladeDataProperties.TimesSelectedAsSquadLeader)
		fmt.Printf("\t\tKillstreakAirKills = %v\n", accoladeDataProperties.KillstreakAirKills)
		fmt.Printf("\t\tAssaults = %v\n", accoladeDataProperties.Assaults)
		fmt.Printf("\t\tFragKills = %v\n", accoladeDataProperties.FragKills)
		fmt.Printf("\t\tKillstreakEmergencyAirdropUsed = %v\n", accoladeDataProperties.KillstreakEmergencyAirdropUsed)
		fmt.Printf("\t\tCaptures = %v\n", accoladeDataProperties.Captures)
		fmt.Printf("\t\tKillstreakChopperSupportKills = %v\n", accoladeDataProperties.KillstreakChopperSupportKills)
		fmt.Printf("\t\tSpawnSelectBase = %v\n", accoladeDataProperties.SpawnSelectBase)
		fmt.Printf("\t\tNoKill10Deaths = %v\n", accoladeDataProperties.NoKill10Deaths)
		fmt.Printf("\t\tLeastDeaths = %v\n", accoladeDataProperties.LeastDeaths)
		fmt.Printf("\t\tKillstreakSentryGunKills = %v\n", accoladeDataProperties.KillstreakSentryGunKills)
		fmt.Printf("\t\tLongestTimeSpentOnWeapon = %v\n", accoladeDataProperties.LongestTimeSpentOnWeapon)
		fmt.Printf("\t\tLowerRankedKills = %v\n", accoladeDataProperties.LowerRankedKills)
		fmt.Printf("\t\tTrophySystemHits = %v\n", accoladeDataProperties.TrophySystemHits)
		fmt.Printf("\t\tClutchRevives = %v\n", accoladeDataProperties.ClutchRevives)
		fmt.Printf("\t\tLowestAvgAltitude = %v\n", accoladeDataProperties.LowestAvgAltitude)
		fmt.Printf("\t\tPickups = %v\n", accoladeDataProperties.Pickups)
		fmt.Printf("\t\tPistolKills = %v\n", accoladeDataProperties.PistolKills)
		fmt.Printf("\t\tReloads = %v\n", accoladeDataProperties.Reloads)
	}

	return nil
}

// Separate function for sending GamerStats request
func sendGamerStats(c *gocod.Client, gamer *gocod.Gamer) error {
	gamerStatsResp, err := c.GamerStats(context.Background(), myTitleMw, gamer, myGameType)
	if err != nil {
		return err
	}
	fmt.Printf("\nGamerStats Status = %v\n", gamerStatsResp.Status)

	fmt.Printf("Title = %v\n", gamerStatsResp.Data.Title)
	fmt.Printf("Platform = %v\n", gamerStatsResp.Data.Platform)
	fmt.Printf("Username = %v\n", gamerStatsResp.Data.Username)
	fmt.Printf("Type = %v\n", gamerStatsResp.Data.Type)
	fmt.Printf("Level = %v\n", gamerStatsResp.Data.Level)
	fmt.Printf("MaxLevel = %v\n", gamerStatsResp.Data.MaxLevel)
	fmt.Printf("LevelXpRemainder = %v\n", gamerStatsResp.Data.LevelXpRemainder)
	fmt.Printf("LevelXpGained = %v\n", gamerStatsResp.Data.LevelXpGained)
	fmt.Printf("Prestige = %v\n", gamerStatsResp.Data.Prestige)
	fmt.Printf("PrestigeID = %v\n", gamerStatsResp.Data.PrestigeID)
	fmt.Printf("MaxPrestige = %v\n", gamerStatsResp.Data.MaxPrestige)
	fmt.Printf("TotalXp = %v\n", gamerStatsResp.Data.TotalXp)
	fmt.Printf("ParagonRank = %v\n", gamerStatsResp.Data.ParagonRank)
	fmt.Printf("ParagonID = %v\n", gamerStatsResp.Data.ParagonID)
	fmt.Printf("S = %v\n", gamerStatsResp.Data.S)
	fmt.Printf("P = %v\n", gamerStatsResp.Data.P)

	// ALL - gamer stats summary
	allProperties := gamerStatsResp.Data.Lifetime.All.Properties
	fmt.Printf("ALL Gamer Stats Summary\n")

	fmt.Printf("\tRecordLongestWinStreak = %v\n", allProperties.RecordLongestWinStreak)
	fmt.Printf("\tRecordXpInAMatch = %v\n", allProperties.RecordXpInAMatch)
	fmt.Printf("\tAccuracy = %v\n", allProperties.Accuracy)
	fmt.Printf("\tLosses = %v\n", allProperties.Losses)
	fmt.Printf("\tTotalGamesPlayed = %v\n", allProperties.TotalGamesPlayed)
	fmt.Printf("\tScore = %v\n", allProperties.Score)
	fmt.Printf("\tWinLossRatio = %v\n", allProperties.WinLossRatio)
	fmt.Printf("\tTotalShots = %v\n", allProperties.TotalShots)
	fmt.Printf("\tBestScoreXp = %v\n", allProperties.BestScoreXp)
	fmt.Printf("\tGamesPlayed = %v\n", allProperties.GamesPlayed)
	fmt.Printf("\tBestSquardKills = %v\n", allProperties.BestSquardKills)
	fmt.Printf("\tBestSguardWave = %v\n", allProperties.BestSguardWave)
	fmt.Printf("\tBestConfirmed = %v\n", allProperties.BestConfirmed)
	fmt.Printf("\tDeaths = %v\n", allProperties.Deaths)
	fmt.Printf("\tWins = %v\n", allProperties.Wins)
	fmt.Printf("\tBestSquardCrates = %v\n", allProperties.BestSquardCrates)
	fmt.Printf("\tKdRatio = %v\n", allProperties.KdRatio)
	fmt.Printf("\tBestAssists = %v\n", allProperties.BestAssists)
	fmt.Printf("\tBestFieldgoals = %v\n", allProperties.BestFieldgoals)
	fmt.Printf("\tBestScore = %v\n", allProperties.BestScore)
	fmt.Printf("\tRecordDeathsInAMatch = %v\n", allProperties.RecordDeathsInAMatch)
	fmt.Printf("\tScorePerGame = %v\n", allProperties.ScorePerGame)
	fmt.Printf("\tBestSPM = %v\n", allProperties.BestSPM)
	fmt.Printf("\tBestKillChains = %v\n", allProperties.BestKillChains)
	fmt.Printf("\tRecordKillsInAMatch = %v\n", allProperties.RecordKillsInAMatch)
	fmt.Printf("\tSuicides = %v\n", allProperties.Suicides)
	fmt.Printf("\tWlRatio = %v\n", allProperties.WlRatio)
	fmt.Printf("\tCurrentWinStreak = %v\n", allProperties.CurrentWinStreak)
	fmt.Printf("\tBestMatchBonusXp = %v\n", allProperties.BestMatchBonusXp)
	fmt.Printf("\tBestMatchXp = %v\n", allProperties.BestMatchXp)
	fmt.Printf("\tBestSguardWeaponLevel = %v\n", allProperties.BestSguardWeaponLevel)
	fmt.Printf("\tBestKD = %v\n", allProperties.BestKD)
	fmt.Printf("\tKills = %v\n", allProperties.Kills)
	fmt.Printf("\tBestKillsAsInfected = %v\n", allProperties.BestKillsAsInfected)
	fmt.Printf("\tBestReturns = %v\n", allProperties.BestReturns)
	fmt.Printf("\tBestStabs = %v\n", allProperties.BestStabs)
	fmt.Printf("\tBestKillsAsSurvivor = %v\n", allProperties.BestKillsAsSurvivor)
	fmt.Printf("\tTimePlayedTotal = %v\n", allProperties.TimePlayedTotal)
	fmt.Printf("\tBestDestructions = %v\n", allProperties.BestDestructions)
	fmt.Printf("\tHeadshots = %v\n", allProperties.Headshots)
	fmt.Printf("\tBestRescues = %v\n", allProperties.BestRescues)
	fmt.Printf("\tAssists = %v\n", allProperties.Assists)
	fmt.Printf("\tTies = %v\n", allProperties.Ties)
	fmt.Printf("\tRecordKillStreak = %v\n", allProperties.RecordKillStreak)
	fmt.Printf("\tBestPlants = %v\n", allProperties.BestPlants)
	fmt.Printf("\tMisses = %v\n", allProperties.Misses)
	fmt.Printf("\tBestDamage = %v\n", allProperties.BestDamage)
	fmt.Printf("\tBestSetbacks = %v\n", allProperties.BestSetbacks)
	fmt.Printf("\tBestTouchdowns = %v\n", allProperties.BestTouchdowns)
	fmt.Printf("\tScorePerMinute = %v\n", allProperties.ScorePerMinute)
	fmt.Printf("\tBestDeaths = %v\n", allProperties.BestDeaths)
	fmt.Printf("\tBestMedalXp = %v\n", allProperties.BestMedalXp)
	fmt.Printf("\tBestDefends = %v\n", allProperties.BestDefends)
	fmt.Printf("\tBestSquardRevives = %v\n", allProperties.BestSquardRevives)
	fmt.Printf("\tBestKills = %v\n", allProperties.BestKills)
	fmt.Printf("\tBestDefuses = %v\n", allProperties.BestDefuses)
	fmt.Printf("\tBestCaptures = %v\n", allProperties.BestCaptures)
	fmt.Printf("\tHits = %v\n", allProperties.Hits)
	fmt.Printf("\tBestKillStreak = %v\n", allProperties.BestKillStreak)
	fmt.Printf("\tBestDenied = %v\n", allProperties.BestDenied)

	// Data.Lifetime.Mode is of type map[string]GamerStatsResponseDataLifetimeModeAnon
	mode := gamerStatsResp.Data.Lifetime.Mode
	for k, v := range mode {
		fmt.Printf("Mode (Key) = %v:\n", k)

		fmt.Printf("\tKills = %v\n", v.Kills)
		fmt.Printf("\tScore = %v\n", v.Score)
		fmt.Printf("\tTimePlayed = %v\n", v.TimePlayed)
		fmt.Printf("\tKdRatio = %v\n", v.KdRatio)
		fmt.Printf("\tSetBacks = %v\n", v.SetBacks)
		fmt.Printf("\tScorePerMinute = %v\n", v.ScorePerMinute)
		fmt.Printf("\tStabs = %v\n", v.Stabs)
		fmt.Printf("\tDeaths = %v\n", v.Deaths)
	}

	// Data.Lifetime.ItemData is of type map[string]map[string]GamerStatsResponseDataLifetimeItemDataAnon
	itemData := gamerStatsResp.Data.Lifetime.ItemData
	for k, v := range itemData {
		fmt.Printf("Item Category (Key) = %v:\n", k)

		for k2, v2 := range v {
			fmt.Printf("\tItem (Key) %v:\n", k2)

			fmt.Printf("\t\tHits = %v\n", v2.Hits)
			fmt.Printf("\t\tKills = %v\n", v2.Kills)
			fmt.Printf("\t\tKdRatio = %v\n", v2.KdRatio)
			fmt.Printf("\t\tHeadshots = %v\n", v2.Headshots)
			fmt.Printf("\t\tAccuracy = %v\n", v2.Accuracy)
			fmt.Printf("\t\tShots = %v\n", v2.Shots)
			fmt.Printf("\t\tDeaths = %v\n", v2.Deaths)
		}
	}

	// Data.Lifetime.ScorestreakData. is of type map[string]GamerStatsResponseDataLifetimeScorestreakDataLethalScorestreakDataAnon
	lethalScorestreakData := gamerStatsResp.Data.Lifetime.ScorestreakData.LethalScorestreakData
	for k, v := range lethalScorestreakData {
		fmt.Printf("Lethal Core Streak (Key) = %v:\n", k)

		fmt.Printf("\tExtraStat1 = %v\n", v.ExtraStat1)
		fmt.Printf("\tUses = %v\n", v.Uses)
		fmt.Printf("\tAwardedCount = %v\n", v.AwardedCount)
	}

	// Data.Lifetime.ScorestreakData.SupportScorestreakData. is of type map[string]GamerStatsResponseDataLifetimeScorestreakDataSupportScorestreakDataAnon
	supportScorestreakData := gamerStatsResp.Data.Lifetime.ScorestreakData.SupportScorestreakData
	for k, v := range supportScorestreakData {
		fmt.Printf("Support Core Streak (Key) = %v:\n", k)

		fmt.Printf("\tExtraStat1 = %v\n", v.ExtraStat1)
		fmt.Printf("\tUses = %v\n", v.Uses)
		fmt.Printf("\tAwardedCount = %v\n", v.AwardedCount)
	}

	// Accolade Data
	accoladeDataProperties := gamerStatsResp.Data.Lifetime.AccoladeData.Properties
	fmt.Printf("Accolade Data\n")

	fmt.Printf("\tClassChanges = %v\n", accoladeDataProperties.ClassChanges)
	fmt.Printf("\tHighestAvgAltitude = %v\n", accoladeDataProperties.HighestAvgAltitude)
	fmt.Printf("\tKillsFromBehind = %v\n", accoladeDataProperties.KillsFromBehind)
	fmt.Printf("\tLmgDeaths = %v\n", accoladeDataProperties.LmgDeaths)
	fmt.Printf("\tRiotShieldDamageAbsorbed = %v\n", accoladeDataProperties.RiotShieldDamageAbsorbed)
	fmt.Printf("\tFlashbangHits = %v\n", accoladeDataProperties.FlashbangHits)
	fmt.Printf("\tMeleeKills = %v\n", accoladeDataProperties.MeleeKills)
	fmt.Printf("\tTagsLargestBank = %v\n", accoladeDataProperties.TagsLargestBank)
	fmt.Printf("\tShotgunKills = %v\n", accoladeDataProperties.ShotgunKills)
	fmt.Printf("\tSniperDeaths = %v\n", accoladeDataProperties.SniperDeaths)
	fmt.Printf("\tTimeProne = %v\n", accoladeDataProperties.TimeProne)
	fmt.Printf("\tKillstreakWhitePhosphorousKillsAssists = %v\n", accoladeDataProperties.KillstreakWhitePhosphorousKillsAssists)
	fmt.Printf("\tShortestLife = %v\n", accoladeDataProperties.ShortestLife)
	fmt.Printf("\tDeathsFromBehind = %v\n", accoladeDataProperties.DeathsFromBehind)
	fmt.Printf("\tHigherRankedKills = %v\n", accoladeDataProperties.HigherRankedKills)
	fmt.Printf("\tMostAssists = %v\n", accoladeDataProperties.MostAssists)
	fmt.Printf("\tLeastKills = %v\n", accoladeDataProperties.LeastKills)
	fmt.Printf("\tTagsDenied = %v\n", accoladeDataProperties.TagsDenied)
	fmt.Printf("\tKillstreakWheelsonKills = %v\n", accoladeDataProperties.KillstreakWheelsonKills)
	fmt.Printf("\tSniperHeadshots = %v\n", accoladeDataProperties.SniperHeadshots)
	fmt.Printf("\tKillstreakJuggernautKills = %v\n", accoladeDataProperties.KillstreakJuggernautKills)
	fmt.Printf("\tSmokesUsed = %v\n", accoladeDataProperties.SmokesUsed)
	fmt.Printf("\tAvengerKills = %v\n", accoladeDataProperties.AvengerKills)
	fmt.Printf("\tDecoyHits = %v\n", accoladeDataProperties.DecoyHits)
	fmt.Printf("\tKillstreakCarePackageUsed = %v\n", accoladeDataProperties.KillstreakCarePackageUsed)
	fmt.Printf("\tMolotovKills = %v\n", accoladeDataProperties.MolotovKills)
	fmt.Printf("\tGasHits = %v\n", accoladeDataProperties.GasHits)
	fmt.Printf("\tComebackKills = %v\n", accoladeDataProperties.ComebackKills)
	fmt.Printf("\tLmgHeadshots = %v\n", accoladeDataProperties.LmgHeadshots)
	fmt.Printf("\tSmgDeaths = %v\n", accoladeDataProperties.SmgDeaths)
	fmt.Printf("\tCarrierKills = %v\n", accoladeDataProperties.CarrierKills)
	fmt.Printf("\tDeployableCoverUsed = %v\n", accoladeDataProperties.DeployableCoverUsed)
	fmt.Printf("\tThermiteKills = %v\n", accoladeDataProperties.ThermiteKills)
	fmt.Printf("\tArKills = %v\n", accoladeDataProperties.ArKills)
	fmt.Printf("\tC4Kills = %v\n", accoladeDataProperties.C4Kills)
	fmt.Printf("\tSuicides = %v\n", accoladeDataProperties.Suicides)
	fmt.Printf("\tClutch = %v\n", accoladeDataProperties.Clutch)
	fmt.Printf("\tSurvivorKills = %v\n", accoladeDataProperties.SurvivorKills)
	fmt.Printf("\tKillstreakGunshipKills = %v\n", accoladeDataProperties.KillstreakGunshipKills)
	fmt.Printf("\tTimeSpentAsPassenger = %v\n", accoladeDataProperties.TimeSpentAsPassenger)
	fmt.Printf("\tReturns = %v\n", accoladeDataProperties.Returns)
	fmt.Printf("\tSmgHeadshots = %v\n", accoladeDataProperties.SmgHeadshots)
	fmt.Printf("\tLauncherDeaths = %v\n", accoladeDataProperties.LauncherDeaths)
	fmt.Printf("\tOneShotOneKills = %v\n", accoladeDataProperties.OneShotOneKills)
	fmt.Printf("\tAmmoBoxUsed = %v\n", accoladeDataProperties.AmmoBoxUsed)
	fmt.Printf("\tSpawnSelectSquad = %v\n", accoladeDataProperties.SpawnSelectSquad)
	fmt.Printf("\tWeaponPickups = %v\n", accoladeDataProperties.WeaponPickups)
	fmt.Printf("\tPointBlankKills = %v\n", accoladeDataProperties.PointBlankKills)
	fmt.Printf("\tTagsCaptured = %v\n", accoladeDataProperties.TagsCaptured)
	fmt.Printf("\tKillstreakGroundKills = %v\n", accoladeDataProperties.KillstreakGroundKills)
	fmt.Printf("\tDistanceTraveledInVehicle = %v\n", accoladeDataProperties.DistanceTraveledInVehicle)
	fmt.Printf("\tLongestLife = %v\n", accoladeDataProperties.LongestLife)
	fmt.Printf("\tStunHits = %v\n", accoladeDataProperties.StunHits)
	fmt.Printf("\tSpawnSelectFlag = %v\n", accoladeDataProperties.SpawnSelectFlag)
	fmt.Printf("\tShotgunHeadshots = %v\n", accoladeDataProperties.ShotgunHeadshots)
	fmt.Printf("\tBombDefused = %v\n", accoladeDataProperties.BombDefused)
	fmt.Printf("\tSnapshotHits = %v\n", accoladeDataProperties.SnapshotHits)
	fmt.Printf("\tNoKillsWithDeath = %v\n", accoladeDataProperties.NoKillsWithDeath)
	fmt.Printf("\tKillstreakAUAVAssists = %v\n", accoladeDataProperties.KillstreakAUAVAssists)
	fmt.Printf("\tKillstreakPersonalUAVKills = %v\n", accoladeDataProperties.KillstreakPersonalUAVKills)
	fmt.Printf("\tTacticalInsertionSpawns = %v\n", accoladeDataProperties.TacticalInsertionSpawns)
	fmt.Printf("\tLauncherKills = %v\n", accoladeDataProperties.LauncherKills)
	fmt.Printf("\tSpawnSelectVehicle = %v\n", accoladeDataProperties.SpawnSelectVehicle)
	fmt.Printf("\tMostKillsLeastDeaths = %v\n", accoladeDataProperties.MostKillsLeastDeaths)
	fmt.Printf("\tMostKills = %v\n", accoladeDataProperties.MostKills)
	fmt.Printf("\tDefends = %v\n", accoladeDataProperties.Defends)
	fmt.Printf("\tTimeSpentAsDriver = %v\n", accoladeDataProperties.TimeSpentAsDriver)
	fmt.Printf("\tBombDetonated = %v\n", accoladeDataProperties.BombDetonated)
	fmt.Printf("\tArHeadshots = %v\n", accoladeDataProperties.ArHeadshots)
	fmt.Printf("\tTimeOnPoint = %v\n", accoladeDataProperties.TimeOnPoint)
	fmt.Printf("\tLmgKills = %v\n", accoladeDataProperties.LmgKills)
	fmt.Printf("\tKillstreakUAVAssists = %v\n", accoladeDataProperties.KillstreakUAVAssists)
	fmt.Printf("\tCarepackagesCaptured = %v\n", accoladeDataProperties.CarepackagesCaptured)
	fmt.Printf("\tMostKillsLongestStreak = %v\n", accoladeDataProperties.MostKillsLongestStreak)
	fmt.Printf("\tKillstreakCruiseMissileKills = %v\n", accoladeDataProperties.KillstreakCruiseMissileKills)
	fmt.Printf("\tLongestStreak = %v\n", accoladeDataProperties.LongestStreak)
	fmt.Printf("\tDestroyedKillstreaks = %v\n", accoladeDataProperties.DestroyedKillstreaks)
	fmt.Printf("\tHipfireKills = %v\n", accoladeDataProperties.HipfireKills)
	fmt.Printf("\tStimDamageHealed = %v\n", accoladeDataProperties.StimDamageHealed)
	fmt.Printf("\tSkippedKillcams = %v\n", accoladeDataProperties.SkippedKillcams)
	fmt.Printf("\tLeastAssists = %v\n", accoladeDataProperties.LeastAssists)
	fmt.Printf("\tMostMultikills = %v\n", accoladeDataProperties.MostMultikills)
	fmt.Printf("\tHighestRankedKills = %v\n", accoladeDataProperties.HighestRankedKills)
	fmt.Printf("\tKillstreakAirstrikeKills = %v\n", accoladeDataProperties.KillstreakAirstrikeKills)
	fmt.Printf("\tDistanceTravelled = %v\n", accoladeDataProperties.DistanceTravelled)
	fmt.Printf("\tKillstreakKills = %v\n", accoladeDataProperties.KillstreakKills)
	fmt.Printf("\tSemtexKills = %v\n", accoladeDataProperties.SemtexKills)
	fmt.Printf("\tPenetrationKills = %v\n", accoladeDataProperties.PenetrationKills)
	fmt.Printf("\tExplosionsSurvived = %v\n", accoladeDataProperties.ExplosionsSurvived)
	fmt.Printf("\tHighestMultikill = %v\n", accoladeDataProperties.HighestMultikill)
	fmt.Printf("\tArDeaths = %v\n", accoladeDataProperties.ArDeaths)
	fmt.Printf("\tLongshotKills = %v\n", accoladeDataProperties.LongshotKills)
	fmt.Printf("\tProximityMineKills = %v\n", accoladeDataProperties.ProximityMineKills)
	fmt.Printf("\tTagsMegaBanked = %v\n", accoladeDataProperties.TagsMegaBanked)
	fmt.Printf("\tMostKillsMostHeadshots = %v\n", accoladeDataProperties.MostKillsMostHeadshots)
	fmt.Printf("\tFirstInfected = %v\n", accoladeDataProperties.FirstInfected)
	fmt.Printf("\tKillstreakCUAVAssists = %v\n", accoladeDataProperties.KillstreakCUAVAssists)
	fmt.Printf("\tThrowingKnifeKills = %v\n", accoladeDataProperties.ThrowingKnifeKills)
	fmt.Printf("\tExecutionKills = %v\n", accoladeDataProperties.ExecutionKills)
	fmt.Printf("\tLastSurvivor = %v\n", accoladeDataProperties.LastSurvivor)
	fmt.Printf("\tReconDroneMarks = %v\n", accoladeDataProperties.ReconDroneMarks)
	fmt.Printf("\tDeadSilenceKills = %v\n", accoladeDataProperties.DeadSilenceKills)
	fmt.Printf("\tRevengeKills = %v\n", accoladeDataProperties.RevengeKills)
	fmt.Printf("\tInfectedKills = %v\n", accoladeDataProperties.InfectedKills)
	fmt.Printf("\tKillEnemyTeam = %v\n", accoladeDataProperties.KillEnemyTeam)
	fmt.Printf("\tSniperKills = %v\n", accoladeDataProperties.SniperKills)
	fmt.Printf("\tKillstreakCluserStrikeKills = %v\n", accoladeDataProperties.KillstreakCluserStrikeKills)
	fmt.Printf("\tMeleeDeaths = %v\n", accoladeDataProperties.MeleeDeaths)
	fmt.Printf("\tTimeWatchingKillcams = %v\n", accoladeDataProperties.TimeWatchingKillcams)
	fmt.Printf("\tKillstreakTankKills = %v\n", accoladeDataProperties.KillstreakTankKills)
	fmt.Printf("\tNoKillNoDeath = %v\n", accoladeDataProperties.NoKillNoDeath)
	fmt.Printf("\tShotgunDeaths = %v\n", accoladeDataProperties.ShotgunDeaths)
	fmt.Printf("\tKillstreakChopperGunnerKills = %v\n", accoladeDataProperties.KillstreakChopperGunnerKills)
	fmt.Printf("\tShotsFired = %v\n", accoladeDataProperties.ShotsFired)
	fmt.Printf("\tStoppingPowerKills = %v\n", accoladeDataProperties.StoppingPowerKills)
	fmt.Printf("\tPistolPeaths = %v\n", accoladeDataProperties.PistolPeaths)
	fmt.Printf("\tKillstreakShieldTurretKills = %v\n", accoladeDataProperties.KillstreakShieldTurretKills)
	fmt.Printf("\tTimeCrouched = %v\n", accoladeDataProperties.TimeCrouched)
	fmt.Printf("\tNoDeathsFromBehind = %v\n", accoladeDataProperties.NoDeathsFromBehind)
	fmt.Printf("\tBombPlanted = %v\n", accoladeDataProperties.BombPlanted)
	fmt.Printf("\tSetbacks = %v\n", accoladeDataProperties.Setbacks)
	fmt.Printf("\tSmgKills = %v\n", accoladeDataProperties.SmgKills)
	fmt.Printf("\tClaymoreKills = %v\n", accoladeDataProperties.ClaymoreKills)
	fmt.Printf("\tKills10NoDeaths = %v\n", accoladeDataProperties.Kills10NoDeaths)
	fmt.Printf("\tPistolHeadshots = %v\n", accoladeDataProperties.PistolHeadshots)
	fmt.Printf("\tKillstreakVTOLJetKills = %v\n", accoladeDataProperties.KillstreakVTOLJetKills)
	fmt.Printf("\tHeadshots = %v\n", accoladeDataProperties.Headshots)
	fmt.Printf("\tMostDeaths = %v\n", accoladeDataProperties.MostDeaths)
	fmt.Printf("\tAdsKills = %v\n", accoladeDataProperties.AdsKills)
	fmt.Printf("\tEmpDroneHits = %v\n", accoladeDataProperties.EmpDroneHits)
	fmt.Printf("\tDefenderKills = %v\n", accoladeDataProperties.DefenderKills)
	fmt.Printf("\tLauncherHeadshots = %v\n", accoladeDataProperties.LauncherHeadshots)
	fmt.Printf("\tTimesSelectedAsSquadLeader = %v\n", accoladeDataProperties.TimesSelectedAsSquadLeader)
	fmt.Printf("\tKillstreakAirKills = %v\n", accoladeDataProperties.KillstreakAirKills)
	fmt.Printf("\tAssaults = %v\n", accoladeDataProperties.Assaults)
	fmt.Printf("\tFragKills = %v\n", accoladeDataProperties.FragKills)
	fmt.Printf("\tKillstreakEmergencyAirdropUsed = %v\n", accoladeDataProperties.KillstreakEmergencyAirdropUsed)
	fmt.Printf("\tCaptures = %v\n", accoladeDataProperties.Captures)
	fmt.Printf("\tKillstreakChopperSupportKills = %v\n", accoladeDataProperties.KillstreakChopperSupportKills)
	fmt.Printf("\tSpawnSelectBase = %v\n", accoladeDataProperties.SpawnSelectBase)
	fmt.Printf("\tNoKill10Deaths = %v\n", accoladeDataProperties.NoKill10Deaths)
	fmt.Printf("\tLeastDeaths = %v\n", accoladeDataProperties.LeastDeaths)
	fmt.Printf("\tKillstreakSentryGunKills = %v\n", accoladeDataProperties.KillstreakSentryGunKills)
	fmt.Printf("\tLongestTimeSpentOnWeapon = %v\n", accoladeDataProperties.LongestTimeSpentOnWeapon)
	fmt.Printf("\tLowerRankedKills = %v\n", accoladeDataProperties.LowerRankedKills)
	fmt.Printf("\tTrophySystemHits = %v\n", accoladeDataProperties.TrophySystemHits)
	fmt.Printf("\tClutchRevives = %v\n", accoladeDataProperties.ClutchRevives)
	fmt.Printf("\tLowestAvgAltitude = %v\n", accoladeDataProperties.LowestAvgAltitude)
	fmt.Printf("\tPickups = %v\n", accoladeDataProperties.Pickups)
	fmt.Printf("\tPistolKills = %v\n", accoladeDataProperties.PistolKills)
	fmt.Printf("\tReloads = %v\n", accoladeDataProperties.Reloads)

	return nil
}
