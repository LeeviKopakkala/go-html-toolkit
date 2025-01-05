# go-html-toolkit

`go-html-toolkit` is a lightweight Go package for working with HTML documents. It provides functions to load HTML from files or URLs and search for elements using tags, attributes, or IDs.

---

## Installation

Ensure your Go environment is set up and simply import `htmlutils` into your project. Use it as a local package or include it as part of your module.

---

## Features

- **Load HTML from File**: Parse an HTML file into a DOM structure.
- **Load HTML from URL**: Fetch and parse HTML content from a URL.
- **Search DOM Elements**:
  - By tag name.
  - By attributes.
  - By ID.

---

## Usage

### Parse HTML from File

```go
package main

import (
    "fmt"
    "log"
    "github.com/LeeviKopakkala/go-html-toolkit/htmlutils"
)

func main() {
    doc, err := htmlutils.FileToHtml("example.html")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("HTML parsed successfully:", doc)
}
```

### Parse HTML from URL

```go
doc, err := htmlutils.UrlToHtml("https://example.com")
if err != nil {
    log.Fatal(err)
}
fmt.Println("HTML from URL:", doc)
```

### Find an Element by Tag

```go
element, err := htmlutils.FindElementByTag(doc, "p")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Found element:", element)

```

### Find an Element by Attribute

```go
element, err := htmlutils.FindElementByAttr(doc, "class", "example")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Found element by attribute:", element)

```

### Find an Element by ID

```go
element, err := htmlutils.FindElementById(doc, "header")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Found element by ID:", element)

```

## Error handling

Functions in htmlutils return errors for invalid input or if the requested element is not found. Always handle these errors appropriately.

## Contributing

Feel free to contribute by submitting issues or pull requests on GitHub.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
