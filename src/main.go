package main

import (
	"embed"
	_ "embed"
	"fmt"
	"github.com/johngenerator/src/types"
	"html/template"
	"log"
	"net/http"
	"strconv"
)
const characterHtml = "html/templates/character.html"
var n types.Names
var s types.Skills
var o types.Obsessions

//go:embed html/index.html
var indexHTML []byte

//go:embed html/templates/*
var templates embed.FS

//go:embed assets/names.json
var namesJson []byte

//go:embed assets/obsessions.json
var obsJson []byte

//go:embed assets/skills.json
var skillsJson []byte


func characterHandler(w http.ResponseWriter, r *http.Request) {
	var skillCount int
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error creating character", 500)
	}
	if r.Form["skills"] != nil {
		skillCount, _ = strconv.Atoi(r.Form["skills"][0])
	}
	c := types.CreateNewCharacter(n,s,o, skillCount)
	parsedTemplate, err := template.ParseFS(templates, characterHtml)
	if err != nil {
		log.Fatal("unable to parse ", err)
	}
	err = parsedTemplate.Execute(w, c)
	if err != nil {
		log.Println("error creating john ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, string(indexHTML))
}


func main() {
	n = types.LoadNames(namesJson)
	s = types.LoadSkills(skillsJson)
	o = types.LoadObsessions(obsJson)
	http.HandleFunc("/character", characterHandler)
	http.HandleFunc("/", index)
	// start server
	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}
