package crawl

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


func NewXPathHelper(htmlContent string) (*XPathHelper, error ){
	content := []byte(htmlContent)
	doc,err := htmlquery.Parse(bytes.NewReader(content))
	if err != nil {
		return nil,fmt.Errorf("failed to parse HTML: %w", err)
	}
	return &XPathHelper{htmlContent: doc},nil
}

func (x *XPathHelper) ExtractByXPath(xpath string) (string, error) {
	node := htmlquery.FindOne(x.htmlContent, xpath)
	if node == nil {
		return "", fmt.Errorf("no element found for XPath: %s", xpath)
	}

	return htmlquery.OutputHTML(node, true), nil
}


func (x *XPathHelper) ExtractFirstNonEmptyByXPath(xpath string) (string, error) {
	nodes := htmlquery.Find(x.htmlContent, xpath)
	for _, node := range nodes {
		if content := htmlquery.OutputHTML(node, true); content != "" {
			return content, nil
		}
	}
	return "", fmt.Errorf("no non-empty element found for XPath: %s", xpath)
}

func (x *XPathHelper) ExtractAllNonEmptyByXPath(xpath string) ([]string, error) {
	nodes := htmlquery.Find(x.htmlContent, xpath)
	var results []string

	for _, node := range nodes {
		if content := htmlquery.OutputHTML(node, true); content != "" {
			results = append(results, content)
		}
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no non-empty elements found for XPath: %s", xpath)
	}

	return results, nil
}
