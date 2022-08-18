// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/generated/two"
	"entgo.io/ent/dialect/sql"
)

// Two is the model entity for the Two schema.
type Two struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Something holds the value of the "something" field.
	Something string `json:"something,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Two) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case two.FieldID:
			values[i] = new(sql.NullInt64)
		case two.FieldSomething:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Two", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Two fields.
func (t *Two) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case two.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case two.FieldSomething:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field something", values[i])
			} else if value.Valid {
				t.Something = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Two.
// Note that you need to call Two.Unwrap() before calling this method if this Two
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Two) Update() *TwoUpdateOne {
	return (&TwoClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Two entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Two) Unwrap() *Two {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("generated: Two is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Two) String() string {
	var builder strings.Builder
	builder.WriteString("Two(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("something=")
	builder.WriteString(t.Something)
	builder.WriteByte(')')
	return builder.String()
}

// Twos is a parsable slice of Two.
type Twos []*Two

func (t Twos) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
