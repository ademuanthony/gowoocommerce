package main

import (
	"fmt"
	"github.com/ademuanthony/gowoocommerce"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	client := gowoocommerce.NewRestfulClient("http://shop.betterlifeglobal.org/wp-json/wc/v2/",
		"ck_8b79e61cf2b701833588771b6146f1dc71ab359c", "cs_113d48168e37ae0089ee0e399c4389cdab53524b")

	wcOrder := gowoocommerce.NewWcOrder(client)
	order, err := wcOrder.GetOne(200)
	if err != nil {
		fmt.Println(err)
	} else {
		spew.Dump(order)
	}

}
