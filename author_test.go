package main_test

import (
  . "github.com/voxsim/readitlater2github"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "sort"
)

var _ = Describe("Author", func() {
  var (
    author1  Author
    author2  Author
  )

  BeforeEach(func() {
    author1 = Author("name1")
    author2 = Author("name2")
  })

  It("can be sorted by name", func() {
    sortedList := []Author{author1, author2}
    notSortedList := []Author{author2, author1}

    sort.Sort(ByName(notSortedList))

    Expect(notSortedList).To(Equal(sortedList))
  })

  It("prints name when it is converted in string", func() {
    Expect(author1.ToString()).To(Equal("## name1\n"))
  })
})
