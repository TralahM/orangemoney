// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/docallback": {
            "post": {
                "description": "Accept a callback request",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Callbacks"
                ],
                "summary": "Accept a callback request",
                "parameters": [
                    {
                        "description": "DoCallback",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.DoCallback"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/main.DoCallbackResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/dochecktrans": {
            "post": {
                "description": "Check the balance of a specified account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CheckTransaction"
                ],
                "summary": "Do a Check Balance transaction.",
                "parameters": [
                    {
                        "description": "DoCheckTrans",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.DoCheckTrans"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/main.DoCheckTransResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/dom2m": {
            "post": {
                "description": "Perform a Merchant to Merchant Transaction.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "B2B"
                ],
                "summary": "Do M2M Transaction",
                "parameters": [
                    {
                        "description": "DoM2M",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.DoM2M"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/main.DoM2MResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/dom2s": {
            "post": {
                "description": "Perform a Merchant to Subscriber transaction.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "B2C"
                ],
                "summary": "Do M2S Transaction.",
                "parameters": [
                    {
                        "description": "DoM2S",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.DoM2S"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/main.DoM2SResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/dos2m": {
            "post": {
                "description": "Perform a Subscriber to Merchant Transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "C2B"
                ],
                "summary": "Do an S2M transaction.",
                "parameters": [
                    {
                        "description": "DoS2M",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.DoS2M"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/main.DoS2MResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/sendsms": {
            "post": {
                "description": "Send an SMS to a Subscriber.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SendSMS"
                ],
                "summary": "Send an SMS",
                "parameters": [
                    {
                        "description": "SendSMS",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.SendSMS"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/main.SendSMSResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/tcheckbal": {
            "post": {
                "description": "Check the balance of a specified account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Do a Check Balance transaction.",
                "parameters": [
                    {
                        "description": "TcheckBal",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.TcheckBal"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/main.TcheckBalResponse"
                        }
                    }
                }
            }
        },
        "/swagger.json": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "internal"
                ],
                "summary": "Get API Swagger Definition",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.DoCallback": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "mermsisdn": {
                    "type": "string"
                },
                "partnID": {
                    "type": "string"
                },
                "resultCode": {
                    "type": "string"
                },
                "resultDesc": {
                    "type": "string"
                },
                "subsmsisdn": {
                    "type": "string"
                },
                "transid": {
                    "type": "string"
                }
            }
        },
        "main.DoCallbackResponse": {
            "type": "object",
            "properties": {
                "PartnID": {
                    "type": "string"
                },
                "ResultCode": {
                    "type": "string"
                },
                "ResultDesc": {
                    "type": "string"
                }
            }
        },
        "main.DoCheckTrans": {
            "type": "object",
            "properties": {
                "mermsisdn": {
                    "type": "string"
                },
                "partnID": {
                    "type": "string"
                },
                "transid": {
                    "type": "string"
                }
            }
        },
        "main.DoCheckTransResponse": {
            "type": "object",
            "properties": {
                "ResultCode": {
                    "type": "string"
                },
                "ResultDesc": {
                    "type": "string"
                },
                "TransID": {
                    "type": "string"
                }
            }
        },
        "main.DoM2M": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "mermsisdn1": {
                    "type": "string"
                },
                "mermsisdn2": {
                    "type": "string"
                },
                "partnID": {
                    "type": "string"
                },
                "transid": {
                    "type": "string"
                }
            }
        },
        "main.DoM2MResponse": {
            "type": "object",
            "properties": {
                "PartnID": {
                    "type": "string"
                },
                "ResultCode": {
                    "type": "string"
                },
                "ResultDesc": {
                    "type": "string"
                },
                "SystemID": {
                    "type": "string"
                },
                "TransID": {
                    "type": "string"
                }
            }
        },
        "main.DoM2S": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "mermsisdn": {
                    "type": "string"
                },
                "partnID": {
                    "type": "string"
                },
                "subsmsisdn": {
                    "type": "string"
                },
                "transid": {
                    "type": "string"
                }
            }
        },
        "main.DoM2SResponse": {
            "type": "object",
            "properties": {
                "PartnID": {
                    "type": "string"
                },
                "ResultCode": {
                    "type": "string"
                },
                "ResultDesc": {
                    "type": "string"
                },
                "SystemID": {
                    "type": "string"
                },
                "TransID": {
                    "type": "string"
                }
            }
        },
        "main.DoS2M": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "callbackurl": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "mermsisdn": {
                    "type": "string"
                },
                "messageS2m": {
                    "type": "string"
                },
                "partnID": {
                    "type": "string"
                },
                "subsmsisdn": {
                    "type": "string"
                },
                "transid": {
                    "type": "string"
                }
            }
        },
        "main.DoS2MResponse": {
            "type": "object",
            "properties": {
                "PartnID": {
                    "type": "string"
                },
                "ResultCode": {
                    "type": "string"
                },
                "ResultDesc": {
                    "type": "string"
                },
                "SystemID": {
                    "type": "string"
                },
                "TransID": {
                    "type": "string"
                }
            }
        },
        "main.SendSMS": {
            "type": "object",
            "properties": {
                "flash": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "msisdn": {
                    "type": "string"
                },
                "partnID": {
                    "type": "string"
                },
                "partnName": {
                    "type": "string"
                },
                "sender": {
                    "type": "string"
                }
            }
        },
        "main.SendSMSResponse": {
            "type": "object",
            "properties": {
                "Message": {
                    "type": "string"
                },
                "Resultdesc": {
                    "type": "string"
                },
                "Resutcode": {
                    "type": "string"
                }
            }
        },
        "main.TcheckBal": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "partnID": {
                    "type": "string"
                },
                "subsmsisdn": {
                    "type": "string"
                },
                "transid": {
                    "type": "string"
                }
            }
        },
        "main.TcheckBalResponse": {
            "type": "object",
            "properties": {
                "Balance": {
                    "type": "string"
                },
                "PartnID": {
                    "type": "string"
                },
                "ResultCode": {
                    "type": "string"
                },
                "ResultDesc": {
                    "type": "string"
                },
                "TransID": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
