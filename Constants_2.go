package main
import "fmt"

func main() {
  const secondsinminutes =60
  const minutesinhours=60
  const secondsinhours= secondsinminutes * minutesinhours
  fmt.Println("The seconds in hours is ",secondsinhours)
}
