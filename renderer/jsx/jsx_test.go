package jsx

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
)

func TestJsx(t *testing.T) {
	md := renderMD(DefaultRenderer())

	if err := renderDoc(md); err != nil {
		t.Error(err.Error())
	}
}

func TestHtml(t *testing.T) {
	md := renderMD(nil)

	if err := renderDoc(md); err != nil {
		t.Error(err.Error())
	}
}

func renderMD(rend renderer.Renderer) goldmark.Markdown {
	rOpts := mdOptions()

	if rend != nil {
		rOpts = append(
			rOpts,
			goldmark.WithRenderer(rend),
		)
	}

	return goldmark.New(rOpts...)
}

func renderDoc(md goldmark.Markdown) error {

	bs, err := ioutil.ReadFile("test2.md")
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := md.Convert(bs, &buf); err != nil {
		return err
	}

	fmt.Println(buf.String())
	// fmt.Printf("%s\n", buf.String())

	return nil
}

func mdOptions() []goldmark.Option {
	return []goldmark.Option{
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),

		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),

		goldmark.WithExtensions(
			extension.NewTable(),
			extension.Linkify,
			extension.Typographer,
		),
	}
}
