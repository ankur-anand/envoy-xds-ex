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
	// Name is the logical name for the virtual host as well for cluster
	Name      string     `json:"name"`
	Upstreams []Upstream `json:"upstreams"`
	// domains for the virtual host
	Domains []string `json:"domains"`
	Prefix  string   `json:"prefix"`
	// https://blog.turbinelabs.io/circuit-breaking-da855a96a61d
	Retries uint32        `json:"retries"`
	Timeout time.Duration `json:"timeout"`
	// The maximum number of connections that Envoy
	// will make to the upstream cluster. If not specified, the default is 1024.
	// we have to support.
	// For HTTP/1.1 connections, use max_connections.
	// For HTTP/2 connections, use max_requests.
	MaxConnection uint32 `json:"maxConnection"`
}
