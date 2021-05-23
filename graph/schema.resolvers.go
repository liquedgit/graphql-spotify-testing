package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"

	"github.com/abenbyy/spotify-graphql/graph/generated"
	"github.com/abenbyy/spotify-graphql/graph/model"
)

func (r *queryResolver) Artist(ctx context.Context, name string) (*model.Artist, error) {
	url := "https://api.spotify.com/v1/search?q=" + url2.QueryEscape(name) + "&type=artist&limit=1&offset=0"

	bearer := "Bearer " + CLIENT_TOKEN

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result map[string]interface{}
	json.Unmarshal(body, &result)
	artist := result["artists"].(map[string]interface{})["items"].([]interface{})[0].(map[string]interface{})
	image := artist["images"].([]interface{})[0].(map[string]interface{})["url"]

	url2 := "https://api.spotify.com/v1/artists/" + artist["id"].(string) + "/albums"
	req, err = http.NewRequest("GET", url2, nil)
	req.Header.Add("Authorization", bearer)

	resp, err = client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var result2 map[string]interface{}
	json.Unmarshal(body, &result2)

	albumsJSON := result2["items"].([]interface{})
	var albums []*model.Album

	for i:=0 ; i<len(albumsJSON);i++{
		album := albumsJSON[i].(map[string]interface{})
		albumIMG := album["images"].([]interface{})[0].(map[string]interface{})["url"]

		albums = append(albums, &model.Album{
			ID:    album["id"].(string),
			Name:  album["name"].(string),
			Image: albumIMG.(string),
		})
	}

	res := &model.Artist{
		ID:    artist["id"].(string),
		Name:  artist["name"].(string),
		Image: image.(string),
		Albums: albums,
	}

	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var CLIENT_TOKEN = "BQDgWnrPNf_ybG6HswReLuEsBqLQETR0-R0-Rn0mQauyi2O3MgeEwjNGuRKECTVHx4SvfqwCB9HJDXOgp0E"

type mutationResolver struct{ *Resolver }
