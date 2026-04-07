package elements

// The ProposedEnd (MeetingRegistrationResponseObjectType) element specifies an attendee's proposed new end time for a meeting.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/proposedend-meetingregistrationresponseobjecttype
import (
	"encoding/xml"
	"time"
)

type ProposedEndMeetingRegistrationResponseObjectType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (P *ProposedEndMeetingRegistrationResponseObjectType) SetForMarshal() {
	P.XMLName.Local = "ProposedEnd"
}

func (P *ProposedEndMeetingRegistrationResponseObjectType) GetSchema() *Schema {
	return &SchemaTypes
}
