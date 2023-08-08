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
        "/student": {
            "get": {
                "description": "Show Student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "Show Student",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Student identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ShowStudentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a job Student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "Update Student",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Student Identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Student data to Update",
                        "name": "Student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateStudentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.UpdateStudentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new  Student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "Create Student",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateStudentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.CreateStudentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a new  Student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "Delete student",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Student identification",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.DeleteStudentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/students": {
            "get": {
                "description": "List all students",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Students"
                ],
                "summary": "List students",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ListStudentsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.CreateStudentRequest": {
            "type": "object",
            "properties": {
                "lab": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "period": {
                    "type": "string"
                }
            }
        },
        "handler.CreateStudentResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.StudentResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.DeleteStudentResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.StudentResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ListStudentsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.StudentResponse"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.ShowStudentResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.StudentResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateStudentRequest": {
            "type": "object",
            "properties": {
                "Lab": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "period": {
                    "type": "string"
                }
            }
        },
        "handler.UpdateStudentResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/schemas.StudentResponse"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.StudentResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deteledAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lab": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "period": {
                    "type": "string"
                },
                "updatedAt": {
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
