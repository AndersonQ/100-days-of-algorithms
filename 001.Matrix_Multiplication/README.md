Matrix Multiplication
======================

A simple iterative implementation of
[matrix multiplication](https://en.wikipedia.org/wiki/Matrix_multiplication_algorithm#Iterative_algorithm)

## Build

```
$ go build
```

## Run

```
$ ./001.Matrix_Multiplication < input
```

### Input format

* The size of matrix A separated by space
* Each line of A separated by space
* The size of B separated by space
* Each line of B separated by space

**Example**:
```
A:
  2
  2
B:
  3 3 3
  
Result:
  6 6 6
  6 6 6
```
1 2
2
2
3 1
3 3 3

## Test
```
$ go test ./...
```
