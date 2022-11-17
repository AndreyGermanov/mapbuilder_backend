package mapbuilder

import (
	"os"
	"testing"
)

func TestLoadFromJSON(t *testing.T) {
	country := CreateCountryShape()
	err := country.loadFromJSON("CAN", SCALE_10M)
	if err != nil {
		t.Errorf("Shape load error: '%s'\n", err.Error())
	}
	svg := country.GetSVG(TransformParams{scaleWidth: 1000},
		CountryParams{styles: map[string]string{"fill": "none", "stroke": "black", "stroke-width": "1"}})
	os.WriteFile("country.svg", []byte(svg), os.ModePerm)
}
