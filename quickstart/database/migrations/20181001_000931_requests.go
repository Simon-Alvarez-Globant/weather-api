package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Requests_20181001_000931 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Requests_20181001_000931{}
	m.Created = "20181001_000931"

	migration.Register("Requests_20181001_000931", m)
}

// Run the migrations
func (m *Requests_20181001_000931) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE weather(id INT AUTO_INCREMENT PRIMARY KEY, city VARCHAR(512), response TEXT, timestamp DATETIME)`)
	m.SQL(`CREATE INDEX idx_url ON weather(city)`)
}

// Reverse the migrations
func (m *Requests_20181001_000931) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`DROP TABLE weather`)
}
