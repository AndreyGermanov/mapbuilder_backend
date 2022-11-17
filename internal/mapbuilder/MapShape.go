package mapbuilder

import (
	"fmt"
	"strings"
)

type MapShape struct {
	countries map[string]*CountryShape
	ShapeParams
}

func CreateMapShape() *MapShape {
	result := &MapShape{countries: map[string]*CountryShape{},
		ShapeParams: ShapeParams{left: 99999999.0, top: 99999999.0, bottom: -99999999.9, right: -999999999.9},
	}
	return result
}

func (m *MapShape) addCountry(isoCode string, scale string) error {
	if _, ok := m.countries[isoCode]; ok {
		return fmt.Errorf("Country already exists '%s'", isoCode)
	}
	if country, ok := CountryShapes[scale+"_"+isoCode]; !ok {
		return fmt.Errorf("Country with isoCode %s in scale %s does not exist in database", isoCode, scale)
	} else {
		m.countries[isoCode] = country
		if country.left < m.left {
			m.left = country.left
		}
		if country.top < m.top {
			m.top = country.top
		}
		if country.right > m.right {
			m.right = country.right
		}
		if country.bottom > m.bottom {
			m.bottom = country.bottom
		}
		m.width = m.right - m.left
		m.height = m.bottom - m.top
	}
	return nil
}

func (m *MapShape) GetSVG(params TransformParams, countryParams CountryParams) string {
	params = PrepareTransformParams(params, m.ShapeParams)
	result := strings.Builder{}
	result.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	result.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="%f %f %f %f">`,
		0.0, 0.0, m.width*params.scaleX, m.height*params.scaleY),
	)
	for _, country := range m.countries {
		result.WriteString(country.GetSVGPath(params, countryParams))
	}
	result.WriteString("</svg>")
	return result.String()
}
