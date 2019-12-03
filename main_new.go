package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"runtime"
)

func main() {
	var urlStr string
	countStr := make(chan int, 5)
	totalCh := make(chan int, 5)
	
	/*считываем строку ввода*/
	fmt.Print("Введите URL: ")
	fmt.Fscan(os.Stdin, &urlStr)
	
	/*переводим считаную строку в срез*/
	urlArray := strings.Split(urlStr, "\\n")
	
	/*высчитываем длину среза*/
	k:=len(urlArray)
	
	/*установим количество потоков*/
	runtime.GOMAXPROCS(k)
	
	/*анонимная функция высчитывает кол-во вхождений в теле ответа*/
	go func(urlArr []string) {
		for _, url := range urlArr {
			resp, err := http.Get(url)
			if err != nil {
				countStr <- 0
				fmt.Println(err)
				continue
			}
			b, err := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				countStr <- 0
				fmt.Println(err)
				continue
			}
			fmt.Println("Count for ", url, ":", strings.Count(string(b), "Go"))
			countStr <- strings.Count(string(b), "Go")
		}
	}(urlArray)
	/*анонимная функция высчитывает общее кол-во вхождений*/
	go func() {
		total := 0	
		for i := 0; i < k; i++ {
			countRes := <-countStr
			total += countRes
		}
		totalCh <- total
	} ()
	fmt.Println("Total: ", <- totalCh)
}