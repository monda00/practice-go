package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var (
	//query = flag.String("query", "Nuphy Halo65", "Search term")
	//query = flag.String("query", "Nuphy Air60", "Search term")
	query = flag.String("query", "HHKB professional hybrid type-s", "Search term")
	maxResults = flag.Int64("max-results", 5, "Max Youtube results")
)


func getApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to read env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	return apiKey
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message + ": %v", err.Error())
	}
}

func searchByKeyword(service *youtube.Service, part string) {
	call := service.Search.List([]string{part}).
					Q(*query).
					Type("video").
					Order("viewCount").
					MaxResults(*maxResults)
	response, err := call.Do()
	handleError(err, "")

	fmt.Println(response)
	fmt.Println("-----------")
	fmt.Println(response.Items[0].Snippet.Title)
	fmt.Println("-----------")
	fmt.Println(response.PageInfo.TotalResults)
}

func main() {
	ctx := context.Background()
	apiKey := getApiKey()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	handleError(err, "Error creating YouTube client")

	searchByKeyword(service, "snippet")
}