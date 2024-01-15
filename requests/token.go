package requests

// FileterType: Supported url query string for query postgres
type FilterType int

const (
	Eq    FilterType = iota // Equal
	Gt                      // Greater than
	Gte                     // Greater than or equal
	Lt                      // Lower than
	Lte                     // Lower than or equal
	Neq                     // <> || != (Not Equal)
	Like                    // LIKE
	ILike                   // ILIKE
	In                      // IN
	Is                      // IS
	Fts                     // @@ -> to_tsquery
	PlFts                   // @@ -> plainto_tsquery
	Cs                      // @>
	Cd                      // <@
)
