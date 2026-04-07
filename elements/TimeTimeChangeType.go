package elements

// The Time element describes the time when the time changes between standard time and daylight saving time.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/time-timechangetype
import (
	"encoding/xml"
	"time"
)

type TimeTimeChangeType struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (T *TimeTimeChangeType) SetForMarshal() {
	T.XMLName.Local = "Time"
}

func (T *TimeTimeChangeType) GetSchema() *Schema {
	return &SchemaTypes
}

func (T *TimeTimeChangeType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	// Parse using the TimeOnly layout: "15:04:05"
	t, err := time.Parse(time.TimeOnly, s)
	if err != nil {
		return err
	}

	T.TEXT = t

	return nil
}
