package strings

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_should_return_success_when_hashPrefix(t *testing.T) {
	var str = "@ssss"
	fmt.Println("str has prefix @: ", strings.HasPrefix(str, "@"))
}

func Test_should_return_success_when_split(t *testing.T) {
	var str = "@ssss"
	s := strings.Split(str, "@")
	fmt.Print("s[0]: ", s[0])
	fmt.Print("s[1]: ", s[1])
}

func Test_should_return_success_when_trim(t *testing.T) {
	var str = "@ssss   "
	fmt.Println("len before trim: ", len(str))
	str = strings.TrimSpace(str)
	fmt.Println("len after trim: ", len(str))
}

func Test_should_return_success_string2int(t *testing.T) {
	var str = "123"
	fmt.Println(strconv.Atoi(str))

}

func Test_should_return_success_string2uint(t *testing.T) {
	var str = "123"
	v, _ := strconv.ParseUint(str, 10, 0)
	fmt.Println(uint(v))

}

//map key string 是否 words case sensitive
func Test_should_return_success_words_sensitive(t *testing.T) {
	m1 := make(map[string]string)
	m1["aaa"] = "aaa"
	m1["bbb"] = "bbb"
	m1["aaA"] = "aaA"

	fmt.Println(m1["aaA"])
}
