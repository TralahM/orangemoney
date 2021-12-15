package orangesdk

import (
	"encoding/json"
	"encoding/xml"
	"html/template"
	"io"
	"log"
	"os"
)

// Request interface is a Data Transfer Object interface
type Request interface {
	XML(io.Writer) error
	Type() RType
}

// Response interface is a Data Return Object interface
type Response interface {
	JSON(io.Writer) error
}

// DoM2S payload Merchant To Subscriber
type DoM2S struct {
	Text       string `xml:",chardata" json:"-"`
	Subsmsisdn string `xml:"subsmsisdn"`
	PartnID    string `xml:"PartnId"`
	Mermsisdn  string `xml:"mermsisdn"`
	Transid    string `xml:"transid"`
	Currency   string `xml:"currency"`
	Amount     string `xml:"amount"`
}

// XML Writes the DoM2S struct info into xml
func (data *DoM2S) XML(wr io.Writer) error {
	logger := log.New(os.Stdout, "drcorangedtos: ", log.Ldate|log.Ltime|log.Lshortfile)
	t, err := template.New("dom2s").Parse(m2sreq)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	err = t.Execute(wr, data)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

// m2SREQUEST struct
type m2SREQUEST struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	Soapenv string   `xml:"soapenv,attr"`
	Ser     string   `xml:"ser,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text  string `xml:",chardata" json:"-"`
		DoM2S struct {
			Text       string `xml:",chardata" json:"-"`
			Subsmsisdn string `xml:"subsmsisdn"`
			PartnID    string `xml:"PartnId"`
			Mermsisdn  string `xml:"mermsisdn"`
			Transid    string `xml:"transid"`
			Currency   string `xml:"currency"`
			Amount     string `xml:"amount"`
		} `xml:"doM2S"`
	} `xml:"Body"`
}

// Dto Returns the data transfer object of the request
func (r *m2SREQUEST) Dto() Request {
	var o DoM2S = r.Body.DoM2S
	return &o
}

// DoM2SResponse struct
type DoM2SResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	SystemID   string `xml:"systemId"`
	TransID    string `xml:"transId"`
}

// JSON writes to the writer the json string of the object.
func (r DoM2SResponse) JSON(wr io.Writer) error {
	out, err := json.MarshalIndent(r, " ", "  ")
	if err != nil {
		return err
	}
	wr.Write(out)
	return nil
}

// JSON writes to the writer the json string of the object.
func (r TcheckBalResponse) JSON(wr io.Writer) error {
	out, err := json.MarshalIndent(r, " ", "  ")
	if err != nil {
		return err
	}
	wr.Write(out)
	return nil
}

// m2sRESPONSE struct
type m2sRESPONSE struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text          string `xml:",chardata" json:"-"`
		DoM2SResponse struct {
			Text   string `xml:",chardata" json:"-"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text       string `xml:",chardata" json:"-"`
				PartnID    string `xml:"partnId"`
				ResultCode string `xml:"resultCode"`
				ResultDesc string `xml:"resultDesc"`
				SystemID   string `xml:"systemId"`
				TransID    string `xml:"transId"`
			} `xml:"return"`
		} `xml:"doM2SResponse"`
	} `xml:"Body"`
}

// DoS2M Request struct
type DoS2M struct {
	Text        string `xml:",chardata" json:"-"`
	Subsmsisdn  string `xml:"subsmsisdn"`
	PartnID     string `xml:"PartnId"`
	Mermsisdn   string `xml:"mermsisdn"`
	Transid     string `xml:"transid"`
	Currency    string `xml:"currency"`
	Amount      string `xml:"amount"`
	Callbackurl string `xml:"callbackurl"`
	MessageS2m  string `xml:"message_s2m"`
}

