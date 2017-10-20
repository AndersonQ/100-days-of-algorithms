Matrix Addition
======================

A wee program to add tow matrices

## Build

```
$ go build
```

## Run

```
$ ./014.Matrix.Addition < input
```

### Input format

* The size of matrix A separated by space
* Each line of A separated by space
* The size of B separated by space
* Each line of B separated by space

**Example**:
```
A:
  1 1
  1 1
B:
  1 1
  1 1
  
Result:
  2 2
  2 2

```
```
input file: 2 2 1 1 1 1 2 2 1 1 1 1
```


## Test
```
$ go test -cover ./...
```
