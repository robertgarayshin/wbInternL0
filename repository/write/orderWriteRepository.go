package write

import (
	"database/sql"
	"wbInternL0/models"
)

func NewOrder(db *sql.DB, order models.Order) error {
	// Метод для записи нового заказа в БД
	query := "INSERT INTO orders (orderuid, tracknumber, entry, name, phone, zip, city, address," +
		"region, email, transaction, requestid, currency, provider, amount, paymentdt, bank, deliverycost, goodstotal, " +
		"customfee, locale, internalsignature, customerid, deliveryservice, shardkey, smid, datecreated, oofshard) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, " +
		"$23, $24, $25, $26, $27, $28)"
	_, err := db.Exec(query,
		order.OrderUid, order.TrackNumber, order.Entry, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip,
		order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email, order.Payment.Transaction,
		order.Payment.RequestId, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt,
		order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee, order.Locale,
		order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated,
		order.OofShard)
	if err != nil {
		return err
	}
	for i := range order.Items {
		_, err = db.Exec("INSERT INTO order_item(chrtid, tracknumber, price, rid, name, sale, size, totalprice, nmid, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
			order.Items[i].ChrtId, order.Items[i].TrackNumber, order.Items[i].Price, order.Items[i].Rid,
			order.Items[i].Name, order.Items[i].Sale, order.Items[i].Size, order.Items[i].TotalPrice, order.Items[i].NmId,
			order.Items[i].Brand, order.Items[i].Status)
		if err != nil {
			return err
		}
	}
	return nil
}
