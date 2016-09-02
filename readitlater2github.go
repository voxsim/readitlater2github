package main

import "encoding/csv"
import "fmt"
import "os"

func main() {
  parameters := os.Args[1:]

  tags := []string{"Archive"}
  default_author := Author("Not Categorized")
  instapaper := Instapaper{tags, default_author}

  header := "A list of useful articles and videos generated from my Instapaper archived list\n\n# Links\nWARNING: some of these topics are in italian language\n"
  footer := "\n# Are you a newbie? Do you need a study path?\nYou can enjoy [the joebew42' study path](https://github.com/joebew42/study-path) :)\n"
  report := Report{header, footer}

  inputFile, err := os.Open(parameters[0])
  if(err != nil) {
    fmt.Printf("inputFile error: %s", err)
    os.Exit(1)
  }

  outputFile, err := os.OpenFile(parameters[1], os.O_WRONLY, os.ModeAppend)
  if(err != nil) {
    fmt.Printf("inputFile error: %s", err)
    os.Exit(1)
  }

  csvFile := csv.NewReader(inputFile)

  authors, articles := instapaper.Parse(csvFile)
  output := report.generate(authors, articles)
  outputFile.WriteString(output)

  outputFile.Sync()

  defer outputFile.Close()
}
