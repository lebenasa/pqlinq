package musics

import "time"

type (
	// ReleaseDateFilter is filter expression builder for field ReleaseDate.
	ReleaseDateFilter struct{}
)

var ReleaseDate ReleaseDateFilter

// Eq builds release_date = value filter.
func (f ReleaseDateFilter) Eq(value time.Time) (exp Filter) {
	return Filter{
		expression: "release_date = ?",
		args:       []interface{}{value},
	}
}

// Neq builds release_date <> value filter.
func (f ReleaseDateFilter) Neq(value time.Time) (exp Filter) {
	return Filter{
		expression: "release_date <> ?",
		args:       []interface{}{value},
	}
}

// Gt builds release_date > value filter.
func (f ReleaseDateFilter) Gt(value time.Time) (exp Filter) {
	return Filter{
		expression: "release_date > ?",
		args:       []interface{}{value},
	}
}

// Lt builds release_date < value filter.
func (f ReleaseDateFilter) Lt(value time.Time) (exp Filter) {
	return Filter{
		expression: "release_date < ?",
		args:       []interface{}{value},
	}
}

// Gte builds release_date >= value filter.
func (f ReleaseDateFilter) Gte(value time.Time) (exp Filter) {
	return Filter{
		expression: "release_date >= ?",
		args:       []interface{}{value},
	}
}

// Lte builds release_date <= value filter.
func (f ReleaseDateFilter) Lte(value time.Time) (exp Filter) {
	return Filter{
		expression: "release_date <= ?",
		args:       []interface{}{value},
	}
}

// Equal builds release_date = value filter.
func (f ReleaseDateFilter) Equal(value time.Time) (exp Filter) {
	return f.Eq(value)
}

// NotEqual builds release_date <> value filter.
func (f ReleaseDateFilter) NotEqual(value time.Time) (exp Filter) {
	return f.Neq(value)
}

// GreaterThan builds release_date > value filter.
func (f ReleaseDateFilter) GreaterThan(value time.Time) (exp Filter) {
	return f.Gt(value)
}

// LessThan builds release_date < value filter.
func (f ReleaseDateFilter) LessThan(value time.Time) (exp Filter) {
	return f.Lt(value)
}

// GreaterThanEqual builds release_date >= value filter.
func (f ReleaseDateFilter) GreaterThanEqual(value time.Time) (exp Filter) {
	return f.Gte(value)
}

// LessThanEqual builds release_date <= value filter.
func (f ReleaseDateFilter) LessThanEqual(value time.Time) (exp Filter) {
	return f.Lte(value)
}

// Between builds release_date > min AND release_date < max filter.
func (f ReleaseDateFilter) Between(min, max time.Time) (exp Filter) {
	return Filter{
		expression: "release_date > ? AND release_date < ?",
		args:       []interface{}{min, max},
	}
}

// Range builds release_date >= min AND release_date <= max filter.
func (f ReleaseDateFilter) Range(min, max time.Time) (exp Filter) {
	return Filter{
		expression: "release_date >= ? AND release_date <= ?",
		args:       []interface{}{min, max},
	}
}

// Name SQL field of ReleaseDate.
func (f ReleaseDateFilter) Name() string {
	return "release_date"
}
