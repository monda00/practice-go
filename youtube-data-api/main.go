package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
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

func searchByKeyword(service *youtube.Service, part string, query string) {
	call := service.Search.List([]string{part}).
					Q(query).
					Type("video").
					Order("relevance").
					MaxResults(10)
	response, err := call.Do()

	handleError(err, "")

	fmt.Printf("%s: %d\n", query, response.PageInfo.TotalResults)
	fmt.Println("===== Movies =====")

	items := response.Items
	for _, item := range items {
		fmt.Println(item.Snippet.Title)
	}
	fmt.Println("")

}

func main() {
	ctx := context.Background()
	apiKey := getApiKey()

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))

	handleError(err, "Error creating YouTube client")

	/*
	queries := [...] string{
		`"HHKB Professional HYBRID Type-S"`,
		`"HHKB Professional HYBRID"-"Type-S"`,
		`"HHKB Professional Classic"-"HYBRID"`,
		`"NuPhy Air75"`,
		`"NuPhy Air96"`,
		`"NuPhy Air60"`,
		`"NuPhy Field75"`,
		`"NuPhy Halo96"`,
		`"NuPhy Halo75"`,
		`"NuPhy Halo65"`,
		`"MX Mechanical"-"Mini for Mac"`,
		`"MX Mechanical Mini for Mac"`,
		`"MX Keys S"`,
		`"MX Keys Mini"-"for Mac"`,
		`"MX Keys for Mac"`,
		`"MX Keys Mini for Mac"`,
		`"ERGO K860"`,
		`"Craft" "Logicool"`,
		`"Signature K855"`,
		`"POP Keys" "Logicool"`,
		`"Keychron K1"-"SE"`,
		`"Keychron K1 SE"`,
		`"Keychron K3"`,
		`"Keychron K5"-"SE"`,
		`"Keychron K5 SE"`,
	}
	*/
	queries := [...] string{
		`"HHKB Professional HYBRID Type-S"`,
		`allintitle: "HHKB Professional HYBRID Type-S"`,
		`"HHKB Professional HYBRID"-"Type-S"`,
		`"HHKB Professional HYBRID" -"Type-S"`,
		`"HHKB Professional HYBRID" NOT "Type-S"`,
		`allintitle: "HHKB Professional HYBRID"-"Type-S"`,
		`"HHKB Professional Classic"-"HYBRID"`,
		`allintitle: "HHKB Professional Classic" -"HYBRID"`,
		`allintitle: "HHKB Professional Classic" NOT "HYBRID"`,
		`"MX Mechanical"-"Mini for Mac"`,
		`"MX Mechanical" -"Mini for Mac"`,
		`"MX Mechanical" NOT "Mini for Mac"`,
		`"MX Mechanical Mini for Mac"`,
		`allintitle: "MX Mechanical Mini for Mac"`,
		`"MX Keys S"`,
		`"MX Keys Mini"-"for Mac"`,
		`"MX Keys for Mac"`,
		`"MX Keys Mini for Mac"`,
		`"Keychron K1"-"SE"`,
		`allintitle: "Keychron K1"-"SE"`,
		`allintitle: "Keychron K1" -"SE"`,
		`allintitle: "Keychron K1" NOT "SE"`,
		`"Keychron K1 SE"`,
		`"Keychron K3"`,
		`"Keychron K5"-"SE"`,
		`"Keychron K5 SE"`,
	}

	for _, query := range queries {
		searchByKeyword(service, "snippet", query)
	}
}