package main

import (
	"embed"
	"html/template"
	"log"
)

//go:embed template/*
var templatesFS embed.FS

var tmpl *template.Template

func initTemplates() {
	var err error
	tmpl, err = template.ParseFS(templatesFS, "template/*.html")
	if err != nil {
		log.Fatalf("failed to parse templates: %v", err)
	}
}
