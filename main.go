package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type MyRecord map[string]string
type GrouppedDataType map[string]MyRecord

type Config struct {
	Source string
	Destination string
    ResultFilenameSuffix string
}

const (
	    Cbonds = "CbondsInterRates"
        MxRefinitive = "MXRefinitv"
        MoexCutOff = "MOEX_CUT_OFF"
        MoexOff = "MOEX_OFFICIAL"
        NsdOff = "NSD_OFFICIAL"
)

var ReportName string
var ReportFieldsMapper map[string]string
var ReportKeyFields = []string{}


func main() {
	var filename_param string
	params := os.Args[1:]
	if len(params) > 0 {
		filename_param = params[1]
	}

	cfg, err := LoadConfig("cfg_example.json")
	if err != nil {
		log.Fatalf("Error when reading config file: %v", err)
	}

	path_src := fmt.Sprintf("%v/%v", cfg.Source, filename_param)
	dst_filename, _ := strings.CutSuffix(filename_param, ".csv")
	path_dst := fmt.Sprintf("%v/%v%v.csv", cfg.Destination, dst_filename, cfg.ResultFilenameSuffix)

	fmt.Println(filename_param)
	ReportName, err = FindReportName(filename_param)
	if err != nil {
		log.Fatalf("Report name not found: %v", err)
	}
	ReportFieldsMapper, err = FindFieldMapper(ReportName)
	if err != nil {
		log.Fatalf("Report fields' mapper not found: %v", err)
	}
	ReportKeyFields, err = SelectReportKeys(ReportName)
	if err != nil {
		log.Fatalf("Report key not found: %v", err)
	}


	log.Printf("Started working on \"%v\" ...", filename_param)
	file, err := os.Open(path_src)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = FindDelimiter(ReportName)
	headers := []string{}
	GrouppedData := GrouppedDataType{}

	i := 0
	doubles := 0
	for ; true; {
		i += 1
		rec, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("Error when opening source file: %v\n", err)
		}
		if i == 1 {
			headers = rec
			continue
		}
		rec_ := MyRecord(ToMap(rec, headers))
		unique_key := KeyValues(ReportKeyFields, rec_)
		_, ok := GrouppedData[unique_key]
		if ok {
			doubles += 1
			// fmt.Printf("doubled key found: %v\n", unique_key)
			// fmt.Printf("Previous row: %v\n", GrouppedData[unique_key])
			// fmt.Printf("Current row: %v\n\n", rec_)
		}
		GrouppedData[unique_key] = rec_
		// if doubles > 20 {
		// 	// fmt.Println(GrouppedData)
		// 	log.Fatalf("STOP\n")
		// }
	}
	CreateCsv(GrouppedData, headers, &cfg, path_dst)
	log.Printf("Total rows: %v, doubles: %v\n", i, doubles)
}
