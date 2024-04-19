package account

import "fmt"

type AccountDetailStmt struct{}

func (a AccountDetailStmt) UpdateAField(fieldName string, fieldValue string, id int) string {
	return fmt.Sprintf("update account.account set %s = %s where id = %d;", fieldName, fieldValue, id)
}

func (a AccountDetailStmt) DeleteById(id int) string {
	return fmt.Sprintf("delete from account.account where id = %d;", id)
}
