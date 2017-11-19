package musics

type (
	// AlbumFilter is filter expression builder for field Album.
	AlbumFilter struct{}
)

var Album AlbumFilter

// Eq builds album = value filter.
func (f AlbumFilter) Eq(value string) (exp Filter) {
	return Filter{
		expression: "album = ?",
		args:       []interface{}{value},
	}
}

// Neq builds album <> value filter.
func (f AlbumFilter) Neq(value string) (exp Filter) {
	return Filter{
		expression: "album <> ?",
		args:       []interface{}{value},
	}
}

// Gt builds album > value filter.
func (f AlbumFilter) Gt(value string) (exp Filter) {
	return Filter{
		expression: "album > ?",
		args:       []interface{}{value},
	}
}

// Lt builds album < value filter.
func (f AlbumFilter) Lt(value string) (exp Filter) {
	return Filter{
		expression: "album < ?",
		args:       []interface{}{value},
	}
}

// Gte builds album >= value filter.
func (f AlbumFilter) Gte(value string) (exp Filter) {
	return Filter{
		expression: "album >= ?",
		args:       []interface{}{value},
	}
}

// Lte builds album <= value filter.
func (f AlbumFilter) Lte(value string) (exp Filter) {
	return Filter{
		expression: "album <= ?",
		args:       []interface{}{value},
	}
}

// Equal builds album = value filter.
func (f AlbumFilter) Equal(value string) (exp Filter) {
	return f.Eq(value)
}

// NotEqual builds album <> value filter.
func (f AlbumFilter) NotEqual(value string) (exp Filter) {
	return f.Neq(value)
}

// GreaterThan builds album > value filter.
func (f AlbumFilter) GreaterThan(value string) (exp Filter) {
	return f.Gt(value)
}

// LessThan builds album < value filter.
func (f AlbumFilter) LessThan(value string) (exp Filter) {
	return f.Lt(value)
}

// GreaterThanEqual builds album >= value filter.
func (f AlbumFilter) GreaterThanEqual(value string) (exp Filter) {
	return f.Gte(value)
}

// LessThanEqual builds album <= value filter.
func (f AlbumFilter) LessThanEqual(value string) (exp Filter) {
	return f.Lte(value)
}

// Between builds album > min AND album < max filter.
func (f AlbumFilter) Between(min, max string) (exp Filter) {
	return Filter{
		expression: "album > ? AND album < ?",
		args:       []interface{}{min, max},
	}
}

// Range builds album >= min AND album <= max filter.
func (f AlbumFilter) Range(min, max string) (exp Filter) {
	return Filter{
		expression: "album >= ? AND album <= ?",
		args:       []interface{}{min, max},
	}
}

// Name SQL field of Album.
func (f AlbumFilter) Name() string {
	return "album"
}
