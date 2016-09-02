package main

import "fmt"

type Article struct {
  Title string
  Url string
}

func (article Article) String() string {
  return fmt.Sprintf("- [%s](%s)\n", article.Title, article.Url)
}

type ByTitle []Article

func (articles ByTitle) Len() int {
  return len(articles)
}

func (articles ByTitle) Swap(i, j int) {
  articles[i], articles[j] = articles[j], articles[i]
}

func (articles ByTitle) Less(i, j int) bool {
  return articles[i].Title < articles[j].Title
}
