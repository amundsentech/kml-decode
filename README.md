# Kml Decode

Plenty of libraries encode kml, but none decode kml, likely due to it being xml which can use pre-made parsers.  The advantage of a special kmldecoder is kml files have known schemas, are the default google earth file types, and can be commonly used as inputs for other uses such as viewing in 3D (eg. MineAR).  As kml xml is verbose and repetative, ***Kml Decode*** makes decoding/unmarshaling them pretty useful.

## Usage

````KmlDecode(kmlbuf, &kml)````

Reads a kml file into a kml struct.  Decodes one/more point, line or polygon geometries, with/out z elevation values, with/out extended attributes.

## Note!

All of the extended attributes are by default type STRING.  This parser does _not_ create new struct fields encoded as specified by Schema.SimpleField.Name and Schema.SimpleField.Type.  


THEREFORE the user of this object must create the new field types according to the Schema, then populate with the values for each Placemark associated in its SimpleData.Key/Value struct.
