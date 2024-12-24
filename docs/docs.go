// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/posts/{project_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List all posts of a project by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "List all posts of a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/post.Post"
                            }
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "410": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/posts/{project_id}/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new post with the given title, text content, image links, video links, is idea and scheduled at",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Create a new post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Post creation request",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.createPostRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/post.Post"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "409": {
                        "description": "Post already exists",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/posts/{project_id}/{post_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a post by its id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Get a post by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post ID",
                        "name": "post_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.Post"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "410": {
                        "description": "Post not found",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/projects": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "List all projects that the user is a member of",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "List all projects",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/project.Project"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new project with the given name and description",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Create a new project",
                "parameters": [
                    {
                        "description": "Project creation request",
                        "name": "project",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.createProjectRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/project.Project"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "409": {
                        "description": "Project already exists",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a project by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Get a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/project.Project"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "410": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/projects/{project_id}/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add a user to a project by their email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "projects"
                ],
                "summary": "Add a user to a project",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "project_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "410": {
                        "description": "Project not found",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Register a new user with username, password and email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User creation request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.createUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Login with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/users/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logout the currently authenticated user",
                "tags": [
                    "users"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/users/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get information about the currently authenticated user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        },
        "/users/roles": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all application roles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get all roles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.AppRole"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Assign an application role to a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Assign role to user",
                "parameters": [
                    {
                        "description": "Assign role request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.assignRoleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Remove an application role from a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Remove role from user",
                "parameters": [
                    {
                        "description": "Remove role request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.assignRoleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/errors.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.APIError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "details": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "handlers.assignRoleRequest": {
            "type": "object",
            "properties": {
                "role_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "handlers.createPostRequest": {
            "type": "object",
            "properties": {
                "image_links": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_idea": {
                    "type": "boolean"
                },
                "scheduled_at": {
                    "type": "string"
                },
                "text_content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "video_links": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "handlers.createProjectRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handlers.createUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "handlers.loginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "post.Post": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image_links": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "is_idea": {
                    "type": "boolean"
                },
                "project_id": {
                    "type": "string"
                },
                "scheduled_at": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "text_content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "video_links": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "project.Project": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "idea_queue": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                },
                "post_queue": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "session.Session": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deviceFingerprint": {
                    "type": "string"
                },
                "expiresAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "string"
                }
            }
        },
        "user.AppRole": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.LoginResponse": {
            "type": "object",
            "properties": {
                "session": {
                    "$ref": "#/definitions/session.Session"
                },
                "user": {
                    "$ref": "#/definitions/user.UserResponse"
                }
            }
        },
        "user.UserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "roles": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/user.AppRole"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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
	Schemes:          []string{"http", "https"},
	Title:            "OpenCM API",
	Description:      "This is the API server for OpenCM",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
