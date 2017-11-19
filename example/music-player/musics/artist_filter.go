package musics

type (
	// ArtistFilter is filter expression builder for field Artist.
	ArtistFilter struct{}
)

var Artist ArtistFilter

// Eq builds artist = value filter.
func (f ArtistFilter) Eq(value string) (exp Filter) {
	return Filter{
		expression: "artist = ?",
		args:       []interface{}{value},
	}
}

// Neq builds artist <> value filter.
func (f ArtistFilter) Neq(value string) (exp Filter) {
	return Filter{
		expression: "artist <> ?",
		args:       []interface{}{value},
	}
}

// Gt builds artist > value filter.
func (f ArtistFilter) Gt(value string) (exp Filter) {
	return Filter{
		expression: "artist > ?",
		args:       []interface{}{value},
	}
}

// Lt builds artist < value filter.
func (f ArtistFilter) Lt(value string) (exp Filter) {
	return Filter{
		expression: "artist < ?",
		args:       []interface{}{value},
	}
}

// Gte builds artist >= value filter.
func (f ArtistFilter) Gte(value string) (exp Filter) {
	return Filter{
		expression: "artist >= ?",
		args:       []interface{}{value},
	}
}

// Lte builds artist <= value filter.
func (f ArtistFilter) Lte(value string) (exp Filter) {
	return Filter{
		expression: "artist <= ?",
		args:       []interface{}{value},
	}
}

// Equal builds artist = value filter.
func (f ArtistFilter) Equal(value string) (exp Filter) {
	return f.Eq(value)
}

// NotEqual builds artist <> value filter.
func (f ArtistFilter) NotEqual(value string) (exp Filter) {
	return f.Neq(value)
}

// GreaterThan builds artist > value filter.
func (f ArtistFilter) GreaterThan(value string) (exp Filter) {
	return f.Gt(value)
}

// LessThan builds artist < value filter.
func (f ArtistFilter) LessThan(value string) (exp Filter) {
	return f.Lt(value)
}

// GreaterThanEqual builds artist >= value filter.
func (f ArtistFilter) GreaterThanEqual(value string) (exp Filter) {
	return f.Gte(value)
}

// LessThanEqual builds artist <= value filter.
func (f ArtistFilter) LessThanEqual(value string) (exp Filter) {
	return f.Lte(value)
}

// Between builds artist > min AND artist < max filter.
func (f ArtistFilter) Between(min, max string) (exp Filter) {
	return Filter{
		expression: "artist > ? AND artist < ?",
		args:       []interface{}{min, max},
	}
}

// Range builds artist >= min AND artist <= max filter.
func (f ArtistFilter) Range(min, max string) (exp Filter) {
	return Filter{
		expression: "artist >= ? AND artist <= ?",
		args:       []interface{}{min, max},
	}
}

// Name SQL field of Artist.
func (f ArtistFilter) Name() string {
	return "artist"
}
