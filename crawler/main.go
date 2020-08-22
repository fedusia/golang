package main

import (
	"bufio"
	"compress/gzip"
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

func main() {
	var zones [1]string
	var err error
	zones[0] = "ru"

	zoneFile := fmt.Sprintf("%s_zone.gz", zones[0])
	zoneFileDecompressed := "zone_ru"

	err := DownloadFile(zoneFile, zoneUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = DecompressFile(zoneFile, zoneFileDecompressed)
	if err != nil {
		log.Fatal(err)
	}

	Domains := LoadData(zoneFileDecompressed)
	err = CheckTLSVersion(Domains)
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

func CheckTLSVersion(Domains []Domain) error {
	var err error
	var wg sync.WaitGroup
	for i := 1; i < len(Domains); i++ {
		wg.Add(1)
		go SendRequest(Domains[i])
	}
	wg.Wait()
	return err
}

func SendRequest(domain Domain) error {
	url := fmt.Sprintf("https://%s", domain.Name)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(resp.TLS)
	return err
}
