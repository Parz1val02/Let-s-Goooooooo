module example.com/hello

go 1.20

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.12.0 // indirect
	rsc.io/sampler v1.99.99 // indirect
)
