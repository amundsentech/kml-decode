package kmldecode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const (
	//testing datasets
	points = "tests/Keno 250 West Points kml only.kml"
	lines  = "tests/Keno Master Flight Lines.kml"
	shapes = "tests/Keno Master Grid Outlines.kml"
)

func TestKML2Struct(t *testing.T) {

	// build a map of the testing data and inputs
	data := make(map[string]string)
	data[points] = points
	data[lines] = lines
	data[shapes] = shapes

	for item, fileloc := range data {

		// psa
		fmt.Printf("starting in on %v\n", item)

		// prase the inputDetails from origin
		kmlfile, err := os.Open(fileloc)
		if err != nil {
			t.Errorf(err.Error())
		}

		// load file as byte
		kmlbyte, _ := ioutil.ReadAll(kmlfile)
		kmlbuf := bytes.NewBuffer(kmlbyte)

		// this is the test!!!
		var kml KML

		KMLDecode(kmlbuf, &kml)

		kmljson, err := json.Marshal(kml)
		if err != nil {
			fmt.Println("error:", err)
		}

		//fmt.Printf("Inbound reads like:\n%v\n", string(kmlbyte))
		// what does it look like?
		fmt.Printf("KML struct as json reads:\n%v\n", string(kmljson)[0:1000])
	}
}
