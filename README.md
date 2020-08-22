#Kml Decode#

Plenty of other libraries encode to kml, but none decode kml, likely due to it being xml.  The advantage of a decoder is kml files have known schemas, are the defacto google earth file, and can be commonly used as inputs for viewing in 3D (MineAR etc).

##Usage##

// Reads a kml file into a struct


kml.Umarshal(byteValue, &kml) *KML

##func Encode##

````KmlDecode(kmlbuf, &kml)````

Decodes an point, line or polygon geometry, with/out z elevation values, with/out extended attributes.

##Note!##

All of the extended attributes are by default STRING.  This parser does not create new struct fields encoded as specified by Schema.SimpleField.Name and Schema.SimpleField.Type.  


THEREFORE the user of this object must create the new field types according to the Schema, then populate with the values for each Placemark associated in its SimpleData.Key/Value struct.
