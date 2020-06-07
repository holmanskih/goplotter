package dichotomy

import (
	"fmt"
)

func f(x float64) float64 {
	return x*x*x - x*x + 2
}

func Dichotomy(a, b float64) (float64, error) {
	if f(a)*f(b) >= 0 {
		return 0, fmt.Errorf("wrong interal")
	}
	c := a

	for (b - a) >= 0.01 {
		c = (a + b) / 2
		if f(c) == 0 {
			break
		}

		if f(c)*f(a) < 0 {
			b = c
		} else {
			a = c
		}
	}

	return c, nil
}

/*
# Prints root of func(x)
# with error of EPSILON
def bisection(a,b):

    if (func(a) * func(b) >= 0):
        print("You have not assumed right a and b\n")
        return

    c = a
    while ((b-a) >= 0.01):

        # Find middle point
        c = (a+b)/2

        # Check if middle point is root
        if (func(c) == 0.0):
            break

        # Decide the side to repeat the steps
        if (func(c)*func(a) < 0):
            b = c
        else:
            a = c

    print("The value of root is : ","%.4f"%c)

# Driver code
# Initial values assumed
a =-200
b = 300
bisection(a, b)

# This code is contributed
# by Anant Agarwal.
*/
