package internal

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("sqlite", dataSourceName)
	if err != nil {
		log.Fatal("Cannot open database:", err)
	}

	createTables := `
    CREATE TABLE IF NOT EXISTS subscriptions (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        city TEXT NOT NULL,
        frequency TEXT NOT NULL,
        confirmed BOOLEAN DEFAULT 0,
        token TEXT
    );
    CREATE TABLE IF NOT EXISTS weather (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        city TEXT NOT NULL UNIQUE,
        temp_c REAL NOT NULL,
        weather TEXT NOT NULL
    );
    `
	_, err = DB.Exec(createTables)
	if err != nil {
		log.Fatal("Cannot create tables:", err)
	}

	initWeather := []struct {
		City    string
		TempC   float64
		Weather string
	}{
		{"Київ", 20.5, "Сонячно"},
		{"Львів", 18.0, "Хмарно"},
		{"Одеса", 22.3, "Ясно"},
		{"Харків", 19.7, "Дощ"},
		{"Дніпро", 21.1, "Сонячно"},
		{"Запоріжжя", 20.0, "Ясно"},
		{"Вінниця", 18.5, "Хмарно"},
		{"Івано-Франківськ", 17.2, "Дощ"},
		{"Тернопіль", 17.8, "Сонячно"},
		{"Чернівці", 18.3, "Ясно"},
	}

	for _, w := range initWeather {
		_, err := DB.Exec(
			`INSERT OR IGNORE INTO weather (city, temp_c, weather) VALUES (?, ?, ?)`,
			w.City, w.TempC, w.Weather,
		)
		if err != nil {
			log.Printf("Cannot insert weather for %s: %v", w.City, err)
		}
	}
}
