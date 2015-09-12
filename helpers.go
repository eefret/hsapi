package hsapi

//NewAllCards creates a new AllCardsConfig
//no mandatory parameters required default locale is enUS
func NewAllCards() AllCardsConfig {
	return AllCardsConfig{
	}
}

//NewCardBacks creates a new CardBacksConfig
//no mandatory parameters required default locale is enUS
func NewCardBacks() CardBacksConfig {
	return CardBacksConfig{
	}
}

//NewCardSearch creates a new CardSearchConfig
//name parameter is mandatory default locale is enUS
func NewCardSearch(name string) CardSearchConfig {
	return CardSearchConfig{
		Name: name,
	}
}

//NewCardsBySet creates a new CardsBySetConfig
//set parameter is mandatory default locale is enUS
func NewCardsBySet(set Set) CardsBySetConfig {
	return CardsBySetConfig{
		Set: set,
	}
}

//NewCardsByClass creates a new CardsByClassConfig
//set parameter is mandatory default locale is enUS
func NewCardsByClass(class Class) CardsByClassConfig {
	return CardsByClassConfig{
		Class: class,
	}
}

//NewCardsByFaction creates a new CardsByFactionConfig
//faction parameter is mandatory default locale is enUS
func NewCardsByFaction(faction Faction) CardsByFactionConfig {
	return CardsByFactionConfig{
		Faction: faction,
	}
}

//NewCardsByQuality creates a new CardsByQualityConfig
//quality parameter is mandatory default local is enUS
func NewCardsByQuality(quality Quality) CardsByQualityConfig {
	return CardsByQualityConfig{
		Quality: quality,
	}
}

//NewCardsByRace creates a new CardsByQualityConfig
//race parameter is mandatory default local is enUS
func NewCardsByRace(race Race) CardsByRaceConfig {
	return CardsByRaceConfig{
		Race: race,
	}
}

//NewCardsByType creates a new CardsByTypeConfig
//type parameter is mandatory default local is enUS
func NewCardsByType(typ Type) CardsByTypeConfig {
	return CardsByTypeConfig{
		Type: typ,
	}
}

//NewGetCard creates a new GetCardConfig
//name parameter is mandatory, should be exact name or ID default local is enUS
func NewGetCard(id string) GetCardConfig {
	return GetCardConfig{
		Name: id,
	}
}

//NewInfo creates a new InfoConfig
//no mandatory params default locale is enUS
func NewInfo() InfoConfig {
	return InfoConfig{}
}

//NewCardImage creates a new CardImageConfig
//CardID is mandatory default locale is enUS Gold is default false
func NewCardImage(cardID string) CardImageConfig {
	return CardImageConfig{
		CardID: cardID,
		Gold: false,
		Locale: EnUS,
	}
}

//NewCardSound creates a new CardSoundConfig
//CardID and SoundType are mandatory, default local is enUS and
//default extension is MP3
func NewCardSound(cardID string, soundType SoundType) CardSoundConfig {
	return CardSoundConfig{
		CardID: cardID,
		Type: soundType,
		Extension: MP3,
		Locale: EnUS,
	}
}