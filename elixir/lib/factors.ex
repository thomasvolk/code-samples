defmodule Factors do
  def factors(n) when n < 1 do
     []
  end
  def factors(n) do
     factors(n, round(:math.sqrt(n)) + 1, [])
  end
  def properDivisors(n) do
     Enum.drop factors(n), -1
  end
  defp factors(n, i, results) when i > 0 do
    case rem(n, i) do
      0 -> factors(n, i - 1, results ++ [i] ++ [round(n/i)])
      _ -> factors(n, i - 1, results)
    end
  end
  defp factors(_, i, results) when i < 1 do
     Enum.sort(Enum.uniq(results))
  end
end
