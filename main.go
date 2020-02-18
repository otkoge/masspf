package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var domains = make(chan string)
var show_no_mx *bool
var pool_size *int

func domain_check(domain string) error {
	_, err := net.LookupMX(domain)
	if err != nil {
		if *show_no_mx {
			result := fmt.Sprintf("%s - no MX record", domain)
			fmt.Println(result)
		}
		return nil
	}
	txt_record, err := net.LookupTXT(domain)
	if err != nil {
		result := fmt.Sprintf("%s - no TXT record", domain)
		fmt.Println(result)
		return nil
	}
	for _, txt := range txt_record {
		if strings.HasPrefix(txt, "v=spf1") {
			result := fmt.Sprintf("%s %s", domain, txt)
			fmt.Println(result)
			return nil
		}
	}
	result := fmt.Sprintf("%s - no SPF record", domain)
	fmt.Println(result)
	return nil

}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for domain := range domains {
		_ = domain_check(domain)
	}
}

func create_workerpool() {
	var wg sync.WaitGroup
	for i := 0; i < *pool_size; i++ {
		wg.Add(1)
		fmt.Println(i)
		go worker(&wg)
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domains <- sc.Text()
	}
	close(domains)
	wg.Wait()
}

func main() {
	show_no_mx = flag.Bool("snm", false, "Prints domains that have no MX record set")
	pool_size = flag.Int("p", 20, "Size of the workers pool.")
	flag.Parse()
	create_workerpool()
}
