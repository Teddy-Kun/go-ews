package elements

// The StartTime (ReminderMessageDataType) element specifies the starting time of the item that the reminder is for.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/starttime-remindermessagedatatype
import (
	"encoding/xml"
	"time"
)

type StartTimeReminderMessageDataType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *StartTimeReminderMessageDataType) SetForMarshal() {
	S.XMLName.Local = "StartTime"
}

func (S *StartTimeReminderMessageDataType) GetSchema() *Schema {
	return &SchemaTypes
}
