#!/usr/bin/env python3
import math

def factors(n):
    factors = []
    for i in range(1, int(math.sqrt(n)) + 1):
        if(n % i == 0 and i != n):
            factors.append(i)
            converse = int(n / i)
            if(converse != i and converse != n):
                factors.append(converse)
    return sorted(factors)

assert factors(0) == []
assert factors(1) == []
assert factors(2) == [1]
assert factors(3) == [1]
assert factors(4) == [1, 2]
assert factors(5) == [1]
assert factors(6) == [1, 2, 3]
assert factors(7) == [1]
assert factors(8) == [1, 2, 4]
assert factors(9) == [1, 3]
assert factors(10) == [1, 2, 5]
assert factors(11) == [1]
assert factors(12) == [1, 2, 3, 4, 6]
assert factors(13) == [1]
assert factors(28) == [1, 2, 4, 7, 14]

assert sum(factors(6)) == 6
assert sum(factors(28)) == 28
