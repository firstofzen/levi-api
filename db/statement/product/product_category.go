package product

import "fmt"

type ProductCategoryStmt struct{}

func (p ProductCategoryStmt) Add(id int, l1 string, l2 string, l3 string) string {
	return fmt.Sprintf("insert into product.cateogry(id , l1, l2, l3) values (%d, '%s', '%s', '%s')", id, l1, l2, l3)
}
func (p ProductCategoryStmt) Delete(id int) string {
	return fmt.Sprintf("delete from product.cateogry where id = %d", id)
}