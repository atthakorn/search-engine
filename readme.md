# Search Engine in Go 

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)
[![Build Status](https://travis-ci.com/atthakorn/web-scraper.svg?branch=master)](https://travis-ci.com/atthakorn/web-scraper) 


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
