module github.com/mindtastic/cli

go 1.18

require (
	github.com/mindtastic/cli/client v0.0.0-00010101000000-000000000000
	golang.org/x/oauth2 v0.0.0-20220622183110-fd043fe589d2 // indirect
)

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/mindtastic/cli/client => ./client
