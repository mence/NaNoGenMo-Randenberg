# NaNoGenMo 2014: Randenberg

This is my first attempt at [NaNoGenMo 2014](https://github.com/dariusk/NaNoGenMo-2014), an attempt to spend the month of November 2014 writing code that generates a novel of 50k+ words. It's based on [NaNoWriMo](http://nanowrimo.org/), the National Novel Writing Month.

This novel is generated by randomly selecting lines from the 5 most popular books (as well as [Moby Dick](http://www.gutenberg.org/ebooks/2701)) in [Project Gutenberg](https://www.gutenberg.org/ebooks/search/?sort_order=downloads).

It's written in [Go](http://golang.org/). Assuming you have go installed, you can run it at the command line with:

`go run randenberg.go`

### TODO

* Only pick enough sentences to reach the word count (ie it should only pick sentences that comprise 50k worth of words)
* Add some random paragraphing