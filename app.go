package main

import (
	"encoding/json"
	"fmt"
	"log"

	firego "gopkg.in/zabawaba99/firego.v1"
)

func main() {
	f := firego.New("https://futsalicious-f7b04.firebaseio.com", nil)

	usersRef, err := f.Ref("messages")
	if err != nil {
		log.Println(err)
	}

	v := map[string]interface{}{}
	if err := usersRef.Value(&v); err != nil {
		log.Println(err)
	}
	// b, err := json.MarshalIndent(v, "", "  ")
	// fmt.Print(string(b))

	var arr []interface{}
	for _, val := range v {
		arr = append(arr, val)
	}
	b, err := json.MarshalIndent(arr, "", "  ")
	fmt.Print(string(b))

	notifications := make(chan firego.Event)
	if err := f.Watch(notifications); err != nil {
		log.Fatal(err)
	}

	defer f.StopWatching()
	for event := range notifications {
		b, _ := json.MarshalIndent(event, "", "  ")
		fmt.Print(string(b))
	}
	fmt.Printf("Notifications have stopped")

	v = map[string]interface{}{
		"id":        1234,
		"message":   "jegejekgejek",
		"create_by": 321,
	}
	if _, err := usersRef.Push(v); err != nil {
		log.Println(err)
	}

}
