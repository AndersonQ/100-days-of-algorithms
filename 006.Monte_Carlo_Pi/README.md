Monte Carlo &pi;
==============

Using the [Monte Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method) to calculate Pi

The basic idea of Mont Carlo is:
 1. Define a domain of possible inputs
 2. Generate inputs randomly from a probability distribution over the domain
 3. Perform a deterministic computation on the inputs
 4. Aggregate the results

Thus the algorithm to calculate Pi is:

1. Given a circle inscribed in a square with side length 1.
2. randomly draw `n` points (x, y) | x,y in [0.0,1.0)
3. Count the amount of points inside the circle
4. <code>&pi;</code> = 4 * (points inside the circle) / `n`

A point `(x, y)` is inside the circle if:

<code>x<sup>2</sup></code> + <code>y<sup>2</sup></code> <= 1

## Build

```
$ go build
```

## Run

```
$ ./006.Monte_Carlo_Pi -n 100000000
```


References
----------
* https://en.wikipedia.org/wiki/Monte_Carlo_method
* http://www.eveandersson.com/pi/monte-carlo-circle
* https://medium.com/100-days-of-algorithms/day-9-monte-carlo-%CF%80-7ae010743bde

