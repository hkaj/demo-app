package main

import (
	"expvar"
	"io"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Person is the dummy object used in our app
type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

// SuccessCount is the number of successful requests
var SuccessCount = expvar.NewInt("demo.requests.success")

// FailureCount is the number of failed requests
var FailureCount = expvar.NewInt("demo.requests.failures")

func hello(w http.ResponseWriter, req *http.Request) {
	db, err := sqlx.Connect("postgres", "host=postgres user=foo password=bar dbname=app sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	p := Person{}
	db.Get(&p, "SELECT * FROM person OFFSET floor(random() * (SELECT count(*) FROM person)) LIMIT 1 ;")
	db.Close()
	res := "Hello, " + p.FirstName + " " + p.LastName + "\n"
	// Person{FirstName: "Haissam", LastName: "Kaj"}
	SuccessCount.Add(1)
	io.WriteString(w, res)
}

func fail(w http.ResponseWriter, req *http.Request) {
	db, err := sqlx.Connect("postgres", "host=postgres user=foo password=bar dbname=app sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	p := Person{}
	var res string
	// Try and get in non-existant table
	err = db.Get(&p, "SELECT * FROM place OFFSET floor(random() * (SELECT count(*) FROM person)) LIMIT 1 ;")
	db.Close()
	if err != nil {
		res = "Hello... unknown person :)\n"
		FailureCount.Add(1)
	} else {
		res = "Hello, " + p.FirstName + " " + p.LastName + "\n"
		SuccessCount.Add(1)
	}
	// Person{FirstName: "Haissam", LastName: "Kaj"}
	io.WriteString(w, res)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/fail", fail)
	SuccessCount.Set(0)
	FailureCount.Set(0)
	// metrics on http://localhost:8080/debug/vars
	go http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":80", nil)
}
