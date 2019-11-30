package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"strconv"
)

func main() {	
	countStr := make(chan string, 5)	 	
	total := 0
	for _, url := range os.Args[1:] {
		go getCountStr(url, countStr)
		countRes := <-countStr
		fmt.Println("Count for ", url, ":", countRes)
		countResInt, err := strconv.Atoi(countRes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}		
		total += countResInt
	}	
	fmt.Println("Total: ", total)		
}
/**
 * Считает количество вхождений строки в теле ответа запроса  
 */
func getCountStr(url string, countStr chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		countStr <- fmt.Sprint(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		countStr <- fmt.Sprint(err)
		return
	}	
	countStr <- fmt.Sprint(strings.Count(string(b), "Go"))	
}
