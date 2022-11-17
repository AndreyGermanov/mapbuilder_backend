package mapbuilder

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var CountryShapes = map[string]*CountryShape{}

var CountriesMetadata = map[string]map[string]map[string]string{}

var basePath = ""

func LoadCountriesDB() error {
	scales, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}
	for _, scale := range scales {
		if !scale.IsDir() {
			continue
		}
		if _, ok := CountriesMetadata[scale.Name()]; !ok {
			CountriesMetadata[scale.Name()] = map[string]map[string]string{}
		}
		countryFiles, err := os.ReadDir(path.Join(basePath, scale.Name()))
		if err != nil {
			fmt.Printf("Error reading path '%s'\n", path.Join(basePath, scale.Name()))
			continue
		}
		for _, countryFile := range countryFiles {
			if countryFile.IsDir() {
				continue
			}
			LoadCountry(scale.Name(), countryFile.Name())
		}
	}
	return nil
}

func LoadCountry(scale, countryFile string) {
	fileParts := strings.Split(countryFile, ".")
	if len(fileParts) != 2 {
		fmt.Printf("Incorrect country GeoJSON file '%s'\n", countryFile)
		return
	}
	isoCode := fileParts[0]
	country := CreateCountryShape()
	err := country.loadFromJSON(isoCode, scale)
	if err != nil {
		fmt.Printf("Could not JSON-decode country from file '%s'\n", path.Join(basePath, scale, countryFile))
		return
	}
	if _, ok := CountriesMetadata[scale][country.continent]; !ok {
		CountriesMetadata[scale][country.continent] = map[string]string{}
	}
	CountriesMetadata[scale][country.continent][country.isoCode] = country.name
	CountryShapes[scale+"_"+country.isoCode] = country
}

func init() {
	wd, _ := os.Getwd()
	basePath = path.Join(wd, "geodata")
}
