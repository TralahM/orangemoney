package orangesdk

import (
	"bytes"
	"encoding/xml"
)

// RType defines a Request Type
type RType int

// Constants
const (
	RTS2M RType = iota
	RTM2S
	RTM2M
	RTCHKBAL
	RTCHKTXN
	RTS2S
	RTSMS
	RTCBK
)

// String satisfies the Stringer interface on RType
func (t RType) String() string {
	switch t {
	case RTS2M:
		return "Subscriber To Merchant"
	case RTM2S:
		return "Merchant To Subscriber"
	case RTM2M:
		return "Merchant to Merchant"
	case RTCHKBAL:
		return "Check Balance"
	case RTCHKTXN:
		return "Check Transaction Status"
	case RTS2S:
		return "Subscriber To Subscriber"
	case RTSMS:
		return "Send SMS"
	case RTCBK:
		return "Do Callback"
	default:
		return "Unknown Request"
	}
}

// Executor defines the Executor interface
type Executor interface {
	Execute(Request) Response
}

// ExecutorFunc is an Executor
type ExecutorFunc func(Request) Response

// Execute satisfies the Executor interface
func (f ExecutorFunc) Execute(r Request) Response {
	return f(r)
}

// S2M,M2S,M2M,Callback,CheckBal,CheckTrans

// Execute satisfies the Executor interface on @APIClient
func (sdk *APIClient) Execute(r Request) Response {
	var data bytes.Buffer
	r.XML(&data)
	xmlBytes, err := sdk.Post(&data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	return BuildResponse(r.Type(), xmlBytes)
}

// S2M performs a subscriber to merchant operation.
func (sdk *APIClient) S2M(r Request) *DoS2MResponse {
	var data bytes.Buffer
	r.XML(&data)
	xmlBytes, err := sdk.Post(&data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	var o DoS2MResponse
	if err := xml.Unmarshal(xmlBytes, &o); err != nil {
		sdk.logger.Fatal(err)
	}
	sdk.logger.Printf("%s\n", string(xmlBytes))
	sdk.logger.Printf("%#v\n", o)
	return &o
}

// M2S performs a merchant to subscriber operation.
func (sdk *APIClient) M2S(r Request) *DoM2SResponse {
	var data bytes.Buffer
	r.XML(&data)
	xmlBytes, err := sdk.Post(&data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	var o DoM2SResponse
	if err := xml.Unmarshal(xmlBytes, &o); err != nil {
		sdk.logger.Fatal(err)
	}
	sdk.logger.Printf("%s\n", string(xmlBytes))
	sdk.logger.Printf("%#v\n", o)
	return &o
}

// M2M performs a merchant to merchant operation.
func (sdk *APIClient) M2M(r Request) *DoM2MResponse {
	var data bytes.Buffer
	r.XML(&data)
	xmlBytes, err := sdk.Post(&data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	var o DoM2MResponse
	if err := xml.Unmarshal(xmlBytes, &o); err != nil {
		sdk.logger.Fatal(err)
	}
	sdk.logger.Printf("%s\n", string(xmlBytes))
	sdk.logger.Printf("%#v\n", o)
	return &o
}

// Callback performs a webservice callbackurl on a transaction.
func (sdk *APIClient) Callback(r Request) *DoCallbackResponse {
	var data bytes.Buffer
	r.XML(&data)
	xmlBytes, err := sdk.Post(&data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	var o DoCallbackResponse
	if err := xml.Unmarshal(xmlBytes, &o); err != nil {
		sdk.logger.Fatal(err)
	}
	sdk.logger.Printf("%s\n", string(xmlBytes))
	sdk.logger.Printf("%#v\n", o)
	return &o
}

// CheckBal checks the balance.
func (sdk *APIClient) CheckBal(r Request) *TcheckBalResponse {
	var data bytes.Buffer
	r.XML(&data)
	xmlBytes, err := sdk.Post(&data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	var o TcheckBalResponse
	if err := xml.Unmarshal(xmlBytes, &o); err != nil {
		sdk.logger.Fatal(err)
	}
	sdk.logger.Printf("%s\n", string(xmlBytes))
	sdk.logger.Printf("%#v\n", o)
	return &o
}

// CheckTrans checks the status of a completed transaction.
func (sdk *APIClient) CheckTrans(r Request) *DoCheckTransResponse {
	var data bytes.Buffer
	r.XML(&data)
	xmlBytes, err := sdk.Post(&data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	var o DoCheckTransResponse
	if err := xml.Unmarshal(xmlBytes, &o); err != nil {
		sdk.logger.Fatal(err)
	}
	sdk.logger.Printf("%s\n", string(xmlBytes))
	sdk.logger.Printf("%#v\n", o)
	return &o
}

// SendSMS send an sms message to a specified subscriber.
func (sdk *APIClient) SendSMS(r Request) *SendSMSResponse {
	var data bytes.Buffer
	r.XML(&data)
	xmlBytes, err := sdk.Post(&data)
	if err != nil {
		sdk.logger.Fatalln(err)
	}
	var o SendSMSResponse
	if err := xml.Unmarshal(xmlBytes, &o); err != nil {
		sdk.logger.Fatal(err)
	}
	sdk.logger.Printf("%s\n", string(xmlBytes))
	sdk.logger.Printf("%#v\n", o)
	return &o
}
