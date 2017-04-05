package main

import (
	"log"

	"github.com/microo8/plgo"
)

//ConcatAll concatenates all values of an column in a given table
func ConcatAll(tableName, colName string) string {
	logger := plgo.NewErrorLogger("", log.Ltime|log.Lshortfile)
	db, err := plgo.Open()
	if err != nil {
		logger.Fatalf("Cannot open DB: %s", err)
	}
	defer db.Close()
	query := "select " + colName + " from " + tableName
	stmt, err := db.Prepare(query, nil)
	if err != nil {
		logger.Fatalf("Cannot prepare query statement (%s): %s", query, err)
	}
	rows, err := stmt.Query()
	if err != nil {
		logger.Fatalf("Query (%s) error: %s", query, err)
	}
	var ret string
	for rows.Next() {
		var val string
		rows.Scan(&val)
		ret += val
	}
	return ret
}
