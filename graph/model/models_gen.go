// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Album struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Artist struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Image  string   `json:"image"`
	Albums []*Album `json:"albums"`
}