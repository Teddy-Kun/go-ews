package elements

// The SubmittedTime element represents the time that the message entered the server.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/submittedtime
import (
	"encoding/xml"
	"time"
)

type SubmittedTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *SubmittedTime) SetForMarshal() {
	S.XMLName.Local = "SubmittedTime"
}

func (S *SubmittedTime) GetSchema() *Schema {
	return &SchemaTypes
}
