package kata
import(
  "strings"
  "strconv"
  "fmt"
)

func HighAndLow(in string) string {
 var tmpH, tmpL int
  for i, s := range strings.Fields(in) {
    n, _:= strconv.Atoi(string(s))
    if i == 0 {
      tmpH = n
      tmpL = n
    }
    if n > tmpH {
      tmpH = n
    }
    if n < tmpL {
      tmpL = n
    }
  }
  return fmt.Sprintf("%d %d", tmpH, tmpL)
}  