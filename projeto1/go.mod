module projeto1

require (
	github.com/jeffotoni/gcolor v1.0.9
	internal/crypt v0.0.0
	internal/fmts v0.0.0
)

replace (
	internal/crypt => ./internal/crypt
	internal/fmts => ./internal/fmts
)

go 1.16
