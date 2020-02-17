<p align="center">
  <a href="#">
    <img alt="jive-search logo" src="https://github.com/imthaghost/makescraper/blob/master/docs/media/logo.jpg"> 
  </a>
</p>

<br>

<p align="center">
makescraper is a simple API for scraping website data. Just pass in the url and keywords you are looking for and have the scraper return you json.
</p>

<br>
<p align="center">
   <a href="https://goreportcard.com/badge/github.com/imthaghost/makescraper"><img src="https://goreportcard.com/badge/github.com/imthaghost/makescraper"></a>
   <a href="#">
    <img src="https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg"alt="gitmoji-changelog">
  </a>
</p>
<br>

![form](docs/media/form.png)

### 📚 Table of Contents

1. [Project Structure](#project-structure)
2. [Installation](#installation)
3. [Deliverables](#deliverables)
4. [Resources](#resources)

## Project Structure

```bash
📂 makescraper
├── LICENSE
├── README.md
├── assets
│   ├── css
│   │   └── styles.css
│   ├── fonts
│   │
│   ├── img
│   │   └── thumbnail.jpg
│   ├── js
│   │   └── scripts.min.js
│   └── svg
├── config
├── controllers
│   └── url.go
├── crawler
│   └── scrape.go
├── docs
│   └── media
│       ├── form.png
│       └── logo.jpg
├── go.mod
├── go.sum
├── json
│   └── serilization.go
├── models
│   └── site.go
├── parser
├── server.go
└── templates
    └── index.html
```

## 🚀 Installation

### Homebrew

```bash

```

### Deliverables

#### Scraping

-   [x] **IMPORTANT**: Complete the Web Scraper Workflow worksheet distributed in class.
-   [x] Create a `struct` to store your data.
-   [x] Refactor the `c.OnHTML` callback on line `16` to use the selector(s) you tested while completing the worksheet.
-   [x] Print the data you scraped to `stdout`.

##### Stretch Challenges

-   [ ] Add more fields to your `struct`. Extract multiple data points from the website. Print them to `stdout` in a readable format.

#### Serializing & Saving

-   [x] Serialize the `struct` you created to JSON. Print the JSON to `stdout` to validate it.
-   [x] Write scraped data to a file named `output.json`.
-   [x] **Add, commit, and push to GitHub**.

##### Stretch Challenges

-   [ ] TBA 02/10!

## Resources

### Lesson Plans

-   [**BEW 2.5** - Scraping the Web](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/WebScraping.md): Concepts and examples covered in class related to web scraping and crawling.

### Example Code

#### Scraping

-   [**Colly** - Docs](http://go-colly.org/docs/): Check out the sidebar for 20+ examples!
-   [**Ali Shalabi** - Syntax-Helper](https://github.com/alishalabi/syntax-helper): Command line interface to help generate proper code syntax, pulled from the Golang documentation.

#### Serializing & Saving

-   [JSON to Struct](https://mholt.github.io/json-to-go/): Paste any JSON data and convert it into a Go structure that will support storing that data.
-   [GoByExample - JSON](https://gobyexample.com/json): Covers Go's built-in support for JSON encoding and decoding to and from built-in and custom data types (structs).
-   [GoByExample - Writing Files](https://gobyexample.com/writing-files): Covers creating new files and writing to them.
