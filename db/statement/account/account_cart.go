package account

import "fmt"

type AccountCartStmt struct{}

func (a AccountCartStmt) GetById(id int) string {
	return fmt.Sprintf("select * from account.cart where account_id = %d;", id)
}

func (a AccountCartStmt) UpdateAField(fieldName string, fieldValue int, id int) string {
	return fmt.Sprintf("update account.cart set %s = %d where id = %d;", fieldName, fieldValue, id)
}
func (a AccountCartStmt) Delete(id int) string {
	return fmt.Sprintf("delete from account.cart where id = %d;", id)
}
