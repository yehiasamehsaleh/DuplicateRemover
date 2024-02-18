package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func hashLine(line string) string {
	hasher := sha256.New()
	hasher.Write([]byte(line))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func compareAndDeduplicateHashing(file1, file2 string) ([]string, []string) {
	seenHashes := make(map[string]bool)
	uniqueLines := []string{}
	repeatedLines := []string{}

	data1, err := ioutil.ReadFile(file1)
	if err != nil {
		fmt.Println("Error reading file1:", err)
		return nil, nil
	}

	lines1 := strings.Split(string(data1), "\n")

	data2, err := ioutil.ReadFile(file2)
	if err != nil {
		fmt.Println("Error reading file2:", err)
		return nil, nil
	}

	lines2 := strings.Split(string(data2), "\n")

	for _, line := range lines1 {
		hash := hashLine(line)
		if seenHashes[hash] {
			repeatedLines = append(repeatedLines, line)
		} else {
			seenHashes[hash] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	for _, line := range lines2 {
		hash := hashLine(line)
		if seenHashes[hash] {
			repeatedLines = append(repeatedLines, line)
		} else {
			seenHashes[hash] = true
			uniqueLines = append(uniqueLines, line)
		}
	}

	return uniqueLines, repeatedLines
}

func main() {
	file1 := "D://file1.txt"
	file2 := "D://file2.txt"

	uniqueLines, repeatedLines := compareAndDeduplicateHashing(file1, file2)

	// Print repeated lines
	fmt.Println("Repeated lines:")
	for _, line := range repeatedLines {
		fmt.Println(line)
	}

	// Write deduplicated content to file3.txt with error handling
	err := os.WriteFile("D://file3.txt", []byte(strings.Join(uniqueLines, "\n")), 0644)
	if err != nil {
		fmt.Println("Error writing to file3.txt:", err)
	} else {
		fmt.Println("Deduplicated content written to file3.txt")
	}
}
