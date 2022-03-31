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
		done := make(chan bool)
		users, errors := importer.Users(done)
		for {
			select {
			case user := <-users:
				err := exporter.WriteUser(user)
				if err != nil {
					log.Println(err)
				}
			case err := <-errors:
				log.Println(err)
			case <-done:
				log.Println("users done")
				wg.Done()
				return
			}
		}
	}()

	go func() {
		done := make(chan bool)
		groups, errors := importer.Groups(done)
		for {
			select {
			case group := <-groups:
				err := exporter.WriteGroup(group)
				if err != nil {
					log.Println(err)
				}
			case err := <-errors:
				log.Println(err)
			case <-done:
				log.Println("groups done")
				wg.Done()
				return
			}
		}
	}()

	// Add more types as per the reader/writer interfaces

	wg.Wait()
}
