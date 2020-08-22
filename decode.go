package kmldecode

import (
	"bytes"
	"encoding/xml"
	"strconv"
	"strings"
)

// the complete array of all items
type Kml struct {
	XMLName  xml.Name `xml:"kml" json:"kml"`
	//XMLNS    string   `xml:"xmlns,attr" json:"xmlns,attr"`
	Document Document `xml:"Document" json:"Document"`
}

// the Document struct, containing Schema
// and Folders
type Document struct {
	//XMLName xml.Name `xml:"Document" json:"Document"`
	Id     string `xml:"id,attr" json:"id,attr"`
	Schema Schema `xml:"Schema" json:"Schema"`
	Folder Folder `xml:"Folder" json:"Folder"`
}

// Schema describes the types of the extended attributes
type Schema struct {
	Name        string        `xml:"name,attr" json:"name,attr"`
	ID          string        `xml:"id,attr" json:"id,attr"`
	SimpleField []SimpleField `xml:"SimpleField" json:"SimpleField"`
}

type SimpleField struct {
	Name string `xml:"name,attr" json:"name,attr"`
	Type string `xml:"type,attr" json:"type,attr"`
}

// the Folder struct holds all the Placemarks (drawing items)
type Folder struct {
	Name       string      `xml:"name" json:"name"`
	Placemarks []Placemark `xml:"Placemark"`
}

type Placemark struct {
	Name          string        `xml:"name" json:"name"`
	Style         Style         `xml:"Style" json:"Style"`
	MultiGeometry MultiGeometry `xml:"MultiGeometry" json:"MultiGeometry"`
	Point         Point         `xml:"Point"`
	ExtendedData  ExtendedData  `xml:"ExtendedData" json:"ExtendedData"`
}

// Styles are the colors pre assigned
// The following are all STYLES
type Style struct {
	LineStyle LineStyle `xml:"LineStyle" json:"LineStyle"`
	PolyStyle PolyStyle `xml:"PolyStyle" json:"PolyStyle"`
}

type LineStyle struct {
	Color string `xml:"color" json:"color"`
}

type PolyStyle struct {
	Fill string `xml:"fill" json:"fill"`
}

// GEOMS
// MultiGeometry is the container
type MultiGeometry struct {
	LineString LineString `xml:"LineString" json:"LineString"`
	Polygon    Polygon    `xml:"Polygon" json:"Polygon"`
}

type Point struct {
	StringCoords string    `xml:"coordinates" json:"-"`
	Coordinates  []float64 `json:"coordinates"`
}

type LineString struct {
	StringCoords string      `xml:"coordinates" json:"-"`
	Coordinates  [][]float64 `json:"coordinates"`
}

type Polygon struct {
	OuterBoundary OuterBoundary `xml:"outerBoundaryIs" json:"outerBoundaryIs"`
}

type OuterBoundary struct {
	LinearRing LinearRing `xml:"LinearRing" json:"LinearRing"`
}

type LinearRing struct {
	StringCoords string      `xml:"coordinates" json:"-"`
	Coordinates  [][]float64 `json"coordinates"`
}

// EXTENDED ATTRIBUTES
// added to kmls that weren't in original schema
type ExtendedData struct {
	SchemaData SchemaData `xml:"SchemaData" json:"SchemaData"`
}

type SchemaData struct {
	Schema string	`xml:"schemaUrl,attr" json:"schemaUrl,attr"`
	SimpleData []SimpleData `xml:"SimpleData" json:"SimpleData"`
}

type SimpleData struct {
	Key   string `xml:"name,attr" json:"name,attr"`
	Value string `xml:",chardata" json:",chardata"`
}

// unravel the xml into a single KML struct
// all the code basically sets up the structs
// the rest is just using xml to parse
func KmlDecode(f *bytes.Buffer, kml *Kml) {

	// xml.Unmarshal(f, kml)
	d := xml.NewDecoder(f)

	// place all the xml where it belongs
	d.Decode(kml)

	// change coords from string >> [][]float64
	for i, geom := range kml.Document.Folder.Placemarks {

		// if point
		if len(geom.Point.StringCoords) > 0 {
			str := geom.Point.StringCoords

			payload := coordStringDecode(str)

			kml.Document.Folder.Placemarks[i].Point.Coordinates = payload[0]
		}

		// if linestring
		if len(geom.MultiGeometry.LineString.StringCoords) > 0 {
			str := geom.MultiGeometry.LineString.StringCoords

			payload := coordStringDecode(str)

			kml.Document.Folder.Placemarks[i].MultiGeometry.LineString.Coordinates = payload
		}

		// if polygon
		if len(geom.MultiGeometry.Polygon.OuterBoundary.LinearRing.StringCoords) > 0 {
			str := geom.MultiGeometry.Polygon.OuterBoundary.LinearRing.StringCoords

			payload := coordStringDecode(str)

			kml.Document.Folder.Placemarks[i].MultiGeometry.Polygon.OuterBoundary.LinearRing.Coordinates = payload
		}
	}
}

// coordStringDecode performs the parsing and number conversion
func coordStringDecode(str string) [][]float64 {
	temp := strings.Split(str, " ")

	var payload [][]float64

	for _, coord := range temp {

		xyz := strings.Split(coord, ",")

		x, _ := strconv.ParseFloat(string(xyz[0]), 64)
		y, _ := strconv.ParseFloat(string(xyz[1]), 64)

		floatcoord := []float64{x, y}

		// test if a third coord (elevation) is present!!
		if len(xyz) > 2 {
			z, _ := strconv.ParseFloat(string(xyz[2]), 64)
			floatcoord = append(floatcoord, z)
		}

		payload = append(payload,floatcoord)
	}

	return payload
}
