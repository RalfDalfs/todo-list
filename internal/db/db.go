package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

// Шаблон для создания БД если она отсутствует
const schema = `CREATE TABLE IF NOT EXISTS scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL,
    title VARCHAR(255) NOT NULL,
    comment TEXT NOT NULL,
    repeat VARCHAR(255)
);


CREATE INDEX IF NOT EXISTS scheduler_date ON scheduler(date);
`

// Init открывает базу данных, создаёт её по шаблону если она не создана
func Init(dbFile string) error {
	_, err := os.Stat(dbFile)
	dbExists := !os.IsNotExist(err)
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("Не удалось открыть БД:%v", err)
	}
	if !dbExists {
		fmt.Println("База данных не найдена, создается новая")
		_, err = db.Exec(schema)
		if err != nil {
			return fmt.Errorf("Ошибка при создании таблицы: %v", err)
		}
	}
	return nil
}

func Close() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Printf("Ошибка при закрытии БД: %v", err)
		}
	}
}
