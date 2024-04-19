package statement

import "strings"

func Transaction(str ...string) string {
	var rs strings.Builder
	rs.WriteString("begin; ")
	for _, s := range str {
		rs.WriteString(s)
	}
	rs.WriteString(" commit;")
	return rs.String()
}
