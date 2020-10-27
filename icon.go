package icon

import (
	"fmt"
	"io/ioutil"
)

func getIcon() []byte {
	b, err := ioutil.ReadFile("clock.ico")
	if err != nil {
		fmt.Print(err)
	}
	return b
}
