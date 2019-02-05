
<p align="center">
  <img src="docs/logo.svg" title="Mulungu logo" width="200" height="200">
</p>

# Mulungu

Mulungu is a service for creating and managing organizational charts, built in [Go](https://golang.org/).

## Features
- Easy API to manage charts and employees
- Low memory consumption due to the usage of memory pointers
- Fast operations, even with many nodes
- Uses memory cache by default

## Tests 
This project uses Behavior Specifications tests, in BDDish style.  
To run the tests, please run:

```
$ go test -v ./...
```

<p align="center">
  <img src="docs/tests.png" title="Sample of screen showing results of behavioral tests">
</p>

## What is the meaning of the name Mulungu?
[*Erythrina Mulungu*](https://en.wikipedia.org/wiki/Erythrina_mulungu) is a native brazilian tree with a distinct red flowers, and an interesting branching pattern. 