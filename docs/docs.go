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
        "/v1/account/id": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "GetAccountByID",
                "operationId": "GetAccountByID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        },
                        "headers": {
                            "X-Request-Id": {
                                "type": "string",
                                "description": "The unique id with this request"
                            }
                        }
                    },
                    "default": {
                        "description": "Return if any error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        },
                        "headers": {
                            "X-Request-Id": {
                                "type": "string",
                                "description": "The unique id with this request"
                            }
                        }
                    }
                }
            }
        },
        "/v1/account/list": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "GetAccountList",
                "operationId": "GetAccountList",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The Request PageNo",
                        "name": "pageNo",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "The Request PageSize",
                        "name": "pageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        },
                        "headers": {
                            "X-Request-Id": {
                                "type": "string",
                                "description": "The unique id with this request"
                            }
                        }
                    },
                    "default": {
                        "description": "Return if any error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        },
                        "headers": {
                            "X-Request-Id": {
                                "type": "string",
                                "description": "The unique id with this request"
                            }
                        }
                    }
                }
            }
        },
        "/v1/account/mobile": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "GetAccountByMobile",
                "operationId": "GetAccountByMobile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The user mobile",
                        "name": "mobile",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        },
                        "headers": {
                            "X-Request-Id": {
                                "type": "string",
                                "description": "The unique id with this request"
                            }
                        }
                    },
                    "default": {
                        "description": "Return if any error",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        },
                        "headers": {
                            "X-Request-Id": {
                                "type": "string",
                                "description": "The unique id with this request"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
