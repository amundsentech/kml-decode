#Kml Decode#

Plenty of other libraries encode to kml, but none decode kml, likely due to it being xml.  The advantage of a decoder is kml files have known schemas, are the defacto google earth file, and can be commonly used as inputs for viewing in 3D (MineAR etc).

##Usage##

// Reads a kml file into a struct


kml.Umarshal(byteValue, &kml) *KML

##func Encode##

```func Decode(f byteValue) (kml.Element, error)```

Decodes an arbitrary geometry.
