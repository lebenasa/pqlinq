package musics

type (
	// TitleFilter is filter expression builder for field Title.
	TitleFilter struct{}
)

var Title TitleFilter

// Eq builds title = value filter.
func (f TitleFilter) Eq(value string) (exp Filter) {
	return Filter{
		expression: "title = ?",
		args:       []interface{}{value},
	}
}

// Neq builds title <> value filter.
func (f TitleFilter) Neq(value string) (exp Filter) {
	return Filter{
		expression: "title <> ?",
		args:       []interface{}{value},
	}
}

// Gt builds title > value filter.
func (f TitleFilter) Gt(value string) (exp Filter) {
	return Filter{
		expression: "title > ?",
		args:       []interface{}{value},
	}
}

// Lt builds title < value filter.
func (f TitleFilter) Lt(value string) (exp Filter) {
	return Filter{
		expression: "title < ?",
		args:       []interface{}{value},
	}
}

// Gte builds title >= value filter.
func (f TitleFilter) Gte(value string) (exp Filter) {
	return Filter{
		expression: "title >= ?",
		args:       []interface{}{value},
	}
}

// Lte builds title <= value filter.
func (f TitleFilter) Lte(value string) (exp Filter) {
	return Filter{
		expression: "title <= ?",
		args:       []interface{}{value},
	}
}

// Equal builds title = value filter.
func (f TitleFilter) Equal(value string) (exp Filter) {
	return f.Eq(value)
}

// NotEqual builds title <> value filter.
func (f TitleFilter) NotEqual(value string) (exp Filter) {
	return f.Neq(value)
}

// GreaterThan builds title > value filter.
func (f TitleFilter) GreaterThan(value string) (exp Filter) {
	return f.Gt(value)
}

// LessThan builds title < value filter.
func (f TitleFilter) LessThan(value string) (exp Filter) {
	return f.Lt(value)
}

// GreaterThanEqual builds title >= value filter.
func (f TitleFilter) GreaterThanEqual(value string) (exp Filter) {
	return f.Gte(value)
}

// LessThanEqual builds title <= value filter.
func (f TitleFilter) LessThanEqual(value string) (exp Filter) {
	return f.Lte(value)
}

// Between builds title > min AND title < max filter.
func (f TitleFilter) Between(min, max string) (exp Filter) {
	return Filter{
		expression: "title > ? AND title < ?",
		args:       []interface{}{min, max},
	}
}

// Range builds title >= min AND title <= max filter.
func (f TitleFilter) Range(min, max string) (exp Filter) {
	return Filter{
		expression: "title >= ? AND title <= ?",
		args:       []interface{}{min, max},
	}
}

// Name SQL field of Title.
func (f TitleFilter) Name() string {
	return "title"
}
