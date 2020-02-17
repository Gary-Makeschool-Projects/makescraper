package json

import (
	"encoding/json"
	"fmt"
)

// Serialize takes in a bye stream and serializes it to JSON
func Serialize(x: bool, y:string, z:int) {
	aBoolValue, _ := json.Marshal(true)
	fmt.Println(string(aBoolValue))

	anIntValue, _ := json.Marshal(1)
	fmt.Println(string(anIntValue))

	aFloatValue, _ := json.Marshal(2.34)
	fmt.Println(string(aFloatValue))

	aStringValue, _ := json.Marshal("BEW 2.5")
	fmt.Println(string(aStringValue))
}
