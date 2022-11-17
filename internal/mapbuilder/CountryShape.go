package mapbuilder

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CountryShape struct {
	ShapeParams
	isoCode   string
	continent string
	name      string
	scale     string
	polygons  []Polygon
}

var polygonsCount = 0
var pointsCount = 0

type CountryParams struct {
	styles map[string]string
}

func CreateCountryShape() *CountryShape {
	result := &CountryShape{
		ShapeParams: ShapeParams{left: 99999999.0, top: 99999999.0, right: -99999999.9, bottom: -999999999.9},
		polygons:    []Polygon{},
	}
	return result
}

func RequestCountry(isoCode string, scale string) ([]byte, error) {
	return os.ReadFile(basePath + "/" + scale + "/" + isoCode + ".geojson")
}

func (shape *CountryShape) loadFromJSON(isoCode string, scale string) error {
	data, err := RequestCountry(isoCode, scale)
	if err != nil {
		return err
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return err
	}
	properties, ok := result["properties"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("Could not load properties of country '%s'.", isoCode)
	}
	shape.name, ok = properties["name"].(string)
	if !ok {
		return fmt.Errorf("Could not load country name for '%s'.", isoCode)
	}
	shape.continent = properties["continent"].(string)
	if !ok {
		return fmt.Errorf("Could not load country continent for '%s'.", isoCode)
	}
	geometry, ok := result["geometry"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("Could not load geometry of country '%s'.", isoCode)
	}
	if geometry["type"].(string) == "MultiPolygon" {
		polygonIfs := geometry["coordinates"].([]interface{})
		for _, polygonIf := range polygonIfs {
			shape.loadPolygonFromJSON(polygonIf.([]interface{})[0].([]interface{}))
		}
	} else {
		shape.loadPolygonFromJSON(geometry["coordinates"].([]interface{})[0].([]interface{}))
	}
	shape.width = shape.right - shape.left
	shape.height = shape.bottom - shape.top
	shape.isoCode = isoCode
	shape.scale = scale
	return nil
}

func (shape *CountryShape) loadPolygonFromJSON(polygonIf []interface{}) {
	polygon := Polygon{}
	for _, pointIf := range polygonIf {
		x := pointIf.([]interface{})[0].(float64)
		y := -1.0 * pointIf.([]interface{})[1].(float64)
		if x < shape.left {
			shape.left = x
		}
		if x > shape.right {
			shape.right = x
		}
		if y < shape.top {
			shape.top = y
		}
		if y > shape.bottom {
			shape.bottom = y
		}
		polygon = append(polygon, []float64{x, y})
		pointsCount++
	}
	if len(polygon) > 0 {
		shape.polygons = append(shape.polygons, polygon)
	}
	polygonsCount++
}

func (shape *CountryShape) GetSVG(params TransformParams, countryParams CountryParams) string {
	params = PrepareTransformParams(params, shape.ShapeParams)
	result := strings.Builder{}
	result.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	result.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" viewBox="%f %f %f %f">`,
		0.0, 0.0, shape.width*params.scaleX, shape.height*params.scaleY),
	)
	result.WriteString(shape.GetSVGPath(params, countryParams))
	result.WriteString("</svg>")
	return result.String()
}

func (shape *CountryShape) GetSVGPath(params TransformParams, countryParams CountryParams) string {
	result := strings.Builder{}
	result.WriteString(`<path `)
	if params.offsetX != 0 {
		result.WriteString(fmt.Sprintf(`offsetX="%s" `, strconv.FormatFloat(params.offsetX, 'f', -1, 64)))
	}
	if params.offsetY != 0 {
		result.WriteString(fmt.Sprintf(`offsetY="%s" `, strconv.FormatFloat(params.offsetY, 'f', -1, 64)))
	}
	if params.scaleX != 1 {
		result.WriteString(fmt.Sprintf(`scaleX="%s" `, strconv.FormatFloat(params.scaleX, 'f', -1, 64)))
	}
	if params.scaleY != 1 {
		result.WriteString(fmt.Sprintf(`scaleY="%s" `, strconv.FormatFloat(params.scaleY, 'f', -1, 64)))
	}
	result.WriteString(`id="` + shape.isoCode + `" `)
	if countryParams.styles != nil {
		result.WriteString(`style="`)
		for name, value := range countryParams.styles {
			result.WriteString(name + ":" + value + ";")
		}
		result.WriteString(`" `)
	}
	result.WriteString(`d="`)
	for _, polygon := range shape.polygons {
		result.WriteString("M ")
		for _, point := range polygon {
			result.WriteString(
				fmt.Sprintf("%s,%s ",
					strconv.FormatFloat((point[0]-params.offsetX)*params.scaleX, 'f', -1, 64),
					strconv.FormatFloat((point[1]-params.offsetY)*params.scaleY, 'f', -1, 64)),
			)
		}
		result.WriteString("Z ")
	}
	result.WriteString(`"/>`)
	return result.String()
}
