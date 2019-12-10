package cluster

import (
	"time"

	"github.com/golang/protobuf/ptypes"

	apiV2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
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

// ADDR for port and ip address
type ADDR struct {
	Address string
	Port    uint32
}

// MakeStaticCluster  creates a cluster using static resources.
func MakeStaticCluster(name string, addrs []ADDR) *apiV2.Cluster {
	lbEndpoints := make([]*endpoint.LbEndpoint, 0)

	for _, addr := range addrs {
		// endpoint creation
		hostIdentifier := &endpoint.LbEndpoint_Endpoint{Endpoint: &endpoint.
			Endpoint{Address: &core.Address{
			Address: &core.Address_SocketAddress{SocketAddress: &core.SocketAddress{
				Protocol: core.SocketAddress_TCP,
				Address:  addr.Address,
				PortSpecifier: &core.SocketAddress_PortValue{PortValue: addr.
					Port},
			}},
		}}}

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
		ClusterName: name,
		Endpoints:   endpoints,
	}

	timeout := 1 * time.Second

	// Cluster using static endpoints of type
	// V2.Cluster_Static
	return &apiV2.Cluster{
		Name:                          name,
		AltStatName:                   name,
		ClusterDiscoveryType:          &apiV2.Cluster_Type{Type: apiV2.Cluster_STATIC},
		EdsClusterConfig:              nil,
		ConnectTimeout:                ptypes.DurationProto(timeout),
		PerConnectionBufferLimitBytes: nil, // default 1MB
		LbPolicy:                      apiV2.Cluster_ROUND_ROBIN,
		LoadAssignment:                clusterLoadAssigment,
	}
}
