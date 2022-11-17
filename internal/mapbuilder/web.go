package mapbuilder

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type MapRequest struct {
	Scale     string   `json:"scale"`
	Width     float64  `json:"width"`
	Height    float64  `json:"height"`
	Countries []string `json:"countries"`
}

func RunWebServer() {
	LoadCountriesDB()
	server := http.Server{Addr: "0.0.0.0:6001", Handler: http.HandlerFunc(handle)}

	server.ListenAndServe()
}

func handle(w http.ResponseWriter, r *http.Request) {
	SetupCORS(w, r)
	switch r.URL.String() {
	case "/metadata":
		metadata, err := getMetadata()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(metadata)
	case "/map":
		if r.Method != "POST" {
			w.WriteHeader(200)
			return
		}
		svg, err := getMap(r)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(svg)
	}
}

func SetupCORS(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, sid")
	header.Set("Content-Type", "application/json")
}

func getMetadata() ([]byte, error) {
	return json.Marshal(CountriesMetadata)
}

func getMap(r *http.Request) ([]byte, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var request MapRequest
	err = json.Unmarshal(data, &request)
	if err != nil {
		return nil, err
	}
	mapShape := CreateMapShape()
	for _, isoCode := range request.Countries {
		mapShape.addCountry(isoCode, request.Scale)
	}
	svg := mapShape.GetSVG(TransformParams{scaleWidth: request.Width, scaleHeight: request.Height},
		CountryParams{styles: map[string]string{"stroke": "black", "fill": "none"}},
	)
	svg = strings.Replace(svg, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>", "", -1)
	return []byte(svg), nil
}
