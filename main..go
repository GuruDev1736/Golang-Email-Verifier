package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Print("This is the email verifier tool project \n")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("domain , hasMX , hasSPF , SPFRecord , hasDMARC , DMARCRecord")
	fmt.Println("Enter the Email : ")

	for scanner.Scan() {
		CheckDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read form the input %v \n", err)
	}
}

func CheckDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var SPFRecord, DMARCRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error : %v", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	TXTRecord, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error : %v", err)
	}

	for _, record := range TXTRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			SPFRecord = record
			break
		}
	}

	dmarc, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error : %v ", err)
	}

	for _, dmarcRecord := range dmarc {
		strings.HasPrefix(dmarcRecord, "v=DMARC1")
		hasDMARC = true
		DMARCRecord = dmarcRecord
		break
	}

	fmt.Printf("The Domain is : %v\n", domain)
	fmt.Printf("The Domain has MX : %v\n", hasMX)
	fmt.Printf("The Domain has SPF is : %v\n", hasSPF)
	fmt.Printf("The Domain has SPFRecord : %v\n", SPFRecord)
	fmt.Printf("The Domain has DMARC : %v\n", hasDMARC)
	fmt.Printf("The Domain has DMARCRecord : %v\n", DMARCRecord)

}
