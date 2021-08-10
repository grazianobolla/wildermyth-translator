package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var code_regex, _ = regexp.Compile(`\<.*?\>`)
var line_regex, _ = regexp.Compile(`=.*`)

func main() {
	var file_path_flag = flag.String("file", "", "File path")
	flag.Parse()

	process_file(*file_path_flag)
}

func process_file(file_path string) {
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if (len(line) > 0 && line[0] == '#') || len(line) <= 0 {
			continue
		}

		str := parse_line(line_regex.FindString(line)[1:])

		res := line_regex.ReplaceAllString(line, "="+str)
		fmt.Println(res)
	}
}

func parse_line(line string) string {
	prev_str := code_regex.FindAllString(line, -1)
	line = code_regex.ReplaceAllString(line, "(he)")
	res := translate_text(line)

	//put back code stuff
	index := 0
	for {
		if index >= len(prev_str) {
			break
		}

		res = strings.Replace(res, "(él)", prev_str[index], 1)
		index++
	}

	return res
}

func translate_text(text string) string {
	// str, err := translator_client.Translate([]string{"AMAZING"}, deepl.English, deepl.Spanish)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(str)
	return "translated_text"
}
