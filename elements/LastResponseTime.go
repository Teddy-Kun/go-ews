package elements

// The LastResponseTime element represents the date and time of the latest response received.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/lastresponsetime
import (
	"encoding/xml"
	"time"
)

type LastResponseTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (L *LastResponseTime) SetForMarshal() {
	L.XMLName.Local = "LastResponseTime"
}

func (L *LastResponseTime) GetSchema() *Schema {
	return &SchemaTypes
}
