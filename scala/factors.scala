import scala.math

def factors(n: Int) = (1 until (math.sqrt(n).toInt + 1)).filter( (i) => n % i == 0).map( i => List(i, n / i)).flatten.toList.sorted
def properDivisors(n: Int) = factors(n).reverse.drop(1).reverse

val factorsTestData = Seq(
    ( 0, Seq())
  , ( 1, Seq(1))
  , ( 2, Seq(1, 2))
  , ( 3, Seq(1, 3))
  , ( 4, Seq(1, 2, 4))
  , ( 5, Seq(1, 5))
  , ( 6, Seq(1, 2, 3, 6))
  , ( 7, Seq(1, 7))
  , ( 8, Seq(1, 2, 4, 8))
  , ( 9, Seq(1, 3, 9))
  , (10, Seq(1, 2, 5, 10))
  , (11, Seq(1, 11))
  , (12, Seq(1, 2, 3, 4, 6, 12))
  , (13, Seq(1, 13))
  , (28, Seq(1, 2, 4, 7, 14, 28))
)

factorsTestData.foreach( n => assert( factors(n._1) == n._2, s"factors(${n._1}) != ${n._2} (actual: ${factors(n._1)})" ))

assert( properDivisors(0) == Seq() )
assert( properDivisors(1) == Seq() )
assert( properDivisors(2) == Seq(1) )
assert( properDivisors(3) == Seq(1) )
assert( properDivisors(4) == Seq(1, 2) )
assert( properDivisors(6) == Seq(1, 2, 3) )

assert( properDivisors(6).foldLeft(0)(_ + _) == 6 )
assert( properDivisors(28).foldLeft(0)(_ + _) == 28 )
