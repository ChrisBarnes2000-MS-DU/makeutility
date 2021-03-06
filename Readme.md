# MakeUtility Project [![Go Report Card](https://goreportcard.com/badge/github.com/Christopher-MakeSchool/makesite)](https://goreportcard.com/report/github.com/Christopher-MakeSchool/makesite)

## [Proposal](Proposal.md): Golang Lichess Package

A Go implementation of [Lichess](https://lichess.org)'s [API](https://lichess.org/api)

## 📚 Table of Contents

- [MakeUtility Project ![Go Report Card](https://goreportcard.com/report/github.com/Christopher-MakeSchool/makesite)](#makeutility-project-)
  - [Golang Lichess Package](#golang-lichess-package)
  - [📚 Table of Contents](#-table-of-contents)
  - [Project Structure](#project-structure)
  - [Getting Started](#getting-started)
    - [Usage](#usage)

## Project Structure

```bash
📂 makeutility
├── 📂 .github
    ├── 📂 ISSUE_TEMPLATE
        ├── bug_report.md
        ├── feature_request.md
    ├── CODE_OF_CONDUCT.md
    ├── CODEOWNERS
├── 📂 cmd
    ├── root.go
    ├── version.go
├── 📂 docs
    ├── club-flyer.html
├── 📂 images
    ├── Rubric.png
├── 📂 models
    ├── club-flyer.go
├── 📂 tmpl
    ├── club-flyer.tmpl
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── makeutility.go
├── Proposal.md
└── Readme.md
```

## 💻 Local Development

## Getting Started

- First, Create an API Token from [here](https://lichess.org/account/oauth/token/create)
- Second, [fork this repo](https://github.com/Christopher-MakeSchool/makeutility/fork),
- Third, run these commands to clone it locally and get started

```zsh
# Clone and CD into/Open this project
$ git clone git@github.com:YOUR_GITHUB_USERNAME/makeutility.git && cd makeutility
# Create a .env file to hold our secrets
$ touch .env > LICHESS_TOKEN=YOUR_API_TOKEN
# Download & Install the dependancies. Then Compile the program
$ go build
# Run the program locally
$ go run makeutility.go
```

## Resources

- [**BEW 2.5**: Project #1 - SSGs](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/SSGProject)
- [**BEW 2.5**: Files & Directories](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/FilesDirectories)
- [**BEW 2.5**: Files & Directories](https://make-school-courses.github.io/BEW-2.5-Strongly-Typed-Languages/#/Lessons/**3rdPartyLibs**)
- [Go Template CheeatSheet](https://curtisvermeeren.github.io/2017/09/14/Golang-Templates-Cheatsheet)
- [Intro to go testing](https://tutorialedge.net/golang/intro-testing-in-go/)
- [Intro to go benchmark testing](https://tutorialedge.net/golang/benchmarking-your-go-programs/)
- [Runes in golang](https://www.geeksforgeeks.org/rune-in-golang/)
- [Default Vaules for Struct Fields in golang](https://www.geeksforgeeks.org/how-to-assign-default-value-for-struct-field-in-golang/)

### Example Code

- [**Go By Example**: Reading Files](https://gobyexample.com/reading-files)
- [**Go By Example**: Writing Files](https://gobyexample.com/writing-files)
- [**Go By Example**: Panic](https://gobyexample.com/panic)
- [**GopherAcademy**: Using Go Templates](https://blog.gopheracademy.com/advent-2017/using-go-templates/)
- [**rapid7.com**: Building a Simple CLI Tool with Golang](https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/)

## Contributors

Anyone is welcome to contribute!

<table>
  <tr>
    <td align="center"><a href="https://github.com/chrisbarnes7404"><img src="https://avatars3.githubusercontent.com/u/25515082?s=460&u=d6d50a936b3e64d2e3d071574891a81faa33d0cb&v=4" width="75px;" alt="Chris Barnes"/><br /><sub><b>Chirs Barnes</b></sub></a><br /><a href="https://github.com/tempor1s/msconsole/commits?author=chrisbarnes7404" title="Code">💻</a></td>
  </tr>
</table>
