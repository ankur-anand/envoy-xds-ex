## xDS Envoy Examples

Envoy discovers its various dynamic resources via the filesystem or
by querying one or more management servers.

Collectively, these discovery services and their corresponding APIs are
referred to as xDS.

The xds API is a communication protocol defined by Envoy for data interaction between the control plane and the data plane.

Concept	Full name	description
LDS	Listener Discovery Service	
RDS	Route Discovery Service	
CDS	Cluster Discovery Service	
EDS	Endpoint Discovery Service	
SDS	Service Discovery Service	Renamed EDS
ADS	Aggregated Discovery Service	
HDS	Health Discovery Service	
SDS	Secret Discovery Service	
MS	Metric Service	
RLS	Rate Limit Service	
xDS		The collective name of the above APIs

https://github.com/envoyproxy/data-plane-api/blob/master/API_OVERVIEW.md

The dynamic configuration in envoyProxy can be pushed through the Management
 Server that implements data-plane-api.
 
 **The Resource discovery is uncertain if there are multiple discovery
  channels**. Use ADS to ensure their delivery order.
  
  **Whenever the listener changes, wait for the existing connection to be
   emptied** or wait for the drain to timeout before applying the latest
    configuration.

