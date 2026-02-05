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
    "updateMeRequest": {
      "type": "object",
      "properties": {
        "full_name": {"type": "string"},
        "phone": {"type": "string"},
        "address": {"type": "string"},
        "profile_pic": {"type": "string"}
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
