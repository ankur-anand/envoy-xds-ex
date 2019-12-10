## xDS Envoy Examples

Envoy discovers its various dynamic resources via the filesystem or
by querying one or more management servers.

Collectively, these discovery services and their corresponding APIs are
referred to as xDS.

The xds API is a communication protocol defined by Envoy for data interaction between the control plane and the data plane.

The following are the parts of Envoy’s runtime model we can configure dynamically through xDS:

Listeners Discovery Service API - LDS to publish ports on which to listen for traffic
Endpoints Discovery Service API- EDS for service discovery,
Routes Discovery Service API- RDS for traffic routing decisions
Clusters Discovery Service- CDS for backend services to which we can route traffic
Secrets Discovery Service - SDS for distributing secrets (certificates and keys)
xDS		The collective name of the above APIs

https://github.com/envoyproxy/data-plane-api/blob/master/API_OVERVIEW.md

The dynamic configuration in envoyProxy can be pushed through the Management
 Server that implements data-plane-api.
 
 **The Resource discovery is uncertain if there are multiple discovery
  channels**. Use ADS to ensure their delivery order.
  
  **Whenever the listener changes, wait for the existing connection to be
   emptied** or wait for the drain to timeout before applying the latest
    configuration.

