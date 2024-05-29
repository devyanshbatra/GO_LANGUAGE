package main
import "fmt"

func main() {
  const messagelength = 30
  const messagelengthallowed = 20 
  fmt.Println("Trying to send a message of length",messagelength)
  if(messagelength<messagelengthallowed){
      fmt.Println("message is sent and length of message was %v",messagelength)
    }
    
    else{
      fmt.Println("Message not sent")
    }
}
