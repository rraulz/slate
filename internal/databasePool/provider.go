package databasePool

import "github.com/google/wire"

var ProviderSetDatabasePool = wire.NewSet(NewDatabase)
