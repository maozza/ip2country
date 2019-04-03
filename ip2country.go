package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	geoip2 "github.com/oschwald/geoip2-golang"
)

var dbFile = "GeoIP2-Country.mmdb"

func handler(w http.ResponseWriter, r *http.Request) {
	db, err := geoip2.Open(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var ipAddr = r.URL.Path[1:]
	if ipAddr == "" {
		errorHandler(w, r, http.StatusNotFound, "IP  is Null")
		return
	}

	ip := net.ParseIP(ipAddr)
	if ip == nil {
		errorHandler(w, r, http.StatusNotFound, "Not Valid IP")
		return
	}
	record, err := db.City(ip)
	if err != nil {
		errorHandler(w, r, 500, "Internal Server Error")
		return
	}
	resp := make(map[string]string, 2)
	resp["country"] = record.Country.Names["en"]
	resp["code"] = record.Country.IsoCode
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Fatal(err)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, msg string) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, msg)
	}
}
