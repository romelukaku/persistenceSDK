package mapper

import (
	"github.com/persistenceOne/persistenceSDK/schema/helpers/base"
)

var Mapper = base.NewMapper(
	ModuleName,
	generateKey,
	assetPrototype,
	registerCodec,
)