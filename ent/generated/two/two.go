// Code generated by ent, DO NOT EDIT.

package two

const (
	// Label holds the string label denoting the two type in the database.
	Label = "two"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSomething holds the string denoting the something field in the database.
	FieldSomething = "something"
	// Table holds the table name of the two in the database.
	Table = "table_two"
)

// Columns holds all SQL columns for two fields.
var Columns = []string{
	FieldID,
	FieldSomething,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// SomethingValidator is a validator for the "something" field. It is called by the builders before save.
	SomethingValidator func(string) error
)