package main

import (
  "sort"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSortAuthorsByName(testRunner *testing.T) {
  sortedList := []Author{Author("name1"), Author("name2")}
  notSortedList := []Author{Author("name2"), Author("name1")}

  sort.Sort(ByName(notSortedList))

  assert.Equal(testRunner, notSortedList, sortedList, "they should be equal")
}
