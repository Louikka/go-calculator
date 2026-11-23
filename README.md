# go-calculator

Terminal calculator (yes, seriously) written in Go language.


Allowed operations are *addition* (`+`), *subtraction*/*negation* (`-`), *multiplication* (`*`), *division* (`/`), *raising to a power* (`^`) and *factorial* (`!`).


In addition, various functions and constants are also availible.
- Constants
    - `PI` — the number ***π***, approximately equal to 3.14159.
    - `E` — the number ***e*** (Euler's number), approximately equal to 2.71828.
    - `PHI` — the number ***φ*** (golden ratio), approximately equal to 1.61803.
    - `PSI` — the number ***ψ*** (supergolden ratio), approximately equal to 1.46557.
- Functions
    - `SIN(x)` — sine function, argument should be expressed in radian measure.
    - `COS(x)` — cosine function, argument should be expressed in radian measure.
    - `TAN(x)` — tangent function, argument should be expressed in radian measure.
    - `ATAN(x)` — arctangent (in radians) function.
    - `ABS(x)` — absolute value function, `ABS(x) = |x|`.
    - `LOG(x, y)` — logarithm of `y` to the base `x`.
    - `LN(x)` — natural logarithm, same as `LOG(E, x)`.
    - `SQRT(x)` — square root function.
    - `CBRT(x)` — cubic root function.
    - `ROUND(x)` — rounds to the nearest integer.
    - `RAND()` — returns random floating-point number in the half-open interval [0, 1).
    - `SUM(i=n..m, x)` — summation function, where `i` is index variable, `n` is the lower limit of the range, `m` is the upper limit of the range (both are inclusive), `x` is an expression where usage of the index variable are permitted.
    - `PROD(i=n..m, x)` — product (multiplication) function, where `i` is index variable, `n` is the lower limit of the range, `m` is the upper limit of the range (both are inclusive), `x` is an expression where usage of the index variable are permitted.


The program also supports flags.
- `-r` loops program until an error or stop command (`Q`, `QUIT`, `STOP`, `END`) is encountered.


## Building

```sh
$ go build
```


## ToDo

- add more tests
