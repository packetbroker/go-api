// Copyright © 2021 The Things Industries B.V.

package mappingapi

import (
	_ "embed"
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
	return tmpl.Execute(w, struct {
		Server,
		TokenURL string
	}{
		Server:   server,
		TokenURL: tokenURL,
	})
}
