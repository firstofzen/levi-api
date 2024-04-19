package trade

import "fmt"

type OrderShip struct{}

func (o OrderShip) Add(oId int, price float64, shippingUnit string) string {
	return fmt.Sprintf("insert into trade.order_ship(order_id, price, shipping_unit) values (%d, %f, '%s');", oId, price, shippingUnit)
}
func (o OrderShip) UpdateState(state string, oId int, des string) string {
	return fmt.Sprintf("update trade.order_ship set shipping_state = '%s', last_update = current_timestamp and description = '%s' where order_id = %d;", state, des, oId)
}
func (o OrderShip) Delete(oId int) string {
	return fmt.Sprintf("delete from trade.order_ship where order_id = %d;", oId)
}
