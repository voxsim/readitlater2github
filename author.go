package main

import "fmt"

type Author string

func (author Author) ToString() string {
  return fmt.Sprintf("## %s\n", author)
}

type ByName []Author

func (author ByName) Len() int           { return len(author) }
func (author ByName) Swap(i, j int)      { author[i], author[j] = author[j], author[i] }
func (author ByName) Less(i, j int) bool { return author[i] < author[j] }
