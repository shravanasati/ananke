# ananke

A HTML to markdown converter.

Powered by [html2md](https://github.com/shravanasati/ananke/tree/master/html2md).

<!-- todo include example screenshots -->

### Usage

ananke can read input from STDIN as well as from the arguments passed to it. If multiple arguments are passed, they are concatenated.

Read a HTML file and print the Markdown output:
```sh
cat index.html | ananke
```

Read a HTML file and write a new Markdown file:
```sh
cat index.html | ananke > index.md
```

Read a HTML file, print the output as well as write it to a file:
```sh
cat index.html | ananke | tee /dev/tty index.md
```

Read HTML from a URL and print the output:
```sh
curl --no-progress-meter -L https://wikipedia.org/wiki/Anime | ananke
```