package main

import "github.com/defaulterrr/iot3/processing/internal/app"

// type DHT struct {
// 	Temperature float32   `json:"temperature"`
// 	Humidity    float32   `json:"humidity"`
// 	Created     time.Time `json:"created"`
// }

// type Camera struct {
// 	PeopleCount uint32 `xml:"peoplecount"`
// }

// type L1 struct {
// 	DHTMetrics    DHT
// 	CameraMetrics Camera
// }

func main() {
	app.Run()

	// l1 := L1{}

	// dht, err := http.Get("http://localhost:8091/dht")
	// if err != nil {
	// 	log.Fatalf("Error while getting DHT: %v", err)
	// }
	// defer dht.Body.Close()

	// err = json.NewDecoder(dht.Body).Decode(&l1.DHTMetrics)
	// if err != nil {
	// 	log.Fatalf("Error while decoding DHT: %v", err)
	// }

	// camera, err := http.Get("http://localhost:8092/camera")
	// if err != nil {
	// 	log.Fatalf("Error while getting Camera: %v", err)
	// }
	// defer camera.Body.Close()

	// err = xml.NewDecoder(camera.Body).Decode(&l1.CameraMetrics)
	// if err != nil {
	// 	log.Fatalf("Error while decoding Camera: %v", err)
	// }

	// f, err := os.OpenFile("index.html", os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Fatalf("Error while opening file: %v", err)
	// }

	// tmpl, err := template.ParseFiles("template.html")
	// if err != nil {
	// 	log.Fatalf("Error while parsing file: %v", err)
	// }

	// err = tmpl.Execute(f, l1)
	// if err != nil {
	// 	log.Fatalf("Error while creating template: %v", err)
	// }

}
