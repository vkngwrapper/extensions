package extensions

//go:generate go run ./library/cmd
//go:generate go run go.uber.org/mock/mockgen -source=./library/iface.go -destination=./library/mocks/library.go
