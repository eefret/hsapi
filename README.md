# Hearthstone API [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/eefret/hsapi) [![Build Status](https://travis-ci.org/eefret/hsapi.svg?branch=master)](https://travis-ci.org/eefret/hsapi)
Golang Bindings for the hearthstoneapi.com API

## Summary
Unoficial [hearthstoneapi](http://hearthstoneapi.com/) Go Client Library

## Instalation

```shell
go get github.com/eefret/hsapi
```

## Documentation
For the general api docs please see the [Official Mashape Docs](https://market.mashape.com/omgvamp/hearthstone). For details on all the functionality on this library, see the [GoDoc](https://godoc.org/github.com/eefret/hsapi) documentation.

## Usage Example

```go
import(
    "github.com/eefret/hsapi"
    "fmt"
)

  // Search Example
  func main() {
  	api := hsapi.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN") // Get yours in mashape
	api.Debug = true // Will print debug outputs
	config := hsapi.NewCardSearch("tirion fordring") // or any incomplete search
	config.Collectible = true // To narrow the search a little more
	resp, _ := api.Search(config)
	fmt.Println(resp[0].Name) // Tirion Fordring
	
	c := hsapi.NewCardImage(resp[0].CardID)
	c.Gold = true
	resp, _ := api.CardImage(c)
	f, _ := os.Create("Tirion_Fordring.gif")
	defer f.Close()
	f.Write(resp.Image) // Writing []byte into the file

	//Now featuring card sounds
	c2 := hs.NewCardSound("EX1_383", hs.PLAY)
    c2.Locale = hs.EsES
    c2.Extension = hs.MP3
    resp, _ := api.CardSound(c2)
	f, _ := os.Create("Tirion_Fordring_Play.mp3")
	defer f.Close()
	f.Write(resp.Sound) // Writing []byte into the file
  }
```
