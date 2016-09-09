package main

import (
  "encoding/csv"
  "flag"
  "fmt"
  "os"
  "strings"
)

func checked(file *os.File, err error) *os.File {
  if(err != nil) {
    fmt.Printf("Error: %s", err)
    os.Exit(1)
  }
  return file
}

func processing_command_line(args []string) (*os.File, []string, Author) {
  var input_file string
  var tags string
  var default_author string

  cmdFlags := flag.NewFlagSet("event", flag.ContinueOnError)
  cmdFlags.StringVar(&input_file, "input", "", "Input CSV File")
  cmdFlags.StringVar(&input_file, "i", "", "Input CSV File")
  cmdFlags.StringVar(&tags, "tags", "Archive", "A comma-separated list of selected tag")
  cmdFlags.StringVar(&tags, "t", "Archive", "A comma-separated list of selected tag")
  cmdFlags.StringVar(&default_author, "defaultauthor", "Not Categorized", "A default author used if an author is not recognized")
  cmdFlags.StringVar(&default_author, "d", "Not Categorized", "A default author used if an author is not recognized")

  if err := cmdFlags.Parse(args); err != nil {
    fmt.Printf("Error processing command line: %s", err)
    os.Exit(1)
  }

  return checked(os.Open(input_file)), strings.Split(tags, ","), Author(default_author)
}

func main() {
  args := os.Args[1:]

  input_file, tags, default_author := processing_command_line(args);

  instapaper := Instapaper{tags, default_author}

  header := "A list of awesome articles and videos generated from my Instapaper archived list on Software Design, Testing, Public Speaking, etc.\n\nWARNING: some of them are in italian language\n"
  footer := "\n# Are you a newbie? Do you need a study path?\nYou can enjoy [the joebew42' study path](https://github.com/joebew42/study-path) :)\n"
  report := Report{header, footer}

  csv_file := csv.NewReader(input_file)

  authors, articles := instapaper.Parse(csv_file)
  output := report.generate(authors, articles)
  fmt.Printf(output)
}
