package hsapi

//Locale is the type that defines the language for the endpoints
type Locale int

//Locales objects
const (
	DeDE Locale = iota + 1
	EnGB
	EnUS
	EsES
	EsMX
	FrFR
	ItIT
	KoKR
	PtBR
	RuRU
	ZhCN
	ZhTW
)

var locales = [...]string{
	"",
	"deDE",
	"enGB",
	"enUS",
	"esES",
	"esMX",
	"frFR",
	"itIT",
	"koKR",
	"plPL",
	"ptBR",
	"ruRU",
	"zhCN",
	"zhTW",
}

func (l Locale) String() string {
	return locales[l]
}

//Set is the type that defines cardsets in some of the endpoints
type Set int

//Set objects
const (
	Basic Set = iota + 1
	Classic
	Credits
	Naxxramas
	Debug
	GoblinsvsGnomes
	Missions
	Promotion
	Reward
	System
	BlackrockMountain
	HeroSkins
	TavernBrawl
	TheGrandTournament
)

var sets = [...]string{
	"",
	"Basic",
	"Classic",
	"Credits",
	"Naxxramas",
	"Debug",
	"Goblins vs Gnomes",
	"Missions",
	"Promotion",
	"Reward",
	"System",
	"Blackrock Mountain",
	"Hero Skins",
	"Tavern Brawl",
	"The Grand Tournament",
}

func (s Set) String() string {
	return sets[s]
}

//Class is the enum to define the classes
type Class int

//Class objects
const (
	Druid Class = iota + 1
	Hunter
	Mage
	Paladin
	Priest
	Rogue
	Shaman
	Warlock
	Warrior
	Dream
)

var classes = [...]string{
	"",
	"Druid",
	"Hunter",
	"Mage",
	"Paladin",
	"Priest",
	"Rogue",
	"Shaman",
	"Warlock",
	"Warrior",
	"Dream",
}

func (c Class) String() string {
	return classes[c]
}

//Faction is the enum for the factions
type Faction int

//Faction objects
const (
	Horde Faction = iota + 1
	Alliance
	Neutral
)

var factions = [...]string{
	"",
	"Horde",
	"Alliance",
	"Neutral",
}

func (f Faction) String() string {
	return factions[f]
}

//Quality is the enum for the qualities
type Quality int

//Quality objects
const (
	Free Quality = iota + 1
	Common
	Rare
	Epic
	Legendary
)

var qualities = [...]string{
	"",
	"Free",
	"Common",
	"Rare",
	"Epic",
	"Legendary",
}

func (q Quality) String() string {
	return qualities[q]
}

//Race is the enum for the race
type Race int

//Race objects
const (
	Demon Race = iota + 1
	Dragon
	Mech
	Murloc
	Beast
	Pirate
	Totem
)

var races = [...]string{
	"",
	"Demon",
	"Dragon",
	"Mech",
	"Murloc",
	"Beast",
	"Pirate",
	"Totem",
}

func (r Race) String() string {
	return races[r]
}

//Type is the enum for the type
type Type int

//Type objects
const (
	Hero Type = iota + 1
	Minion
	Spell
	Enchantment
	Weapon
	HeroPower
)

var types = [...]string{
	"",
	"Hero",
	"Minion",
	"Spell",
	"Enchantment",
	"Weapon",
	"Hero Power",
}

func (t Type) String() string {
	return types[t]
}

//FileExtension is the enum for the extensions
type FileExtension int

//FileExtension objects
const (
	JPG FileExtension = iota + 1
	JPEG
	GIF
	PNG
	MP3
	OGG
)

var extensions = [...]string{
	"",
	"jpg",
	"jpeg",
	"gif",
	"png",
	"mp3",
	"ogg",
}

func (f FileExtension) String() string {
	return extensions[f]
}

//SoundType is the enum for the extensions
type SoundType int

//SoundType objects
const (
	PLAY SoundType = iota + 1
	ATTACK
	DEATH
	TRIGGER
)

var soundtypes = [...]string{
	"",
	"Play",
	"Attack",
	"Death",
	"Trigger",
}

func (s SoundType) String() string {
	return soundtypes[s]
}
