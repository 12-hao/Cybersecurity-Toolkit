package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./Nmap.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS scan_results (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		ip TEXT NOT NULL,
		port INT NOT NULL,
		vulnerability TEXT NOT NULL,
		severity TEXT NOT NULL,
		timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);` // 确保这里没有多余的逗号

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println("创建表失败:", err)
		return
	}
	fmt.Println("创建成功")
}
