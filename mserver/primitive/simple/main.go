package main

//type myEndpoint struct {
//	version string
//	name    string
//	address string
//	port    uint32
//}
//
//type NodeConfig struct {
//	// Identifies a specific Envoy instance.
//	// The node identifier is presented to the management server,
//	// which may use this identifier to distinguish per Envoy configuration
//	// for serving.
//	node      *envoycore.Node
//	endpoints []envoycache.Resource // //[]*api.ClusterLoadAssignment
//	clusters  []envoycache.Resource //[]*api.Cluster
//	routes    []envoycache.Resource //[]*api.RouteConfiguration
//	listeners []envoycache.Resource //[]*api.Listener
//}
//
////implement cache.NodeHash
//func (n NodeConfig) ID(node *core.Node) string {
//	return node.GetId()
//}
//
//func AddClusterWithStaticStaticEndpoint(n *NodeConfig) {
//	// `CDS：Upstream Cluster`
//	n.clusters = append(n.clusters, cluster)
//}
//
//func AddClusterWithDynamicEndpoint(n *NodeConfig) {
//	// `EDS：Upstream Server`
//
//	n.endpoints = append(n.endpoints, endpoint)
//	n.clusters = append(n.clusters, cluster)
//}
//
//func AddListenerWithStaticRoute(n *NodeConfig) {
//	// `LDS: Listener`
//
//	n.listeners = append(n.listeners, listener)
//}
//
//func AddListenerWithDynamicRoute(n *NodeConfig) {
//	// `RDS：Route`
//	n.listeners = append(n.listeners, listener)
//	n.routes = append(n.routes, route)
//}
//
//func createEndpoint(endp myEndpoint) *envoyapi.ClusterLoadAssignment {
//	clusterLoadAssignment := &envoyapi.ClusterLoadAssignment{
//		ClusterName: endp.name,
//		Endpoints: []*envoyendpoint.LocalityLbEndpoints{{
//			LbEndpoints: []*envoyendpoint.LbEndpoint{{
//				HostIdentifier: &envoyendpoint.LbEndpoint_Endpoint{
//					Endpoint: &envoyendpoint.Endpoint{
//						Address: &envoycore.Address{
//							Address: &envoycore.Address_SocketAddress{
//								SocketAddress: &envoycore.SocketAddress{
//									Address:       endp.address,
//									PortSpecifier: &envoycore.SocketAddress_PortValue{PortValue: endp.port},
//								},
//							},
//						},
//					},
//				},
//			}},
//		}},
//	}
//	return clusterLoadAssignment
//}

// Update_SnapshotCache () fills
// node_config into snapshotCache,
// and specifies the version number of the
// configuration during filling. As long as the configuration number changes,
//it is considered as a configuration that needs to be updated. There is no concept of increment or rollback.
//func UpdateSnapshotCache(s envoycache.SnapshotCache, n *NodeConfig, version string) {
//	err := s.SetSnapshot(n.ID(n.node), envoycache.NewSnapshot(version,
//		n.endpoints, n.clusters, n.routes, n.listeners))
//	if err != nil {
//		glog.Error(err)
//	}
//}

//func run(listen string, endp myEndpoint) error {
//
//	// If you set the result of xDS as a cache, it will return as a nice xDS API.
//	snapshotCache := envoycache.NewSnapshotCache(false, hash{}, nil)
//
//	// one is to store all configured caches,
//	// and the other is a callback function that will be called when
//	// processing envoy requests:
//	server := envoyserver.NewServer(context.TODO(), snapshotCache, Callback{})
//	// All FetchXXX functions
//	// (functions that handle envoy requests) are ultimately called,
//	// Fetch()and their implementation is as follows:
//	// cache.CacheIs an interface, and variables that implement this interface can be used as
//	// SnapshotCache implements the SetSnapshot()interface:
//	// the second parameter Snapshot is the full configuration on the
//	// corresponding node:
//	err := snapshotCache.SetSnapshot("cluster.local/node0", createSnapshot(endp))
//	if err != nil {
//		return err
//	}
//
//	grpcServer := grpc.NewServer()
//
//	envoyapi.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
//
//	lsn, err := net.Listen("tcp", listen)
//	if err != nil {
//		return err
//	}
//	go func() {
//		log.Printf("start grpc server version:%s", endp.version)
//		grpcServer.Serve(lsn)
//	}()
//
//	quit := make(chan os.Signal)
//	signal.Notify(quit, os.Interrupt)
//	<-quit
//	log.Println("stopping grpc server...")
//
//	grpcServer.Stop()
//	//grpcServer.GracefulStop()
//
//	return nil
//}
//
//func main() {
//	var listen string
//	flag.StringVar(&listen, "listen", ":20000", "listen port")
//	flag.Parse()
//
//	log.Printf("Starting server with -listen=%s", listen)
//
//	end1 := myEndpoint{
//		version: "0",
//		name:    "cluster-1",
//		address: "127.0.0.1",
//		port:    3001,
//	}
//	end2 := myEndpoint{
//		version: "1",
//		name:    "cluster-2",
//		address: "127.0.0.1",
//		port:    3002,
//	}
//}
