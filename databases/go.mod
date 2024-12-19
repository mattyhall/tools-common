module github.com/couchbase/tools-common/databases

go 1.21

replace github.com/mattn/go-sqlite3 v1.14.24 => github.com/mattyhall/go-sqlite3 v0.0.0-20241218151426-637bb55262e7

require (
	github.com/couchbase/tools-common/sync/v2 v2.0.1
	github.com/mattn/go-sqlite3 v1.14.24
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
