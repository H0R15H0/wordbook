package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

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

type FetchedEvent struct {
	Text string `json:"text"`
}

func HandleRequest(ctx context.Context, e FetchedEvent) {
	text := e.Text

	_, err := client.Page.Create(context.Background(), &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: wordbookDatabase,
		},
		Properties: notionapi.Properties{
			wordbookDatabaseColumnIDs.word: notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: &notionapi.Text{Content: text}},
				},
			},
			wordbookDatabaseColumnIDs.sourceUrl: notionapi.URLProperty{
				URL: "https://google.com",
			},
			wordbookDatabaseColumnIDs.wikipediaUrl: notionapi.URLProperty{
				URL: fmt.Sprintf("https://ja.wikipedia.org/w/index.php?search=%s&ns0=1", url.QueryEscape(text)),
			},
		},
	})

	if err != nil {
		log.Println("text: ", text)
		log.Println(err)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
