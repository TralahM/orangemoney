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
                    "login"
                ],
                "summary": "Do M2M Transaction",
                "parameters": [
                    {
                        "description": "DoM2M",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.DoM2M"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
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
                    "login"
                ],
                "summary": "Do M2S Transaction.",
                "parameters": [
                    {
                        "description": "DoM2S",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.DoM2S"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
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
                    "login"
                ],
                "summary": "Do an S2M transaction.",
                "parameters": [
                    {
                        "description": "DoS2M",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.DoS2M"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.DoS2MResponse"
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
                "partnID": {
                    "type": "string"
                },
                "resultCode": {
                    "type": "string"
                },
                "resultDesc": {
                    "type": "string"
                },
                "systemID": {
                    "type": "string"
                },
                "transID": {
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
                "partnID": {
                    "type": "string"
                },
                "resultCode": {
                    "type": "string"
                },
                "resultDesc": {
                    "type": "string"
                },
                "systemID": {
                    "type": "string"
                },
                "transID": {
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
                "partnID": {
                    "type": "string"
                },
                "resultCode": {
                    "type": "string"
                },
                "resultDesc": {
                    "type": "string"
                },
                "systemID": {
                    "type": "string"
                },
                "transID": {
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
