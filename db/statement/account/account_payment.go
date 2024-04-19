package account

import "fmt"

type AccountPaymentStmt struct{}

func (a AccountPaymentStmt) Update(id int, bank string, number string) string {
	return fmt.Sprintf("update account.account_payment set bank = '%s', number = '%s' where id = %d;", bank, number, id)
}
func (a AccountPaymentStmt) DeleteById(id int) string {
	return fmt.Sprintf("delete from account.account_payment where id = %d;", id)
}
