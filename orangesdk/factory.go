package orangesdk

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
)

// BuildRequest is a helper function to quickly build Request objects
func BuildRequest(rt RType, jsonBytes []byte) Request {
	logger := log.New(os.Stdout, "drcorangefactory: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(rt.String(), string(jsonBytes))
	switch rt {
	case RTS2M:
		var o DoS2M
		if err := json.Unmarshal(jsonBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTM2S:
		var o DoM2S
		if err := json.Unmarshal(jsonBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTM2M:
		var o DoM2M
		if err := json.Unmarshal(jsonBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTCHKBAL:
		var o TcheckBal
		if err := json.Unmarshal(jsonBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTCHKTXN:
		var o DoCheckTrans
		if err := json.Unmarshal(jsonBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTS2S:
		var o DoS2S
		if err := json.Unmarshal(jsonBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTSMS:
		var o SendSMS
		if err := json.Unmarshal(jsonBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTCBK:
		var o DoCallback
		if err := json.Unmarshal(jsonBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	default:
		return nil
	}
}

// BuildResponse is a helper function to quickly build Response objects
func BuildResponse(rt RType, xmlBytes []byte) Response {
	logger := log.New(os.Stdout, "drcorangefactory: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(rt.String(), string(xmlBytes))
	switch rt {
	case RTS2M:
		var o DoS2MResponse
		if err := xml.Unmarshal(xmlBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTM2S:
		var o DoM2SResponse
		if err := xml.Unmarshal(xmlBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTM2M:
		var o DoM2MResponse
		if err := xml.Unmarshal(xmlBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTCHKBAL:
		var o TcheckBalResponse
		if err := xml.Unmarshal(xmlBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTCHKTXN:
		var o DoCheckTransResponse
		if err := xml.Unmarshal(xmlBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTS2S:
		var o DoS2SResponse
		if err := xml.Unmarshal(xmlBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTSMS:
		var o SendSMSResponse
		if err := xml.Unmarshal(xmlBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	case RTCBK:
		var o DoCallbackResponse
		if err := xml.Unmarshal(xmlBytes, &o); err != nil {
			logger.Fatal(err)
		}
		return &o
	default:
		return nil
	}
}
