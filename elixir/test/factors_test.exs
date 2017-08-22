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
  end
end
