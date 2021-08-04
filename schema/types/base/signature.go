/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package base

import (
	"encoding/base64"

	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/tendermint/tendermint/crypto"
)

var _ types.Signature = (*signature)(nil)

func (baseSignature signature) String() string {
	return base64.URLEncoding.EncodeToString(baseSignature.Bytes())
}
func (baseSignature signature) Bytes() []byte   { return baseSignature.SignatureBytes }
func (baseSignature signature) GetID() types.ID { return baseSignature.ID }
func (baseSignature signature) Verify(pubKey crypto.PubKey, bytes []byte) bool {
	return pubKey.VerifySignature(bytes, baseSignature.Bytes())
}
func (baseSignature signature) GetValidityHeight() types.Height {
	return baseSignature.ValidityHeight
}
func (baseSignature signature) HasExpired(height types.Height) bool {
	return baseSignature.GetValidityHeight().Compare(height) > 0
}

func NewSignature(id types.ID, signatureBytes []byte, validityHeight types.Height) types.Signature {
	return signature{
		ID:             id,
		SignatureBytes: signatureBytes,
		ValidityHeight: validityHeight,
	}
}
