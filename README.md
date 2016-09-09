This program generate a markdown file from a csv export file of my favorite Read it later

## Why I create that?
I was reading and watching a lot of videos, but after a while I understand that I was not rigorous and my knowledge it was like a puzzle with big holes. I created this program to generate a file and push it every time I can; this permits me to see what I really do in the last week, month, year and maybe understand what people I enjoy to follow, found other speakers when I am looking for some conferences and so on.
If you study and have passion I sincerely suggest to do something similar and review often your study growth.

## How to compile it (after setting your gopath properly)
``` sh
go install
```

## How to test it
``` sh
go test
```

## How to use it
If you download the csv in your Downloads directory and you want to store the markdown file in ~/Projects/awesome-topics/README.md, do that:

``` sh
readitlater2github -i ~/Downloads/instapaper-export.csv > ~/Projects/awesome-topics/README.md
```

## Do you like it?
If the answer is yes, please star it.

## Open issues

* everything is not tested due to my lack of experience of golang testing framework
* vendor directory not local in the directory
* it does not support the pipe command for input
* it does not support pocket, readability, etc.

Do you see something else? Please send me a tweet to [@simonvocella](http://twitter.com/simonvocella) or open a issue!
