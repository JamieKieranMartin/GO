package main

import (
	"net/http"
	"encoding/json"
	"bytes"
	"log"
	"io/ioutil"
	"math/rand"
	"time"
	"fmt"
	"strings"
)

type GenericAlarm struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type int `json:"type"`
}


func main() {

	values := map[string]string{ "name": "Alarm", "type": "0", "startedBy": "00000000-0000-0000-0000-000000000000" }

	jsonValue, _ := json.Marshal(values)

	res, err := http.Post("http://localhost:5000/api/Alarms", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	data := GenericAlarm{}
	json.Unmarshal(body, &data)

	log.Println(data)

	a := []string{data.Id}

	time.Sleep(time.Duration(time.Second * time.Duration(rand.Intn(5))))


	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "http://localhost:5000/api/Alarms?closedBy=00000000-0000-0000-0000-000000000000", strings.NewReader(strings.Join(a, ",")))
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewDecoder(req.Body).Decode("[]")

	fmt.Println(req)
	req.Header.Add("Content-Type", "application/json;charset=utf-8");

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}	

	fmt.Println(resp)
}