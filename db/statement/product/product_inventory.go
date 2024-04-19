package product

import "fmt"

type ProductInventoryStmt struct{}

func (p ProductInventoryStmt) Add(id int) string {
	return fmt.Sprintf("insert into product.inventory(product_id) values (%d);", id)
}
func (p ProductInventoryStmt) UpdateField(id int, fieldName string, fieldValue int) string {
	return fmt.Sprintf("update product.inventory set %s = %d where product_id = %d;", fieldName, fieldValue, id)
}
func (p ProductInventoryStmt) Delete(id int) string {
	return fmt.Sprintf("delete from product.inventory where product_id = %d;", id)
}
