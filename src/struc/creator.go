package struc

import (
	"encoding/json"
	"log"
	"net/http"
)

type Creator struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Total   int `json:"total"`
		Count   int `json:"count"`
		Results []struct {
			ID         int    `json:"id"`
			FirstName  string `json:"firstName"`
			MiddleName string `json:"middleName"`
			LastName   string `json:"lastName"`
			Suffix     string `json:"suffix"`
			FullName   string `json:"fullName"`
			Modified   string `json:"modified"`
			Thumbnail  struct {
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

func FetchCreator(lastName string) Creator {
	resp, err := http.Get("https://gateway.marvel.com:443/v1/public/creators?lastName=" + lastName + "&ts=2&apikey=f204cdf734b24a3e74364d4161c65516&hash=111cb7aebcb27fb71c4d8a6703b833a1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//Create a variable of the same type as our model
	var cResp Creator

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("ooopsss! an error occurred, please try again")
	}

	return cResp
}
