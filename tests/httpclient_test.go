package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
)

func TestHttpClient(t *testing.T) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			resp, err := http.Get("http://localhost:8090")
			if err != nil {
				fmt.Println(err.Error())
			} else {
				bt, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("索引", idx, string(bt))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
