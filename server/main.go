package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jomei/notionapi"
)

const (
	wordbookDatabase = "e1a6f31ab3724ca6bc7aa23d9985e4b9"
)

var (
	wordbookDatabaseColumnIDs = struct {
		word         string
		sourceUrl    string
		wikipediaUrl string
	}{
		"title",
		"u%3Fqn",
		"U%3BeC",
	}
	token  = os.Getenv("NOTION_INTEGRATION_TOKEN")
	client = notionapi.NewClient(notionapi.Token(token))
)

type requestBody struct {
	Text string `json:"text"`
}

func HandleRequest(ctx context.Context, e events.LambdaFunctionURLRequest) {
	var r requestBody
	json.Unmarshal([]byte(e.Body), &r)

	_, err := client.Page.Create(ctx, &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: wordbookDatabase,
		},
		Properties: notionapi.Properties{
			wordbookDatabaseColumnIDs.word: notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: &notionapi.Text{Content: r.Text}},
				},
			},
			wordbookDatabaseColumnIDs.sourceUrl: notionapi.URLProperty{
				URL: "https://google.com",
			},
			wordbookDatabaseColumnIDs.wikipediaUrl: notionapi.URLProperty{
				URL: fmt.Sprintf("https://ja.wikipedia.org/w/index.php?search=%s&ns0=1", url.QueryEscape(r.Text)),
			},
		},
	})

	if err != nil {
		log.Println("text: ", r.Text)
		log.Println(err)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
