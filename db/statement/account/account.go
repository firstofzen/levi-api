package account

import "fmt"

type AccountStmt struct {
}

func (a AccountStmt) Init(email string, name string, avatarUrl string) string {
	return fmt.Sprintf("with a_id as (insert into account.account(email, name, avatar_url) values ('%s', '%s', '%s') returning account_id)"+
		"insert into account.detail(account_id) select * from a_id;"+
		"insert into account.payment(account_id) select * from a_id;"+
		"insert into account.social(account_id) select * from a_id;"+
		"insert into account.cart(account_id) select * from a_id;"+
		"insert into account.booth(account_id) select * from a_id;", email, name, avatarUrl)
}

func (a AccountStmt) GetById(id int) string {
	return fmt.Sprintf("select * from account.account inner join account.detail using(account_id) inner join account.payment using(account_id) inner join account.social using(account_id) where account_id = %d;", id)
}
func (a AccountStmt) GetByEmail(email string) string {
	return fmt.Sprintf("select * from account.account inner join account.detail using(account_id) inner join account.payment using(account_id) inner join account.social using(account_id) where account_id = '%s';", email)
}
func (a AccountStmt) UpdateField(fieldName string, fieldValue string, id int) string {
	return fmt.Sprintf("update account.account set %s = %s where id = %d;", fieldName, fieldValue, id)
}
func (a AccountStmt) CheckExistByEmail(email string) string {
	return fmt.Sprintf("select exists(select 1 from account.account where email = %s);", email)
}
func (a AccountStmt) Delete(id int) string {
	return fmt.Sprintf("delete from account.account where id = %d;", id)
}
