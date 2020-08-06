/*
 Copyright [2019] - [2020], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package maintainer

import (
	"github.com/cosmos/cosmos-sdk/codec"
	assetsMapper "github.com/persistenceOne/persistenceSDK/modules/assets/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/mapper"
	"github.com/persistenceOne/persistenceSDK/schema"
)

func registerCodec(codec *codec.Codec) {
	codec.RegisterConcrete(queryRequest{}, QueryRoute+"/"+"request", nil)
	codec.RegisterConcrete(queryResponse{}, QueryRoute+"/"+"response", nil)
}

var packageCodec = codec.New()

func init() {
	registerCodec(packageCodec)
	schema.RegisterCodec(packageCodec)
	assetsMapper.Mapper.RegisterCodec(packageCodec)
	mapper.Mapper.RegisterCodec(packageCodec)
	packageCodec.Seal()
}