// XML Writes the DoS2M struct info into xml
func (data *DoS2M) XML(wr io.Writer) error {
	logger := log.New(os.Stdout, "drcorangedtos: ", log.Ldate|log.Ltime|log.Lshortfile)
	t, err := template.New("dos2m").Parse(s2mreq)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	err = t.Execute(wr, data)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

// s2MREQUEST struct
type s2MREQUEST struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	Soapenv string   `xml:"soapenv,attr"`
	Ser     string   `xml:"ser,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text  string `xml:",chardata" json:"-"`
		DoS2M struct {
			Text        string `xml:",chardata" json:"-"`
			Subsmsisdn  string `xml:"subsmsisdn"`
			PartnID     string `xml:"PartnId"`
			Mermsisdn   string `xml:"mermsisdn"`
			Transid     string `xml:"transid"`
			Currency    string `xml:"currency"`
			Amount      string `xml:"amount"`
			Callbackurl string `xml:"callbackurl"`
			MessageS2m  string `xml:"message_s2m"`
		} `xml:"doS2M"`
	} `xml:"Body"`
}

// Dto Returns the data transfer object of the request
func (r *s2MREQUEST) Dto() Request {
	var o DoS2M = r.Body.DoS2M
	return &o
}

// DoS2MResponse struct
type DoS2MResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	SystemID   string `xml:"systemId"`
	TransID    string `xml:"transId"`
}

// JSON writes to the writer the json string of the object.
func (r DoS2MResponse) JSON(wr io.Writer) error {
	out, err := json.MarshalIndent(r, " ", "  ")
	if err != nil {
		return err
	}
	wr.Write(out)
	return nil
}

// JSON writes to the writer the json string of the object.
func (r DoM2MResponse) JSON(wr io.Writer) error {
	out, err := json.MarshalIndent(r, " ", "  ")
	if err != nil {
		return err
	}
	wr.Write(out)
	return nil
}

// JSON writes to the writer the json string of the object.
func (r DoS2SResponse) JSON(wr io.Writer) error {
	out, err := json.MarshalIndent(r, " ", "  ")
	if err != nil {
		return err
	}
	wr.Write(out)
	return nil
}

// JSON writes to the writer the json string of the object.
func (r DoCallbackResponse) JSON(wr io.Writer) error {
	out, err := json.MarshalIndent(r, " ", "  ")
	if err != nil {
		return err
	}
	wr.Write(out)
	return nil
}

// JSON writes to the writer the json string of the object.
func (r DoCheckTransResponse) JSON(wr io.Writer) error {
	out, err := json.MarshalIndent(r, " ", "  ")
	if err != nil {
		return err
	}
	wr.Write(out)
	return nil
}

// JSON writes to the writer the json string of the object.
func (r SendSMSResponse) JSON(wr io.Writer) error {
	out, err := json.MarshalIndent(r, " ", "  ")
	if err != nil {
		return err
	}
	wr.Write(out)
	return nil
}

// s2mRESPONSE struct
type s2mRESPONSE struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text          string `xml:",chardata" json:"-"`
		DoS2MResponse struct {
			Text   string `xml:",chardata" json:"-"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text       string `xml:",chardata" json:"-"`
				PartnID    string `xml:"partnId"`
				ResultCode string `xml:"resultCode"`
				ResultDesc string `xml:"resultDesc"`
				SystemID   string `xml:"systemId"`
				TransID    string `xml:"transId"`
			} `xml:"return"`
		} `xml:"doS2MResponse"`
	} `xml:"Body"`
}

// DoS2S struct
type DoS2S struct {
	Text        string `xml:",chardata" json:"-"`
	Subsmsisdn  string `xml:"subsmsisdn"`
	PartnID     string `xml:"PartnId"`
	Subsmsisdn2 string `xml:"subsmsisdn2"`
	Transid     string `xml:"transid"`
	Currency    string `xml:"currency"`
	Amount      string `xml:"amount"`
	Callbackurl string `xml:"callbackurl"`
	MessageS2s  string `xml:"message_s2s"`
}

