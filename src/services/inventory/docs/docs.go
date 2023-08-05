// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/broker": {
            "get": {
                "description": "Retrieve a paginated list of cluster that the user has access",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Broker"
                ],
                "summary": "Retrieve a list of rabbitmq clusters registered",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "name": "PageNumber",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "name": "PageSize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/contracts.PaginatedResult-entities_BrokerEntity"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new \u003cb\u003eRabbitMQ\u003c/b\u003e broker. The credential provider must be valid and the cluster operational",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Broker"
                ],
                "summary": "Register a new RabbitMQ Broker",
                "parameters": [
                    {
                        "description": "Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contracts.CreateBrokerRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/broker/exists": {
            "get": {
                "description": "Check if exists an rabbitmq cluster with host es",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Broker"
                ],
                "summary": "Verify if exists a rabbitmqcluster",
                "parameters": [
                    {
                        "type": "string",
                        "name": "Host",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "Port",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/broker/{brokerId}": {
            "get": {
                "description": "Retrieve a single rabbitmq cluster",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Broker"
                ],
                "summary": "Retrieve a single rabbitmq cluster",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of a broker to be retrived",
                        "name": "brokerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.BrokerEntity"
                        }
                    }
                }
            },
            "delete": {
                "description": "Soft delete a broker will not completly erase from database, but will not show up anymore in the",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Broker"
                ],
                "summary": "Soft delete a broker",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id of a broker to be soft deleted",
                        "name": "brokerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "contracts.CreateBrokerRequest": {
            "type": "object",
            "required": [
                "description",
                "host",
                "name",
                "password",
                "port",
                "user"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "contracts.PaginatedResult-entities_BrokerEntity": {
            "type": "object",
            "properties": {
                "pageNumber": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.BrokerEntity"
                    }
                },
                "totalItems": {
                    "type": "integer"
                }
            }
        },
        "entities.BrokerEntity": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
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
