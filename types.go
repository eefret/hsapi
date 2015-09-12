package hsapi


//==========================================RESPONSE STRUCTS

//AllCardResponse is a struct for allcards method
type AllCardResponse struct {
	Basic                []Card `json:"Basic"`
	BlackrockMountain   []Card `json:"Blackrock Mountain"`
	Classic              []Card `json:"Classic"`
	Credits              []Card `json:"Credits"`
	Debug                []Card `json:"Debug"`
	GoblinsvsGnomes    []Card `json:"Goblins vs Gnomes"`
	HeroSkins           []Card `json:"Hero Skins"`
	Missions             []Card `json:"Missions"`
	Naxxramas            []Card `json:"Naxxramas"`
	Promotion            []Card `json:"Promotion"`
	Reward               []Card `json:"Reward"`
	System               []Card `json:"System"`
	TavernBrawl         []Card `json:"Tavern Brawl"`
	TheGrandTournament []Card `json:"The Grand Tournament"`
}


//CardsResponse is an array of cards which is returned by almost every endpoint

//Card is the main struct for cards
type Card struct {
	Artist      string `json:"artist"`
	Attack      int    `json:"attack"`
	CardID      string `json:"cardId"`
	CardSet     string `json:"cardSet"`
	Collectible bool   `json:"collectible"`
	Cost        int    `json:"cost"`
	Elite       bool   `json:"elite"`
	Faction     string `json:"faction"`
	Flavor      string `json:"flavor"`
	Health      int    `json:"health"`
	Img         string `json:"img"`
	PlayerClass string `json:"playerClass"`
	ImgGold     string `json:"imgGold"`
	Locale      string `json:"locale"`
	Mechanics   []struct {
		Name string `json:"name"`
	} `json:"mechanics"`
	Name   string `json:"name"`
	Rarity string `json:"rarity"`
	Text   string `json:"text"`
	Type   string `json:"type"`
}

//InfoResponse is the response returned by the info endpoint
type InfoResponse struct {
	Classes   []string `json:"classes"`
	Factions  []string `json:"factions"`
	Locales   []string `json:"locales"`
	Patch     string   `json:"patch"`
	Qualities []string `json:"qualities"`
	Races     []string `json:"races"`
	Sets      []string `json:"sets"`
	Types     []string `json:"types"`
}

//CardBack is the response returned by the cardbacks endpoint
type CardBack struct {
	CardBackID   string `json:"cardBackId"`
	Description  string `json:"description"`
	Enabled      bool   `json:"enabled"`
	HowToGet     string `json:"howToGet"`
	Img          string `json:"img"`
	ImgAnimated  string `json:"imgAnimated"`
	Locale       string `json:"locale"`
	Name         string `json:"name"`
	SortCategory string `json:"sortCategory"`
	SortOrder    string `json:"sortOrder"`
	Source       string `json:"source"`
}

//CardImageResponse is the response returned by the CardImage endpoint
type CardImageResponse struct {
	Image []byte
	ContentLength int64
	Extension FileExtension
}

//CardSoundResponse is the response returned by the CardSound endpoint
type CardSoundResponse struct {
	Sound []byte
	ContentLength int64
	Extension FileExtension
	Type SoundType
}

//==========================================REQUEST STRUCTS

//AllCardsConfig configuration of the AllCards method
type AllCardsConfig struct {
	Attack int
	Collectible bool
	Cost int
	Durability int
	Health int
	Locale Locale
}

//CardBacksConfig configuration of the CardBacks method
type CardBacksConfig struct {
	Locale Locale
}

//CardSearchConfig configuration of the CardSearch method
type CardSearchConfig struct {
	Name string
	Collectible bool
	Locale Locale
}

//BaseConfig is the base for most of the configuration types
type BaseConfig struct{
	Attack int
	Collectible bool
	Cost int
	Durability int
	Health int
	Locale Locale
}

//CardsBySetConfig configuration of the CardsBySet method
type CardsBySetConfig struct {
	Set Set
	BaseConfig
}

//CardsByClassConfig configuration of the CardsByClassConfig method
type CardsByClassConfig struct {
	Class Class
	BaseConfig
}

//CardsByFactionConfig configuration of the CardsByFaction method
type CardsByFactionConfig struct {
	Faction Faction
	BaseConfig
}

//CardsByQualityConfig configuration of the CardsByQuality method
type CardsByQualityConfig struct {
	Quality Quality
	BaseConfig
}

//CardsByRaceConfig configuration of the CardsByRace method
type CardsByRaceConfig struct {
	Race Race
	BaseConfig
}

//CardsByTypeConfig configuration of the CardsByType method
type CardsByTypeConfig struct {
	Type Type
	BaseConfig
}

//GetCardConfig configuration of the GetCard method
type GetCardConfig struct {
	Name string
	Collectible bool
	Locale Locale
}

//InfoConfig configuration of the Info method
type InfoConfig struct {
	Locale Locale
}

//CardImageConfig configuration of the CardImage method
type CardImageConfig struct {
	CardID string
	Gold bool
	Locale Locale
}

//CardSoundConfig configudation of the GetCardSound method
type CardSoundConfig struct {
	CardID string
	Type SoundType
	Locale Locale
	Extension FileExtension
}