package main

import (
	"embed"
	"github.com/BenSlabbert/gist-app/cmd"
	"log"
)

// GitCommit is set during compilation
var GitCommit string

//go:embed web/public/*
var website embed.FS

func main() {
	log.Printf("GitCommit: %s", GitCommit)
	// files in embedded are available as web/public/
	dir, err := website.ReadDir("web/public")
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range dir {
		log.Println(entry.Name())
	}

	cmd.Execute()
}
