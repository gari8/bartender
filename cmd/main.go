package main

import (
	"bartender"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewDatabaseConnection() (*sql.DB, error) {
	conn, err := sql.Open("postgres", "host=db user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type Input struct {
	ID uint32 `db:"id"`
	Text string `db:"text"`
}

func main() {
	//db := gorp.DbMap{}
	_, _ = NewDatabaseConnection()
	inputs := []Input{
		{
			ID: 1,
			Text: "hi honey",
		},
		{
			ID: 2,
			Text: "hi honey!",
		},
		{
			ID: 3,
			Text: "hi honey!!",
		},
		{
			ID: 4,
			Text: "hi honey!!!",
		},
		{
			ID: 5,
			Text: "hi honey!!!!",
		},
		{
			ID: 6,
			Text: "hi honey!!!!!",
		},
	}
	query := bartender.ShakeCocktails("sample", inputs)
	//query2 := bartender.ShakeOneCocktail("sample", Input{
	//	ID:   1,
	//	Text: "hi honey",
	//})
	fmt.Println(query)
	//fmt.Println(query2)
	//db.Insert(inputs)
}
