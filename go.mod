module example.com/main

go 1.24.5

replace example.com/matrix => ./matrix

replace example.com/polyfit => ./polyfit

require (
	example.com/matrix v0.0.0-00010101000000-000000000000
	example.com/polyfit v0.0.0-00010101000000-000000000000
)
