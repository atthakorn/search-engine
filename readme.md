# Search Engine in Go 

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![Build Status](https://travis-ci.com/atthakorn/search-engine.svg?branch=master)](https://travis-ci.com/atthakorn/search-engine) 


This is experimental project baked with [go-colly](https://github.com/gocolly/colly) and  [blevesearch](https://github.com/blevesearch) to build search engine in `Go`.  


### Prerequisite

To use package in [blevex](https://github.com/blevesearch/blevex) (blevesearch's extension), 
it is required to install C dependency as following steps.

*This is verified against `Debian GNU/Linux 8.11 (jessie)*

```shell
$ sudo apt-get install libleveldb-dev libstemmer-dev libicu-dev build-essential
$ cd ~
$ git clone https://github.com/blevesearch/cld2.git
$ cd cld2/internal/
$ ./compile_libs.sh
$ sudo cp *.so /usr/local/lib
```


Minimum requirement for ICU analyzer and tokenizer, libicu-dev only is required. 
Thus, go ahead and simply install it by issuing this command in prompt

```shell
$ sudo apt-get install libicu-dev
```   


### Installation

Use `git clone` or `go get` to download project to your go workspace in `$GOPATH` then run `dep ensure` to initialise project.

```shell

$ go get github.com/atthakorn/search-engine
$ cd $GOPATH/src/github.com/atthakorn/search-engine
$ dep ensure
```

### Config

Here is the list of parameter you can find in `.config.yml`


```yaml

# sites to crawl
entryPoint:
- https://en.wikipedia.org/wiki/Main_Page


# max depth for crawler to follow
maxDepth: 2

# max worker
parallelism: 1

# random delay (second)
delay: 1


# path to store data scraped data from crawler
dataPath: "data/crawl"

# path to store indexed data
indexPath: "data/index"


# http address, 0.0.0.0:8080 or :8080, it means listening all ipv4 address in local machine
httpAddress: ":8080"
```

### Crawling & Indexing

To crawling websites, just `cd` to project root and run

```shell
$ go run cmd/crawl/main.go
``` 

Once, crawl complete, the scraped data will be kept at `./data/crawl/*.json`


Now you can index crawl data by issuing following command

```shell
$ go run cmd/index/main.go
```

The data will be indexed by boltdb, the file will be located at `./data/index/*`


### Starting Up Http Server
 
 Search Engine comes with `simple search application`  , you can start http server by following command
 
 ```shell
$ go run cmd/http/main.go
```

and you can open application via browser at `http://localhost:8080`