package elements

// The ReminderNextTime element specifies the date and time for the next reminder.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/remindernexttime
import (
	"encoding/xml"
	"time"
)

type ReminderNextTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (R *ReminderNextTime) SetForMarshal() {
	R.XMLName.Local = "ReminderNextTime"
}

func (R *ReminderNextTime) GetSchema() *Schema {
	return &SchemaTypes
}
