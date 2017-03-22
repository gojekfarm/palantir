package goconfig

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
		driver:          getStringOrPanic("db_driver"),
		url:             getStringOrPanic("db_url"),
		slaveUrl:        getStringOrPanic("db_slave_url"),
		maxConn:         getIntOrPanic("db_max_conn"),
		idleConn:        getIntOrPanic("db_idle_conn"),
		connMaxLifetime: time.Duration(getIntOrPanic("db_conn_max_lifetime")) * time.Second,
	}
}
