package main

type Report struct {
  header string
  footer string
}

func (report Report) generate(authors []Author, bibliography map[Author][]Article) string {
  content := report.header
  for _, author := range authors {
    content += "\n"
    content += author.ToString()
    articles := bibliography[author]
    for _, article := range articles {
      content += article.ToString()
    }
  }
  content += report.footer
  return content
}
