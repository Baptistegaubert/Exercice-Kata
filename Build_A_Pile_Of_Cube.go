package kata

func FindNb(m int) (t int) {
var n = 0;
  for t=0; t < m;n++{
    
    t += n*n*n
  }
  if t == m {
    return n-1
  }
  return -1
}
