package gowoocommerce

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

type Order struct {
	Id                 int        `json:"id"`
	ParentId           int        `json:"parent_id"`
	Number             string     `json:"number"`
	OrderKey           string     `json:"order_key"`
	CreatedVia         string     `json:"created_via"`
	Version            string     `json:"version"`
	Status             string     `json:"status"`
	Currency           string     `json:"currency"`
	DateCreated        string     `json:"date_created"`
	DateCreatedGmt     string     `json:"date_created_gmt"`
	ModifiedDate       string     `json:"modified_date"`
	ModifiedDateGmt    string     `json:"modified_date_gmt"`
	DiscountTotal      float64    `json:"discount_total,string"`
	DiscountTax        float64    `json:"discount_tax,string"`
	ShippingTotal      float64    `json:"shipping_total,string"`
	ShippingTax        float64    `json:"shipping_tax,string"`
	CartTax            float64    `json:"cart_tax,string"`
	TotalTax           float64    `json:"total_tax,string"`
	Total              float64    `json:"total,string"`
	PriceIncludeTax    bool       `json:"price_include_tax"`
	CustomerId         int        `json:"customer_id"`
	CustomerAddressId  int        `json:"customer_address_id"`
	CustomerUserAgent  string     `json:"customer_user_agent"`
	CustomerNote       string     `json:"customer_note"`
	Billing            struct{}   `json:"billing"`
	Shipping           struct{}   `json:"shipping"`
	PaymentMethod      string     `json:"payment_method"`
	PaymentMethodTitle string     `json:"payment_method_title"`
	TransactionId      string     `json:"transaction_id"`
	DatePaid           string     `json:"date_paid"`
	DatePainGmt        string     `json:"date_pain_gmt"`
	DateCompleted      string     `json:"date_completed"`
	DateCompledGmt     string     `json:"date_compled_gmt"`
	CartHash           string     `json:"cart_hash"`
	MetaData           []struct{} `json:"meta_data"`
	LineItems          []LineItem `json:"line_items"`
	TaxLines           []struct{} `json:"tax_lines"`
	ShippingLines      []struct{} `json:"shipping_lines"`
	FeeLines           []struct{} `json:"fee_lines"`
	CouponLines        []struct{} `json:"coupon_lines"`
	Links              struct{}   `json:"_links"`
}

type LineItem struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	ProductId   int     `json:"product_id"`
	VariationId int     `json:"variation_id"`
	Quantity    int     `json:"quantity"`
	SubTotal    float64 `json:"sub_total"`
	Total       float64 `json:"total"`
	Sku         string  `json:"sku"`
	Price       float64 `json:"price"`
}

type wcOrder struct {
	client *restfulClient
}

func NewWcOrder(client *restfulClient) wcOrder {
	return wcOrder{client: client}
}

func (this wcOrder) Endpoint() string {
	return "orders"
}

func (this wcOrder) extract(record map[string]interface{}) Order {
	order := Order{
		Id:       record["id"].(int),
		Version:  record["version"].(string),
		Status:   record["status"].(string),
		CartHash: record["cart_hash"].(string),
		Number:   record["number"].(string),
	}
	order.Total, _ = strconv.ParseFloat(record["total"].(string), 10)
	return order
}

func (this wcOrder) GetAll(params map[string]string) ([]Order, error) {
	resString, err := this.client.request(this.Endpoint(), RequestMethods.GET, nil, params)
	var orders []Order
	if err != nil {
		return orders, err
	}

	var out []map[string]interface{}
	err = json.Unmarshal([]byte(resString), &out)
	if err != nil {
		fmt.Println("\n", err, "\n")
		return orders, err
	}

	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           &orders,
		WeaklyTypedInput: true,
		TagName:          "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return orders, err
	}

	err = decoder.Decode(out)

	//err = mapstructure.WeakDecode(out, &orders)
	return orders, err
}

func (this wcOrder) GetOne(id int) (Order, error) {
	var order Order

	resString, err := this.client.request(fmt.Sprintf("%s/%d", this.Endpoint(), id), RequestMethods.GET, nil, nil)
	if err != nil {
		return order, err
	}

	//fmt.Printf("\n\n", resString, "\n\n")

	var out map[string]interface{}
	err = json.Unmarshal([]byte(resString), &out)
	if err != nil {
		fmt.Println("\n", err, "\n")
		return order, err
	}

	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           &order,
		WeaklyTypedInput: true,
		TagName:          "json",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return order, err
	}

	err = decoder.Decode(out)

	return order, nil
}
