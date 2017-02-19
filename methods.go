package hsapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eefret/hsapi/sounds"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

//All constants
const (
	BaseURL       = "https://omgvamp-hearthstone-v1.p.mashape.com/"
	INFO          = "info"
	CARDS         = "cards"
	GetCARD       = "cards/%v"
	CardBACK      = "cardbacks"
	CardSEARCH    = "cards/search/%v"
	CardBySET     = "cards/sets/%v"
	CardByCLASS   = "cards/classes/%v"
	CardByRACE    = "cards/races/%v"
	CardByQUALITY = "cards/qualities/%v"
	CardByTYPE    = "cards/types/%v"
	CardByFACTION = "cards/factions/%v"
	SoundsURL     = "http://wowimg.zamimg.com/hearthhead/sounds%v/VO_%v_%v_%v.%v"
	//First %v is locale with a / before and in low caps if exists if not ''
	//Second %v is CardID
	//Third %v is SoundType
	//Fourth %v is Index (01-05)
	//Last %v is Extension (ogg, mp3)
)

//ErrMultipleCards is thrown when a single card is required but the query returns several
var ErrMultipleCards = errors.New("Multiple Cards obtained in a single card context.")

//==========================================GENERAL METHODS

func (hs *HsAPI) makeRequest(endpoint string, keyparam string, params url.Values) ([]byte, error) {
	url, err := url.Parse(BaseURL)
	if keyparam != "" {
		url.Path += fmt.Sprintf(endpoint, keyparam)
	} else {
		url.Path += endpoint
	}
	url.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header.Set("X-Mashape-Key", hs.token)
	req.Header.Set("Accept", "application/json")
	resp, err := hs.client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if hs.Debug {
		log.Println(url)
		log.Println(string(body))
	}

	if resp.StatusCode == http.StatusOK {
		return body, nil
	}
	return nil, err
}

func (hs *HsAPI) parseValues(config interface{}, exclude int) url.Values {
	u := url.Values{}
	t := reflect.TypeOf(config)
	v := reflect.ValueOf(config)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if i == exclude { //if excluded pass
			continue
		}

		//if its a struct add all his fields and continue
		if t.Field(i).Type.Kind() == reflect.Struct {
			for k, v := range hs.parseValues(field.Interface(), -1) {
				u.Add(k, v[0])
			}
			continue
		}

		//If its a boolean parse to int
		if t.Field(i).Type.Kind() == reflect.Bool && field.Interface() == true {
			u.Add(strings.ToLower(t.Field(i).Name), "1")
			continue
		}

		//if its not Zero Value add
		if field.Interface() != reflect.Zero(field.Type()).Interface() {
			u.Add(strings.ToLower(t.Field(i).Name), fmt.Sprintf("%v", field.Interface()))
		}
		if hs.Debug {
			log.Println(field.Interface())
			log.Println(reflect.Zero(field.Type()).Interface())
			log.Println(t.Field(i).Name)
		}
	}
	return u
}

func (hs *HsAPI) simpleGetBodyRequest(URL string) (*http.Response, error) {
	resp, err := hs.client.Get(URL)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, err
	}
	return resp, nil

}

//==========================================API METHODS

//Info obtains info about the API
func (hs *HsAPI) Info(config InfoConfig) InfoResponse {
	v := hs.parseValues(config, -1)
	b, _ := hs.makeRequest(INFO, "", v)
	var response InfoResponse
	json.Unmarshal(b, &response)
	return response
}

