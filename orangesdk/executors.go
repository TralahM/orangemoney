package orangesdk

import (
	"bytes"
)

// RType defines a Request Type
type RType int

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
func (sdk *APIClient) S2M(r Request) Response {
	x := sdk.Execute(r)
	return x
}

// M2S performs a merchant to subscriber operation.
func (sdk *APIClient) M2S(r Request) DoM2SResponse {
	var res DoM2SResponse
	x := sdk.Execute(r)
	res = x.(DoM2SResponse)
	return res
}

// M2M performs a merchant to merchant operation.
func (sdk *APIClient) M2M(r Request) DoM2MResponse {
	var res DoM2MResponse
	x := sdk.Execute(r)
	res = x.(DoM2MResponse)
	return res
}

// Callback performs a webservice callbackurl on a transaction.
func (sdk *APIClient) Callback(r Request) DoCallbackResponse {
	var res DoCallbackResponse
	x := sdk.Execute(r)
	res = x.(DoCallbackResponse)
	return res
}

// CheckBal checks the balance.
func (sdk *APIClient) CheckBal(r Request) TcheckBalResponse {
	var res TcheckBalResponse
	x := sdk.Execute(r)
	res = x.(TcheckBalResponse)
	return res
}

// CheckTrans checks the status of a completed transaction.
func (sdk *APIClient) CheckTrans(r Request) DoCheckTransResponse {
	var res DoCheckTransResponse
	x := sdk.Execute(r)
	res = x.(DoCheckTransResponse)
	return res
}

// SendSMS send an sms message to a specified subscriber.
func (sdk *APIClient) SendSMS(r Request) SendSMSResponse {
	var res SendSMSResponse
	x := sdk.Execute(r)
	res = x.(SendSMSResponse)
	return res
}
