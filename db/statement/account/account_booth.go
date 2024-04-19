package account

import "fmt"

type AccountBoothStmt struct{}

func (a AccountBoothStmt) GetById(id int) string {
	return fmt.Sprintf("select * from account.booth where id = %d;", id)
}

func (a AccountBoothStmt) UpdateField(id int, fieldName string, fieldValue string) string {
	return fmt.Sprintf("update account.booth set %s = %s where id = %d;", fieldName, fieldValue, id)
}
func (a AccountBoothStmt) UpdateArrayField(id int, isInsert bool, fieldName string, fieldValue string) string {
	if isInsert {
		return fmt.Sprintf("update account.booth set %s = array_append(%s, %s) where id = %d", fieldName, fieldName, fieldValue, id)
	}
	return fmt.Sprintf("update account.booth set %s = array_remove(%s, %s) where id = %d", fieldName, fieldName, fieldValue, id)
}
func (a AccountBoothStmt) Delete(id int) string {
	return fmt.Sprintf("delete from account.booth where id = %d;", id)
}