//AllCards obtain all cards in Hearthstone
func (hs *HsAPI) AllCards(config AllCardsConfig) (AllCardResponse, error) {
	v := hs.parseValues(config, -1)
	b, err := hs.makeRequest(CARDS, "", v)
	if err == nil {
		var response AllCardResponse
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return AllCardResponse{}, err
}

//CardBacks obtain all cards backs in Hearthstone
func (hs *HsAPI) CardBacks(config CardBacksConfig) ([]CardBack, error) {
	v := hs.parseValues(config, -1)
	b, err := hs.makeRequest(CardBACK, "", v)
	if err == nil {
		var response []CardBack
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []CardBack{}, err
}

//GetCard obtain a single card
//Exact Name or ID is mandatory, please use ID for more accurate
func (hs *HsAPI) GetCard(config GetCardConfig) ([]Card, error) {
	v := hs.parseValues(config, 0)
	b, err := hs.makeRequest(GetCARD, config.Name, v)
	if err == nil {
		var response []Card
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []Card{}, err
}

//Search search for a card given a name query
//the Name query is Mandatory
func (hs *HsAPI) Search(config CardSearchConfig) ([]Card, error) {
	v := hs.parseValues(config, 0)
	b, err := hs.makeRequest(CardSEARCH, config.Name, v)
	if err == nil {
		var response []Card
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []Card{}, err
}

//CardsBySet obtains all cards of the given Set
//Set param is mandatory
func (hs *HsAPI) CardsBySet(config CardsBySetConfig) ([]Card, error) {
	v := hs.parseValues(config, 0)
	b, err := hs.makeRequest(CardBySET, config.Set.String(), v)
	if err == nil {
		var response []Card
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []Card{}, err
}

//CardsByClass obtains all cards by class
//Class param is mandatory
func (hs *HsAPI) CardsByClass(config CardsByClassConfig) ([]Card, error) {
	v := hs.parseValues(config, 0)
	b, err := hs.makeRequest(CardByCLASS, config.Class.String(), v)
	if err == nil {
		var response []Card
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []Card{}, err
}

//CardsByFaction obtains all cards by faction
//Faction param is mandatory
func (hs *HsAPI) CardsByFaction(config CardsByFactionConfig) ([]Card, error) {
	v := hs.parseValues(config, 0)
	b, err := hs.makeRequest(CardByFACTION, config.Faction.String(), v)
	if err == nil {
		var response []Card
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []Card{}, err
}

//CardsByQuality obtains all cards by quality
//Quality param is mandatory
func (hs *HsAPI) CardsByQuality(config CardsByQualityConfig) ([]Card, error) {
	v := hs.parseValues(config, 0)
	b, err := hs.makeRequest(CardByQUALITY, config.Quality.String(), v)
	if err == nil {
		var response []Card
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []Card{}, err
}

//CardsByRace obtains all cards by race
//Race param is mandatory
func (hs *HsAPI) CardsByRace(config CardsByRaceConfig) ([]Card, error) {
	v := hs.parseValues(config, 0)
	b, err := hs.makeRequest(CardByRACE, config.Race.String(), v)
	if err == nil {
		var response []Card
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []Card{}, err
}

//CardsByType obtains all cards by type
//Type param is mandatory
func (hs *HsAPI) CardsByType(config CardsByTypeConfig) ([]Card, error) {
	v := hs.parseValues(config, 0)
	b, err := hs.makeRequest(CardByTYPE, config.Type.String(), v)
	if err == nil {
		var response []Card
		err = json.Unmarshal(b, &response)
		if err == nil {
			return response, nil
		}
	}
	return []Card{}, err
}

//CardImage returns an Image given a Card
func (hs *HsAPI) CardImage(config CardImageConfig) (CardImageResponse, error) {
	c := NewGetCard(config.CardID)
	c.Locale = config.Locale

	cards, err := hs.GetCard(c)
	if err != nil {
		return CardImageResponse{}, err
	}
	if len(cards) > 1 {
		return CardImageResponse{}, ErrMultipleCards
	}
	card := cards[0]

	var URL string
	var extension FileExtension
	if config.Gold {
		URL = card.ImgGold
		extension = GIF
	} else {
		URL = card.Img
		extension = PNG
	}

	resp, err := hs.simpleGetBodyRequest(URL)
	defer resp.Body.Close()
	image, err := ioutil.ReadAll(resp.Body)
	length, err := strconv.ParseInt(resp.Header.Get("Content-Length"),
		10, 64)

	if err != nil {
		return CardImageResponse{}, err
	}

	return CardImageResponse{
		Image:         image,
		ContentLength: length,
		Extension:     extension,
	}, nil
}

func (hs *HsAPI) CardSound(config CardSoundConfig) (CardSoundResponse, error) {
	soundsMap, err := sounds.New()
	if err != nil {
		return CardSoundResponse{}, err
	}

	var URL string
	var soundsArray []string
	var ok bool
	//first get url from the map
	switch config.Locale {
	case EnUS:
		soundsArray, ok = soundsMap.En[config.CardID][strings.ToLower(config.Type.String())]
	case EnGB:
		soundsArray, ok = soundsMap.En[config.CardID][strings.ToLower(config.Type.String())]
	case DeDE:
		soundsArray, ok = soundsMap.De[config.CardID][strings.ToLower(config.Type.String())]
	case EsES:
		soundsArray, ok = soundsMap.Es[config.CardID][strings.ToLower(config.Type.String())]
	case EsMX:
		soundsArray, ok = soundsMap.Es[config.CardID][strings.ToLower(config.Type.String())]
	case FrFR:
		soundsArray, ok = soundsMap.Fr[config.CardID][strings.ToLower(config.Type.String())]
	case ItIT:
		soundsArray, ok = soundsMap.It[config.CardID][strings.ToLower(config.Type.String())]
	case PtBR:
		soundsArray, ok = soundsMap.Pt[config.CardID][strings.ToLower(config.Type.String())]
	case RuRU:
		soundsArray, ok = soundsMap.Ru[config.CardID][strings.ToLower(config.Type.String())]
	}

	if !ok {
		return CardSoundResponse{}, errors.New("card not found")
	}

	if config.Extension == OGG {
		URL = soundsArray[0]
	} else if config.Extension == MP3 {
		URL = soundsArray[1]
	}

	resp, err := hs.simpleGetBodyRequest(URL)

	defer resp.Body.Close()
	sound, err := ioutil.ReadAll(resp.Body)
	length, err := strconv.ParseInt(resp.Header.Get("Content-Length"),
		10, 64)

	if err != nil {
		fmt.Println(err)
		return CardSoundResponse{}, err
	}

	return CardSoundResponse{
		Sound:         sound,
		ContentLength: length,
		Extension:     config.Extension,
		Type:          config.Type,
	}, nil
}
