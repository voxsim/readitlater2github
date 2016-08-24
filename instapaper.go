package main

import "encoding/csv"
import "fmt"
import "sort"
import "strings"
import "reflect"

func in_array(val interface{}, array interface{}) bool {
  switch reflect.TypeOf(array).Kind() {
  case reflect.Slice:
    s := reflect.ValueOf(array)

    for i := 0; i < s.Len(); i++ {
      if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
        return true
      }
    }
  }

  return false
}

type Instapaper struct {
  tags []string
  default_author Author
}

func (instapaper Instapaper) retrieveAuthorAndTitle(title string) (Author, string) {
  if strings.Contains(title, "|") {
    fields := strings.Split(title, "|")
    return Author(fields[0]), fields[1]
  }

  return instapaper.default_author, title
}

func (instapaper Instapaper) Parse(file *csv.Reader) ([]Author, map[Author][]Article) {
  bibliographies := make(map[Author][]Article)
  bibliographies[instapaper.default_author] = []Article{}

  for record, err := file.Read(); err == nil; record, err = file.Read() {
    fmt.Printf("processing: %s\n", record)

    url, title, _, tag := record[0], record[1], record[2], record[3]
    author, title := instapaper.retrieveAuthorAndTitle(title)

    if in_array(tag, instapaper.tags) {
      if bibliographies[author] == nil {
        bibliographies[author] = []Article {}
      }
      bibliographies[author] = append(bibliographies[author], Article{title, url})
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
