package main

import (
	"encoding/xml"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Camera struct {
	PeopleCount uint32 `xml:"peoplecount"`
}

func GetDHTMetrics() Camera {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return Camera{
		PeopleCount: r1.Uint32() % 50,
	}
}

func main() {
	http.HandleFunc("/camera", func(w http.ResponseWriter, r *http.Request) {
		metrics, err := xml.Marshal(GetDHTMetrics())
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/xml")
		w.Write(metrics)
	})

	log.Fatal(http.ListenAndServe(":8092", nil))
}
