package elements

// The Country element represents the country or region for a given physical address.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/countryorregion
import "encoding/xml"

type CountryOrRegion struct {
	XMLName xml.Name
	TEXT    string `xml:",chardata"`
}

func (C *CountryOrRegion) SetForMarshal() {
	C.XMLName.Local = "CountryOrRegion"
}

func (C *CountryOrRegion) GetSchema() *Schema {
	return &SchemaTypes
}