// XML Writes the DoS2S struct info into xml
func (data *DoS2S) XML(wr io.Writer) error {
	logger := log.New(os.Stdout, "drcorangedtos: ", log.Ldate|log.Ltime|log.Lshortfile)
	t, err := template.New("dos2s").Parse(s2sreq)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	err = t.Execute(wr, data)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

// s2SREQUEST struct
type s2SREQUEST struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	Soapenv string   `xml:"soapenv,attr"`
	Ser     string   `xml:"ser,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text  string `xml:",chardata" json:"-"`
		DoS2S struct {
			Text        string `xml:",chardata" json:"-"`
			Subsmsisdn  string `xml:"subsmsisdn"`
			PartnID     string `xml:"PartnId"`
			Subsmsisdn2 string `xml:"subsmsisdn2"`
			Transid     string `xml:"transid"`
			Currency    string `xml:"currency"`
			Amount      string `xml:"amount"`
			Callbackurl string `xml:"callbackurl"`
			MessageS2s  string `xml:"message_s2s"`
		} `xml:"doS2S"`
	} `xml:"Body"`
}

// Dto Returns the data transfer object of the request
func (r *s2SREQUEST) Dto() Request {
	var o DoS2S = r.Body.DoS2S
	return &o
}

// DoS2SResponse struct
type DoS2SResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	SystemID   string `xml:"systemId"`
	TransID    string `xml:"transId"`
}

// s2sRESPONSE struct
type s2sRESPONSE struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text          string `xml:",chardata" json:"-"`
		DoS2SResponse struct {
			Text   string `xml:",chardata" json:"-"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text       string `xml:",chardata" json:"-"`
				PartnID    string `xml:"partnId"`
				ResultCode string `xml:"resultCode"`
				ResultDesc string `xml:"resultDesc"`
				SystemID   string `xml:"systemId"`
				TransID    string `xml:"transId"`
			} `xml:"return"`
		} `xml:"doS2SResponse"`
	} `xml:"Body"`
}

// DoM2M struct
type DoM2M struct {
	Text       string `xml:",chardata" json:"-"`
	Mermsisdn1 string `xml:"mermsisdn1"`
	PartnID    string `xml:"PartnId"`
	Mermsisdn2 string `xml:"mermsisdn2"`
	Transid    string `xml:"transid"`
	Currency   string `xml:"currency"`
	Amount     string `xml:"amount"`
}

// XML Writes the DoM2M struct info into xml
func (data *DoM2M) XML(wr io.Writer) error {
	logger := log.New(os.Stdout, "drcorangedtos: ", log.Ldate|log.Ltime|log.Lshortfile)
	t, err := template.New("dom2m").Parse(m2mreq)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	err = t.Execute(wr, data)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

// m2MREQUEST struct
type m2MREQUEST struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	Soapenv string   `xml:"soapenv,attr"`
	Ser     string   `xml:"ser,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text  string `xml:",chardata" json:"-"`
		DoM2M struct {
			Text       string `xml:",chardata" json:"-"`
			Mermsisdn1 string `xml:"mermsisdn1"`
			PartnID    string `xml:"PartnId"`
			Mermsisdn2 string `xml:"mermsisdn2"`
			Transid    string `xml:"transid"`
			Currency   string `xml:"currency"`
			Amount     string `xml:"amount"`
		} `xml:"doM2M"`
	} `xml:"Body"`
}

// Dto Returns the data transfer object of the request
func (r *m2MREQUEST) Dto() Request {
	var o DoM2M = r.Body.DoM2M
	return &o
}

// DoM2MResponse struct
type DoM2MResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	SystemID   string `xml:"systemId"`
	TransID    string `xml:"transId"`
}

// m2mRESPONSE struct
type m2mRESPONSE struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text          string `xml:",chardata" json:"-"`
		DoM2MResponse struct {
			Text   string `xml:",chardata" json:"-"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text       string `xml:",chardata" json:"-"`
				PartnID    string `xml:"partnId"`
				ResultCode string `xml:"resultCode"`
				ResultDesc string `xml:"resultDesc"`
				SystemID   string `xml:"systemId"`
				TransID    string `xml:"transId"`
			} `xml:"return"`
		} `xml:"doM2MResponse"`
	} `xml:"Body"`
}

