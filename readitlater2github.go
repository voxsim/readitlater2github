package main

import "golang.org/x/net/html"
import "os"

func main() {
  parameters := os.Args[1:]
  inputFile, _ := os.Open(parameters[0])
  outputFile, _ := os.OpenFile(parameters[1], os.O_WRONLY, os.ModeAppend)

  tags := []string{"Archive"}
  instapaper := Instapaper{inputFile, tags}

  header := "A list of useful articles and videos generated from my Instapaper archived list\n\n# Links\nWARNING: some of these topics are in italian language\n"
  footer := "\n# Are you a newbie? Do you need a study path?\nYou can enjoy [the joebew42' study path](https://github.com/joebew42/study-path) :)\n"
  readme := Readme{header, footer}

  input, _ := html.Parse(inputFile)
  authors, articles := instapaper.Parse(input)
  output := readme.generate(authors, articles)
  outputFile.WriteString(output)
}
