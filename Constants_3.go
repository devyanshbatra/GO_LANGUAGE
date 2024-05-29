package main
import "fmt"

func main() {
  const name="Devyansh batra"
  const openrate="2.5"
  msg:=fmt.Sprintf("HI %s your open rate is %v",
  name,
  openrate,)
  fmt.Println(msg)
}
