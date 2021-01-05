package error

import (
	"fmt"
	"net/http"
	"testing"
)

func TestOnly(t *testing.T) {

}
func CertificateLowerService(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}
	//生成要访问的url
	url := "https://www.baidu.com"
	//提交请求
	//bf := make([]byte, 1023)
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	//reqest.Header.Add("Cookie", "xxxxxx")
	//reqest.Header.Add("User-Agent", "xxx")
	//reqest.Header.Add("X-Requested-With", "xxxx")
	reqest.Header.Add("TESTST", "xxxx")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	DDl := response.Header.Get("RETURN")

	fmt.Println(DDl)
	w.Header().Set("bbbbbbbbbbbbb","ddddddddd")
	defer response.Body.Close()
	return
}
func CertificateLowerService2(w http.ResponseWriter, r *http.Request) {

	client := &http.Client{}
	//生成要访问的url
	url := "http://127.0.0.1:8890/bbbb"
	//提交请求
	//bf := make([]byte, 1023)
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	reqest.Header.Add("Cookie", "xxxxxx")
	reqest.Header.Add("User-Agent", "xxx")
	reqest.Header.Add("X-Requested-With", "xxxx")
	reqest.Header.Add("TESTST", "xxxx")

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	DDl := response.Header.Get("RETURN")
	fmt.Println(DDl)
	w.Header().Set("testtttt","ddddddddd")
	defer response.Body.Close()
	return
}

func NewSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	//todo 安恒对接
	mux.HandleFunc("/test", CertificateLowerService)

	return mux
}

func NewSOAPServer(addr string) *http.Server {
	mux := NewSOAPMux()
	server := &http.Server{
		Handler: mux,
		Addr:    addr,
	}
	return server
}

func TestXXX(t *testing.T) {

	addr := fmt.Sprintf("0.0.0.0:8090")
	s := NewSOAPServer(addr)

	if err := s.ListenAndServe(); err != nil {
	}
	//tests := []struct {
	//	name string
	//}{
	//	// TODO: test cases
	//}
	//for _, test := range tests {
	//	t.Run(test.name, func(t *testing.T) {
	//
	//	})
	//}

}
