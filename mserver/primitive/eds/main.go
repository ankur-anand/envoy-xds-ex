package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v2/endpoint"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/server"
	"google.golang.org/grpc"
)

// NodeHash interface implementation. Implements a hash function that returns a character string from an Envoy identifier.
type hash struct{}

func (hash) ID(node *core.Node) string {
	if node == nil {
		return "unknown"
	}
	return node.Cluster + "/" + node.Id
}

var upstreams = map[string][]struct {
	Address string
	Port    uint32
}{
	// address
	"random_payload_cluster1": {{"127.0.0.1", 3001}, {"127.0.0.1", 3002}},
}

// Returns a snapshot. The structure is the same as the Protocol Buffer definition.
func defaultSnapshot() cache.Snapshot {
	var resources []cache.Resource
	for cluster, ups := range upstreams {
		eps := make([]*endpoint.LocalityLbEndpoints, len(ups))
		for i, up := range ups {
			eps[i] = &endpoint.LocalityLbEndpoints{
				LbEndpoints: []*endpoint.LbEndpoint{{
					HostIdentifier: &endpoint.LbEndpoint_Endpoint{
						Endpoint: &endpoint.Endpoint{
							Address: &core.Address{
								Address: &core.Address_SocketAddress{
									SocketAddress: &core.SocketAddress{
										Address:       up.Address,
										PortSpecifier: &core.SocketAddress_PortValue{PortValue: up.Port},
									},
								},
							},
						},
					},
				}},
			}
		}
		assignment := &api.ClusterLoadAssignment{
			ClusterName: cluster,
			Endpoints:   eps,
		}
		resources = append(resources, assignment)
	}

	return cache.NewSnapshot("0.0", resources, nil, nil, nil, nil)
}

func run(listen string) error {
	// If the result of xDS is set as a cache, it will be nicely returned as xDS API.
	snapshotCache := cache.NewSnapshotCache(false, hash{}, nil)
	server := xds.NewServer(context.Background(), snapshotCache, nil)
	//Remember the hash value returned by NodeHash and the snapshot of its settings as a cache
	err := snapshotCache.SetSnapshot("cluster.local/node0", defaultSnapshot())
	if err != nil {
		return err
	}

	// Launch gRCP server and provide API
	grpcServer := grpc.NewServer()
	api.RegisterEndpointDiscoveryServiceServer(grpcServer, server)

	lsn, err := net.Listen("tcp", listen)
	if err != nil {
		return err
	}
	return grpcServer.Serve(lsn)
}

func main() {
	var listen string
	flag.StringVar(&listen, "listen", ":20000", "listen port")
	flag.Parse()

	log.Printf("Starting server with -listen=%s", listen)

	err := run(listen)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
