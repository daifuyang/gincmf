/**
** @创建时间: 2020/11/4 11:46 下午
** @作者　　: return
** @描述　　:
 */
package migrate

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_Region(t *testing.T) {
	f, err := os.Open("./region.sql")
	if err != nil {
		fmt.Println("err",err)
	}
	bytes,_ := ioutil.ReadAll(f)
	result := string(bytes)
	result = strings.ReplaceAll(result,"{prefix}","test")
}