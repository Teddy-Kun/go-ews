package elements

// The CompleteDate element represents the date on which an item was completed.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/completedate
import (
	"encoding/xml"
	"time"
)

type CompleteDate struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (C *CompleteDate) SetForMarshal() {
	C.XMLName.Local = "CompleteDate"
}

func (C *CompleteDate) GetSchema() *Schema {
	return &SchemaTypes
}
