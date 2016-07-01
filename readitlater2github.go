package main

import "fmt"
import "os"

func generateReadme(authors []Author, articles map[Author][]Article, header string, footer string) string {
  content := header
  for _, author := range authors {
    content += fmt.Sprintf("\n## %s\n", author)
    article := articles[author]
    for _, article := range article {
      content += fmt.Sprintf("- [%s](%s)\n", article.title, article.url)
    }
  }
  content += footer
  return content
}

func main() {
  parameters := os.Args[1:]
  exportFile, _ := os.Open(parameters[0])
  tags := []string{"Archive"}

  instapaper := Instapaper{exportFile, tags}
  authors, articles := instapaper.Parse()

  readmeFile, _ := os.OpenFile(parameters[1], os.O_WRONLY, os.ModeAppend)
  header := "A list of useful articles and videos generated from my Instapaper archived list\n\n# Links\nWARNING: some of these topics are in italian language\n"
  footer := "\n# Are you a newbie? Do you need a study path?\nYou can enjoy [the joebew42' study path](https://github.com/joebew42/study-path) :)\n"
  content := generateReadme(authors, articles, header, footer)
  readmeFile.WriteString(content)
}
