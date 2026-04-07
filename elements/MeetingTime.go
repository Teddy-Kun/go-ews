package elements

// The MeetingTime element represents a suggested meeting time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/meetingtime
import (
	"encoding/xml"
	"time"
)

type MeetingTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (M *MeetingTime) SetForMarshal() {
	M.XMLName.Local = "MeetingTime"
}

func (M *MeetingTime) GetSchema() *Schema {
	return &SchemaTypes
}
