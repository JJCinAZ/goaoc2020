package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Passport struct {
	fields map[string]string
}

type Field struct {
	Field    string
	Req      bool
	RegMatch *regexp.Regexp
}

func main() {
	passports := processData(input)
	fmt.Println("Part1:")
	p1Fields := []Field{
		{"byr", true, nil},
		{"iyr", true, nil},
		{"eyr", true, nil},
		{"hgt", true, nil},
		{"hcl", true, nil},
		{"ecl", true, nil},
		{"pid", true, nil},
		{"cid", false, nil},
	}
	fmt.Println(countValid(passports, p1Fields))

	fmt.Println("Part2:")
	p2Fields := []Field{
		{"byr", true, regexp.MustCompile(`^(19[2-9][0-9])|(200[0-2])$`)},
		{"iyr", true, regexp.MustCompile(`^(201[0-9])|(2020)$`)},
		{"eyr", true, regexp.MustCompile(`^(202[0-9])|(2030)$`)},
		{"hgt", true, regexp.MustCompile(`^((1[5-8][0-9])cm$|^(19[0-3])cm$)|^((59|6[0-9]|7[0-6])in$)`)},
		{"hcl", true, regexp.MustCompile(`^#[0-9a-f]{6}$`)},
		{"ecl", true, regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)},
		{"pid", true, regexp.MustCompile(`^[0-9]{9}$`)},
		{"cid", false, nil},
	}
	fmt.Println(countValid(passports, p2Fields))

}

func countValid(passports []Passport, fieldset []Field) int {
	valid := 0
	for _, p := range passports {
		if isValid(p, fieldset) {
			valid++
		}
	}
	return valid
}

func isValid(p Passport, fieldset []Field) bool {
	for _, f := range fieldset {
		if f.Req {
			if value, found := p.fields[f.Field]; !found {
				return false
			} else {
				if f.RegMatch != nil && !f.RegMatch.MatchString(value) {
					fmt.Printf("%s:%s\n", f.Field, value)
					return false
				}
			}
		}
	}
	return true
}

func processData(input []string) []Passport {
	var (
		p Passport
	)
	passports := make([]Passport, 0)
	p.fields = make(map[string]string)
	for _, l := range input {
		if len(l) == 0 {
			if len(p.fields) > 0 {
				passports = append(passports, p)
			}
			p.fields = make(map[string]string)
			continue
		}
		for _, element := range strings.Split(l, " ") {
			if parts := strings.SplitN(element, ":", 2); len(parts) > 0 {
				if _, exists := p.fields[parts[0]]; exists {
					fmt.Printf("duplicate field %s\n", parts[0])
				}
				p.fields[parts[0]] = parts[1]
			} else {
				fmt.Printf("syntax error: %s\n", element)
			}
		}
	}
	if len(p.fields) > 0 {
		passports = append(passports, p)
	}
	fmt.Printf("Read %d passports, first PID: %s, last PID: %s\n", len(passports), passports[0].fields["pid"],
		passports[len(passports)-1].fields["pid"])
	return passports
}
