package main

import (
	"log"

	"github.com/mkraft/harbour"
	"github.com/mkraft/mattermostjsonl"
	"github.com/mkraft/slackexport"
)

func main() {
	var importer harbour.Reader
	var exporter harbour.Writer

	// based on some input we select a concreate reader and writer
	importer = &slackexport.Reader{}
	exporter = &mattermostjsonl.Writer{}

	userIterator, err := importer.UserIterator()
	if err != nil {
		log.Fatalln(err)
	}

	for userIterator.HasNext() {
		user, err := userIterator.Next()
		if err != nil {
			log.Println(err)
		}
		err = exporter.WriteUser(user)
		if err != nil {
			log.Println(err)
		}
	}

	groupIterator, err := importer.GroupIterator()
	if err != nil {
		log.Println(err)
	}

	for groupIterator.HasNext() {
		group, err := groupIterator.Next()
		if err != nil {
			log.Println(err)
		}
		err = exporter.WriteGroup(group)
		if err != nil {
			log.Println(err)
		}
	}
}
