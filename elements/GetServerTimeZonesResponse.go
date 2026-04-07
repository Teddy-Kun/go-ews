package elements

// The GetServerTimeZonesResponse element defines a response to a GetServerTimeZones operation request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/getservertimezonesresponse
import "encoding/xml"

type GetServerTimeZonesResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (G *GetServerTimeZonesResponse) SetForMarshal() {
	G.XMLName.Local = "GetServerTimeZonesResponse"
}

func (G *GetServerTimeZonesResponse) GetSchema() *Schema {
	return &SchemaMessages
}
