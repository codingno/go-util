package print

import (
	"bytes"
	"fmt"
	"os"

	"github.com/bytedance/sonic"
)

func PrettyPrint(args ...any) {

  if len(args) == 0 {
    panic("Pretty print need at least one argument")
  }

	spacer := "X"

  if len(args) > 1 {
    spacer = fmt.Sprintf("%v", args[1])
  }

  spacers := spacer

	for i := 0; i < 50; i++ {
		spacers += spacer
	}

	b, err := sonic.ConfigDefault.MarshalIndent(args[0], "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(spacers)
	fmt.Println(string(b))
	fmt.Println(spacers)

}

func GetPrinted(fn func()) string {
    originalStdout := os.Stdout

    r, w, _ := os.Pipe()
    os.Stdout = w

    fn()

    w.Close()

    os.Stdout = originalStdout

    var buf bytes.Buffer
    buf.ReadFrom(r)
    capturedOutput := buf.String()

    return capturedOutput
}
