package autoconfig

import (
	"strconv"
	"encoding"
)


type Database struct {
	Url      string
	Username string
	Password string
	Port     string
	Dbname   string
	Charset  string
	Location string
	MaxIdle  int
	MaxOpen  int
	Timeout  int
}

type Redis struct {
	Address  string
	Port     int
	Db       int
	Password string
	Sentinel struct {
		Master string   `yaml:"master"`
		Nodes  []string `yaml: "nodes,flow"`
	}
	Pool struct {
		MaxIdle     int `yaml:"maxIdle"`
		PoolTimeout int `yaml:"poolTimeout"`
	}
}

type Mail struct {
	Server	string
	Port	int
	UserName	string
	Password	string
	SenderName	string
	SenderMail	string
	Encoding	string
}

type AppConfig struct {
	Database Database
	Redis    Redis
}

func (d Database) Dns() string {

	url := d.Username + ":" + d.Password + "@tcp(" + d.Url + ":" + d.Port + ")/" + d.Dbname;

	if d.Charset != "" || d.Location != "" {
		url += "?"

		if d.Charset != "" {
			url += "charset=" + d.Charset
		}

		if d.Location != "" {
			url += "&loc=" + d.Location
		}
	}

	if d.Timeout != 0 {
		url += "&timeout=" + strconv.Itoa(d.Timeout)
	}

	return url
}

func (r Redis) ToUrl() string {
	return r.Address + ":" + strconv.Itoa(r.Port)
}
