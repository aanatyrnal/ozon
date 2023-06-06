package requests

import (
	"fmt"
	//"strconv"
	//"bytes"
	//"context"
	"io/ioutil"
	//"os"
	"bytes"
	"net/http"
	"ozon/internal/utils/tools"
	"strings"
	"time"
	// write file
	//"strconv"
	//"os"
)

type DumpResponse struct {
	Code int
	Body string
	Id   int
}

type DumpResponseWithOptions struct {
	Code int
	Body string
	Id   int
}

type UrlWithMethod struct {
	Url    string
	Method string
	Body   string
}

// type RequestUrl struct {
// 	Id string
// 	Url string
// 	Options interface{}
// }

// type Block struct {
//   Try     func()
//   Catch   func(Exception)
//   Finally func()
// }

// type Exception interface{}

// func Throw(up Exception) {
//   //panic(up)
// }

// func (tcf Block) Do() {
//   if tcf.Finally != nil {

//       defer tcf.Finally()
//   }
//   if tcf.Catch != nil {
//       defer func() {
//           if r := recover(); r != nil {
//               tcf.Catch(r)
//           }
//       }()
//   }
//   tcf.Try()
// }

func my_req(client http.Client, url string, id int) DumpResponse {
	// client:=http.Client{
	//   Timeout: 5 * time.Second,
	// }

	req, err := http.NewRequest("GET", url, nil) //"http://country.io/capital.json", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("no request", err)
	}
	//fmt.Println(resp.Header)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("no body", err) //error
	} else {
		//fmt.Println(string(body))
		return DumpResponse{Code: int(resp.StatusCode), //strconv.Itoa(resp.StatusCode),
			Body: string(body),
			Id:   id}
	}
	return DumpResponse{Code: 100, Id: id}
}

func my_req_method(client http.Client, url string, method string, body string, id int) DumpResponse {
	if strings.ToUpper(method) != "POST" {
		method = "GET"
	}
	bodyByte := []byte{}
	if len(body) > 0 {
		bodyByte = []byte(body)
	}

	//
	// d1 := []byte(url)
	// err := os.WriteFile("log_"+strconv.Itoa(id)+".txt", d1, 0644)
	// if err!=nil { fmt.Println(err) }
	//
	// save file log
	// fn:=strings.Split(url, "/")
	// myid:=fn[len(fn)-1]
	// if 	myid =="id130" {
	// 	f,buff,err:=FileContextInit("logs/log_"+myid+".txt") //strconv.Itoa(id)+".txt")
	// 	if err!=nil {}
	// 	WriteFileBufferListNewline(buff, []string{url})
	// 	FileContextClose(f, buff)
	// }

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyByte)) //nil) //"http://country.io/capital.json", nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36")
	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("no request", err)
	}
	//fmt.Println(resp.Header)
	defer resp.Body.Close()
	read_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("no body", err) //error
	} else {
		//fmt.Println(string(body))
		return DumpResponse{Code: int(resp.StatusCode), //strconv.Itoa(resp.StatusCode),
			Body: string(read_body),
			Id:   id}
	}
	return DumpResponse{Code: 100, Id: id}
}

func MakeRequest(client http.Client, url string, id int, ch chan<- DumpResponse) {

	tools.Block{
		Try: func() {
			//fmt.Println("I tried")
			ch <- my_req(client, url, id)
			// if resp.Body != "" {
			//   //fmt.Print(resp.Code, resp.Body)
			//   ch <- resp //fmt.Sprintf("%.2f elapsed with response length: %d %s", len(resp.Body), url)
			// }
			//Throw("Oh,...sh...")
		},
		Catch: func(e interface{}) {
			ch <- DumpResponse{Code: 99, Id: id}

			//fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			//fmt.Println("Finally...")
		},
	}.Do()

	// resp:=Try( func(){return my_req(client, url)}, func(e){fmt.Print(e)})
	// if resp.Body != "" {
	//   //fmt.Print(resp.Code, resp.Body)
	//   ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", len(resp.Body), url)
	// }
}

func MakeRequestOptions(client http.Client, url string, method string, body string, id int, ch chan<- DumpResponse) {

	tools.Block{
		Try: func() {
			ch <- my_req_method(client, url, method, body, id)
		},
		Catch: func(e interface{}) {
			ch <- DumpResponse{Code: 99, Id: id} //fmt.Printf("Caught %v\n", e)
		},
		Finally: func() { //fmt.Println("Finally...")
		},
	}.Do()
}

