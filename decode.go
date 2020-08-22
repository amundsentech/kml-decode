package kmldecode

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// the complete array of all items
type Kml struct {
	Document Document `xml:"Document"`
}

// the Document struct, containing Schema
// and Folders
type Document struct {
	XMLName xml.Name      `xml:"Document"`
	Schema  []SimpleField `xml:"Schema"`
	Folder  Folder        `xml:"Folder"`
}

// Schema describes the types of the extended attributes
type Schema struct {
	XMLName     xml.Name      `xml:"Schema"`
	SimpleField []SimpleField `xml:"SimpleField"`
}

type SimpleField struct {
	XMLName xml.Name `xml:"SimpleField"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type"`
}

// the Folder struct holds all the Placemarks (drawing items)
type Folder struct {
	XMLName   xml.Name    `xml:"Folder"`
	Name      string      `xml:"name"`
	Placemark []Placemark `xml:"Placemark"`
}

// Placemarks are the root level of the geom
type Placemark struct {
	XMLName       xml.Name     `xml:"Placemark"`
	Name          string       `xml:"name"`
	Style         Style        `xml:"Style"`
	MultiGeometry []Geom       `xml:"MultiGeometry"`
	ExtendedData  ExtendedData `xml:"ExtendedData"`
}

// Styles are the colors pre assigned
type Style struct {
	XMLName   xml.Name  `xml:"Style"`
	LineStyle LineStyle `xml:"LineStyle"`
	PolyStyle PolyStyle `xml:"PolyStyle"`
}

type LineStyle struct {
	XMLName xml.Name `xml:"LineStyle"`
	Color   string   `xml:"color"`
}

type PolyStyle struct {
	XMLName xml.Name `xml:"PolyStyle"`
	Fill    string   `xml:"fill"`
}

// MultiGeometry describes various geom types
type MultiGeometry struct {
	XMLName    xml.Name   `xml:"MultiGeometry"`
	LineString Linestring `xml:"LineString"`
	Polygon    Polygon    `xml:"Polygon"`
}

type LineString struct {
	XMLName     xml.Name    `xml:"LineString"`
	Coordinates [][]float64 `xml:"coordinates"`
}

type Polygon struct {
	XMLName       xml.Name      `xml:"Polygon"`
	OuterBoundary OuterBoundary `xml:"outerBoundaryIs"`
}

type OuterBoundary struct {
	XMLName    xml.Name   `xml:"outerBoundaryIs"`
	LinearRing LinearRing `xml:"LinearRing"`
}

type LinearRing struct {
	XMLName     xml.Name    `xml:"LinearRing"`
	Coordinates [][]float64 `xml:"coordinates"`
}

// ExtendedData are attributes added to kmls that weren't in original schema
type ExtendedData struct {
	XMLName    xml.Name     `xml:"ExtendedData"`
	SchemaData []SchemaData `xml:"SchemaData"`
}

type SchemaData struct {
	XMLName    xml.Name     `xml:"SchemaData"`
	SimpleData []SimpleData `xml:"SimpleData"`
}

type SimpleData struct {
	XMLName xml.Name `xml:"SimpleData"`
	Key     string   `xml:"name"`
	Value   string   `xml:"SumpleData"`
}

// unravel the xml into a single KML struct
// all the code basically sets up the structs
// the rest is just using xml to parse
func Decode(f bytevalue) (*Kml, error) {

	var kml Kml

	_, error := xml.UnMarshal(f, kml)

	if err != nil {
		return kml, err
	}

	return kml, _
}
