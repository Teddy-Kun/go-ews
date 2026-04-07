package elements

// The StartDate element represents the start date of an item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startdate
import (
	"encoding/xml"
	"time"
)

type StartDate struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *StartDate) SetForMarshal() {
	S.XMLName.Local = "StartDate"
}

func (S *StartDate) GetSchema() *Schema {
	return &SchemaTypes
}
