package elements

// The DueDate element represents the date an item is due.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/duedate
import (
	"encoding/xml"
	"time"
)

type DueDate struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *DueDate) SetForMarshal() {
	D.XMLName.Local = "DueDate"
}

func (D *DueDate) GetSchema() *Schema {
	return &SchemaTypes
}
