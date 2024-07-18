package print

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrettyPrint(t *testing.T) {
  t.Run("normal pretty print", func(t *testing.T) {
    printed := GetPrinted(func() {
      PrettyPrint("testing")
    })
    expected := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n\"testing\"\nXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX\n"
    assert.Equal(t, expected, printed, "printed std out is not equal")
  })

  t.Run("panic pretty print", func(t *testing.T) {
    defer func() {
      if r := recover(); r == nil {
        t.Errorf("The function did not panic")
      }
    }()
    PrettyPrint()
  })

  t.Run("spacer pretty print", func(t *testing.T) {
    printed := GetPrinted(func() {
      PrettyPrint("testing", "8")
    })
    expected := "888888888888888888888888888888888888888888888888888\n\"testing\"\n888888888888888888888888888888888888888888888888888\n"
    assert.Equal(t, expected, printed, "printed std out is not equal")
  })
}

func TestGetPrinted(t *testing.T) {
  printed := GetPrinted(func() {
    fmt.Println("testing")
  }) 

  assert.Equal(t, "testing\n", printed, "printed std out is not equal")
}
