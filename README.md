<p align="center"><img src="./img/logo.png" width="350"></p>

[![Build Status](https://travis-ci.org/cyruzin/golang-tmdb.svg?branch=master)](https://travis-ci.org/cyruzin/golang-tmdb) [![Build status](https://ci.appveyor.com/api/projects/status/vv76pwj3n4jssuhh?svg=true)](https://ci.appveyor.com/project/cyruzin/golang-tmdb) [![Coverage Status](https://coveralls.io/repos/github/cyruzin/golang-tmdb/badge.svg?branch=master&service=github)](https://coveralls.io/github/cyruzin/golang-tmdb?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/cyruzin/golang-tmdb)](https://goreportcard.com/report/github.com/cyruzin/golang-tmdb) [![GoDoc](https://godoc.org/github.com/cyruzin/golang-tmdb?status.svg)](https://godoc.org/github.com/cyruzin/golang-tmdb) [![Sourcegraph](https://sourcegraph.com/github.com/cyruzin/golang-tmdb/-/badge.svg)](https://sourcegraph.com/github.com/cyruzin/golang-tmdb?badge) [![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

This is a Golang wrapper for working with TMDb API. It aims to support version 3. 

An API Key is required. To register for one, head over to [themoviedb.org](https://www.themoviedb.org/settings/api).

This product uses the TMDb API but is not endorsed or certified by TMDb.

## Requirements

- Go 1.11.x or higher. We aim to support the latest supported versions of go.

## Installation

```sh
go get -u github.com/cyruzin/golang-tmdb
```
## Usage

To get started, import the `tmdb` package and initiate the client:

```go
import "github.com/cyruzin/golang-tmdb"

tmdbClient, err := tmdb.Init("YOUR_APIKEY")

if err != nil {
    fmt.Println(err)
}
    
movie, err := tmdbClient.GetMovieDetails(297802, nil)
```

With optional params:

```go
options := make(map[string]string)
options["language"] = "pt-BR"
options["append_to_response"] = "credits,images"
movie, err := tmdbClient.GetMovieDetails(297802, options)
```

For more examples, [click here](https://github.com/cyruzin/golang-tmdb/tree/master/examples).
