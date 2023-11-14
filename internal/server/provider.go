package server

import "github.com/google/wire"

var ProviderSetHTTPServer = wire.NewSet(NewHTTPServer)
