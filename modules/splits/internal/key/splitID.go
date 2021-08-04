/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	codecUtilities "github.com/persistenceOne/persistenceSDK/utilities/codec"
)

type splitID struct {
	OwnerID   types.ID `json:"ownerID" valid:"required~required field ownerID missing"`
	OwnableID types.ID `json:"ownableID" valid:"required~required field ownableID missing"`
}

func (splitID splitID) Size() int {
	panic("implement me")
}

func (splitID splitID) MarshalTo(i []byte) (int, error) {
	panic("implement me")
}

func (splitID splitID) Unmarshal(i []byte) error {
	panic("implement me")
}

var _ types.ID = (*splitID)(nil)
var _ helpers.Key = (*splitID)(nil)

func (splitID splitID) Bytes() []byte {
	return append(
		splitID.OwnerID.Bytes(),
		splitID.OwnableID.Bytes()...)
}
func (splitID splitID) String() string {
	var values []string
	values = append(values, splitID.OwnerID.String())
	values = append(values, splitID.OwnableID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (splitID splitID) Compare(id types.ID) int {
	return bytes.Compare(splitID.Bytes(), id.Bytes())
}
func (splitID splitID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(splitID.Bytes())
}
func (splitID) RegisterCodec(codec *codec.LegacyAmino) {
	codecUtilities.RegisterXPRTConcrete(codec, module.Name, splitID{})
}
func (splitID splitID) IsPartial() bool {
	return len(splitID.OwnableID.Bytes()) == 0
}
func (splitID splitID) Equals(key helpers.Key) bool {
	return splitID.Compare(splitIDFromInterface(key)) == 0
}

func NewSplitID(ownerID types.ID, ownableID types.ID) types.ID {
	return splitID{
		OwnerID:   ownerID,
		OwnableID: ownableID,
	}
}
