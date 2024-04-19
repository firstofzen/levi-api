package product

import "fmt"

type ProductSkuStmt struct{}

func (p ProductSkuStmt) Get(pId int) string {
	return fmt.Sprintf("select * from product.sku where product_id = %d;", pId)
}

func (p ProductSkuStmt) Add(pId int, size int, name string, price float64) string {
	return fmt.Sprintf("insert into product.sku(product_id, size, name, price) values (%d, %d, '%s', %f);", pId, size, name, price)
}

func (p ProductSkuStmt) Delete(pId int) string {
	return fmt.Sprintf("delete from product.sku where product_id = %d;", pId)
}
