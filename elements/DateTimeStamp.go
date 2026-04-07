package elements

// The DateTimeStamp element indicates the date and time that an instance of a calendar object was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/datetimestamp
import (
	"encoding/xml"
	"time"
)

type DateTimeStamp struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (D *DateTimeStamp) SetForMarshal() {
	D.XMLName.Local = "DateTimeStamp"
}

func (D *DateTimeStamp) GetSchema() *Schema {
	return &SchemaTypes
}
