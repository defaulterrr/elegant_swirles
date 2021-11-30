package main

import (
	"math/rand"
	"time"

	"github.com/defaulterrr/iot3/dht/internal/app"
)

type DHT struct {
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	Created     time.Time `json:"created"`
}

func GetDHTMetrics() DHT {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return DHT{
		Temperature: 20 + r1.Float32()*10,
		Humidity:    20 + r1.Float32()*30,
		Created:     time.Now(),
	}
}

func main() {
	// http.HandleFunc("/dht", func(w http.ResponseWriter, r *http.Request) {
	// 	metrics, err := json.Marshal(GetDHTMetrics())
	// 	if err != nil {
	// 		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// 		return
	// 	}

	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write(metrics)
	// })

	// log.Fatal(http.ListenAndServe(":8091", nil))

	app.Run()
}
