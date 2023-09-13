package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/joho/godotenv"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type KeyboardScore struct {
	Product string
	Score int64
}

type KeyboardScoreList []KeyboardScore

func (ks KeyboardScoreList) Len() int { return len(ks) }
func (ks KeyboardScoreList) Swap(i, j int) { ks[i], ks[j] = ks[j], ks[i] }
func (ks KeyboardScoreList) Less(i, j int) bool { return ks[i].Score < ks[j].Score }

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

func getTotalResults(service *youtube.Service, part string, query string) int64 {
	call := service.Search.List([]string{part}).
					Q(query).
					Type("video").
					Order("relevance").
					MaxResults(10)
	response, err := call.Do()

	handleError(err, "")

	fmt.Printf("%s: %d\n", query, response.PageInfo.TotalResults)
	/*
	fmt.Println("===== Movies =====")

	items := response.Items
	for _, item := range items {
		fmt.Println(item.Snippet.Title)
	}
	fmt.Println("")
	*/

	totalResults := response.PageInfo.TotalResults

	return totalResults
}

func createKeyboardRanking(keyboardScoreMap map[string]int64) {
	kr := make(KeyboardScoreList, len(keyboardScoreMap))
	i := 0
	for k, v := range keyboardScoreMap {
		kr[i] = KeyboardScore{k, v}
		i++
	}
	sort.Sort(sort.Reverse(kr))
	for i, v := range kr {
		fmt.Println(i, v)
	}
}

func main() {
	ctx := context.Background()
	apiKey := getApiKey()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	handleError(err, "Error creating YouTube client")

	queries := map[string]string{
		"HHKB Professional HYBRID Type-S": `"HHKB Professional HYBRID Type-S"`,
		"HHKB Professional HYBRID": `"HHKB Professional HYBRID" -"Type-S"`,
		"HHKB Professional Classic": `"HHKB Professional Classic" -"HYBRID"`,
		"NuPhy Air75": `"NuPhy Air75"`,
		"NuPhy Air96": `"NuPhy Air96"`,
		"NuPhy Air60": `"NuPhy Air60"`,
		"NuPhy Field75": `"NuPhy Field75"`,
		"NuPhy Halo96": `"NuPhy Halo96"`,
		"NuPhy Halo75": `"NuPhy Halo75"`,
		"NuPhy Halo65": `"NuPhy Halo65"`,
		"MX Mechanical": `"MX Mechanical" -"Mini for Mac"`,
		"MX Mechanical Mini for Mac": `"MX Mechanical Mini for Mac"`,
		"MX Keys S": `"MX Keys S"`,
		"MX Keys Mini": `"MX Keys Mini" -"for Mac"`,
		"MX Keys for Mac": `"MX Keys for Mac"`,
		"MX Keys Mini for Mac": `"MX Keys Mini for Mac"`,
		"ERGO K860": `"ERGO K860"`,
		"Craft": `"Craft" "Logicool"`,
		"Signature K855": `"Signature K855"`,
		"POP Keys": `"POP Keys" "Logicool"`,
		"Keychron K1": `"Keychron K1" -"SE"`,
		"Keychron K1 SE": `"Keychron K1 SE"`,
		"Keychron K3": `"Keychron K3"`,
		"Keychron K5": `"Keychron K5" -"SE"`,
		"Keychron K5 SE": `"Keychron K5 SE"`,
		"Keychron K7": `"Keychron K7"`,
		"Keychron S1": `"Keychron S1"`,
		"Keychron K1 Pro": `"Keychron K1 Pro"`,
		"Keychron K3 Pro": `"Keychron K3 Pro"`,
		"Keychron K5 Pro": `"Keychron K5 Pro"`,
		"Keychron K7 Pro": `"Keychron K7 Pro"`,
		"Keychron K9 Pro": `"Keychron K9 Pro"`,
	}

	keyboardScoreMap := make(map[string]int64)
	for product, query := range queries {
		keyboardScoreMap[product] = getTotalResults(service, "snippet", query)
	}

	createKeyboardRanking(keyboardScoreMap)
}