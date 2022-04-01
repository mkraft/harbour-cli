package mattermostjsonl

import (
	"encoding/json"
	"fmt"
	"log"

	"example.com/mattermost"
)

type Writer struct{}

func (m *Writer) WriteUser(user *mattermost.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
	return nil
}

func (m *Writer) WriteGroup(group *mattermost.Group) error {
	data, err := json.Marshal(group)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
	return nil
}
