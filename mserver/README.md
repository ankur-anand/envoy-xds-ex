It should be noted that go-control-plane is not a manager server, it is a code library (Library) that implements data-plane-api .

When developing a manager server on the basis of go-control-plane, you only need to consider the access to configuration data, and do not need to consider the details of communication with eonvy.

