package musics

type (
	// RatingFilter is filter expression builder for field Rating.
	RatingFilter struct{}
)

var Rating RatingFilter

// Eq builds rating = value filter.
func (f RatingFilter) Eq(value float64) (exp Filter) {
	return Filter{
		expression: "rating = ?",
		args:       []interface{}{value},
	}
}

// Neq builds rating <> value filter.
func (f RatingFilter) Neq(value float64) (exp Filter) {
	return Filter{
		expression: "rating <> ?",
		args:       []interface{}{value},
	}
}

// Gt builds rating > value filter.
func (f RatingFilter) Gt(value float64) (exp Filter) {
	return Filter{
		expression: "rating > ?",
		args:       []interface{}{value},
	}
}

// Lt builds rating < value filter.
func (f RatingFilter) Lt(value float64) (exp Filter) {
	return Filter{
		expression: "rating < ?",
		args:       []interface{}{value},
	}
}

// Gte builds rating >= value filter.
func (f RatingFilter) Gte(value float64) (exp Filter) {
	return Filter{
		expression: "rating >= ?",
		args:       []interface{}{value},
	}
}

// Lte builds rating <= value filter.
func (f RatingFilter) Lte(value float64) (exp Filter) {
	return Filter{
		expression: "rating <= ?",
		args:       []interface{}{value},
	}
}

// Equal builds rating = value filter.
func (f RatingFilter) Equal(value float64) (exp Filter) {
	return f.Eq(value)
}

// NotEqual builds rating <> value filter.
func (f RatingFilter) NotEqual(value float64) (exp Filter) {
	return f.Neq(value)
}

// GreaterThan builds rating > value filter.
func (f RatingFilter) GreaterThan(value float64) (exp Filter) {
	return f.Gt(value)
}

// LessThan builds rating < value filter.
func (f RatingFilter) LessThan(value float64) (exp Filter) {
	return f.Lt(value)
}

// GreaterThanEqual builds rating >= value filter.
func (f RatingFilter) GreaterThanEqual(value float64) (exp Filter) {
	return f.Gte(value)
}

// LessThanEqual builds rating <= value filter.
func (f RatingFilter) LessThanEqual(value float64) (exp Filter) {
	return f.Lte(value)
}

// Between builds rating > min AND rating < max filter.
func (f RatingFilter) Between(min, max float64) (exp Filter) {
	return Filter{
		expression: "rating > ? AND rating < ?",
		args:       []interface{}{min, max},
	}
}

// Range builds rating >= min AND rating <= max filter.
func (f RatingFilter) Range(min, max float64) (exp Filter) {
	return Filter{
		expression: "rating >= ? AND rating <= ?",
		args:       []interface{}{min, max},
	}
}

// Name SQL field of Rating.
func (f RatingFilter) Name() string {
	return "rating"
}
