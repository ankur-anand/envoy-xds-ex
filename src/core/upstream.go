package core

import "time"

// Upstream is composed form of Virtual Host.
type Upstream struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     uint32 `json:"port"`
	PortName string `json:"portName"`
}

// Cluster compact form of an Envoy cluster and multiple upstream
type Cluster struct {
	Name      string        `json:"name"`
	Upstreams []Upstream    `json:"upstreams"`
	Domains   []string      `json:"domains"`
	Prefix    string        `json:"prefix"`
	Retries   uint32        `json:"retries"`
	Timeout   time.Duration `json:"timeout"`
}
