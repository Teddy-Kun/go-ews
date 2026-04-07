package elements

// The AssignedTime element represents the time when a task is assigned to a contact.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/assignedtime
import (
	"encoding/xml"
	"time"
)

type AssignedTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (A *AssignedTime) SetForMarshal() {
	A.XMLName.Local = "AssignedTime"
}

func (A *AssignedTime) GetSchema() *Schema {
	return &SchemaTypes
}
