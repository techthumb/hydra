/*
 * ORY Hydra
 *
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * OpenAPI spec version: latest
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"time"
)

type AuthenticationSession struct {
	AuthenticatedAt time.Time `json:"AuthenticatedAt,omitempty"`

	ID string `json:"ID,omitempty"`

	Subject string `json:"Subject,omitempty"`
}
