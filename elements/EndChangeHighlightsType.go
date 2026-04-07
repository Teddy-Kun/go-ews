package elements

// The End element specifies the changes made to a meeting end time when a meeting update occurs.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/end-changehighlightstype
import (
	"encoding/xml"
	"time"
)

type EndChangeHighlightsType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (E *EndChangeHighlightsType) SetForMarshal() {
	E.XMLName.Local = "End"
}

func (E *EndChangeHighlightsType) GetSchema() *Schema {
	return &SchemaTypes
}
