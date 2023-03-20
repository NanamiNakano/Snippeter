package global

import (
	"github.com/gomarkdown/markdown/html"
	"github.com/nanaminakano/snippeter/models/ent"
)

var Database *ent.Client

var (
	MarkdownRenderer = html.NewRenderer(html.RendererOptions{
		Flags: html.CommonFlags | html.HrefTargetBlank,
	})
)
