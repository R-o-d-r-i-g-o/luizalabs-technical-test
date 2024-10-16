// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/v1/address/{zip-code}": {
            "get": {
                "description": "Get address details using a provided ZIP code. Returns a structured response with address data or error information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Address"
                ],
                "summary": "Retrieve CEP information by ZIP code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Cache control directive (e.g., 'no-cache')",
                        "name": "X-Cache-Control",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "ZIP Code",
                        "name": "zip-code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_features_zipcode.swagGetAddressByZipCodeResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid ZIP code format",
                        "schema": {
                            "$ref": "#/definitions/luizalabs-technical-test_pkg_server.APIErrorResponse"
                        }
                    },
                    "404": {
                        "description": "ZIP code not found",
                        "schema": {
                            "$ref": "#/definitions/luizalabs-technical-test_pkg_server.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/auth/login": {
            "post": {
                "description": "Authenticates the user with the provided credentials and returns a JWT token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticate user and return a JWT token",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_features_auth.PostLoginPayload"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Token generated successfully",
                        "schema": {
                            "$ref": "#/definitions/internal_features_auth.swagAuthenticateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/luizalabs-technical-test_pkg_server.APIErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/luizalabs-technical-test_pkg_server.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/auth/register": {
            "post": {
                "description": "Registers a new user with the provided information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_features_auth.PostRegisterPayload"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User successfully registered"
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/luizalabs-technical-test_pkg_server.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/luizalabs-technical-test_pkg_server.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/v1/health/metrics": {
            "get": {
                "description": "Returns the Prometheus metrics for monitoring",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Expose Prometheus metrics",
                "responses": {
                    "200": {
                        "description": "Metrics",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/health/ping": {
            "get": {
                "description": "Responds with a \"pong\" message to indicate that the service is healthy.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_features_health.swagHealthResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internal_features_auth.AuthenticateUserResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "internal_features_auth.PostLoginPayload": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "internal_features_auth.PostRegisterPayload": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "internal_features_auth.swagAuthenticateUserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal_features_auth.AuthenticateUserResponse"
                }
            }
        },
        "internal_features_health.healthResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "internal_features_health.swagHealthResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal_features_health.healthResponse"
                }
            }
        },
        "internal_features_zipcode.GetAddressByZipCodeResponse": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "neighborhood": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "internal_features_zipcode.swagGetAddressByZipCodeResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/internal_features_zipcode.GetAddressByZipCodeResponse"
                }
            }
        },
        "luizalabs-technical-test_pkg_server.APIErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "error": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
