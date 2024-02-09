package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func doApiCall(url string) (string, error) {
	cocToken := os.Getenv("COC_TOKEN")

	req, err := http.NewRequest("GET", url, nil)
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
		url := params.Get("url")

		if url == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		url = strings.ReplaceAll(url, "#", "%23")

		response, err := doApiCall(url)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println("New request: ", url)

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(response))
	})

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
