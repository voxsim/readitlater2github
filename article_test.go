package main

import (
  "sort"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSortArticlesByTitle(t *testing.T) {
  sortedList := []Article{Article{"title1", ""}, Article{"title2", ""}}
  notSortedList := []Article{Article{"title2", ""}, Article{"title1", ""}}

  sort.Sort(ByTitle(notSortedList))

  assert.Equal(t, notSortedList, sortedList, "they should be equal")
}
