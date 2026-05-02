package handler

import "github.com/google/wire"

// ProviderSet is biz providers.
var HandlerProviderSet = wire.NewSet(NewTransactionUsecase)
