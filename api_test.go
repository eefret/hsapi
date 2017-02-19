package hsapi_test

import (
	"fmt"
	hs "github.com/eefret/hsapi"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestHs(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewInfo()
	resp := api.Info(config)
	assert.NotEmpty(t, resp, "Empty Response")
}

func TestInfo(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewInfo()
	config.Locale = hs.EsES
	resp := api.Info(config)
	assert.NotEmpty(t, resp, "Empty Response")
}

func TestGetCard(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewGetCard("EX1_572")
	config.Collectible = true
	resp, err := api.GetCard(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while instantiating hearthstoneApi")
	config.Collectible = false
	config.Name = "Ysera"
	config.Locale = hs.EsES
	config.Collectible = true
	resp, err = api.GetCard(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing GetCard")
}

func TestAllCards(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewAllCards()
	config.Collectible = true
	config.Attack = 5
	config.Durability = 5
	config.Cost = 5
	resp, err := api.AllCards(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing allcards")
}

func TestCardBacks(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardBacks()
	config.Locale = hs.EsES
	resp, err := api.CardBacks(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing CardBack")
	config.Locale = hs.EsES
	resp, err = api.CardBacks(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing CardBack")
}

func TestSearch(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardSearch("tirion fordring")
	config.Collectible = true
	resp, err := api.Search(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing Search")
	assert.Equal(t, 1, len(resp), "obtained more cards than expected")
	assert.Equal(t, resp[0].Rarity, hs.Legendary.String(), "Wrong card fetched")
	config = hs.NewCardSearch("alexstrasza")
	config.Collectible = true
	resp, err = api.Search(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing Search")
	assert.Equal(t, 2, len(resp), "obtained more cards than expected")
	assert.Equal(t, resp[0].CardID, "AT_071", "Wrong cards fetched")
	assert.Equal(t, resp[1].CardID, "EX1_561", "Wrong cards fetched")
}

func TestSet(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardsBySet(hs.TheGrandTournament)
	config.Collectible = true
	resp, err := api.CardsBySet(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing Search")
	assert.NotEqual(t, 0, len(resp), "obtained no cards ")
	assert.Equal(t, resp[0].Type, hs.Spell.String(), "Wrong card fetched")
}

func TestClass(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardsByClass(hs.Warlock)
	resp, err := api.CardsByClass(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing Search")
	assert.NotEqual(t, 0, len(resp), "obtained no cards ")
	assert.Equal(t, resp[0].Type, hs.Enchantment.String(), "Wrong card fetched")
}

func TestFaction(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardsByFaction(hs.Horde)
	config.Collectible = true
	resp, err := api.CardsByFaction(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing Search")
	assert.NotEqual(t, 0, len(resp), "obtained no cards ")
	assert.Equal(t, resp[0].Faction, hs.Horde.String(), "Wrong card fetched")
}

func TestQuality(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardsByQuality(hs.Legendary)
	config.Collectible = true
	resp, err := api.CardsByQuality(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing Search")
	assert.NotEqual(t, 0, len(resp), "obtained no cards ")
	assert.Equal(t, resp[0].Name, "Patches the Pirate", "Wrong card fetched")
}

func TestRace(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardsByRace(hs.Dragon)
	config.Collectible = true
	resp, err := api.CardsByRace(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing Search")
	assert.NotEqual(t, 0, len(resp), "obtained no cards ")
	assert.Equal(t, resp[0].CardID, "BRM_004", "Wrong card fetched")
}

func TestType(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardsByType(hs.Weapon)
	config.Collectible = true
	resp, err := api.CardsByType(config)
	assert.NotEmpty(t, resp, "Empty Response")
	assert.NoError(t, err, "An error ocurred while testing Search")
	assert.NotEqual(t, 0, len(resp), "obtained no cards ")
	assert.Equal(t, resp[0].PlayerClass, hs.Warrior.String(), "Wrong card fetched")
}

func TestImgCard(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardImage("EX1_383")
	config.Gold = true
	resp, err := api.CardImage(config)
	assert.Equal(t, hs.GIF.String(), resp.Extension.String(), "Wrong extension")
	assert.NoError(t, err, "An error ocurred while testing CardImage")

	if _, err := os.Stat("test_images"); os.IsNotExist(err) {
		os.Mkdir("."+string(filepath.Separator)+"test_images", 0777)
	}
	f, err := os.Create("test_images/Tirion_Fordring.gif")
	assert.NoError(t, err, "An error ocurred while testing CardImage")
	defer f.Close()
	f.Write(resp.Image)
	_, existsErr := os.Stat("test_images/Tirion_Fordring.gif")
	assert.True(t, !os.IsNotExist(existsErr), "File couldnt be created")
	assert.NoError(t, os.Remove("test_images/Tirion_Fordring.gif"), "couldn't remove the file")
}

func TestSoundCard(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN")
	api.Debug = true
	config := hs.NewCardSound("EX1_383", hs.PLAY)
	config.Locale = hs.EsES
	config.Extension = hs.MP3
	resp, err := api.CardSound(config)
	assert.Equal(t, hs.MP3.String(), resp.Extension.String(), "Wrong extension")
	assert.NoError(t, err, "An error ocurred while testing CardImage")
	fmt.Printf("%q\n", resp)
	if _, err := os.Stat("test_sounds"); os.IsNotExist(err) {
		os.Mkdir("."+string(filepath.Separator)+"test_sounds", 0777)
	}
	f, err := os.Create("test_sounds/Tirion_Fordring_Play.mp3")
	assert.NoError(t, err, "An error ocurred while testing CardImage")
	defer f.Close()
	f.Write(resp.Sound)
	_, existsErr := os.Stat("test_sounds/Tirion_Fordring_Play.mp3")
	assert.True(t, !os.IsNotExist(existsErr), "File couldnt be created")
	assert.NoError(t, os.Remove("test_sounds/Tirion_Fordring_Play.mp3"), "couldn't remove the file")
}

func TestErrors(t *testing.T) {
	api := hs.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHIzp1jsnTpw7tNCxEZIN")
	api.Debug = true

	//Type
	config := hs.NewCardsByType(hs.Weapon)
	config.Collectible = true
	resp, err := api.CardsByType(config)
	assert.Empty(t, resp, "Empty Response")
	assert.Error(t, err, "An error ocurred while testing Errors")

	//Race
	config1 := hs.NewCardsByRace(hs.Dragon)
	config1.Collectible = true
	resp1, err1 := api.CardsByRace(config1)
	assert.Empty(t, resp1, "Empty Response")
	assert.Error(t, err1, "An error ocurred while testing Errors")

	//Quality
	config2 := hs.NewCardsByQuality(hs.Legendary)
	config2.Collectible = true
	resp2, err2 := api.CardsByQuality(config2)
	assert.Empty(t, resp2, "Empty Response")
	assert.Error(t, err2, "An error ocurred while testing Errors")

	//Ifo
	config3 := hs.NewInfo()
	config3.Locale = hs.EsES
	resp3 := api.Info(config3)
	assert.NotEmpty(t, resp3, "Empty Response")

	//GetCard
	config4 := hs.NewGetCard("EX1_572")
	config4.Collectible = true
	resp4, err4 := api.GetCard(config4)
	assert.Empty(t, resp4, "Empty Response")
	assert.Error(t, err4, "An error ocurred while testing Errors")

	//AllCards
	config5 := hs.NewAllCards()
	config5.Collectible = true
	config5.Attack = 5
	config5.Durability = 5
	config5.Cost = 5
	resp5, err5 := api.AllCards(config5)
	assert.NotEmpty(t, resp5, "Empty Response")
	assert.Error(t, err5, "An error ocurred while testing Errors")

	//Cardbacks
	config6 := hs.NewCardBacks()
	config6.Locale = hs.EsES
	resp6, err6 := api.CardBacks(config6)
	assert.Empty(t, resp6, "Empty Response")
	assert.Error(t, err6, "An error ocurred while testing Errors")

	//CardSearch
	config7 := hs.NewCardSearch("tirion fordring")
	config7.Collectible = true
	resp7, err7 := api.Search(config7)
	assert.Empty(t, resp7, "Empty Response")
	assert.Error(t, err7, "An error ocurred while testing Errors")

	//By Set
	config8 := hs.NewCardsBySet(hs.TheGrandTournament)
	config8.Collectible = true
	resp8, err8 := api.CardsBySet(config8)
	assert.Empty(t, resp8, "Empty Response")
	assert.Error(t, err8, "An error ocurred while testing Errors")

	//By Class
	config9 := hs.NewCardsByClass(hs.Warlock)
	resp9, err9 := api.CardsByClass(config9)
	assert.Empty(t, resp9, "Empty Response")
	assert.Error(t, err9, "An error ocurred while testing Errors")

	//By Faction
	config10 := hs.NewCardsByFaction(hs.Horde)
	config10.Collectible = true
	resp10, err10 := api.CardsByFaction(config10)
	assert.Empty(t, resp10, "Empty Response")
	assert.Error(t, err10, "An error ocurred while testing Errors")

	//CardImg
	config11 := hs.NewCardImage("EX1383")
	config11.Gold = true
	_, err11 := api.CardImage(config11)
	assert.Error(t, err11, "An error ocurred while testing Errors")

	//CardImg
	config12 := hs.NewCardSound("EX1383", hs.ATTACK)
	config12.Extension = hs.MP3
	_, err12 := api.CardSound(config12)
	assert.Error(t, err12, "An error ocurred while testing Errors")
}
