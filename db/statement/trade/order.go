package trade

import "fmt"

type OrderStmt struct{}

func (o OrderStmt) Add(aId int, pId int, totalAmount float64) string {
	return fmt.Sprintf("insert into trade.order(account_id, product_id, total_amount) values (%d, %d, %f);", aId, pId, totalAmount)
}
func (o OrderStmt) GetByAccountId(accountId int) string {
	return fmt.Sprintf("select * from trade.order inner join order.ship using(order_id) where account_id = %d;", accountId)
}
func (o OrderStmt) UpdateStatus(aId int, pId int, status string) string {
	return fmt.Sprintf("update trade.order set status = '%s', last_update = current_timestamp where account_id = %d and product_id = %d;", status, aId, pId)
}
func (o OrderStmt) Delete(aId int, pId int) string {
	return fmt.Sprintf("delete from trade.order where account_id = %d and product_id = %d;", aId, pId)
}
