package main

import (
	"log"
	"sync"

	"github.com/mkraft/mattermost"
	"github.com/mkraft/mattermostjsonl"
	"github.com/mkraft/slackexport"
)

func main() {
	var importer mattermost.Reader
	var exporter mattermost.Writer

	// based on some input we select a concreate reader and writer
	importer = &slackexport.Reader{}
	exporter = &mattermostjsonl.Writer{}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for user := range importer.Users() {
			err := exporter.WriteUser(user)
			if err != nil {
				log.Println(err)
			}
		}
		wg.Done()
	}()

	go func() {
		for group := range importer.Groups() {
			err := exporter.WriteGroup(group)
			if err != nil {
				log.Println(err)
			}
		}
		wg.Done()
	}()

	wg.Wait()
}
