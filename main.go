package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type coord struct {
	x int
	y int
}

func main() {
	clearScreen()
	dbConnection()
	// Schiffe()

}

func Schiffe() {

	x_range := 20
	y_range := 20
	ships := []Ship{}
	ship := Ship{}
	cood := coord{}

	cood = coord{x: 5, y: 8}
	ship = Ship{length: 4, startPos: cood, orientation: 45}
	ships = append(ships, ship)
	ship = Ship{length: 6, startPos: coord{x: 6, y: 8}, orientation: 45}
	ships = append(ships, ship)
	ship = Ship{length: 2, startPos: coord{x: 2, y: 1}, orientation: 45}
	ships = append(ships, ship)

	game := NewGame(NewPlayground(x_range, y_range), ships)
	game.getPlayground().setValue(4, 4, 4)
	game.displayGame()

	s := fmt.Sprintf("%U", 'Z')
	fmt.Println(s)
}

func dbConnection() {
	dsn := "extern:marco@tcp(192.168.188.193:3306)/privat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// Loggt den Fehler und beendet das Programm, wenn die DSN-Syntax ungültig ist etc.
		log.Fatalf("Fehler beim Vorbereiten der Datenbankverbindung: %v", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		// Loggt den Fehler und beendet das Programm, wenn keine Verbindung hergestellt werden konnte.
		log.Fatalf("Fehler beim Verbindungsaufbau zur Datenbank: %v", err)
	}
	fmt.Println("Erfolgreich mit MariaDB verbunden!")

	rows, err := db.Query("select id from test order by 1 desc")
	// type person struct {
	// 	vorname        string
	// 	nachname       string
	// 	titel          string
	// 	erstellt_durch string
	// }

	//rows, err := db.Query("select Vorname, Nachname, Titel, Geburt, Erstellt_Am, Erstellt_durch from personen order by 1 desc")
	if err != nil {
		//TODO
	}
	defer rows.Close()

	var id int
	text := "Wert: "
	for rows.Next() {
		err := rows.Scan(&id)
		text = fmt.Sprintf("%v %v\t", text, id)
		if err != nil {
			log.Printf("Fehler beim Scannen einer Benutzerzeile: %v", err)
			continue // Überspringe diese Zeile bei Fehler
		}
	}
	fmt.Println(text)
}
