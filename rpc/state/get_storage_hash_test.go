// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package state

import (
	"testing"

	"github.com/Phala-Network/go-substrate-rpc-client/v3/types"
	"github.com/stretchr/testify/assert"
)

func TestState_GetStorageHashLatest(t *testing.T) {
	key := types.NewStorageKey(types.MustHexDecodeString("0x3a636f6465"))
	hash, err := state.GetStorageHashLatest(key)
	assert.NoError(t, err)
	var expected types.Hash
	copy(expected[:], types.MustHexDecodeString(mockSrv.storageHashHex))
	assert.Equal(t, expected, hash)
}

func TestState_GetStorageHash(t *testing.T) {
	key := types.NewStorageKey(types.MustHexDecodeString("0x3a636f6465"))
	hash, err := state.GetStorageHash(key, mockSrv.blockHashLatest)
	assert.NoError(t, err)
	var expected types.Hash
	copy(expected[:], types.MustHexDecodeString(mockSrv.storageHashHex))
	assert.Equal(t, expected, hash)
}
