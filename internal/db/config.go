package db

const (
	DefaultMaxOpenConns = 10
	DefaultMaxIdleConns = 10
)

type Config struct {
	DriverName   string
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
}
