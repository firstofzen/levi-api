package product

import (
	"fmt"
	"time"
)

type ProductSaleStmt struct{}

func (p *ProductSaleStmt) Add(pId int) string {
	return fmt.Sprintf("insert into product.sale(product_id) values (%d);", pId)
}
func (p *ProductSaleStmt) Update(expire time.Time, discount float64, pId int) string {
	return fmt.Sprintf("update product.sale set expire = '%s', discount = %f, create_at = current_timestamp where product_id = %d;", expire.String(), discount, pId)
}
func (p *ProductSaleStmt) Delete(pId int) string {
	return fmt.Sprintf("delete from product.sale where product_id = %d;", pId)
}
