package product

import "fmt"

type ProductRatingStmt struct{}

func (p ProductRatingStmt) Get(pId int) string {
	return fmt.Sprintf("select * from product.rating where product_id = %d;", pId)
}
func (p ProductRatingStmt) Add(pId int, aId int, comment string, pictureUrl string, score float64) string {
	return fmt.Sprintf("insert into product.rating(product_id, account_id, comment, picture_url, score) values (%d, %d, '%s', '%s', %f);", pId, aId, comment, pictureUrl, score)
}
func (p ProductRatingStmt) UpdateField(pId int, aId int, fieldName string, fieldValue string) string {
	return fmt.Sprintf("update product.rating set %s = %s where product_id = %d and account_id = %d;", fieldName, fieldValue, pId, aId)
}
func (p ProductRatingStmt) DeleteByAccountID(pId int, aId int) string {
	return fmt.Sprintf("delete from product.rating where product_id = %d and account_id = %d;", pId, aId)
}
func (p ProductRatingStmt) DeleteByProductID(pId int) string {
	return fmt.Sprintf("delete from product.rating where product_id = %d;", pId)
}
