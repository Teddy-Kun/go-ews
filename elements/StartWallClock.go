package elements

// The StartWallClock element specifies the start time of a meeting in the time zone of the location in which the meeting takes place.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/startwallclock
import (
	"encoding/xml"
	"time"
)

type StartWallClock struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (S *StartWallClock) SetForMarshal() {
	S.XMLName.Local = "StartWallClock"
}

func (S *StartWallClock) GetSchema() *Schema {
	return &SchemaTypes
}