// TcheckBal struct
type TcheckBal struct {
	Text       string `xml:",chardata" json:"-"`
	Subsmsisdn string `xml:"subsmsisdn"`
	Currency   string `xml:"currency"`
	PartnID    string `xml:"PartnId"`
	Transid    string `xml:"transid"`
}

// XML Writes the TcheckBal struct info into xml
func (data *TcheckBal) XML(wr io.Writer) error {
	logger := log.New(os.Stdout, "drcorangedtos: ", log.Ldate|log.Ltime|log.Lshortfile)
	t, err := template.New("tcheckbal").Parse(tcheckbalreq)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	err = t.Execute(wr, data)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

// cHECKBALANCEREQUEST struct
type cHECKBALANCEREQUEST struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	Soapenv string   `xml:"soapenv,attr"`
	Ser     string   `xml:"ser,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text      string `xml:",chardata" json:"-"`
		TcheckBal struct {
			Text       string `xml:",chardata" json:"-"`
			Subsmsisdn string `xml:"subsmsisdn"`
			Currency   string `xml:"currency"`
			PartnID    string `xml:"PartnId"`
			Transid    string `xml:"transid"`
		} `xml:"TcheckBal"`
	} `xml:"Body"`
}

// Dto Returns the data transfer object of the request
func (r *cHECKBALANCEREQUEST) Dto() Request {
	var o TcheckBal = r.Body.TcheckBal
	return &o
}

// TcheckBalResponse struct
type TcheckBalResponse struct {
	Text       string `xml:",chardata" json:"-"`
	Balance    string `xml:"balance"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	TransID    string `xml:"transId"`
}

// checkbalanceRESPONSE struct
type checkbalanceRESPONSE struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text              string `xml:",chardata" json:"-"`
		TcheckBalResponse struct {
			Text   string `xml:",chardata" json:"-"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text       string `xml:",chardata" json:"-"`
				Balance    string `xml:"balance"`
				PartnID    string `xml:"partnId"`
				ResultCode string `xml:"resultCode"`
				ResultDesc string `xml:"resultDesc"`
				TransID    string `xml:"transId"`
			} `xml:"return"`
		} `xml:"TcheckBalResponse"`
	} `xml:"Body"`
}

// DoCheckTrans struct
type DoCheckTrans struct {
	Text      string `xml:",chardata" json:"-"`
	PartnID   string `xml:"PartnId"`
	Mermsisdn string `xml:"mermsisdn"`
	Transid   string `xml:"transid"`
}

// XML Writes the TcheckBal struct info into xml
func (data *DoCheckTrans) XML(wr io.Writer) error {
	logger := log.New(os.Stdout, "drcorangedtos: ", log.Ldate|log.Ltime|log.Lshortfile)
	t, err := template.New("dochecktrans").Parse(checktransreq)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	err = t.Execute(wr, data)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

// cHECKTRANSACTIONSTATUSREQUEST struct
type cHECKTRANSACTIONSTATUSREQUEST struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	Soapenv string   `xml:"soapenv,attr"`
	Ser     string   `xml:"ser,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text         string `xml:",chardata" json:"-"`
		DoCheckTrans struct {
			Text      string `xml:",chardata" json:"-"`
			PartnID   string `xml:"PartnId"`
			Mermsisdn string `xml:"mermsisdn"`
			Transid   string `xml:"transid"`
		} `xml:"doCheckTrans"`
	} `xml:"Body"`
}

// Dto Returns the data transfer object of the request
func (r *cHECKTRANSACTIONSTATUSREQUEST) Dto() Request {
	var o DoCheckTrans = r.Body.DoCheckTrans
	return &o
}

