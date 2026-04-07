package elements

// The GlobalObjectId element is not used.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/globalobjectid
import (
	"encoding/xml"
	"time"
)

type GlobalObjectId struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (G *GlobalObjectId) SetForMarshal() {
	G.XMLName.Local = "GlobalObjectId"
}

func (G *GlobalObjectId) GetSchema() *Schema {
	return &SchemaTypes
}
