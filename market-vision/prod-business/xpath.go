package prodbusiness

import (
	"bytes"
	"fmt"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	// "google.golang.org/api/content/v2"
)

type XPathHelper struct {
	htmlContent *html.Node
}

func NewXPathHelper(htmlContent string) (*XPathHelper, error) {
	content := []byte(htmlContent)
	doc, err := htmlquery.Parse(bytes.NewReader(content))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}
	return &XPathHelper{htmlContent: doc}, nil
}

func (x *XPathHelper) ExtractByXPath(xpath string) string {
	node := htmlquery.FindOne(x.htmlContent, xpath)

	return htmlquery.OutputHTML(node, true)
}

func (x *XPathHelper) ExtractFirstNonEmptyByXPath(xpath string) string {
	nodes := htmlquery.Find(x.htmlContent, xpath)
	for _, node := range nodes {
		content := htmlquery.OutputHTML(node, true)
		return content
	}
	return ""
}

func (x *XPathHelper) ExtractAllNonEmptyByXPath(xpath string) []string {
	nodes := htmlquery.Find(x.htmlContent, xpath)
	var results []string

	for _, node := range nodes {
		if content := htmlquery.OutputHTML(node, true); content != "" {
			results = append(results, content)
		}
	}

	return results
}
