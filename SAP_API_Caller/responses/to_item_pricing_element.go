package responses

type ToItemPricingElement struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			SalesQuotation                 string `json:"SalesQuotation"`
			SalesQuotationItem             string `json:"SalesQuotationItem"`
			PricingProcedureStep           string `json:"PricingProcedureStep"`
			PricingProcedureCounter        string `json:"PricingProcedureCounter"`
			ConditionType                  string `json:"ConditionType"`
			PricingDateTime                string `json:"PricingDateTime"`
			PriceConditionDeterminationDte string `json:"PriceConditionDeterminationDte"`
			ConditionCalculationType       string `json:"ConditionCalculationType"`
			ConditionBaseValue             string `json:"ConditionBaseValue"`
			ConditionRateValue             string `json:"ConditionRateValue"`
			ConditionCurrency              string `json:"ConditionCurrency"`
			ConditionQuantity              string `json:"ConditionQuantity"`
			ConditionQuantityUnit          string `json:"ConditionQuantityUnit"`
			ConditionCategory              string `json:"ConditionCategory"`
			PricingScaleType               string `json:"PricingScaleType"`
			ConditionRecord                string `json:"ConditionRecord"`
			ConditionSequentialNumber      string `json:"ConditionSequentialNumber"`
			TaxCode                        string `json:"TaxCode"`
			ConditionAmount                string `json:"ConditionAmount"`
			TransactionCurrency            string `json:"TransactionCurrency"`
			ConditionClass                 string `json:"ConditionClass"`
			PricingScaleBasis              string `json:"PricingScaleBasis"`
			ConditionScaleBasisValue       string `json:"ConditionScaleBasisValue"`
			ConditionScaleBasisUnit        string `json:"ConditionScaleBasisUnit"`
			ConditionScaleBasisCurrency    string `json:"ConditionScaleBasisCurrency"`
			ConditionIsManuallyChanged     bool   `json:"ConditionIsManuallyChanged"`
		} `json:"results"`
	} `json:"d"`
}
