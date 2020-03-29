package repo

import "log"

func FindRecentFactEntities(restaurantId string) []OrderInvoice {

	orders := FindRecentOrders(restaurantId)
	var orderInvoices []OrderInvoice
	for _, order := range orders {
		subTotalCents := 0
		for _, orderItem := range order.Items {
			itemTotal := orderItem.Quantity * orderItem.PriceCents
			subTotalCents = subTotalCents + itemTotal
			orderItem.TotalAmountCents = itemTotal
		}
		subTotalDec := decimal.NewFromInt(int64(subTotalCents))
		taxRate := decimal.NewFromFloat(TAX_RATE)
		taxDec := taxRate.Mul(subTotalDec)
		total := subTotalDec.Add(taxDec)

		invoice := Invoice{
			OrderId:             order.Id,
			RestaurantId:        order.RestaurantId,
			TaxType:             "HST",
			Status:              "UNPAID",
			SubTotalAmountCents: int64(subTotalCents),
			SubTotalAmount:      subTotalDec.StringFixed(2),
			TaxAmountCents:      taxDec.IntPart(),
			TaxAmount:           subTotalDec.StringFixed(2),
			DiscountCents:       0,
			Discount:            "",
			TotalAmountCents:    total.IntPart(),
			TotalAmount:         total.StringFixed(2),
			TipAmountCents:      0,
			TipAmount:           "",
			DeliveryChargeCents: 0,
			DeliveryCharge:      "",
		}
		orderInvoice := OrderInvoice{
			InvoiceInfo: invoice,
			OrderInfo:   order,
		}
		orderInvoices = append(orderInvoices, orderInvoice)
	}
	return orderInvoices;
}

func SaveFactEntity(order *Order, status string) {
	order.Status = status
	cloudstore.AddWithId("orders", order.Id, order)
	if order.Status != PENDING_CREATION {
		channel := LiveUpdateChannelMap[order.RestaurantId]
		if channel != nil {
			channel <- order
		}
	} else {
		log.Println("Order is pending so not sending update: " , order)
	}
}

func SaveDimensionEntity(order *Order, status string) {
	order.Status = status
	cloudstore.AddWithId("orders", order.Id, order)
	if order.Status != PENDING_CREATION {
		channel := LiveUpdateChannelMap[order.RestaurantId]
		if channel != nil {
			channel <- order
		}
	} else {
		log.Println("Order is pending so not sending update: " , order)
	}
}