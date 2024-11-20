// Package openapi Code generated by swaggo/swag. DO NOT EDIT
package openapi

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
        "/health": {
            "get": {
                "description": "Checks if the API has connectivity to your database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "healthcheck"
                ],
                "summary": "Check if the API is available.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    }
                }
            }
        },
        "/students": {
            "get": {
                "description": "Allows you to search for students by controlling pagination, sorting and filtering results.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Allows you to search the students.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "To identify the tenant",
                        "name": "x-tenant",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "To identify the request",
                        "name": "x-cid",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sortField",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sortOrder",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    }
                }
            },
            "post": {
                "description": "Allows you to create a student after validating duplicate names.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Allows you to create a student.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "To identify the tenant",
                        "name": "x-tenant",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "To identify the request",
                        "name": "x-cid",
                        "in": "header"
                    },
                    {
                        "description": "To create a student",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.StudentCreateInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    }
                }
            }
        },
        "/students/{id}": {
            "get": {
                "description": "Allows you to get the details of a student by ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Allows you to get the details of a student.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "To identify the tenant",
                        "name": "x-tenant",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "To identify the request",
                        "name": "x-cid",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "description": "Student identifier",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.Result": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "dtos.StudentCreateInput": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 200
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Student Swagger Example API",
	Description:      "This is an example API written in Go.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
