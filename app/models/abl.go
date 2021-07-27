package models

import (
	"github.com/jinzhu/gorm"
)

type Abl []struct {
	Index    int      `json:"index"`
	Name     string   `json:"name"`
	FullName string   `json:"full_name"`
	Desc     []string `json:"desc"`
	Skills   []struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	} `json:"skills"`
	URL string `json:"url"`
}

type Abilities struct {
	Index              int `json:"index"`
	Irratibility       int `json:"irr"`
	Bribability        int `json:"bri"`
	RelativeFriendship int `json:"rel"`
	Naivety            int `json:"nai"`
}
type Secret struct {
	Secret   string `json:"secret"`
	Severity int    `json:"severity"`
	Teir     int    `json:"teir"`
}

type romevar struct {
	Name   string `json:"name"`
	Affect int    `json:"affect"`
	Care   int    `json:"care"`
}
type rome struct {
	Name   string `json:"name"`
	Affect int    `json:"affect"`
	Care   int    `json:"care"`
}

type Character struct {
	Name          string    `json:"name"`
	Drinking      bool      `json:"drinking"`
	Smoking       bool      `json:"smoking"`
	Religon       string    `json:"religon"`
	SecretOne     []Secret  `json:"secret_one"`
	SecretTwo     []Secret  `json:"secret_two"`
	SecretThree   []Secret  `json:"secret_three"`
	RomanceOne    rome      `json:"rome_one"`
	RomanceTwo    rome      `json:"rome_two"`
	RomanceVarOne romevar   `json:"rome_var_one"`
	RomanceVarTwo romevar   `json:"rome_var_two"`
	Goals         []goal    `json:"goals"`
	TraitOne      string    `json:"trait_one"`
	TraitTwo      string    `json:"trait_two"`
	Ability       Abilities `json:"abilities"`
}
