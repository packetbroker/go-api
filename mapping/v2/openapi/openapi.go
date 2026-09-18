// SPDX-FileCopyrightText: Copyright 2021 The Things Industries B.V.
// SPDX-License-Identifier: Apache-2.0

// Package mappingapi contains the OpenAPI specification of the Packet Broker Mapping API v2.
package mappingapi

import (
	_ "embed"
	"fmt"
	"io"
	"text/template"
)

var (
	//go:embed openapi.tmpl.json
	json string
	tmpl = template.Must(template.New("openapi").Parse(json))
)

// PathPrefix is the URL path prefix for all operations.
const PathPrefix = "/api/v2"

// WriteOpenAPI writes the OpenAPI specification.
func WriteOpenAPI(w io.Writer, server, tokenURL string) error {
	err := tmpl.Execute(w, struct {
		Server,
		TokenURL string
	}{
		Server:   server,
		TokenURL: tokenURL,
	})
	if err != nil {
		return fmt.Errorf("execute OpenAPI template: %w", err)
	}
	return nil
}
