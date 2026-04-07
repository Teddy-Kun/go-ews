package elements

// The RetentionDate element specifies the last date that an item must be retained.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/retentiondate
import (
	"encoding/xml"
	"time"
)

type RetentionDate struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (R *RetentionDate) SetForMarshal() {
	R.XMLName.Local = "RetentionDate"
}

func (R *RetentionDate) GetSchema() *Schema {
	return &SchemaTypes
}
