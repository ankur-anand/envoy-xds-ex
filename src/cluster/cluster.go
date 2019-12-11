package cluster

import (
	core2 "github.com/ankur-anand/envoy-xds/src/core"
	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/golang/protobuf/ptypes"

	apiV2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	apiCluster "github.com/envoyproxy/go-control-plane/envoy/api/v2/cluster"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
)

// Sample Cluster Configuration
// clusters:
//    - name: nodeusrsvc
//      type: LOGICAL_DNS
//      connect_timeout: 1s
//      load_assignment:
//        cluster_name: nodeusrsvc
//        endpoints:
//          - lb_endpoints:
//              - endpoint:
//                  address:
//                    socket_address:
//                      address: some.ankuranand.in
//                      port_value: 80

// MakeNewCluster creates a cluster of Envoy
func MakeNewCluster(cluster core2.Cluster) *apiV2.
	Cluster {

	// endpoints inside the cluster
	lbEndpoints := make([]*endpoint.LbEndpoint, 0)

	for _, upstream := range cluster.Upstreams {
		hostIdentifier := &endpoint.LbEndpoint_Endpoint{Endpoint: &endpoint.
			Endpoint{Address: newAddress(upstream.Host, upstream.Port)}}

		lbEndpoint := &endpoint.LbEndpoint{
			HostIdentifier: hostIdentifier,
		}

		lbEndpoints = append(lbEndpoints, lbEndpoint)
	}

	// endpoints grouping, consisting of multiple endpoints
	localityLbEndpoints := &endpoint.LocalityLbEndpoints{
		LbEndpoints: lbEndpoints,
	}

	endpoints := make([]*endpoint.LocalityLbEndpoints, 0)
	endpoints = append(endpoints, localityLbEndpoints)

	// cluster multiple endpoints
	clusterLoadAssigment := &apiV2.ClusterLoadAssignment{
		// endpoint is static, cluster name can be empty
		ClusterName: cluster.Name,
		Endpoints:   endpoints,
	}

	return &apiV2.Cluster{
		Name:        cluster.Name,
		AltStatName: cluster.Name,
		// https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/service_discovery#service-discovery
		// Logical DNS  is optimal for large scale web services that must be
		//accessed via DNS.
		ClusterDiscoveryType: &apiV2.Cluster_Type{Type: apiV2.
			Cluster_LOGICAL_DNS},
		DnsLookupFamily:               apiV2.Cluster_V4_ONLY,
		EdsClusterConfig:              nil,
		ConnectTimeout:                ptypes.DurationProto(cluster.Timeout),
		PerConnectionBufferLimitBytes: nil, // default 1MB
		LbPolicy:                      apiV2.Cluster_ROUND_ROBIN,
		LoadAssignment:                clusterLoadAssigment,
		CircuitBreakers: &apiCluster.CircuitBreakers{
			Thresholds: []*apiCluster.
				CircuitBreakers_Thresholds{{MaxRetries: &wrappers.
				UInt32Value{Value: cluster.Retries},
				MaxConnections: &wrappers.UInt32Value{Value: getMaxConnection(cluster.
					MaxConnection)}}},
		},
	}
}

func newAddress(address string, port uint32) *core.Address {
	addr := &core.Address{
		Address: &core.Address_SocketAddress{SocketAddress: &core.
			SocketAddress{
			Protocol:      core.SocketAddress_TCP,
			Address:       address,
			PortSpecifier: &core.SocketAddress_PortValue{PortValue: port}},
		}}
	return addr
}

// getMaxConnection
func getMaxConnection(connection uint32) uint32 {
	if connection <= 0 {
		return 1024
	}
	return connection
}
