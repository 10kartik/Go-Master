module example/hello

go 1.24.2

require (
	example/hello/greetings v0.0.0-00010101000000-000000000000
)

replace example/hello/greetings => ../greetings
