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

	// based on some input we select a concreate reader and writer, for example
	// go run . -from slackexport -to mattermostjsonl
	importer = &slackexport.Reader{}
	exporter = &mattermostjsonl.Writer{}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for result := range importer.Users() {
			if result.Err != nil {
				log.Println(result.Err)
				continue
			}
			err := exporter.WriteUser(result.Val)
			if err != nil {
				log.Println(err)
			}
		}
		wg.Done()
	}()

	go func() {
		for result := range importer.Groups() {
			if result.Err != nil {
				log.Println(result.Err)
			}
			err := exporter.WriteGroup(result.Val)
			if err != nil {
				log.Println(err)
			}
		}
		wg.Done()
	}()

	// Add more types as per the reader/writer interfaces

	wg.Wait()
}
