package context

import (
	"fmt"
	"testing"
)

func TestSendSync(t *testing.T) {
	coreContext := GetContext(MsgCtxTypeChannel)
	coreContext.AddModule("test_dest")
	messsage := "hello"

	go func() {
		coreContext.Send("test_dest", messsage)
		fmt.Printf("send message %v\n", messsage)
	}()

	msg, err := coreContext.Receive("test_dest")
	fmt.Printf("receive msg: %v, error: %v\n", msg, err)
}
