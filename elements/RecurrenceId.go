package elements

// The RecurrenceId element is used to identify a specific instance of a recurring calendar item.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/recurrenceid
import (
	"encoding/xml"
	"time"
)

type RecurrenceId struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (R *RecurrenceId) SetForMarshal() {
	R.XMLName.Local = "RecurrenceId"
}

func (R *RecurrenceId) GetSchema() *Schema {
	return &SchemaTypes
}
