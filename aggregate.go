package squrl

import "fmt"

func Avg(field, table string) string {
	return fmt.Sprintf(`AVG("%v".%v)`, table, field)
}

func Count(field, table string) string {
	return fmt.Sprintf(`COUNT("%v".%v)`, table, field)
}

func Max(field, table string) string {
	return fmt.Sprintf(`MAX("%v".%v)`, table, field)
}

func Min(field, table string) string {
	return fmt.Sprintf(`MIN("%v".%v)`, table, field)
}

func Sum(field, table string) string {
	return fmt.Sprintf(`SUM("%v".%v)`, table, field)
}
