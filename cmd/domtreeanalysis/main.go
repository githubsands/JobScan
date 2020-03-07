// make_http_request.go
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var (
	u   = flag.String("url", "", "Url wanted to be scraped")
	o   = flag.String("test", "", "Activate the html API: Token, Tag, or TagName are valid options")
	p   = flag.String("stdout", "", "Copy to standard output")
	t   = flag.String("tokenize", "", "Tokenize the HTTP response's body")
	tag = flag.String("tag", "", "Tag analysis")
	pa  = flag.String("parse", "", "The data to be parsed, if empty no parsing occurs")
)

func main() {
	flag.Parse()

	if *u == "" {
		fmt.Printf("%v", fmt.Sprintf("Need a solid url to continue\n"))
		os.Exit(1)
	}

	fmt.Printf("Getting resource from: %v\n", *u)
	r, err := http.Get(*u)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	defer r.Body.Close()

	if err = stdOut(p, r.Body); err != nil {
		fmt.Printf("Failed to copy to standard output")
		os.Exit(1)
	}

	// Tag inspecting logic of tool
	if *tag != "" {
		PrintTag(r.Body)
		os.Exit(1)
	}

	// Parsing logic of tool
	if *pa != "" {
		toBeParsed := *pa
		ParseHTML(r.Body, toBeParsed)
		os.Exit(0)
	}

	// z := html.NewTokenizer(r.Body)

	/*
		switch *o {
		case "parse":
			parseHTML(r,
			case "Tag":
				Tag(z)
			case "TagName":
				TagName(z)
		default:
			os.Exit(1)
		}
	*/

	// f(doc)
}

// Copy data from the response to standard output
func stdOut(flag *string, rc io.ReadCloser) error {
	if *flag != "on" {
		return nil
	}

	_, err := io.Copy(os.Stdout, rc)
	if err != nil {
		return err
	}

	return nil
}

// var matcher = func(n *html.Node) bool {
// must check for nil value
// if n.DataAtom == atom.A && n.Parent != nil {
// return scrape.Attr(n.Parent, "class") == "result-info"
// }
// return false
// }

// parseHTML parses the given node at data n.
func ParseHTML(r io.Reader, toBeParsed string) error {
	var parse func(*html.Node, string)
	parse = func(n *html.Node, toBeParsed string) {
		if n.Type == html.ElementNode && n.Data == fmt.Sprintf(toBeParsed) {
			fmt.Printf("%v", n)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parse(c, toBeParsed)
		}
	}

	n, err := html.Parse(r)
	if err != nil {
		return err
	}

	parse(n, toBeParsed)
	return nil
}

/*
func Token(z *html.Tokenizer) {
	f(z)
}
*/

// Tag calls token.  It repeatedly calls next and returns its type. Which parses the next token and returns its type
func PrintTag(r io.Reader) error {
	z := html.NewTokenizer(r)
	for {
		switch z.Next() {
		case html.ErrorToken:
			fmt.Printf("%v", z.Token())
			return z.Err()
		default:
			fmt.Printf("%v", z.Token())
		}
	}
}

/*
// TagName calls text or tagname/tagattr it is the lower the level API of this package
func TagName(z *html.Tokenizer) error {
	depth := 0
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return z.Err()
		case html.TextToken:
			if depth > 0 {
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				emitBytes(z.Text())
			}
		case html.StartTagToken, html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				if tt == html.StartTagToken {
					depth++
				} else {
					depth--
				}
			}
		}
	}
}
*/

func emitToken(z *html.Token) {
	fmt.Printf("%v", z)
}

func emitBytes(b []byte) {
	fmt.Printf("%v", b)
}
