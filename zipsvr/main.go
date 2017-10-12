package main

import (
	"encoding/json"
	"fmt"
	"info344-in-class/zipsvr/handlers"
	"info344-in-class/zipsvr/models"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const zipsPath = "/zips/"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s!", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 { //if not set, default to :443
		addr = ":443"
	}

	tlsKey := os.Getenv("TLSKEY")
	tlsCert := os.Getenv("TLSCERT")
	if len(tlsKey) == 0 || len(tlsCert) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}

	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatalf("Error loading zips: %v", err)
	}
	log.Printf("Loaded %d zips", len(zips))

	cityIndex := models.ZipIndex{}
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}

	mux.Handle(zipsPath, cityHandler)

	fmt.Printf("server is listening at https://%s\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCert, tlsKey, mux))
}
