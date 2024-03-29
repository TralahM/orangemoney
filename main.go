package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/tralahm/orangemoney/docs"

	"github.com/tralahm/orangemoney/orangesdk"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	ServPort       = getEnv("PORT", "8080")
	c2bCallbackUrl = getEnv("CLIENT_C2B_CALLBACK_URL", "https://c2b_vodacash/")
	b2cCallbackUrl = getEnv("CLIENT_B2C_CALLBACK_URL", "https://b2c_vodacash/")
	redisUrl       = getEnv("REDIS_URL", "localhost:6379")
	remoteIPaddr   = getEnv("REMOTEIP", "")
	remotePortaddr = getEnv("REMOTEPORT", "8088")
	authToken      = getEnv("AUTHTOKEN", "")
	serverRootUrl  = getEnv("ROOTURL", "http://localhost:8080")
)

// DoM2SResponse struct
type DoM2SResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId" json:"PartnID"`
	ResultCode string `xml:"resultCode" json:"ResultCode"`
	ResultDesc string `xml:"resultDesc" json:"ResultDesc"`
	SystemID   string `xml:"systemId" json:"SystemID"`
	TransID    string `xml:"transId" json:"TransID"`
}

// DoS2MResponse struct
type DoS2MResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId" json:"PartnID"`
	ResultCode string `xml:"resultCode" json:"ResultCode"`
	ResultDesc string `xml:"resultDesc" json:"ResultDesc"`
	SystemID   string `xml:"systemId" json:"SystemID"`
	TransID    string `xml:"transId" json:"TransID"`
}

// DoS2SResponse struct
type DoS2SResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId" json:"PartnID"`
	ResultCode string `xml:"resultCode" json:"ResultCode"`
	ResultDesc string `xml:"resultDesc" json:"ResultDesc"`
	SystemID   string `xml:"systemId" json:"SystemID"`
	TransID    string `xml:"transId" json:"TransID"`
}

// DoM2MResponse struct
type DoM2MResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId" json:"PartnID"`
	ResultCode string `xml:"resultCode" json:"ResultCode"`
	ResultDesc string `xml:"resultDesc" json:"ResultDesc"`
	SystemID   string `xml:"systemId" json:"SystemID"`
	TransID    string `xml:"transId" json:"TransID"`
}

// TcheckBalResponse struct
type TcheckBalResponse struct {
	Text       string `xml:",chardata" json:"-"`
	Balance    string `xml:"balance" json:"Balance"`
	PartnID    string `xml:"partnId" json:"PartnID"`
	ResultCode string `xml:"resultCode" json:"ResultCode"`
	ResultDesc string `xml:"resultDesc" json:"ResultDesc"`
	TransID    string `xml:"transId" json:"TransID"`
}

// DoCheckTransResponse struct
type DoCheckTransResponse struct {
	Text       string `xml:",chardata" json:"-"`
	ResultCode string `xml:"resultCode" json:"ResultCode"`
	ResultDesc string `xml:"resultDesc" json:"ResultDesc"`
	TransID    string `xml:"transId" json:"TransID"`
}

// DoCallbackResponse struct
type DoCallbackResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId" json:"PartnID"`
	ResultCode string `xml:"resultCode" json:"ResultCode"`
	ResultDesc string `xml:"resultDesc" json:"ResultDesc"`
}

// SendSMSResponse struct
type SendSMSResponse struct {
	Text       string `xml:",chardata" json:"-"`
	Message    string `xml:"message" json:"Message"`
	Resultdesc string `xml:"resultdesc" json:"Resultdesc"`
	Resutcode  string `xml:"resutcode" json:"Resutcode"`
}

// SendSMS Request struct
type SendSMS struct {
	Text      string `xml:",chardata" json:"-"`
	Sender    string `xml:"sender"`
	Msisdn    string `xml:"msisdn"`
	Message   string `xml:"message"`
	Flash     string `xml:"flash"`
	PartnID   string `xml:"PartnId"`
	PartnName string `xml:"PartnName"`
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

// TcheckBal struct
type TcheckBal struct {
	Text       string `xml:",chardata" json:"-"`
	Subsmsisdn string `xml:"subsmsisdn"`
	Currency   string `xml:"currency"`
	PartnID    string `xml:"PartnId"`
	Transid    string `xml:"transid"`
}

// DoCheckTrans struct
type DoCheckTrans struct {
	Text      string `xml:",chardata" json:"-"`
	PartnID   string `xml:"PartnId"`
	Mermsisdn string `xml:"mermsisdn"`
	Transid   string `xml:"transid"`
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

// @title DRC Orange Money Proxy REST API
// @version 1.0
// @description This is a service for interacting with Oranges Money's DRC SOAP Integrated Payment Gateway.
// @termsOfService https://github.com/TralahM
// @contact.name API Support
// @contact.email briantralah@gmail.com
// @license.name GNU GENERAL PUBLIC LICENSE
// @license.url http://www.gnu.org/licenses/
// @host orange.betmondenge.com
// @BasePath /
func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*", "*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
			"X-Forwarded-IP",
			"X-Client-C2BCallbackURL",
			"X-Client-B2CCallbackURL",
			"X-Client-ThirdPartyReference",
			"Environment",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	handler := NewIpgHandler()
	// r.Get("/api/v1/health", handler.Health)
	r.Get("/swagger.json", handler.Swagger)
	// r.Get("/api/v1/ready", handler.Ready)
	r.Get("/", func(writer http.ResponseWriter, req *http.Request) {
		http.Redirect(writer, req, serverRootUrl+"/docs/index.html", http.StatusMovedPermanently)
	})
	r.Post("/api/v1/dos2m", handler.Dos2m)
	r.Post("/api/v1/dom2m", handler.Dom2m)
	r.Post("/api/v1/dom2s", handler.Dom2s)
	r.Post("/api/v1/docallback", handler.Docallback)
	r.Post("/api/v1/dochecktrans", handler.DoCheckTrans)
	r.Post("/api/v1/tcheckbal", handler.TcheckBal)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(serverRootUrl+"/swagger.json")))

	handler.logger.Printf("Server starting on 0.0.0.0:%s\n", ServPort)

	err := http.ListenAndServe("0.0.0.0:"+ServPort, r)
	if err != nil {
		handler.logger.Fatal(err)
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// IpgHandler struct
type IpgHandler struct {
	logger *log.Logger
	cli    *orangesdk.APIClient
}

// NewIpgHandler return an IpgHandler.
func NewIpgHandler() *IpgHandler {
	logger := log.New(os.Stdout, "drcorangeproxy: ", log.Ldate|log.Ltime|log.Lshortfile|log.LstdFlags)
	cli := orangesdk.NewClient(authToken, remoteIPaddr, remotePortaddr)
	return &IpgHandler{
		logger: logger,
		cli:    cli,
	}
}

// Swagger godoc
// @Summary Get API Swagger Definition
// @Tags internal
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /swagger.json [get]
func (ipg *IpgHandler) Swagger(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, req, "docs/swagger.json")
}

