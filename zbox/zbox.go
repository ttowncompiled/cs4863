package main

func preprocess(pattern *string) []int {
  zbox := make([]int, len(*pattern), len(*pattern))
  zbox[0] = -1
  for i := 1; i < len(*pattern); i++ {
    zi := 0
    for j := 0; i+j < len(*pattern) && string((*pattern)[i+j]) == string((*pattern)[j]); j++ {
      zi++
    }
    zbox[i] = zi
  }
  return zbox
}

func match(text *string, pattern *string, zbox []int) []int {
  
}

func main() {
  
}