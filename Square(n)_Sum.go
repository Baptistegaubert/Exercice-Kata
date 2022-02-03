package kata

func SquareSum(numbers []int) (res int) {
  for _, total:= range numbers {
    res += total*total
  }
  return res
}