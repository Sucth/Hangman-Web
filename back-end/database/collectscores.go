package backend

import (
	"database/sql"
	"fmt"
)

type Scores struct {
	playername   string
	pointsnumber int
}

func CollectScores() []Scores {
	var result []Scores
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/hangman")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM players")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var points int
		var pseudo string
		if err := rows.Scan(&pseudo, &points); err != nil {
			fmt.Println(err)
		}
		result = append(result, Scores{
			playername:   pseudo,
			pointsnumber: points,
		})
	}
	return result
}
