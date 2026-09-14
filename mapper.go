package main

import "errors"

/*var SecQuotesMapper = map[string]string{
    "isin": "code_value",
    "source_id": "source_id",
    "description": "source",
    "main_source": "main_source",
    "update_date": "publication_date",
    "price_date": "use_date",
    "price_value": "price",
    "price_method_type_name": "price_unit",
    "iso_code_char": "price_currency_iso",
    "price_type_name": "price_type",
    "short_name": "additional_info_source",
    "fx_source_name": "price_for",
} */

var SecQuotesMapper = map[string]string{
    "isin": "code_value",    
    "source_id": "sourceid",
    "main_source": "mainsource",
    "price": "price_value",
    "price_currency_iso": "pricecurrencyiso",
    "price_for": "fx_source_name",
    "price_type": "price_type_name",
    "price_unit": "price_method_type_name",
    "publication_date": "update_date",
    "source": "description",
    "use_date": "price_date",
    "additional_info_source": "short_name",
}

var FxRatesMapper = map[string]string{
    "code1": "base_isocode_char",
    "iso_num1": "base_isocode_num",
    "code2": "quote_isocode_char",
    "iso_num2": "quote_isocode_num",
    "fx_rate_sub_type": "fxrate_type",
    "publication_date": "publication_date",
    "use_date": "use_date",
    "lot_size": "lot_size",
    "value": "fxrate_value",
    "is_public": "is_public",
}

var RatesMapper = map[string]string{
    "name": "instrument_name",
    "rate_sub_type": "fx_source_name",
    "tenor": "value_string",
    "currency": "iso_code_char",
    "publication_date": "calculation_publication_date",
    "use_date": "calculation_use_date",
    "value": "price_value",
    "rate_type_value1": "code_value",
    "is_public": "is_public",
}

func FindFieldMapper(report_name string) (map[string]string, error) {
    mapper := map[string]map[string]string{
        Cbonds: RatesMapper,
        MxRefinitive: FxRatesMapper,
        MoexCutOff: SecQuotesMapper,
        MoexOff: SecQuotesMapper,
        NsdOff: SecQuotesMapper,
    }
    val, ok := mapper[report_name]
    if !ok {return nil, errors.New("Unkhown report name")}
    return val, nil
}

func FindDelimiter(report_name string) rune {
    switch report_name {
    case MoexOff:
        return rune(';')
    default:
        return ','
    }
}