// DoCheckTransResponse struct
type DoCheckTransResponse struct {
	Text       string `xml:",chardata" json:"-"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	TransID    string `xml:"transId"`
}

// checktransactionstatusRESPONSE struct
type checktransactionstatusRESPONSE struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text                 string `xml:",chardata" json:"-"`
		DoCheckTransResponse struct {
			Text   string `xml:",chardata" json:"-"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text       string `xml:",chardata" json:"-"`
				ResultCode string `xml:"resultCode"`
				ResultDesc string `xml:"resultDesc"`
				TransID    string `xml:"transId"`
			} `xml:"return"`
		} `xml:"doCheckTransResponse"`
	} `xml:"Body"`
}

// DoCallback struct
type DoCallback struct {
	Text       string `xml:",chardata" json:"-"`
	Subsmsisdn string `xml:"subsmsisdn"`
	PartnID    string `xml:"PartnId"`
	Mermsisdn  string `xml:"mermsisdn"`
	Transid    string `xml:"transid"`
	Currency   string `xml:"currency"`
	Amount     string `xml:"amount"`
	ResultCode string `xml:"ResultCode"`
	ResultDesc string `xml:"ResultDesc"`
}

// XML Writes the TcheckBal struct info into xml
func (data *DoCallback) XML(wr io.Writer) error {
	logger := log.New(os.Stdout, "drcorangedtos: ", log.Ldate|log.Ltime|log.Lshortfile)
	t, err := template.New("docallback").Parse(callbackreq)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	err = t.Execute(wr, data)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

// cALLBACKURLREQUEST struct
type cALLBACKURLREQUEST struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	Soapenv string   `xml:"soapenv,attr"`
	Ser     string   `xml:"ser,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text       string `xml:",chardata" json:"-"`
		DoCallback struct {
			Text       string `xml:",chardata" json:"-"`
			Subsmsisdn string `xml:"subsmsisdn"`
			PartnID    string `xml:"PartnId"`
			Mermsisdn  string `xml:"mermsisdn"`
			Transid    string `xml:"transid"`
			Currency   string `xml:"currency"`
			Amount     string `xml:"amount"`
			ResultCode string `xml:"ResultCode"`
			ResultDesc string `xml:"ResultDesc"`
		} `xml:"doCallback"`
	} `xml:"Body"`
}

// Dto Returns the data transfer object of the request
func (r *cALLBACKURLREQUEST) Dto() Request {
	var o DoCallback = r.Body.DoCallback
	return &o
}

// DoCallbackResponse struct
type DoCallbackResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
}

// callbackurlRESPONSE struct
type callbackurlRESPONSE struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text               string `xml:",chardata" json:"-"`
		DoCallbackResponse struct {
			Text   string `xml:",chardata" json:"-"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text       string `xml:",chardata" json:"-"`
				PartnID    string `xml:"partnId"`
				ResultCode string `xml:"resultCode"`
				ResultDesc string `xml:"resultDesc"`
			} `xml:"return"`
		} `xml:"doCallbackResponse"`
	} `xml:"Body"`
}

// SendSMS struct dto
type SendSMS struct {
	Text      string `xml:",chardata" json:"-"`
	Sender    string `xml:"sender"`
	Msisdn    string `xml:"msisdn"`
	Message   string `xml:"message"`
	Flash     string `xml:"flash"`
	PartnID   string `xml:"PartnId"`
	PartnName string `xml:"PartnName"`
}

// XML Writes the TcheckBal struct info into xml
func (data *SendSMS) XML(wr io.Writer) error {
	logger := log.New(os.Stdout, "drcorangedtos: ", log.Ldate|log.Ltime|log.Lshortfile)
	t, err := template.New("sendsms").Parse(sendsmsreq)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	err = t.Execute(wr, data)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

// sENDSMSREQUEST struct
type sENDSMSREQUEST struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	Soapenv string   `xml:"soapenv,attr"`
	Ser     string   `xml:"ser,attr"`
	Header  string   `xml:"Header"`
	Body    struct {
		Text    string `xml:",chardata" json:"-"`
		SendSMS struct {
			Text      string `xml:",chardata" json:"-"`
			Sender    string `xml:"sender"`
			Msisdn    string `xml:"msisdn"`
			Message   string `xml:"message"`
			Flash     string `xml:"flash"`
			PartnID   string `xml:"PartnId"`
			PartnName string `xml:"PartnName"`
		} `xml:"sendSMS"`
	} `xml:"Body"`
}

