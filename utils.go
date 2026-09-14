package main

import (
	"os"
	"strings"
	"log"
	"errors"
)

type MyMap map[string]string

func ToMap(record []string, header []string) (res MyMap) {
	res = MyMap{}
	for i, name := range record {
		res[header[i]] = name
	}
	return
}

func KeyValues(keys []string, data map[string]string) string {
	res := []string{}
	for _, key := range keys {
		translated_key := ReportFieldsMapper[key]
		value, ok := data[translated_key]
		if !ok {
			res = append(res, "")
		} else {
			res = append(res, value)
		}
	}
	return strings.Join(res, ",")
}

func csvRow(row map[string]string, headers []string) string {
	res := []string{}
	for _, name := range headers {
		res = append(res, row[name])
	}
	return strings.Join(res, ",")
}

func CreateCsv(data GrouppedDataType, headers []string, cfg *Config, dst_path string) {
	var s string
	log.Printf("Writing to: %v", dst_path)
	file, err := os.Create(dst_path)
	if err != nil {
		log.Fatalf("Error creating output file: %v\n", err)
	}
	defer file.Close()
	i := 0
	for _, row := range data {
		i += 1
		if i == 1 {
			s = strings.Join(headers, ",")
		} else {
			s = csvRow(row, headers)
		}
		file.WriteString(s + "\n")
	}
}

func FindReportName(filename string) (string, error) {
	if strings.HasPrefix(filename, Cbonds) {return Cbonds, nil}
	if strings.HasPrefix(filename, MxRefinitive) {return MxRefinitive, nil}
	if strings.HasPrefix(filename, MoexCutOff) {return MoexCutOff, nil}
	if strings.HasPrefix(filename, MoexOff) {return MoexOff, nil}
	if strings.HasPrefix(filename, NsdOff) {return NsdOff, nil}
	return "", errors.New("Unknown report name")
}
