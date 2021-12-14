package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "./docs"

	"./orangesdk"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	ServPort       = getEnv("PORT", "8000")
	c2bCallbackUrl = getEnv("CLIENT_C2B_CALLBACK_URL", "https://c2b_vodacash/")
	b2cCallbackUrl = getEnv("CLIENT_B2C_CALLBACK_URL", "https://b2c_vodacash/")
	redisUrl       = getEnv("REDIS_URL", "localhost:6379")
	internalC2B    = "https://ipg.betmondenge.com/api/v1/c2b_callback"
	internalB2C    = "https://ipg.betmondenge.com/api/v1/b2c_callback"
)

// DoM2SResponse struct
type DoM2SResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	SystemID   string `xml:"systemId"`
	TransID    string `xml:"transId"`
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

// DoS2SResponse struct
type DoS2SResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	SystemID   string `xml:"systemId"`
	TransID    string `xml:"transId"`
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

// TcheckBalResponse struct
type TcheckBalResponse struct {
	Text       string `xml:",chardata" json:"-"`
	Balance    string `xml:"balance"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	TransID    string `xml:"transId"`
}

// DoCheckTransResponse struct
type DoCheckTransResponse struct {
	Text       string `xml:",chardata" json:"-"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
	TransID    string `xml:"transId"`
}

// DoCallbackResponse struct
type DoCallbackResponse struct {
	Text       string `xml:",chardata" json:"-"`
	PartnID    string `xml:"partnId"`
	ResultCode string `xml:"resultCode"`
	ResultDesc string `xml:"resultDesc"`
}

// SendSMSResponse struct
type SendSMSResponse struct {
	Text       string `xml:",chardata" json:"-"`
	Message    string `xml:"message"`
	Resultdesc string `xml:"resultdesc"`
	Resutcode  string `xml:"resutcode"`
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
	r.Post("/api/v1/dos2m", handler.Dos2m)
	r.Post("/api/v1/dom2m", handler.Dom2m)
	r.Post("/api/v1/dom2s", handler.Dom2s)
	// r.Post("/api/v2/login", handler.Login)
	// r.Post("/api/v1/c2b", handler.C2B)
	// r.Post("/api/v1/b2c", handler.B2C)
	// r.Post("/api/v1/vodacash_c2b_callback", handler.C2BCallback)
	// r.Post("/api/v1/vodacash_b2c_callback", handler.B2CCallback)
	// r.Post("/api/v1/c2b_callback", handler.C2BCallback)
	// r.Post("/api/v1/b2c_callback", handler.B2CCallback)
	// r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("https://raw.githubusercontent.com/tralahm/drcmpesaproxy/master/docs/swagger.json"))) // url pointing to api definition
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://0.0.0.0:"+ServPort+"/swagger.json")))

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
	logger := log.New(os.Stdout, "drcorangeproxy: ", log.Ldate|log.Ltime|log.Lshortfile)
	cli := orangesdk.NewClient("Bearer c0d9acb4-9e88-5030-83e2-456916dc25dr", "1182", "GLORESVENT", "0896643349", "41.77.222.184", "8088")
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

// respondJSON makes the json response with payload as json format.
func (ipg *IpgHandler) respond(w http.ResponseWriter, status int, payload orangesdk.Response) {
	err := payload.JSON(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf8")
	w.WriteHeader(status)
}

// Dos2m godoc
// @Summary Do an S2M transaction.
// @Description Perform a Subscriber to Merchant Transaction
// @Tags login
// @Accept json
// @Produce json
// @Param payload body DoS2M true "DoS2M"
// @Success 201 {object} DoS2MResponse
// @Router /api/v1/dos2m [post]
func (ipg *IpgHandler) Dos2m(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTS2M, body)
	var xreq orangesdk.Response = ipg.cli.S2M(rq)
	ipg.respond(w, 202, xreq)

}

// Dom2s godoc
// @Summary Do M2S Transaction.
// @Description Perform a Merchant to Subscriber transaction.
// @Tags login
// @Accept json
// @Produce json
// @Param payload body DoM2S true "DoM2S"
// @Success 201 {object} DoM2SResponse
// @Router /api/v1/dom2s [post]
func (ipg *IpgHandler) Dom2s(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTM2S, body)
	var xreq orangesdk.Response = ipg.cli.M2S(rq)
	ipg.respond(w, 202, xreq)

}

// Dom2m godoc
// @Summary Do M2M Transaction
// @Description Perform a Merchant to Merchant Transaction.
// @Tags login
// @Accept json
// @Produce json
// @Param payload body DoM2M true "DoM2M"
// @Success 201 {object} DoM2MResponse
// @Router /api/v1/dom2m [post]
func (ipg *IpgHandler) Dom2m(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ipg.logger.Printf("Error reading body: %v\n", err)
		ipg.respondError(w, http.StatusBadRequest, string([]byte(err.Error())))
		return
	}
	rq := orangesdk.BuildRequest(orangesdk.RTM2M, body)
	var xreq orangesdk.Response = ipg.cli.M2M(rq)
	ipg.respond(w, 202, xreq)

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
