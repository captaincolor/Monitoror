module Monitoror

go 1.14

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/shirou/gopsutil/v3 v3.21.2
	google.golang.org/grpc v1.36.0 // indirect
	github.com/send v0.0.0
)

replace github.com/send => ./send