// respond makes the json response with payload as json format.
func (ipg *IpgHandler) respond(w http.ResponseWriter, status int, payload orangesdk.Response) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	var t bytes.Buffer
	data, err := payload.JSON(&t)
	if err != nil {
		ipg.logger.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(status)
	w.Write([]byte(data))
}

// Dos2m godoc
// @Summary Do an S2M transaction.
// @Description Perform a Subscriber to Merchant Transaction
// @Tags C2B
// @Accept json
// @Produce json
// @Param payload body DoS2M true "DoS2M"
// @Success 202 {object} DoS2MResponse
// @Router /api/v1/dos2m [post]
func (ipg *IpgHandler) Dos2m(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTS2M, body)
	xreq := ipg.cli.S2M(rq)
	ipg.respond(w, http.StatusAccepted, xreq.Dro())

}

// DoCallback godoc
// @Summary Accept a callback request
// @Description Accept a callback request
// @Tags Callbacks
// @Accept json
// @Produce json
// @Param payload body DoCallback true "DoCallback"
// @Success 202 {object} DoCallbackResponse
// @Router /api/v1/docallback [post]
func (ipg *IpgHandler) Docallback(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	ipg.logger.Println(string(body))
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTCBK, body)
	xreq := ipg.cli.Callback(rq)
	ipg.respond(w, http.StatusAccepted, xreq.Dro())

}

// DoCheckTrans godoc
// @Summary Check the status of a completed transaction.
// @Description Check the status of a completed transaction.
// @Tags CheckTransaction
// @Accept json
// @Produce json
// @Param payload body DoCheckTrans true "DoCheckTrans"
// @Success 202 {object} DoCheckTransResponse
// @Router /api/v1/dochecktrans [post]
func (ipg *IpgHandler) DoCheckTrans(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTCHKTXN, body)
	xreq := ipg.cli.CheckTrans(rq)
	ipg.respond(w, http.StatusAccepted, xreq.Dro())
}

// TcheckBal godoc
// @Summary Do a Check Balance transaction.
// @Description Check the balance of a specified account.
// @Tags Balance
// @Accept json
// @Produce json
// @Param payload body TcheckBal true "TcheckBal"
// @Success 202 {object} TcheckBalResponse
// @Router /api/v1/tcheckbal [post]
func (ipg *IpgHandler) TcheckBal(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTCHKBAL, body)
	xreq := ipg.cli.CheckBal(rq)
	ipg.respond(w, http.StatusAccepted, xreq.Dro())

}

// Dom2s godoc
// @Summary Do M2S Transaction.
// @Description Perform a Merchant to Subscriber transaction.
// @Tags B2C
// @Accept json
// @Produce json
// @Param payload body DoM2S true "DoM2S"
// @Success 202 {object} DoM2SResponse
// @Router /api/v1/dom2s [post]
func (ipg *IpgHandler) Dom2s(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTM2S, body)
	xreq := ipg.cli.M2S(rq)
	ipg.respond(w, http.StatusAccepted, xreq.Dro())

}

// Dom2m godoc
// @Summary Do M2M Transaction
// @Description Perform a Merchant to Merchant Transaction.
// @Tags B2B
// @Accept json
// @Produce json
// @Param payload body DoM2M true "DoM2M"
// @Success 202 {object} DoM2MResponse
// @Router /api/v1/dom2m [post]
func (ipg *IpgHandler) Dom2m(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTM2M, body)
	xreq := ipg.cli.M2M(rq)
	ipg.respond(w, http.StatusAccepted, xreq.Dro())

}

// SendSMS godoc
// @Summary Send an SMS
// @Description Send an SMS to a Subscriber.
// @Tags SendSMS
// @Accept json
// @Produce json
// @Param payload body SendSMS true "SendSMS"
// @Success 202 {object} SendSMSResponse
// @Router /api/v1/sendsms [post]
func (ipg *IpgHandler) SendSMS(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTSMS, body)
	xreq := ipg.cli.SendSMS(rq)
	ipg.respond(w, http.StatusAccepted, xreq.Dro())

}

// respondError makes the error response with message as json format.
func (ipg *IpgHandler) respondError(w http.ResponseWriter, code int, message string) {
	ipg.respondJSON(w, code, map[string]string{"error": message})
}

// respondJSON makes the json response with payload as json format.
func (ipg *IpgHandler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
