package fetchsymbols

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func GetSymbols() []string {
	resp, err := http.Get("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//slink := make([]string, 0)
	var f func(*html.Node, bool)
	var g func(*html.Node)
	var h func(*html.Node)
	var j func(*html.Node)
	FirstTable := true
	FirstRow := true
	Slink := make([]string, 0)

	f = func(node *html.Node, firstTable bool) {
		if node.Type == html.ElementNode && node.Data == "table" && firstTable {
			FirstTable = false
			g = func(node *html.Node) {
				if node.Type == html.ElementNode && node.Data == "tr" {
					FirstRow = true
					h = func(node *html.Node) {
						if node.Type == html.ElementNode && node.Data == "td" {
							j = func(node *html.Node) {
								if node.Type == html.ElementNode && node.Data == "a" && FirstRow {
									FirstRow = false
									Slink = append(Slink, node.FirstChild.Data)
								}
								for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
									j(nextNode)
								}
							}
							j(node)
						}
						for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
							h(nextNode)
						}
					}
					h(node)
				}
				for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
					g(nextNode)
				}
			}
			g(node)
		}
		for nextNode := node.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
			f(nextNode, FirstTable)
		}
	}
	f(doc, FirstTable)
	return (Slink)
}
