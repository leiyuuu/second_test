package Mydatabase

import (
	"math/rand"
	"strconv"

	"github.com/leiyuuu/second-test/naive-server/src/loggenerator"
)

var mymap map[string]bool = make(map[string]bool)

func Gen_a_token(username, password interface{}) string {
	loggenerator.Trace("create a new token")
	now_code := strconv.Itoa(rand.Int())
	for mymap[now_code] {
		now_code = strconv.Itoa(rand.Int())
	}
	mymap[now_code] = true
	return now_code
}
