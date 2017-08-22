defmodule FactorsTest do
  use ExUnit.Case
  doctest Factors

  test "factors for 6" do
    assert Factors.factors(0) == []
    assert Factors.factors(1) == [1]
    assert Factors.factors(2) == [1,2]
    assert Factors.factors(3) == [1,3]
    assert Factors.factors(4) == [1,2,4]
    assert Factors.factors(5) == [1,5]
    assert Factors.factors(6) == [1,2,3,6]
    assert Factors.factors(12) == [1, 2, 3, 4, 6, 12]
    assert Factors.factors(13) == [1, 13]
    assert Factors.factors(28) == [1, 2, 4, 7, 14, 28]
  end
end
