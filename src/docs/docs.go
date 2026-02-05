package docs

import "github.com/swaggo/swag"

const doc = `{
  "swagger": "2.0",
  "info": {
    "description": "RentItems API",
    "title": "RentItems API",
    "version": "1.0"
  },
  "securityDefinitions": {
    "BearerAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header",
      "description": "Bearer <token>"
    }
  },
  "basePath": "/api/v1",
  "paths": {
    "/auth/register": {
      "post": {
        "tags": ["Auth"],
        "summary": "Register user",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "parameters": [
          {
            "in": "body",
            "name": "payload",
            "required": true,
            "schema": {
              "$ref": "#/definitions/registerRequest"
            }
          }
        ],
        "responses": {
          "201": {"description": "Created"},
          "400": {"description": "Bad Request"}
        }
      }
    },
    "/auth/login": {
      "post": {
        "tags": ["Auth"],
        "summary": "Login user",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "parameters": [
          {
            "in": "body",
            "name": "payload",
            "required": true,
            "schema": {
              "$ref": "#/definitions/loginRequest"
            }
          }
        ],
        "responses": {
          "200": {"description": "OK"},
          "400": {"description": "Bad Request"},
          "401": {"description": "Unauthorized"}
        }
      }
    },
    "/items": {
      "get": {
        "tags": ["Items"],
        "summary": "List items",
        "produces": ["application/json"],
        "parameters": [
          {"name": "category", "in": "query", "type": "string"},
          {"name": "search", "in": "query", "type": "string"},
          {"name": "limit", "in": "query", "type": "integer"},
          {"name": "offset", "in": "query", "type": "integer"}
        ],
        "responses": {"200": {"description": "OK"}}
      },
      "post": {
        "tags": ["Items"],
        "summary": "Create item",
        "consumes": ["multipart/form-data"],
        "produces": ["application/json"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"name": "title", "in": "formData", "type": "string", "required": true},
          {"name": "description", "in": "formData", "type": "string", "required": true},
          {"name": "price", "in": "formData", "type": "number", "required": true},
          {"name": "price_type", "in": "formData", "type": "string", "required": true},
          {"name": "category", "in": "formData", "type": "string", "required": true},
          {"name": "image", "in": "formData", "type": "file", "required": true}
        ],
        "responses": {"201": {"description": "Created"}, "400": {"description": "Bad Request"}, "401": {"description": "Unauthorized"}}
      }
    },
    "/categories": {
      "get": {
        "tags": ["Categories"],
        "summary": "List predefined categories",
        "produces": ["application/json"],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {"$ref": "#/definitions/categoryResponse"}
            }
          }
        }
      },
      "post": {
        "tags": ["Categories"],
        "summary": "Create category",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"in": "body", "name": "payload", "required": true, "schema": {"$ref": "#/definitions/categoryRequest"}}
        ],
        "responses": {
          "201": {"description": "Created"},
          "400": {"description": "Bad Request"},
          "401": {"description": "Unauthorized"},
          "409": {"description": "Conflict"}
        }
      }
    },
    "/categories/{id}": {
      "put": {
        "tags": ["Categories"],
        "summary": "Update category",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer"},
          {"in": "body", "name": "payload", "required": true, "schema": {"$ref": "#/definitions/categoryRequest"}}
        ],
        "responses": {
          "200": {"description": "OK"},
          "400": {"description": "Bad Request"},
          "401": {"description": "Unauthorized"},
          "404": {"description": "Not Found"},
          "409": {"description": "Conflict"}
        }
      },
      "delete": {
        "tags": ["Categories"],
        "summary": "Delete category",
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer"}
        ],
        "responses": {
          "200": {"description": "OK"},
          "400": {"description": "Bad Request"},
          "401": {"description": "Unauthorized"},
          "404": {"description": "Not Found"}
        }
      }
    },
    "/items/{id}": {
      "get": {
        "tags": ["Items"],
        "summary": "Get item detail",
        "produces": ["application/json"],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer"}
        ],
        "responses": {"200": {"description": "OK"}, "404": {"description": "Not Found"}}
      },
      "put": {
        "tags": ["Items"],
        "summary": "Update item",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer"},
          {"in": "body", "name": "payload", "required": true, "schema": {"$ref": "#/definitions/updateItemRequest"}}
        ],
        "responses": {"200": {"description": "OK"}, "401": {"description": "Unauthorized"}, "403": {"description": "Forbidden"}, "404": {"description": "Not Found"}}
      },
      "delete": {
        "tags": ["Items"],
        "summary": "Delete item",
        "security": [{"BearerAuth": []}],
        "parameters": [
          {"name": "id", "in": "path", "required": true, "type": "integer"}
        ],
        "responses": {"200": {"description": "OK"}, "401": {"description": "Unauthorized"}, "403": {"description": "Forbidden"}, "404": {"description": "Not Found"}}
      }
    },
    "/users/me": {
      "get": {
        "tags": ["Users"],
        "summary": "Get current user profile",
        "produces": ["application/json"],
        "security": [{"BearerAuth": []}],
        "responses": {
          "200": {"description": "OK"},
          "401": {"description": "Unauthorized"},
          "404": {"description": "Not Found"}
        }
      },
      "put": {
        "tags": ["Users"],
        "summary": "Update current user profile",
        "consumes": ["application/json"],
        "produces": ["application/json"],
        "security": [{"BearerAuth": []}],
        "parameters": [
          {
            "in": "body",
            "name": "payload",
            "required": true,
            "schema": {
              "$ref": "#/definitions/updateMeRequest"
            }
          }
        ],
        "responses": {
          "200": {"description": "OK"},
          "400": {"description": "Bad Request"},
          "401": {"description": "Unauthorized"},
          "404": {"description": "Not Found"}
        }
      }
    }
  },
  "definitions": {
    "registerRequest": {
      "type": "object",
      "required": ["full_name", "email", "password", "phone", "address"],
      "properties": {
        "full_name": {"type": "string"},
        "email": {"type": "string"},
        "password": {"type": "string"},
        "phone": {"type": "string"},
        "address": {"type": "string"}
      }
    },
    "loginRequest": {
      "type": "object",
      "required": ["email", "password"],
      "properties": {
        "email": {"type": "string"},
        "password": {"type": "string"}
      }
    },
    "updateItemRequest": {
      "type": "object",
      "properties": {
        "title": {"type": "string"},
        "price": {"type": "number"},
        "is_available": {"type": "boolean"}
      }
    },
    "updateMeRequest": {
      "type": "object",
      "properties": {
        "full_name": {"type": "string"},
        "phone": {"type": "string"},
        "address": {"type": "string"},
        "profile_pic": {"type": "string"}
      }
    },
    "categoryResponse": {
      "type": "object",
      "properties": {
        "id": {"type": "integer"},
        "name": {"type": "string"},
        "slug": {"type": "string"}
      }
    },
    "categoryRequest": {
      "type": "object",
      "required": ["name"],
      "properties": {
        "name": {"type": "string"}
      }
    }
  }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}

func init() {
	swag.Register(swag.Name, &s{})
}
