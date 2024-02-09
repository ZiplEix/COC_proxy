package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Request struct {
	PlayerId string `json:"playerId"`
	Url      string `json:"url"`
}

func doApiCall(request Request) (string, error) {
	cocToken := os.Getenv("COC_TOKEN")

	req, err := http.NewRequest("GET", request.Url+"%23"+request.PlayerId, nil)
	req.Header.Set("Authorization", "Bearer "+cocToken)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	godotenv.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		playerId := params.Get("playerId")
		url := params.Get("url")

		if playerId == "" || url == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		request := Request{PlayerId: playerId, Url: url}
		response, err := doApiCall(request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println("New request: ", request)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
