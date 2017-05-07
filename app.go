package main

import (
	"encoding/json"
	"fmt"
	"log"

	firego "gopkg.in/zabawaba99/firego.v1"
)

func main() {
	f := firego.New("https://futsalicious-f7b04.firebaseio.com", nil)

	// notifications := make(chan firego.Event)
	// if err := f.Watch(notifications); err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.StopWatching()
	// for event := range notifications {
	// 	fmt.Printf("Event %#v\n", event)
	// }
	// fmt.Printf("Notifications have stopped")

	// v := "bar"
	// pushedFirego, err := f.Push(v)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var bar string
	// if err := pushedFirego.Value(&bar); err != nil {
	// 	log.Fatal(err)
	// }

	// prints "https://my-firebase-app.firebaseIO.com/-JgvLHXszP4xS0AUN-nI: bar"
	// fmt.Printf("%s: %s\n", pushedFirego, bar)

	usersRef, err := f.Ref("messages")
	if err != nil {
		log.Println(err)
	}

	v := map[string]interface{}{
		"id":        1234,
		"message":   "lalala",
		"create_by": 321,
	}
	if _, err := usersRef.Push(v); err != nil {
		log.Println(err)
	}

	if err := usersRef.Value(&v); err != nil {
		log.Println(err)
	}
	b, err := json.MarshalIndent(v, "", "  ")
	fmt.Print(string(b))

	var arr []interface{}
	for _, val := range v {
		arr = append(arr, val)
	}
	// b, err := json.MarshalIndent(arr, "", "  ")
	// fmt.Print(b)

	notifications := make(chan firego.Event)
	if err := f.Watch(notifications); err != nil {
		log.Fatal(err)
	}

	defer f.StopWatching()
	for event := range notifications {
		fmt.Printf("Event %#v\n", event)
	}
	fmt.Printf("Notifications have stopped")

}
