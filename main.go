package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID            int
	AccountID     int
	Name          string
	ProfileIconID int
	RevisionDate  int
	SummonerLevel int
}

func main() {
	user := User{}
	key := "RGAPI-a1ec362a-28c2-44da-9419-633948269d18"
	// url := fmt.Sprintf("https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/Lebron?api_key=%s", key)
	url := fmt.Sprintf("https://na1.api.riotgames.com//lol/summoner/v3/summoners/by-name/Lebron?api_key=%s", key)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	response := json.NewDecoder(resp.Body).Decode(&user)
	fmt.Println(response)
	fmt.Println(user)
	fmt.Println(user.ID)

	// alternate way of reading the response, if it's not in JSON format
	// response, err := ioutil.ReadAll(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// resString := string(response)

	// fmt.Println(reflect.TypeOf(resString))
}
