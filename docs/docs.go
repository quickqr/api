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
        "contact": {
            "url": "https://gitlab.com/quick-qr/api/-/issues"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/generate": {
            "post": {
                "description": "# Hello, world",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/png"
                ],
                "summary": "Get user list",
                "parameters": [
                    {
                        "description": "Configuration for QR code generator. Default values are showed below",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/v1_api.generateBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Will return generated QR code as PNG",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1_api.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "v1_api.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "v1_api.generateBody": {
            "type": "object",
            "required": [
                "data"
            ],
            "properties": {
                "backgroundColor": {
                    "type": "string",
                    "example": "ffffff"
                },
                "data": {
                    "type": "string",
                    "example": "Some data to encode"
                },
                "disableBorder": {
                    "type": "boolean",
                    "example": false
                },
                "foregroundColor": {
                    "type": "string",
                    "example": "000000"
                },
                "recoveryLevel": {
                    "type": "string",
                    "enum": [
                        "low",
                        "medium",
                        "high",
                        "highest"
                    ],
                    "example": "medium"
                },
                "size": {
                    "type": "integer",
                    "minimum": 128,
                    "example": 512
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.0",
	Host:             "",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "Quick QR API",
	Description:      "Description for methods avaliable with Quick QR API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
