package musics

import "time"

type (
	// LastPlayedFilter is filter expression builder for field LastPlayed.
	LastPlayedFilter struct{}
)

var LastPlayed LastPlayedFilter

// Eq builds last_played = value filter.
func (f LastPlayedFilter) Eq(value time.Time) (exp Filter) {
	return Filter{
		expression: "last_played = ?",
		args:       []interface{}{value},
	}
}

// Neq builds last_played <> value filter.
func (f LastPlayedFilter) Neq(value time.Time) (exp Filter) {
	return Filter{
		expression: "last_played <> ?",
		args:       []interface{}{value},
	}
}

// Gt builds last_played > value filter.
func (f LastPlayedFilter) Gt(value time.Time) (exp Filter) {
	return Filter{
		expression: "last_played > ?",
		args:       []interface{}{value},
	}
}

// Lt builds last_played < value filter.
func (f LastPlayedFilter) Lt(value time.Time) (exp Filter) {
	return Filter{
		expression: "last_played < ?",
		args:       []interface{}{value},
	}
}

// Gte builds last_played >= value filter.
func (f LastPlayedFilter) Gte(value time.Time) (exp Filter) {
	return Filter{
		expression: "last_played >= ?",
		args:       []interface{}{value},
	}
}

// Lte builds last_played <= value filter.
func (f LastPlayedFilter) Lte(value time.Time) (exp Filter) {
	return Filter{
		expression: "last_played <= ?",
		args:       []interface{}{value},
	}
}

// Equal builds last_played = value filter.
func (f LastPlayedFilter) Equal(value time.Time) (exp Filter) {
	return f.Eq(value)
}

// NotEqual builds last_played <> value filter.
func (f LastPlayedFilter) NotEqual(value time.Time) (exp Filter) {
	return f.Neq(value)
}

// GreaterThan builds last_played > value filter.
func (f LastPlayedFilter) GreaterThan(value time.Time) (exp Filter) {
	return f.Gt(value)
}

// LessThan builds last_played < value filter.
func (f LastPlayedFilter) LessThan(value time.Time) (exp Filter) {
	return f.Lt(value)
}

// GreaterThanEqual builds last_played >= value filter.
func (f LastPlayedFilter) GreaterThanEqual(value time.Time) (exp Filter) {
	return f.Gte(value)
}

// LessThanEqual builds last_played <= value filter.
func (f LastPlayedFilter) LessThanEqual(value time.Time) (exp Filter) {
	return f.Lte(value)
}

// Between builds last_played > min AND last_played < max filter.
func (f LastPlayedFilter) Between(min, max time.Time) (exp Filter) {
	return Filter{
		expression: "last_played > ? AND last_played < ?",
		args:       []interface{}{min, max},
	}
}

// Range builds last_played >= min AND last_played <= max filter.
func (f LastPlayedFilter) Range(min, max time.Time) (exp Filter) {
	return Filter{
		expression: "last_played >= ? AND last_played <= ?",
		args:       []interface{}{min, max},
	}
}

// Name SQL field of LastPlayed.
func (f LastPlayedFilter) Name() string {
	return "last_played"
}
