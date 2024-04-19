package product

import "fmt"

type ProductStmt struct{}

func (p ProductStmt) GetById(id int) string {
	return fmt.Sprintf("select * from product.product inner join product.category using(product_id) inner join product.inventory using(product_id) inner join product.sale using(product_id) where product_id = %d;", id)
}

func (p ProductStmt) GetByCategory(l1 string, l2 string, l3 string) string {
	return fmt.Sprintf("select * from product.product inner join product.category as c using(product_id) inner join product.inventory using(product_id) inner join product.sale using(product_id)  where c.l1 = '%s' and c.l2 = '%s' and c.l3 = '%s'", l1, l2, l3)
}

func (p ProductStmt) Add(accId int, name string, avatarUrl string, pictureDescriptions []string, description string, price float64) string {
	return fmt.Sprintf("insert into product.product(account_id, name, avatar_url, picture_descriptions, description, price, total_score) values (%d, '%s', '%s', '%s', '%s', %f, 0);", accId, name, avatarUrl, pictureDescriptions, description, price)
}

func (p ProductStmt) UpdateField(id int, fieldName string, fieldValue string) string {
	return fmt.Sprintf("update product.product set %s = %s where product_id = %d;", fieldName, fieldValue, id)
}

func (p ProductStmt) UpdateArrayField(id int, isInsert bool, fieldName string, fieldValue string) string {
	return fmt.Sprintf("update product.product set %s = array_append(%s, %s) where product_id = %d;", fieldName, fieldName, fieldValue, id)
}

func (p ProductStmt) Delete(id int) string {
	return fmt.Sprintf("delete from product.product where product_id = %d;", id)
}
