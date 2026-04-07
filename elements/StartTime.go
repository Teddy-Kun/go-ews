package elements

// The StartTime element represents the start of a time span.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttime
import (
	"encoding/xml"
	"time"
)

type StartTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *StartTime) SetForMarshal() {
	S.XMLName.Local = "StartTime"
}

func (S *StartTime) GetSchema() *Schema {
	return &SchemaTypes
}
