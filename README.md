goodlearnbenchmarks
=================

Benchmarking framework for machine learning libraries, and benchmark tests for the [`goodlearn`](https://github.com/amitkgupta/goodlearn) library in particular.

Installation
============

`go get github.com/amitkgupta/goodlearnbenchmarks/...`

Usage
=====

This library will provide a framework for writing benchmark tests for machine learning libraries written in Go.  Instructions on using it for this purpose are forthcoming.

This library will also contain concrete benchmark tests for algorithms implemented in the `goodlearn` project.

```
go install github.com/onsi/ginkgo/ginkgo

cd </path/to/goodlearnbenchmarks>
ginkgo -r
```
