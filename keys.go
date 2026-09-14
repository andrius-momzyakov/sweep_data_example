package main

import "errors"


var keys_sec_quotes = []string{
	"isin",
	"source",
	"source_id",
	"price_type",
	"publication_date",
	"use_date",
	"main_source",
	"additional_info_source",
	"price_unit",
	"price_currency_iso",
}

var keys_fx_rates = []string{
	"code1",
	"code2",
	"fx_rate_sub_type",
	"is_public",
	"publication_date",
	"use_date",
}

var keys_rates = []string{
	"name",
	"rate_sub_type",
	"currency",
	"is_public",
	"publication_date",
	"use_date",
}


func SelectReportKeys(report_name string) ([]string, error) {
	mapper := map[string][]string{
        Cbonds: keys_rates,
        MxRefinitive: keys_fx_rates,
        MoexCutOff: keys_sec_quotes,
        MoexOff: keys_sec_quotes,
        NsdOff: keys_sec_quotes,
	}
	val, ok := mapper[report_name]
	if !ok {return nil, errors.New("Unknown report name")}
	return val, nil
}
