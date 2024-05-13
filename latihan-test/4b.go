package main

import "fmt"

type GoFood struct {
	orders []map[string]interface{}
}

func (g *GoFood) addOrder(product string, qty int, price int) {
	g.orders = append(g.orders, map[string]interface{}{
		"product": product,
		"qty":     qty,
		"price":   price,
	})
}

func (g *GoFood) getTransportFee(location string) int {
	if location == "pelita" {
		return 5000
	}
	return 10000
}

func (g *GoFood) getTotalOrder(location string, discountVoucher int) int {
	totalOrder := 0
	for _, order := range g.orders {
		totalOrder += order["qty"].(int) * order["price"].(int)
	}
	transportFee := g.getTransportFee(location)
	return totalOrder + transportFee - discountVoucher
}

type PremiumGoFood struct {
	*GoFood
}

func (eg *PremiumGoFood) getTransportFee(location string) int {
	baseFee := eg.GoFood.getTransportFee(location)
	if baseFee < 10000 {
		return baseFee + 3000
	}
	return baseFee
}

func (eg *PremiumGoFood) getTotalOrder(location string, discountVoucher int) int {
	totalOrder := 0
	for _, order := range eg.orders {
		totalOrder += order["qty"].(int) * order["price"].(int)
	}
	transportFee := eg.getTransportFee(location)
	return totalOrder + transportFee - discountVoucher
}

func main() {
	test := &PremiumGoFood{GoFood: &GoFood{}}
	test.addOrder("mie ayam", 2, 15000)
	test.addOrder("teh o beng", 2, 5000)
	fmt.Println(test.GoFood.getTotalOrder("pelita", 10000)) //use get transport fee method on gofood object
	fmt.Println(test.getTotalOrder("pelita", 10000))        //use get transport fee method on premiumgofood object
}
