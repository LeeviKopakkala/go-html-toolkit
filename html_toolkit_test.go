package htmlutils_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"golang.org/x/net/html"

	"github.com/LeeviKopakkala/go-html-toolkit"
)

func TestFileToHtml(t *testing.T) {
	validHTMLFile := "test.html"
	nonHTMLFile := "test.txt"
	nonexistentFile := "nonexistent.html"

	// Setup valid HTML file
	os.WriteFile(validHTMLFile, []byte("<html><body><div id='test'>Hello</div></body></html>"), 0644)
	defer os.Remove(validHTMLFile)

	_, err := htmlutils.FileToHtml(validHTMLFile)
	if err != nil {
		t.Errorf("Expected no error for valid HTML file, got: %v", err)
	}

	_, err = htmlutils.FileToHtml(nonHTMLFile)
	if err == nil {
		t.Error("Expected error for non-HTML file, got none")
	}

	_, err = htmlutils.FileToHtml(nonexistentFile)
	if err == nil {
		t.Error("Expected error for nonexistent file, got none")
	}
}

func TestUrlToHtml(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<html><body><p>Test</p></body></html>"))
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	_, err := htmlutils.UrlToHtml(ts.URL)
	if err != nil {
		t.Errorf("Expected no error for valid URL, got: %v", err)
	}

	_, err = htmlutils.UrlToHtml("http://nonexistent.url")
	if err == nil {
		t.Error("Expected error for invalid URL, got none")
	}
}

func TestFindElementByTag(t *testing.T) {
	htmlContent := `<html><body><p>Test</p></body></html>`
	doc, _ := html.Parse(bytes.NewReader([]byte(htmlContent)))

	_, err := htmlutils.FindElementByTag(doc, "p")
	if err != nil {
		t.Errorf("Expected to find <p> tag, got error: %v", err)
	}

	_, err = htmlutils.FindElementByTag(doc, "div")
	if err == nil {
		t.Error("Expected error for non-existent <div> tag, got none")
	}
}

func TestFindElementByAttr(t *testing.T) {
	htmlContent := `<html><body><div id="test">Hello</div></body></html>`
	doc, _ := html.Parse(bytes.NewReader([]byte(htmlContent)))

	_, err := htmlutils.FindElementByAttr(doc, "id", "test")
	if err != nil {
		t.Errorf("Expected to find element with id='test', got error: %v", err)
	}

	_, err = htmlutils.FindElementByAttr(doc, "class", "test")
	if err == nil {
		t.Error("Expected error for non-existent class attribute, got none")
	}
}

func TestFindElementById(t *testing.T) {
	htmlContent := `<html><body><div id="test">Hello</div></body></html>`
	doc, _ := html.Parse(bytes.NewReader([]byte(htmlContent)))

	_, err := htmlutils.FindElementById(doc, "test")
	if err != nil {
		t.Errorf("Expected to find element with id='test', got error: %v", err)
	}

	_, err = htmlutils.FindElementById(doc, "nonexistent")
	if err == nil {
		t.Error("Expected error for non-existent id, got none")
	}
}
