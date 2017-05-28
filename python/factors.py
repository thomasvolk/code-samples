#!/usr/bin/env python3
import math

def factors(n):
    factors = []
    for i in range(1, int(math.sqrt(n)) + 1):
        if(n % i == 0):
            factors.append(i)
            converse = int(n / i)
            if(converse != i):
                factors.append(converse)
    return sorted(factors)

def properDivisors(n):
    return factors(n)[:-1]

def isPerfectNumber(n):
    return n == sum(properDivisors(n))

def isAmicableNumber(n):
    twin = sum(properDivisors(n))
    return n == sum(properDivisors(twin)) and twin != n

def amicableNumbers(start, end):
    return [i for i in range(start, end) if isAmicableNumber(i)]

def perfectNumbers(start, end):
    return [i for i in range(start, end) if isPerfectNumber(i)]

assert factors(0) == []
assert factors(1) == [1]
assert factors(2) == [1, 2]
assert factors(3) == [1, 3]
assert factors(4) == [1, 2, 4]
assert factors(5) == [1, 5]
assert factors(6) == [1, 2, 3, 6]
assert factors(7) == [1, 7]
assert factors(8) == [1, 2, 4, 8]
assert factors(9) == [1, 3, 9]
assert factors(10) == [1, 2, 5, 10]
assert factors(11) == [1, 11]
assert factors(12) == [1, 2, 3, 4, 6, 12]
assert factors(13) == [1, 13]
assert factors(28) == [1, 2, 4, 7, 14, 28]

assert properDivisors(0) == []
assert properDivisors(1) == []
assert properDivisors(2) == [1]
assert properDivisors(3) == [1]
assert properDivisors(4) == [1, 2]
assert properDivisors(13) == [1]
assert sum(properDivisors(6)) == 6
assert sum(properDivisors(28)) == 28

assert amicableNumbers(1, 300) == [220, 284]
assert perfectNumbers(1, 300) == [6, 28]
