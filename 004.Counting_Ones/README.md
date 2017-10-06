Counting Ones
======================

A quite efficient way to count the number of one bits in a positive number

The idea is simple, we will do a bitwise AND of `n` and `n - 1` until `n` becomes zero.
The total number of interactions is the amount of *ones* in the number.

**Examples:**

n = 2, one interaction:
```
n =   2  -  1   =  1
     0b10 - 0b01 = 0b01

    10 (2d)
AND 01 (1d)
  = 00 (0d)

```

n = 3, two interactions:
```
 n = 3   -  1   =  2
    0b11 - 0b01 = 0b01

    11 (3d)
AND 01 (1d)
  = 01 (2d)

n = 2  -  1   =  1
   0b10 - 0b01 = 0b01

    10 (2d)
AND 01 (1d)
  = 00 (0d)

```
