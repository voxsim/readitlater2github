package main_test

import (
  . "github.com/voxsim/readitlater2github"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  "sort"
)

var _ = Describe("Article", func() {
  var (
    article1  Article
    article2  Article
  )

  BeforeEach(func() {
    article1 = Article{
      Title:  "title1",
      Url: "",
    }

    article2 = Article{
      Title:  "title1",
      Url: "",
    }
  })

  It("can be sorted by title", func() {
    sortedList := []Article{article1, article2}
    notSortedList := []Article{article2, article1}

    sort.Sort(ByTitle(notSortedList))

    Expect(notSortedList).To(Equal(sortedList))
  })
})
