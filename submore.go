package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func reverseDomain(domain string) string {
	parts := strings.Split(domain, ".")
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.Join(parts, ".")
}

func main() {
	var filePath string
	var baseDomain string
	flag.StringVar(&filePath, "f", "", "Path to the input file containing subdomains")
	flag.StringVar(&baseDomain, "d", "", "The base domain to enforce scope (e.g., target.com)")
	flag.Parse()

	var reader io.Reader
	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Error: could not open file: %v", err)
		}
		defer file.Close()
		reader = file
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			reader = os.Stdin
		} else {
			fmt.Println("submore to sort and scope domains hierarchically.")
			fmt.Println("\nUsage:")
			fmt.Println("  domsort -f <file> [-d <domain>]")
			fmt.Println("  cat <file> | domsort [-d <domain>]")
			os.Exit(1)
		}
	}

	uniqueDomains := make(map[string]struct{})
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		uniqueDomains[line] = struct{}{}
		domain := line
		for {
			dotIndex := strings.Index(domain, ".")
			if dotIndex == -1 {
				break
			}
			parent := domain[dotIndex+1:]
			if strings.Contains(parent, ".") {
				uniqueDomains[parent] = struct{}{}
				domain = parent
			} else {
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: failed while reading input: %v", err)
	}

	var domains []string
	for domain := range uniqueDomains {
		domains = append(domains, domain)
	}
	sort.SliceStable(domains, func(i, j int) bool {
		reversedI := reverseDomain(domains[i])
		reversedJ := reverseDomain(domains[j])
		return reversedI < reversedJ
	})

	for _, domain := range domains {
		if baseDomain == "" {
			fmt.Println(domain)
			continue
		}
		if domain == baseDomain || strings.HasSuffix(domain, "."+baseDomain) {
			fmt.Println(domain)
		}
	}
}
