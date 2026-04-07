package elements

// The Date element represents the date that contains the suggested meeting times.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/date
import (
	"encoding/xml"
	"time"
)

type Date struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *Date) SetForMarshal() {
	D.XMLName.Local = "Date"
}

func (D *Date) GetSchema() *Schema {
	return &SchemaTypes
}
