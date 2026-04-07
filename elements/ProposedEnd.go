package elements

// The ProposedEnd element specifies the proposed end time of a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedend
import (
	"encoding/xml"
	"time"
)

type ProposedEnd struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (P *ProposedEnd) SetForMarshal() {
	P.XMLName.Local = "ProposedEnd"
}

func (P *ProposedEnd) GetSchema() *Schema {
	return &SchemaTypes
}
