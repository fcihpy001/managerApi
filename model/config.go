package model

type Config struct {
	Datasource Datasource
	Version    string
	Dsn        string
	RedisURL   string
	Server     Server
}

type Server struct {
	Host string
	Port string
}

type Datasource struct {
	DriverName string
	Host       string
	Port       string
	Database   string
	UserName   string
	Password   string
	Charset    string
	Loc        string
}
