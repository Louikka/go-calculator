# go-calculator

Terminal calculator (yes, seriously) written in Go language.

---

Allowed operations are *addition* (`+`), *subtraction* (`-`), *multiplication* (`*`), *division* (`/`) and *raising to a power* (`^`).


In addition, various functions and constants are also availible.
- Constants
    - `PI` — the number ***π***, approximately equal to 3.14159.
    - `E` — the number ***e*** (Euler's number), approximately equal to 2.71828.
    - `PHI` — the number ***φ*** (golden ratio), approximately equal to 1.61803.
- Functions
    - `SIN()` — sine function, argument should be expressed in radian measure.
    - `COS()` — cosine function, argument should be expressed in radian measure.
    - `TAN()` — tangent function, argument should be expressed in radian measure.
    - `ATAN()` — arctangent (in radians) function.
    - `EXP()` — natural exponential function, `EXP(x) = e^x`.
    - `ABS()` — absolute value function, `ABS(x) = |x|`.
    - `LOG()` — decimal logarithm.
    - `LN()` — natural logarithm.
    - `SQRT()` — square root function.

---

The program also supports flags.
- `--with-loop` loops program until an error or stop command (`Q`, `QUIT`, `STOP`, `END`) is encountered.
