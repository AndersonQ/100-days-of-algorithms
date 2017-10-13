No Condition Swap
=================

It is a interesting algorithm to swap two number without use any conditional or boolean operators
The idea is simple and clever:

1. Calculate the difference between `a` and `b`
2. Let `f` be the right shift of `n - 1` bits, where `n` is the size of your int
3. `f` will be `0` if `a > b` or `-1` otherwise
4. Use this clever expression: `a = b * (1 + f) - a * f; b = a * (1 + f) - b *f`

### Examples:

#### a > b:
```
a = 42
b = 21

diff = a  - b
21   = 42 - 21
f = diff    >> 63
0 = 42 - 21 >> 63

a = b  * (1 + f) - a  * f
a = 21 * (1 + 0) - 42 * 0
a = 21 * 1 - 42 * 0
a = 21

b = a  * (1 + f) - b  * f
b = 42 * (1 + 0) - 21 * 0
b = 42 * 1 - 21 * 0
```

#### a < b:
```
a = 21
b = 42

diff = a  - b
-21  = 21 - 42
f  = diff    >> 63
-1 = 21 - 42 >> 63

a = b  * (1 + f)  - a  * f
a = 42 * (1 + -1) - 21 * -1
a = 42 * 0 - 21 * -1
a = 21

b = a  * (1 + f)  - b  * f
b = 21 * (1 + -1) - 42 * -1
b = 21 * 0 - 42 * -1
```

## Build

```
$ go build
```

## Run
```
$ ./009.No.Condition.Swap -a [A_INT] -b [ANOTHER_INT]
```

References
----------

* https://medium.com/100-days-of-algorithms/day-16-no-condition-swap-99ff930de0b4
* https://stackoverflow.com/a/1741104
