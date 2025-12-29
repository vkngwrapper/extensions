package extensions

//go:generate go run ./library/cmd
//go:generate mockgen -source=./library/iface.go -destination=./library/mocks/library.go
