package main

import (
	"assignment3/router"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type Value struct {
	Status Status
}

type Status struct {
	Water int
	Wind  int
}

func main() {
	go generateFile()

	r := router.StartApp()
	r.Run(":8080")
}

func generateValue() Value {
	min := 1
	max := 100
	water := rand.Intn(max-min) + min
	wind := rand.Intn(max-min) + min
	data := Value{
		Status{
			Water: water,
			Wind:  wind,
		},
	}

	return data
}

func generateFile() {
	ticker := time.NewTicker(15 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				data := generateValue()

				file, _ := json.MarshalIndent(data, "", "")

				_ = ioutil.WriteFile("json_file.json", file, 0644)
			}
		}
	}()

	time.Sleep(15 * time.Minute)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
