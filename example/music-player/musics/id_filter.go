package musics

type (
	// IDFilter is filter expression builder for field ID.
	IDFilter struct{}
)

var ID IDFilter

// Eq builds id = value filter.
func (f IDFilter) Eq(value int64) (exp Filter) {
	return Filter{
		expression: "id = ?",
		args:       []interface{}{value},
	}
}

// Neq builds id <> value filter.
func (f IDFilter) Neq(value int64) (exp Filter) {
	return Filter{
		expression: "id <> ?",
		args:       []interface{}{value},
	}
}

// Gt builds id > value filter.
func (f IDFilter) Gt(value int64) (exp Filter) {
	return Filter{
		expression: "id > ?",
		args:       []interface{}{value},
	}
}

// Lt builds id < value filter.
func (f IDFilter) Lt(value int64) (exp Filter) {
	return Filter{
		expression: "id < ?",
		args:       []interface{}{value},
	}
}

// Gte builds id >= value filter.
func (f IDFilter) Gte(value int64) (exp Filter) {
	return Filter{
		expression: "id >= ?",
		args:       []interface{}{value},
	}
}

// Lte builds id <= value filter.
func (f IDFilter) Lte(value int64) (exp Filter) {
	return Filter{
		expression: "id <= ?",
		args:       []interface{}{value},
	}
}

// Equal builds id = value filter.
func (f IDFilter) Equal(value int64) (exp Filter) {
	return f.Eq(value)
}

// NotEqual builds id <> value filter.
func (f IDFilter) NotEqual(value int64) (exp Filter) {
	return f.Neq(value)
}

// GreaterThan builds id > value filter.
func (f IDFilter) GreaterThan(value int64) (exp Filter) {
	return f.Gt(value)
}

// LessThan builds id < value filter.
func (f IDFilter) LessThan(value int64) (exp Filter) {
	return f.Lt(value)
}

// GreaterThanEqual builds id >= value filter.
func (f IDFilter) GreaterThanEqual(value int64) (exp Filter) {
	return f.Gte(value)
}

// LessThanEqual builds id <= value filter.
func (f IDFilter) LessThanEqual(value int64) (exp Filter) {
	return f.Lte(value)
}

// Between builds id > min AND id < max filter.
func (f IDFilter) Between(min, max int64) (exp Filter) {
	return Filter{
		expression: "id > ? AND id < ?",
		args:       []interface{}{min, max},
	}
}

// Range builds id >= min AND id <= max filter.
func (f IDFilter) Range(min, max int64) (exp Filter) {
	return Filter{
		expression: "id >= ? AND id <= ?",
		args:       []interface{}{min, max},
	}
}

// Name SQL field of ID.
func (f IDFilter) Name() string {
	return "id"
}
