package elements

// The Time element describes the time when the time changes between standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/time-timechangetype
import (
	"encoding/xml"
	"time"
)

type ClockTime struct {
	Time time.Time `xml:",chardata"`
}

func (c *ClockTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	// Parse using the TimeOnly layout: "15:04:05"
	t, err := time.Parse(time.TimeOnly, s)
	if err != nil {
		return err
	}

	*c = ClockTime{Time: t}

	return nil
}

type TimeTimeChangeType struct {
	XMLName xml.Name
	TEXT    ClockTime `xml:",chardata"`
}

func (T *TimeTimeChangeType) SetForMarshal() {
	T.XMLName.Local = "t:Time"
}

func (T *TimeTimeChangeType) GetSchema() *Schema {
	return &SchemaTypes
}
