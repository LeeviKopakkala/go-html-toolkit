package htmlutils

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

func FileToHtml(file string) (*html.Node, error) {
	if filepath.Ext(file) != ".html" {
		return nil, fmt.Errorf("file is not an HTML file")
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	doc, err := html.Parse(strings.NewReader(string(bytes)))
	if err != nil {
		return nil, fmt.Errorf("error parsing file content: %w", err)
	}

	return doc, nil
}

func UrlToHtml(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	return doc, nil
}

func FindElementByTag(n *html.Node, tag string) (*html.Node, error) {
	if n.Type == html.ElementNode && n.Data == tag {
		return n, nil
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result, err := FindElementByTag(c, tag); err == nil {
			return result, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}

func FindElementByAttr(n *html.Node, key, value string) (*html.Node, error) {
	for _, a := range n.Attr {
		if a.Key == key && a.Val == value {
			return n, nil
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result, err := FindElementByAttr(c, key, value); err == nil {
			return result, nil
		}
	}
	return nil, fmt.Errorf("element not found")
}

func FindElementById(n *html.Node, value string) (*html.Node, error) {
	el, err := FindElementByAttr(n, "id", value)
	if err == nil {
		return el, nil
	}
	return nil, fmt.Errorf("element not found")
}
