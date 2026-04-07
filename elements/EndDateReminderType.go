package elements

// The EndDate element specifies the end date of the item the reminder is for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/enddate-remindertype
import (
	"encoding/xml"
	"time"
)

type EndDateReminderType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (E *EndDateReminderType) SetForMarshal() {
	E.XMLName.Local = "EndDate"
}

func (E *EndDateReminderType) GetSchema() *Schema {
	return &SchemaTypes
}
