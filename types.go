package squrl

type queryType string

const (
	selectType queryType = "SELECT"
	insertType queryType = "INSERT"
	updateType queryType = "UPDATE"
	deleteType queryType = "DELETE"
)
