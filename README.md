# Hearthstone API [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/eefret/hsapi)
Golang Bindings for the hearthstoneapi.com API

##Summary
Unoficial [hearthstoneapi](http://hearthstoneapi.com/) Go Client Library

##Instalation
```
go get github.com/eefret/hsapi
```

##Documentation
For the general api docs please see the [Official Mashape Docs](https://market.mashape.com/omgvamp/hearthstone). For details on all the functionality on this library, see the [GoDoc](https://godoc.org/github.com/eefret/hsapi) documentation.

##Usage Example

```
import(
"github.com/eefret/hsapi"
"fmt"
)
  //Search Example
  api := hsapi.NewHsAPI("tntkXJyM7EmshBgQYsXtCHHEX8Izp1uHrN1jsnTpw7tNCxEZIN") //get yours in mashape
	api.Debug = true //Will print debug outputs
	config := hsapi.NewCardSearch("tirion fordring") //or any incomplete search
	config.Collectible = true //To narrow the search a little more
	resp, err := api.Search(config) 
	fmt.Println(resp[0].Name) //Tirion Fordring
	
	c := hsapi.NewCardImage(resp[0].CardID)
	c.Gold = true
	resp, err := api.CardImage(c)
	f, err := os.Create("Tirion_Fordring.gif")
	defer f.Close()
	f.Write(resp.Image) //Writing []byte into the file
```
