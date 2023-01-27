package backend

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Addscore(Pseudo string, difficulty string) {
	var PJ int
	type Hangman struct {
		pseudo string
		points int
	}
	var Players []Hangman

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/hangman")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM  players")
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	for row.Next() {
		var points int
		var pseudo string
		if err := row.Scan(&pseudo, &points); err != nil {
			fmt.Println(err)
			return
		}
		Players = append(Players, Hangman{
			pseudo: pseudo,
			points: points,
		})
	}

	found := false
	for i := 0; i < len(Players); i++ {
		if Players[i].pseudo == Pseudo {
			PJ = Players[i].points
			found = true
		}
	}

	if found {
		stmt, err := db.Prepare("UPDATE players SET points = ? WHERE pseudo = ?")
		if err != nil {
			fmt.Println()
		}
		defer stmt.Close()
		if difficulty == "1" {
			result, err := stmt.Exec(PJ+1, Pseudo)
			if err != nil {
				fmt.Println(err)
				fmt.Println(result)
			}
		} else if difficulty == "2" {
			result, err := stmt.Exec(PJ+2, Pseudo)
			if err != nil {
				fmt.Println(err)
				fmt.Println(result)
			}
		} else if difficulty == "3" {
			result, err := stmt.Exec(PJ+3, Pseudo)
			if err != nil {
				fmt.Println(err)
				fmt.Println(result)
			}
		}
	} else {
		_, err := db.Exec("INSERT INTO players (points, pseudo) VALUES (?, ?)", 1, Pseudo)
		if err != nil {
			fmt.Println(err)
		}
	}

}
