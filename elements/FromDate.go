package elements

// The FromDate element specifies the date that the message was sent.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/fromdate
import (
	"encoding/xml"
	"time"
)

type FromDate struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (F *FromDate) SetForMarshal() {
	F.XMLName.Local = "FromDate"
}

func (F *FromDate) GetSchema() *Schema {
	return &SchemaMessages
}
