// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/balance/deposit": {
            "put": {
                "description": "Increase balance of user_id if it exists \u0026 value of increment is positive",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Increase balance",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/balance/show/:id": {
            "get": {
                "description": "Show balance of user if id is correct",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Show balance",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/reservation": {
            "post": {
                "description": "Reserve order if it doesn't exist",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservations"
                ],
                "summary": "Reserve order",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/reservation/accept": {
            "put": {
                "description": "Decrease balance of user, if reservation with given parameters exists",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservations"
                ],
                "summary": "Accept reservation",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8010 for local, balance-db:8010 for docker",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Balance microservice",
	Description:      "Balance service task made for Avito internship",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
