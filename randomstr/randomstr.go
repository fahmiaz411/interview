package randomstr

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type Str []string

func (value Str) Len() int {
	return len(value)
}

func (value Str) Less(i, j int) bool {
	return value[i] < value[j]
}

func (value Str) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func RandomStr() {
	rand.Seed(time.Now().Unix())
    str := Str{}

	for i := 0; i < 10; i++ {
		str = append(str, string(uint8(97 + rand.Intn(127-97+1))))
	}

	fmt.Println(strings.Join(str, ""))

	sort.Sort(str)

	fmt.Println(strings.Join(str, ""))
}