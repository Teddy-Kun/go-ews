package elements

// The DateTime element represents the date and time at which the time zone transition occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetime
import (
	"encoding/xml"
	"time"
)

type DateTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *DateTime) SetForMarshal() {
	D.XMLName.Local = "DateTime"
}

func (D *DateTime) GetSchema() *Schema {
	return &SchemaTypes
}
