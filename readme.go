package main

import "fmt"

type Readme struct {
  header string
  footer string
}

func (readme Readme) generate(authors []Author, articles map[Author][]Article) string {
  content := readme.header
  for _, author := range authors {
    content += fmt.Sprintf("\n## %s\n", author)
    article := articles[author]
    for _, article := range article {
      content += fmt.Sprintf("- [%s](%s)\n", article.title, article.url)
    }
  }
  content += readme.footer
  return content
}
