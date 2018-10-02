package robotname

import (
	"crypto/rand"
	simplerand "math/rand"
	"time"
)

var letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []byte("0123456789")

var globalMap = make(map[string]bool)

type Robot struct {
	currentName string
}

func (r *Robot) Name() string {
	if r.currentName != "" {
		return r.currentName
	}
	randomFactor := make([]byte, 1)
	_, err := rand.Read(randomFactor)
	if err != nil {
		return ""
	}

	simplerand.Seed(time.Now().UnixNano() * int64(randomFactor[0]))

	arr := make([]byte, 5)
	for i := 0; i < 2; i++ {
		arr[i] = letters[simplerand.Intn(len(letters))]
	}
	for i := 2; i < 5; i++ {
		arr[i] = numbers[simplerand.Intn(len(numbers))]
	}
	var newName string = string(arr)

	_, alreadyGlobal := globalMap[newName]
	if alreadyGlobal {
		return r.Name()
	}

	r.currentName = newName
	globalMap[newName] = true

	return newName
}
func (r *Robot) Reset() {
	r.currentName = ""
	_ = r.Name()
}
