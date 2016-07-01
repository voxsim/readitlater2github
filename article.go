package main

type Article struct {
  title string
  url string
}

type ByTitle []Article

func (a ByTitle) Len() int           { return len(a) }
func (a ByTitle) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTitle) Less(i, j int) bool { return a[i].title < a[j].title }
