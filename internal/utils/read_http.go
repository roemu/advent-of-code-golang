package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadHTTP(year, day int, session string) string {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	session = fmt.Sprintf("session=%s", session)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cookie", session)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error from in when making file request", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Error from in statuscode", resp.StatusCode)
	}

	return string(bodyBytes)
}
