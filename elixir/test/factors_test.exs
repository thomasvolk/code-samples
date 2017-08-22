defmodule FactorsTest do
  use ExUnit.Case
  doctest Factors

  test "factors for 6" do
    assert Factors.factors(6) == [1,2,3,6]
  end
end
