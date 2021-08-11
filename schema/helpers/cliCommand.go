/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package helpers

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/spf13/cobra"
)

type CLICommand interface {
	ReadInt64(CLIFlag) int64
	ReadInt(CLIFlag) int
	ReadBool(CLIFlag) bool
	ReadString(CLIFlag) string
	ReadBaseReq(client.Context) base.BaseReq

	CreateCommand(func(command *cobra.Command, args []string) error) *cobra.Command
}
