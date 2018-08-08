package gowoocommerce

type WcObject interface {
	Endpoint() string
	Extract(rawRecord map[string]string) WcObject
}

type WcError struct {
	Data struct {
		Status string `json:"status"`
	} `json:"data"`
}

/*
{
  "customer_note": "",
  "billing": {},
  "shipping": {},
  "payment_method": "tbz_betterpay_gateway",
  "payment_method_title": "Betterpay Payment Gateway",
  "transaction_id": "",
  "date_paid": "2018-08-07T09:35:28",
  "date_paid_gmt": "2018-08-07T09:35:28",
  "date_completed": null,
  "date_completed_gmt": null,
  "cart_hash": "ae0f49edf1f439b776db01fe7caafeba",
  "meta_data": [],
  "line_items": [],
  "tax_lines": [],
  "shipping_lines": [],
  "fee_lines": [],
  "coupon_lines": [],
  "refunds": [],
  "_links": {}
}

*/
