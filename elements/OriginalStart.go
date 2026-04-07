package elements

// The OriginalStart element represents the original start time of a calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/originalstart
import (
	"encoding/xml"
	"time"
)

type OriginalStart struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (O *OriginalStart) SetForMarshal() {
	O.XMLName.Local = "OriginalStart"
}

func (O *OriginalStart) GetSchema() *Schema {
	return &SchemaTypes
}
