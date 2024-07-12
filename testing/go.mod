module github.com/psyark/notion/testing

replace github.com/psyark/notion => ../

go 1.22

require (
	github.com/flytam/filenamify v1.2.0
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/psyark/notion v0.0.0-00010101000000-000000000000
	github.com/samber/lo v1.44.0
)

require (
	github.com/dave/jennifer v1.7.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/text v0.16.0 // indirect
)
