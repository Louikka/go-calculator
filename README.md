# go-calculator

Terminal calculator (yes, seriously) written in Go language.


Allowed operations are *addition* (`+`), *subtraction* (`-`), *multiplication* (`*`), *division* (`/`) and *raising to a power* (`^`).


In addition, various functions and constants are also availible.
- Constants
    - `PI` — the number ***π***, approximately equal to 3.14159.
    - `E` — the number ***e*** (Euler's number), approximately equal to 2.71828.
    - `PHI` — the number ***φ*** (golden ratio), approximately equal to 1.61803.
- Functions
    - `SIN(x)` — sine function, argument should be expressed in radian measure.
    - `COS(x)` — cosine function, argument should be expressed in radian measure.
    - `TAN(x)` — tangent function, argument should be expressed in radian measure.
    - `ATAN(x)` — arctangent (in radians) function.
    - `EXP(x)` — natural exponential function, `EXP(x) = e^x`.
    - `ABS(x)` — absolute value function, `ABS(x) = |x|`.
    - `LOG(x)` — decimal logarithm.
    - `LN(x)` — natural logarithm.
    - `SQRT(x)` — square root function.
    - `SUM(i=START..END, x)` — summation function, where `START` is the lower limit of the range, `END` is the upper limit of the range (both are inclusive).
    - `PROD(i=START..END, x)` — product (multiplication) function, where `START` is the lower limit of the range, `END` is the upper limit of the range (both are inclusive).


The program also supports flags.
- `-r` loops program until an error or stop command (`Q`, `QUIT`, `STOP`, `END`) is encountered.
- `-ast` instead of calculating, outputs generated AST of the expression into a file (`out.txt` by default).
- `-o <FILE>` specifies output file.

## ToDo

- Add more tests.
