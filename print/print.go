package print

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data any, spacer string) {

	spacers := spacer
	for i := 0; i < 50; i++ {
		spacers += spacer
	}

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(spacers)
	fmt.Println(string(b))
	fmt.Println(spacers)

}
