package jsx

import (
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/util"
)

type JsxRenderer struct {
	*html.Renderer
}

func DefaultRenderer() renderer.Renderer {
	return renderer.NewRenderer(
		renderer.WithNodeRenderers(
			util.Prioritized(NewRenderer(), 1000),
		),
	)
}

// NewRenderer returns a new Renderer with given options.
func NewRenderer(opts ...html.Option) renderer.NodeRenderer {
	r := &JsxRenderer{
		Renderer: &html.Renderer{
			Config: html.NewConfig(),
		},
	}

	for _, opt := range opts {
		opt.SetHTMLOption(&r.Config)
	}

	return r
}

func (r *JsxRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	// blocks

	reg.Register(ast.KindThematicBreak, r.renderThematicBreak)

}

func (r *JsxRenderer) renderThematicBreak(w util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	if !entering {
		return ast.WalkContinue, nil
	}
	_, _ = w.WriteString("<hr")
	if n.Attributes() != nil {
		html.RenderAttributes(w, n, html.ThematicAttributeFilter)
	}

	if r.XHTML {
		_, _ = w.WriteString(" />\n")
	} else {
		_, _ = w.WriteString(">\n")
	}

	return ast.WalkContinue, nil
}