func SimpleRequestTimeout(url string, timeout int) {
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	//resp, err :=
	client.Get(url)
	// return resp
}

func Request(url string, timeout int) *DumpResponse {
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
	ch := make(chan DumpResponse)
	//fmt.Println("start")
	go MakeRequest(client, url, 0, ch)
	//fmt.Println("done")
	res := <-ch
	close(ch)
	return &res
}

func RequestOptions(url string, method string, body string, timeout int) *DumpResponse {
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
	ch := make(chan DumpResponse)
	//fmt.Println("start")
	go MakeRequestOptions(client, url, method, body, 0, ch)
	//fmt.Println("done")
	res := <-ch
	close(ch)
	return &res
}

func Gather(urls []string, timeout int) []DumpResponse {
	start := time.Now()
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
	ch := make(chan DumpResponse)
	//fmt.Println("start")
	for i, url := range urls {
		go MakeRequest(client, url, i, ch)
	}
	out := make([]DumpResponse, len(urls)) //[]DumpResponse{} //make([]DumpResponse, len(urls))
	for range urls {
		res := <-ch
		out[res.Id] = res //out=append(out, <-ch)
	}
	close(ch)
	fmt.Printf("%.2fs elapsed %d\n", time.Since(start).Seconds(), len(urls))
	return out
}

func GatherOptions(urls []UrlWithMethod, timeout int) []DumpResponse {
	start := time.Now()
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
	ch := make(chan DumpResponse)
	//fmt.Println("start")
	for i, url := range urls {
		go MakeRequestOptions(client,
			url.Url,
			url.Method,
			url.Body, i, ch)
	}
	out := make([]DumpResponse, len(urls)) //[]DumpResponse{} //make([]DumpResponse, len(urls))
	for range urls {
		res := <-ch
		out[res.Id] = res //out=append(out, <-ch)
	}
	close(ch)
	fmt.Printf("%.2fs elapsed %d\n", time.Since(start).Seconds(), len(urls))
	return out
}

// func Gather(urls []Request, timeout int) []DumpResponse {
// 	start := time.Now()
// 	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
// 	ch := make(chan DumpResponse)
// 	//fmt.Println("start")
// 	for i, url := range urls { go MakeRequest(client, url, i, ch) }
// 	out := make([]DumpResponse, len(urls)) //[]DumpResponse{} //make([]DumpResponse, len(urls))
// 	for range urls {
// 		res := <-ch
// 		out[res.Id] = res //out=append(out, <-ch)
// 	}
// 	close(ch)
// 	fmt.Printf("%.2fs elapsed %d\n", time.Since(start).Seconds(), len(urls))
// 	return out
// }

func CopyUrls(url string, n int) []string {
	urls := make([]string, n)
	for i := 0; i < n; i++ {
		urls[i] = url
	}
	return urls
}

func Test() {

	// func() (v int, e error) {
	// 	defer func() {fmt.Print(e)}() // catch
	// 	defer func() {}() //finaly

	url := "https://jquery.com"
	client := http.Client{Timeout: 5 * time.Second}
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//ctx, _ := context.WithTimeout(context.TODO(), 5 * time.Second)
	//defer cancel()
	urls := []string{}
	for i := 1; i <= 100; i++ {
		urls = append(urls, url)
	}
	start := time.Now()
	ch := make(chan DumpResponse) //string)
	for i, url := range urls {
		go MakeRequest(client, url, i, ch)
	}
	for range urls {
		<-ch
		//fmt.Println(<-ch
	}
	fmt.Println("Hi")
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	//return

	// }()
	// fmt.Print("Hi")
}

func MainTest() {
	// One request
	res := Request("http://yandex.ru", 5)
	fmt.Println("test", res.Code)
	// Many requests
	// res=Gather(CopyUrls("http://yandex.ru", 10),10)
	// fmt.Println(len(res))
	// Test()
}

//res := requests.Request("http://yandex.ru", 5)
//log.Println("test",res.Code)
//res:=requests.Gather(requests.CopyUrls("http://yandex.ru", 10),10)
//log.Println(len(res))
//requests.Test()
