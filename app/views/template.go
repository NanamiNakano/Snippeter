package views

import "embed"

//go:embed template/*.tmpl
var TemplateFS embed.FS
