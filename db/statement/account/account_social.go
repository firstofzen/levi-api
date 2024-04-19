package account

import "fmt"

type AccountSocialStmt struct {
}

func (a AccountSocialStmt) UpdateArrayField(isInsert bool, fieldName string, id int) string {
	if isInsert {
		return fmt.Sprintf("update account.social set %s = array_append(%s, %d)", fieldName, fieldName, id)
	}
	return fmt.Sprintf("update account.social set %s = array_remove(%s, %d)", fieldName, fieldName, id)
}
func (a AccountSocialStmt) Delete(id int) string {
	return fmt.Sprintf("delete from account.social where id = %d;", id)
}
