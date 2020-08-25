package main

import (
	"bufio"
	"compress/gzip"
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type Domain struct {
	Name     string
	Reg      string
	Create   string
	PaidTill string
	FreeDate string
	TLSVer   string
}

// TODO: rewrite this sheet to reflect
func (d Domain) Write() []string {
	s := []string{d.Name, d.TLSVer}
	return s
}

func main() {
	var (
		zones [1]string
		err   error
	)
	const (
		workers = 200
	)
	zones[0] = "ru"

	zoneUrl := fmt.Sprintf("https://partner.r01.ru/zones/%s_domains.gz", zones[0])
	zoneFile := fmt.Sprintf("%s_zone.gz", zones[0])
	zoneFileDecompressed := fmt.Sprintf("%s_zone", zones[0])

	err = DownloadFile(zoneFile, zoneUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = DecompressFile(zoneFile, zoneFileDecompressed)
	if err != nil {
		log.Fatal(err)
	}

	Domains := LoadData(zoneFileDecompressed)
	CheckTLSVersion(Domains, workers)
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	buf := make([]byte, 1024)
	bytesRead := 0

	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	for {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				_, err = out.Write(buf[:n])
				if err != nil {
					log.Fatal(err)
				}
				break
			}
			log.Fatal(err)
		}
		_, err = out.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
		bytesRead += n
	}
	defer resp.Body.Close()
	return err
}

func DecompressFile(gzipped string, ungzipped string) error {
	buf := make([]byte, 1024)
	input, err := os.Open(gzipped)
	if err != nil {
		log.Fatal(err)
	}
	r, err := gzip.NewReader(input)
	if err != nil {
		log.Fatal(err)
	}
	out, _ := os.Create(ungzipped)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				_, err = out.Write(buf[:n])
				if err != nil {
					log.Fatal(err)
				}
				break
			}
			log.Fatal(err)
		}

		fmt.Println(n)
		_, err = out.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}

	}
	return err
}

func LoadData(dataFilePath string) []Domain {
	var Domains []Domain
	fd, err := os.Open(dataFilePath)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(fd)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := strings.Fields(s.Text())
		Domains = append(Domains, Domain{line[0], line[1], line[2], line[3], line[4], ""})
	}
	return Domains
}

func CheckTLSVersion(Domains []Domain, workers int) {
	var (
		wg  sync.WaitGroup
		ch  = make(chan Domain)
		ch2 = make(chan Domain)
	)
	// create threads waiting on channel
	for i := 1; i < workers; i++ {
		wg.Add(1)
		go WrapSendRequest(ch, ch2, &wg)
	}

	// create writer to csv
	wg.Add(1)
	go WriteCSVFile(ch2)

	// Generate data and send to channel
	for i := 1; i < len(Domains); i++ {
		ch <- Domains[i]
	}
	// wait while all goroutines  will finish work
	wg.Wait()
}

func WrapSendRequest(ch chan Domain, ch2 chan Domain, wg *sync.WaitGroup) {
	// read data from channel and do staff
	for domain := range ch {
		domain = SendRequest(domain)
		ch2 <- domain
	}
	// when no data in the channel finish goroutine
	wg.Done()
}
func WriteCSVFile(ch chan Domain) {
	fd, _ := os.Create("./out.csv")
	w := csv.NewWriter(fd)
	for domain := range ch {
		w.Write(domain.Write())
		fmt.Printf("TLSVer: Domain: %s %s TLSVerEnd\n", domain.Name, domain.TLSVer)
	}
}
func SendRequest(domain Domain) Domain {
	var (
		tlsVersion string
	)
	url := fmt.Sprintf("https://%s", domain.Name)
	client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
	resp, err := client.Get(url)
	if err != nil {
		//log.Println(err)
		domain.TLSVer = "Could not connect to by https"
		return domain
	}
	if resp.StatusCode == 200 {
		log.Printf("domain: %s status code: %d", domain.Name, resp.StatusCode)
		if resp.TLS.Version == tls.VersionTLS10 {
			tlsVersion = "TLS 1.0"
		} else if resp.TLS.Version == tls.VersionTLS11 {
			tlsVersion = "TLS 1.1"
		} else if resp.TLS.Version == tls.VersionTLS12 {
			tlsVersion = "TLS 1.2"
		} else if resp.TLS.Version == tls.VersionTLS13 {
			tlsVersion = "TLS 1.3"
		} else if resp.TLS.Version == tls.VersionSSL30 {
			tlsVersion = "SSLv3"
		} else {
			tlsVersion = "unknown"
		}
		domain.TLSVer = tlsVersion
	}
	return domain
}
