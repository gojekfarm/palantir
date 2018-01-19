package palantir

import "time"

type DBConfig struct {
	driver          string
	url             string
	slaveUrl        string
	idleConn        int
	maxConn         int
	connMaxLifetime time.Duration
}

func (self *DBConfig) Driver() string {
	return self.driver
}

func (self *DBConfig) Url() string {
	return self.url
}

func (self *DBConfig) SlaveUrl() string {
	return self.slaveUrl
}

func (self *DBConfig) MaxConn() int {
	return self.maxConn
}

func (self *DBConfig) IdleConn() int {
	return self.idleConn
}

func (self *DBConfig) ConnMaxLifetime() time.Duration {
	return self.connMaxLifetime
}

func LoadDbConf() *DBConfig {
	return &DBConfig{
		driver:          getStringOrPanic("DB_DRIVER"),
		url:             getStringOrPanic("DB_URL"),
		slaveUrl:        getStringOrPanic("DB_SLAVE_URL"),
		maxConn:         getIntOrPanic("DB_MAX_CONN"),
		idleConn:        getIntOrPanic("DB_IDLE_CONN"),
		connMaxLifetime: time.Duration(getIntOrPanic("DB_CONN_MAX_LIFETIME")) * time.Second,
	}
}
