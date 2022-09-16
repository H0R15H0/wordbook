package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jomei/notionapi"
)

const (
	wordbookDatabase = "e1a6f31ab3724ca6bc7aa23d9985e4b9"
)

var wordbookDatabaseColumnIDs = struct {
	word         string
	sourceUrl    string
	wikipediaUrl string
}{
	"title",
	"u%3Fqn",
	"U%3BeC",
}

func main() {
	token := os.Getenv("NOTION_INTEGRATION_TOKEN")
	client := notionapi.NewClient(notionapi.Token(token))

	resp, err := client.Database.Query(context.Background(), wordbookDatabase, &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			RichText: &notionapi.TextFilterCondition{Equals: "hoge"},
			Property: "title",
		},
	})

	if err != nil {
		panic(err)
	}

	idTextColumns := map[string]string{}
	for _, v := range resp.Results[0].Properties {
		id, text := mustStringNotionProperty(v)
		idTextColumns[id] = text
	}

	fmt.Println(fmt.Sprintf("%s, %s, %s",
		idTextColumns[wordbookDatabaseColumnIDs.word],
		idTextColumns[wordbookDatabaseColumnIDs.sourceUrl],
		idTextColumns[wordbookDatabaseColumnIDs.wikipediaUrl],
	))

}

func mustStringNotionProperty(p notionapi.Property) (id string, text string) {
	switch p.(type) {
	case *notionapi.TitleProperty:
		title := p.(*notionapi.TitleProperty)
		id, text = title.ID.String(), title.Title[0].PlainText
	case *notionapi.URLProperty:
		url := p.(*notionapi.URLProperty)
		id, text = url.ID.String(), url.URL
	default:
		panic("Undefined property detected.")
	}
	return id, text
}
