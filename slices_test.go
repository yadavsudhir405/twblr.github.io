package fundamentals_2

import "testing"
import "fmt"

func TestShouldMakeSliceOfLength5Capacity10(t *testing.T) {
  t.Skip()
  
  s := []int{1,2,3,4,5}
  if len(s) != 5 {
    t.Error("Slice length is not 5")
  }
  if cap(s) != 10 {
    t.Error("Slice capacity is not 10")
  }
}

func TestAppendSliceAtoB(t *testing.T) {
  t.Skip()

  a := []int{1}[:]
  b := []int{2,3}[:]
  c := []int{}
  fmt.Println(fmt.Sprintf("%v is len(%d) cap(%d)", a, len(a), cap(a)))
  fmt.Println(fmt.Sprintf("%v is len(%d) cap(%d)", b, len(b), cap(b)))
  fmt.Println(fmt.Sprintf("%v is len(%d) cap(%d)", c, len(c), cap(c)))
  if(len(c) != 3) {
    t.Error("Array not appended properly")
  }
}