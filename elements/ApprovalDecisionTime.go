package elements

// The ApprovalDecisionTime element specifies the time at which the approval decision was made.
// https://learn.microsoft.com/en-us/exchange/client-developer/web-service-reference/approvaldecisiontime
import (
	"encoding/xml"
	"time"
)

type ApprovalDecisionTime struct {
	XMLName xml.Name
	TEXT    time.Time `xml:",chardata"`
}

func (A *ApprovalDecisionTime) SetForMarshal() {
	A.XMLName.Local = "ApprovalDecisionTime"
}

func (A *ApprovalDecisionTime) GetSchema() *Schema {
	return &SchemaTypes
}
