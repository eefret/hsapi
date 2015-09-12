package sounds

import (
	"encoding/json"
	"io/ioutil"
)

//CardSounds is the struct defined in the json
// first map is "ID" then "soundtype" then the string array
// contains both sound format ogg in pos 0, and mp3 in pos 1
type CardSounds struct {
	En map[string]map[string][]string `json:"en"`
	De map[string]map[string][]string `json:"de"`
	Es map[string]map[string][]string `json:"es"`
	Fr map[string]map[string][]string `json:"fr"`
	It map[string]map[string][]string `json:"it"`
	Pt map[string]map[string][]string `json:"pt"`
	Ru map[string]map[string][]string `json:"ru"`
}


func New() (CardSounds, error) {
	content, err := ioutil.ReadFile("sounds/card_sounds.json")
	if err != nil {
		return CardSounds{}, err
	}

	var cardSounds CardSounds
	err = json.Unmarshal(content, &cardSounds)
	if err != nil {
		return CardSounds{}, err
	}

	return cardSounds, nil
}