package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jomei/notionapi"
)

const wordbook_database = "e1a6f31ab3724ca6bc7aa23d9985e4b9"

func main() {
	token := os.Getenv("NOTION_INTEGRATION_TOKEN")
	client := notionapi.NewClient(notionapi.Token(token))

	word, err := client.Database.Query(context.Background(), wordbook_database, &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			RichText: &notionapi.TextFilterCondition{Equals: "hoge"},
			Property: "Word",
		},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(word)
}
