// Copyright 2014 The MOAC-core Authors
// This file is part of the MOAC-core library.
//
// The MOAC-core library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The MOAC-core library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the MOAC-core library. If not, see <http://www.gnu.org/licenses/>.

package sentinel

import (
	"math/big"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/types"
	moaccore "github.com/MOACChain/xchain"
	"github.com/MOACChain/xchain/accounts/abi/bind"
)

type XdefiVaultEvent struct {
	Vault         common.Address
	SourceChainid *big.Int
	MappedChainid *big.Int
	SourceToken   common.Address
	MappedToken   common.Address
	From          common.Address
	Amount        *big.Int
	Tip           *big.Int
	Nonce         *big.Int
	BlockNumber   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

type XdefiVaultEventIterator struct {
	Event    *XdefiVaultEvent      // Event containing the contract specifics and raw log
	contract *bind.BoundContract   // Generic contract to use for unpacking event data
	event    string                // Event name to use for unpacking event data
	logs     chan types.Log        // Log channel receiving the found contract events
	sub      moaccore.Subscription // Subscription for errors, completion and termination
	done     bool                  // Whether the subscription completed delivering logs
	fail     error                 // Occurred error to stop iteration
}

type XdefiVault interface {
}
