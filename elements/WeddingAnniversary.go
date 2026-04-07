package elements

// The WeddingAnniversary element contains the wedding anniversary of a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/weddinganniversary
import (
	"encoding/xml"
	"time"
)

type WeddingAnniversary struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (W *WeddingAnniversary) SetForMarshal() {
	W.XMLName.Local = "WeddingAnniversary"
}

func (W *WeddingAnniversary) GetSchema() *Schema {
	return &SchemaTypes
}
