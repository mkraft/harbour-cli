package mattermostjsonl

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mkraft/harbour"
)

type Writer struct{}

func (m *Writer) WriteUser(user *harbour.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
	return nil
}

func (m *Writer) WriteGroup(group *harbour.Group) error {
	data, err := json.Marshal(group)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data))
	return nil
}
