package elements

// The ReceivedTime element specifies the time at which an item was received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/receivedtime
import (
	"encoding/xml"
	"time"
)

type ReceivedTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (R *ReceivedTime) SetForMarshal() {
	R.XMLName.Local = "ReceivedTime"
}

func (R *ReceivedTime) GetSchema() *Schema {
	return &SchemaTypes
}
