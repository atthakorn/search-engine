# Web Scraper in Go 

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![Build Status](https://travis-ci.com/atthakorn/web-scraper.svg?branch=master)](https://travis-ci.com/atthakorn/web-scraper) 


This is experimental project baked with [go-colly](https://github.com/gocolly/colly) and  [blevesearch](https://github.com/blevesearch) to build web scraper in `Go`.  


### Prerequisite

To use ICU in [blevex](https://github.com/blevesearch/blevex) (blevesearch's extension), it requires to install particular dependencies by following through these steps 

This is verified against `Debian GNU/Linux 8.11 (jessie)`

```
$ sudo apt-get install libleveldb-dev libstemmer-dev libicu-dev build-essential
$ go get -u -v  github.com/blevesearch/bleve
$ cd $GOPATH/src/github.com/blevesearch/bleve/analysis/token
$ git clone https://github.com/blevesearch/cld2.git
$ cd cld2/internal/
$ ./compile_libs.sh
$ sudo cp *.so /usr/local/lib
$ go get -u -v -tags full github.com/blevesearch/bleve
```