// Dto Returns the data transfer object of the request
func (r *sENDSMSREQUEST) Dto() Request {
	var o SendSMS = r.Body.SendSMS
	return &o
}

// SendSMSResponse struct
type SendSMSResponse struct {
	Text       string `xml:",chardata" json:"-"`
	Message    string `xml:"message"`
	Resultdesc string `xml:"resultdesc"`
	Resutcode  string `xml:"resutcode"`
}

// sendsmsRESPONSE struct
type sendsmsRESPONSE struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata" json:"-"`
	S       string   `xml:"S,attr"`
	Body    struct {
		Text            string `xml:",chardata" json:"-"`
		SendSMSResponse struct {
			Text   string `xml:",chardata" json:"-"`
			Ns2    string `xml:"ns2,attr"`
			Return struct {
				Text       string `xml:",chardata" json:"-"`
				Message    string `xml:"message"`
				Resultdesc string `xml:"resultdesc"`
				Resutcode  string `xml:"resutcode"`
			} `xml:"return"`
		} `xml:"sendSMSResponse"`
	} `xml:"Body"`
}

// Dro returns an object satisfying the Response
func (r *s2mRESPONSE) Dro() Response {
	var o DoS2MResponse = r.Body.DoS2MResponse.Return
	return &o
}

// Dro returns an object satisfying the Response
func (r *callbackurlRESPONSE) Dro() Response {
	var o DoCallbackResponse = r.Body.DoCallbackResponse.Return
	return &o
}

// Dro returns an object satisfying the Response
func (r *sendsmsRESPONSE) Dro() Response {
	var o SendSMSResponse = r.Body.SendSMSResponse.Return
	return &o
}

// Dro returns an object satisfying the Response
func (r *checktransactionstatusRESPONSE) Dro() Response {
	var o DoCheckTransResponse = r.Body.DoCheckTransResponse.Return
	return &o
}

// Dro returns an object satisfying the Response
func (r *checkbalanceRESPONSE) Dro() Response {
	var o TcheckBalResponse = r.Body.TcheckBalResponse.Return
	return &o
}

// Dro returns an object satisfying the Response
func (r *m2sRESPONSE) Dro() Response {
	var o DoM2SResponse = r.Body.DoM2SResponse.Return
	return &o
}

// Dro returns an object satisfying the Response
func (r *s2sRESPONSE) Dro() Response {
	var o DoS2SResponse = r.Body.DoS2SResponse.Return
	return &o
}

// Dro returns an object satisfying the Response
func (r *m2mRESPONSE) Dro() Response {
	var o DoM2MResponse = r.Body.DoM2MResponse.Return
	return &o
}

// Type Reports the type of Request it is
func (r DoS2M) Type() RType {
	return RTS2M
}

// Type Reports the type of Request it is
func (r DoS2S) Type() RType {
	return RTS2S
}

// Type Reports the type of Request it is
func (r DoM2M) Type() RType {
	return RTM2M
}

// Type Reports the type of Request it is
func (r DoM2S) Type() RType {
	return RTM2S
}

// Type Reports the type of Request it is
func (r DoCallback) Type() RType {
	return RTCBK
}

// Type Reports the type of Request it is
func (r TcheckBal) Type() RType {
	return RTCHKBAL
}

// Type Reports the type of Request it is
func (r DoCheckTrans) Type() RType {
	return RTCHKTXN
}

// Type Reports the type of Request it is
func (r SendSMS) Type() RType {
	return RTSMS
}
