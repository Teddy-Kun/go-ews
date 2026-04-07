package elements

// The CreatedTime element specifies the time at which the item was created.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/createdtime
import (
	"encoding/xml"
	"time"
)

type CreatedTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (C *CreatedTime) SetForMarshal() {
	C.XMLName.Local = "CreatedTime"
}

func (C *CreatedTime) GetSchema() *Schema {
	return &SchemaTypes
}
