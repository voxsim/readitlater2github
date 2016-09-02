package main

import "fmt"

type Markdown struct {
  header string
  footer string
}

func (markdown Markdown) generate(authors []Author, articles map[Author][]Article) string {
  content := markdown.header
  for _, author := range authors {
    content += fmt.Sprintf("\n## %s\n", author)
    article := articles[author]
    for _, article := range article {
      content += article.String()
    }
  }
  content += markdown.footer
  return content
}
