package client

import "github.com/google/wire"


var ClientProviderSet=wire.NewSet(NewAccountServiceClient,NewAccountClient)