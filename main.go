package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Language tells us the code for the language
type Language struct {
	Name string `json:"name"`
}

// NameInLanguage provides a name in a particular language
type NameInLanguage struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

// FlavorText provides a struct for language-specific pokedex entries
type FlavorText struct {
	FlavorText string   `json:"flavor_text"`
	Language   Language `json:"language"`
}

// Species details a pokemon
type Species struct {
	ID                int64            `json:"id"`
	Names             []NameInLanguage `json:"names"`
	FlavorTextEntries []FlavorText     `json:"flavor_text_entries"`
}

// GetEnglishName returns the English name for the species
func (s *Species) GetEnglishName() string {
	for _, v := range s.Names {
		if v.Language.Name == "en" {
			return v.Name
		}
	}
	return "not found"
}

// GetEnglishFlavorText returns the pokedex details for the species
func (s *Species) GetEnglishFlavorText() string {
	for _, v := range s.FlavorTextEntries {
		if v.Language.Name == "en" {
			return v.FlavorText
		}
	}
	return "not found"
}

func main() {
	numPtr := flag.Int("number", 1, "The number to look up")
	flag.Parse()
	fmt.Println("Looking up #", *numPtr)

	url := "http://pokeapi.co/api/v2/pokemon-species/" + fmt.Sprintf("%d", *numPtr)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	res := Species{}
	json.Unmarshal(body, &res)
	fmt.Println(res.GetEnglishName())
	fmt.Println(res.GetEnglishFlavorText())
}
