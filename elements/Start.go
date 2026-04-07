package elements

// The Start element represents the start of a duration.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/start
import (
	"encoding/xml"
	"time"
)

type Start struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *Start) SetForMarshal() {
	S.XMLName.Local = "Start"
}

func (S *Start) GetSchema() *Schema {
	return &SchemaTypes
}
