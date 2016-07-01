package main

import "golang.org/x/net/html"
import "io"
import "strings"
import "fmt"
import "sort"

type Instapaper struct {
  page io.Reader
  tags []string
}

func collectArticles(nodes []*html.Node) ([]Author, map[Author][]Article) {
  bibliographies := make(map[Author][]Article)
  for _, node := range nodes {
    for li := node.FirstChild; li != nil; li = li.NextSibling {
      a := li.FirstChild
      if a != nil && a.Type == html.ElementNode && a.Data == "a" {
        author := Author("Not Categorized")
        url := a.Attr[0].Val
        title := a.FirstChild.Data
        values := strings.Split(title, "|")
        if len(values) == 2 {
          author = Author(values[0])
          title = values[1]
        }
        fmt.Printf("- (%s, %s, %s)\n", author, url, title)
        if bibliographies[author] == nil {
          bibliographies[author] = []Article {}
        }
        bibliographies[author] = append(bibliographies[author], Article{title, url})
      }
    }
  }
  authors := []Author{}
  for author := range bibliographies {
    authors = append(authors, author)
    sort.Sort(ByTitle(bibliographies[author]))
  }
  sort.Sort(ByName(authors))
  return authors, bibliographies
}

func extrapolateTag(node *html.Node) string {
  if node == nil {
    return ""
  }
  if node.PrevSibling == nil {
    return ""
  }
  if node.PrevSibling.PrevSibling == nil {
    return ""
  }
  if node.PrevSibling.PrevSibling.FirstChild == nil {
    return ""
  }
  return node.PrevSibling.PrevSibling.FirstChild.Data
}

func contains(array []string, elementToSearch string) bool {
  for _, element := range array {
    if element == elementToSearch {
      return true
    }
  }
  return false
}

func collectRootOfEachTag(page *html.Node, whichTags []string) []*html.Node {
  body := page.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling
  roots := []*html.Node{}
  for node := body.FirstChild; node != nil; node = node.NextSibling {
    if node.Type == html.ElementNode && node.Data == "ol" {
      tag := extrapolateTag(node)
      if contains(whichTags, tag) {
        roots = append(roots, node)
      }
    }
  }
  return roots;
}

func (instapaper Instapaper) Parse() ([]Author, map[Author][]Article) {
  dom, _ := html.Parse(instapaper.page)
  roots := collectRootOfEachTag(dom, instapaper.tags)
  authors, articles := collectArticles(roots)
  return authors, articles
}
