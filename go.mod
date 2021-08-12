module example.com/hello

go 1.16

require (
	example.com/error v0.0.0-00010101000000-000000000000
	example.com/mathutil v0.0.0-00010101000000-000000000000
	example.com/strutil v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

replace example.com/mathutil => ./mathutil

replace example.com/strutil => ./strutil

replace example.com/error => ./error
