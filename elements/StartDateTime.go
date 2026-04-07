package elements

// The StartDateTime element specifies the start date and time for a rule or a search.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startdatetime
import (
	"encoding/xml"
	"time"
)

type StartDateTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *StartDateTime) SetForMarshal() {
	S.XMLName.Local = "StartDateTime"
}

func (S *StartDateTime) GetSchema() *Schema {
	return &SchemaMessages
}
