package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var reA *regexp.Regexp = regexp.MustCompile(`(^0[1-9]|[1-2][0-9]|3[0-1]).(0[1-9]|1[0-2]).([0-2][0-9]{3})`) //dd.mm.YYYY и в формате YYYY/mm/dd
var reB *regexp.Regexp = regexp.MustCompile(`(^[0-2][0-9]{3})\/(0[1-9]|1[0-2])\/(0[1-9]|[1-2][0-9]|3[0-1])`)

func main() {
	dates := []string{
		"12.09.1978",
		"31.09.1978",
		"1990/06/10",
		"08.03.2021",
		"12.04.1986",
		"25 dec 1988",
		"2001/05/25",
	}

	for _, d := range dates {
		year, month, day, err := parseDate(d)
		if err != nil {
			fmt.Println("ERROR!", err, "-", d)
			continue
		}

		fmt.Printf("%04d.%02d.%02d\n", year, month, day)
	}
}

var UndefinedDateFormat = errors.New("undefined date format")

func parseDate(date string) (year, month, day int64, err error) {
	// TODO: try dd.mm.YYYY format	
	submatches := reA.FindAllStringSubmatch(date, -1)
	
	if len(submatches) == 0 {
		submatches = reB.FindAllStringSubmatch(date, -1)		
		if len(submatches) == 0 {
			err = UndefinedDateFormat
			return
		}
	year = strToInt(submatches[0][1])
	month = strToInt(submatches[0][2]) 
	 day = strToInt(submatches[0][3]) 
		return
	}
	year = strToInt(submatches[0][3])
	month = strToInt(submatches[0][2])
	day = strToInt(submatches[0][1])
	return
}

func strToInt(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
