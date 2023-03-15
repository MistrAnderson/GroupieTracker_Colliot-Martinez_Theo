package struc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type Character struct {
	Error           bool
	Code            int    `json:"code"`
	AttributionText string `json:"attributionText"`
	Data            struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Total   int `json:"total"`
		Count   int `json:"count"`
		Results []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Thumbnail   struct {
				Path      string `json:"path"`
				Extension string `json:"extension"`
			} `json:"thumbnail"`
			ResourceURI string `json:"resourceURI"`
			Comics      struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"comics"`
			Series struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"series"`
			Stories struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"stories"`
			Events struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"events"`
			Urls []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"urls"`
		} `json:"results"`
	} `json:"data"`
}

func FetchCharacter(name string) Character {

	var newCharacter Character

	// Remplace les espace par des '%20' pour l'URL
	space := regexp.MustCompile(" +")
	correctName := space.ReplaceAllString(name, "%20")

	resp, err := http.Get("https://gateway.marvel.com/v1/public/characters?nameStartsWith=" + correctName + "&limit=60&ts=2&apikey=f204cdf734b24a3e74364d4161c65516&hash=111cb7aebcb27fb71c4d8a6703b833a1")
	if err != nil {
		fmt.Println("Une erreur est survenue pendant la requete")
		newCharacter.Error = true
		return newCharacter
	}
	defer resp.Body.Close()

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&newCharacter); err != nil {
		fmt.Println("Une erreur est survenue pendant le decodage")
		newCharacter.Error = true
		return newCharacter
	}

	return newCharacter
}
