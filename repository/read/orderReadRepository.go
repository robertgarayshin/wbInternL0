package read

import (
	"database/sql"
	"fmt"
	"wbInternL0/models"
)

func ReadAll(db *sql.DB) []models.Order {
	var orders []models.Order
	rows, err := db.Query("select * from orders")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		o := models.Order{}
		err := rows.Scan(&o.OrderUid, &o.TrackNumber, &o.Entry, &o.Delivery.Name, &o.Delivery.Phone, &o.Delivery.Zip,
			&o.Delivery.City, &o.Delivery.Address, &o.Delivery.Region, &o.Delivery.Email, &o.Payment.Transaction,
			&o.Payment.RequestId, &o.Payment.Currency, &o.Payment.Provider, &o.Payment.Amount, &o.Payment.PaymentDt,
			&o.Payment.Bank, &o.Payment.DeliveryCost, &o.Payment.GoodsTotal, &o.Payment.CustomFee, &o.Locale,
			&o.InternalSignature, &o.CustomerId, &o.DeliveryService, &o.Shardkey, &o.SmId, &o.DateCreated, &o.OofShard)
		if err != nil {
			fmt.Println(err)
			continue
		}
		orders = append(orders, o)
	}
	err = rows.Close()
	if err != nil {
		return nil
	}
	for i := range orders {
		var item models.Item
		var items []models.Item
		rows, err = db.Query("select * from order_item where tracknumber = $1", orders[i].TrackNumber)
		for rows.Next() {
			err = rows.Scan(&item.ChrtId, &item.TrackNumber, &item.Price,
				&item.Rid, &item.Name, &item.Sale,
				&item.Size, &item.TotalPrice, &item.NmId,
				&item.Brand, &item.Status)
			items = append(items, item)
		}
		err := rows.Close()
		if err != nil {
			return nil
		}
		orders[i].Items = items
	}
	return orders
}
