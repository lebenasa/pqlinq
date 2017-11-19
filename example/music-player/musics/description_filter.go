package musics

type (
	// DescriptionFilter is filter expression builder for field Description.
	DescriptionFilter struct{}
)

var Description DescriptionFilter

// Eq builds description = value filter.
func (f DescriptionFilter) Eq(value string) (exp Filter) {
	return Filter{
		expression: "description = ?",
		args:       []interface{}{value},
	}
}

// Neq builds description <> value filter.
func (f DescriptionFilter) Neq(value string) (exp Filter) {
	return Filter{
		expression: "description <> ?",
		args:       []interface{}{value},
	}
}

// Gt builds description > value filter.
func (f DescriptionFilter) Gt(value string) (exp Filter) {
	return Filter{
		expression: "description > ?",
		args:       []interface{}{value},
	}
}

// Lt builds description < value filter.
func (f DescriptionFilter) Lt(value string) (exp Filter) {
	return Filter{
		expression: "description < ?",
		args:       []interface{}{value},
	}
}

// Gte builds description >= value filter.
func (f DescriptionFilter) Gte(value string) (exp Filter) {
	return Filter{
		expression: "description >= ?",
		args:       []interface{}{value},
	}
}

// Lte builds description <= value filter.
func (f DescriptionFilter) Lte(value string) (exp Filter) {
	return Filter{
		expression: "description <= ?",
		args:       []interface{}{value},
	}
}

// Equal builds description = value filter.
func (f DescriptionFilter) Equal(value string) (exp Filter) {
	return f.Eq(value)
}

// NotEqual builds description <> value filter.
func (f DescriptionFilter) NotEqual(value string) (exp Filter) {
	return f.Neq(value)
}

// GreaterThan builds description > value filter.
func (f DescriptionFilter) GreaterThan(value string) (exp Filter) {
	return f.Gt(value)
}

// LessThan builds description < value filter.
func (f DescriptionFilter) LessThan(value string) (exp Filter) {
	return f.Lt(value)
}

// GreaterThanEqual builds description >= value filter.
func (f DescriptionFilter) GreaterThanEqual(value string) (exp Filter) {
	return f.Gte(value)
}

// LessThanEqual builds description <= value filter.
func (f DescriptionFilter) LessThanEqual(value string) (exp Filter) {
	return f.Lte(value)
}

// Between builds description > min AND description < max filter.
func (f DescriptionFilter) Between(min, max string) (exp Filter) {
	return Filter{
		expression: "description > ? AND description < ?",
		args:       []interface{}{min, max},
	}
}

// Range builds description >= min AND description <= max filter.
func (f DescriptionFilter) Range(min, max string) (exp Filter) {
	return Filter{
		expression: "description >= ? AND description <= ?",
		args:       []interface{}{min, max},
	}
}

// Name SQL field of Description.
func (f DescriptionFilter) Name() string {
	return "description"
}
