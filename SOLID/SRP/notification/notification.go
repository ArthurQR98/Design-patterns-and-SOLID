package notification

import "fmt"

type Notifier struct{}

func (n *Notifier) Notify(message string) {
	fmt.Println("Sending message...")
	fmt.Println(message)
}
