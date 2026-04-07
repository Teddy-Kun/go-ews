package elements

// The FindMailboxStatisticsByKeywordsResponse element specifies the response to a FindMailboxStatisticsByKeywords request.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/findmailboxstatisticsbykeywordsresponse
import "encoding/xml"

type FindMailboxStatisticsByKeywordsResponse struct {
	XMLName xml.Name

	// The ResponseMessages element contains the response messages for an Exchange Web Services request.
	ResponseMessages *ResponseMessages `xml:"ResponseMessages"`
}

func (F *FindMailboxStatisticsByKeywordsResponse) SetForMarshal() {
	F.XMLName.Local = "FindMailboxStatisticsByKeywordsResponse"
}

func (F *FindMailboxStatisticsByKeywordsResponse) GetSchema() *Schema {
	return &SchemaMessages
}
