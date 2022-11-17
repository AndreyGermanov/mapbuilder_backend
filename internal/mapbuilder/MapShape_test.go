package mapbuilder

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestMapShape_GetSVG(t *testing.T) {
	err := LoadCountriesDB()
	if err != nil {
		t.Errorf("Could not load countries database. Error: '%s'", err.Error())
		return
	}
	fmt.Println(time.Now())
	mapShape := CreateMapShape()
	for _, continent := range CountriesMetadata[SCALE_10M] {
		for isoCode, _ := range continent {
			mapShape.addCountry(isoCode, SCALE_10M)
		}
	}
	svg := mapShape.GetSVG(TransformParams{scaleWidth: 1800}, CountryParams{styles: map[string]string{"stroke": "black", "fill": "none"}})
	os.WriteFile("world_10m.svg", []byte(svg), os.ModePerm)
	fmt.Println(time.Now())
	fmt.Printf("Polygons: %d\n", polygonsCount)
	fmt.Printf("Points: %d\n", pointsCount)
}
