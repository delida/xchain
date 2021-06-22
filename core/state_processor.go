// Copyright 2015 The MOAC-core Authors
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

package core

import (
	"math/big"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/crypto"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/MoacLib/params"
	"github.com/MOACChain/MoacLib/state"
	"github.com/MOACChain/MoacLib/types"
	"github.com/MOACChain/MoacLib/vm"
	"github.com/MOACChain/xchain/consensus"
)

var (
	ADD_QUERY_TASK   = "0x726bbce8"
	CHECK_UPGRADABLE = "0xb81fa3f5"
	DELAYED_SEND     = "0xdef0412f"
	ENROLL           = "0x6111ca21"
	JOIN_SCS         = "0x4ad7595b"
	QUERY_TASK       = "0xaea722e8"
	SEND_TASK        = "0x650aed4f"
	EXEC             = "0xc1c0e9c4"
	INIT             = "0xe1c7392a"

	SYSCALLGASPRICE = big.NewInt(20000000000)
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Process processes the state changes according to the MoacNode rules by running
// the transaction messages using the statedb and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the receipts and logs accumulated during the process and
// returns the amount of gasRemaining that was used in the process. If any of the
// transactions failed to execute due to insufficient gasRemaining it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, cfg vm.Config, liveFlag bool) (types.Receipts, []*types.Log, *big.Int, error) {
	log.Debugf("[core/state_processor.go->Process]")
	var (
		receipts     types.Receipts
		totalUsedGas = big.NewInt(0)
		header       = block.Header()
		allLogs      []*types.Log
		gp           = new(GasPool).AddGas(block.GasLimit())
	)
	var txs []*types.Transaction

	//check if any system contract in the txs already, if yes, remove it.
	//this is to remove any faked sys call from network
	for _, tx := range block.Transactions() {
		txs = append(txs, tx)
	}

	// Iterate over and process the individual transactions
	for i, tx := range txs {
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		receipt, _, err := ApplyTransaction(
			p.config,
			p.bc,
			nil,
			gp,
			statedb,
			header,
			tx,
			totalUsedGas,
			cfg,
			nil,
		)
		if err != nil {
			return nil, nil, nil, err
		}
		receipts = append(receipts, receipt)
		allLogs = append(allLogs, receipt.Logs...)
	}
	p.engine.Finalize(p.bc, header, statedb, txs, block.Uncles(), receipts, liveFlag)

	return receipts, allLogs, totalUsedGas, nil
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gasRemaining used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(
	config *params.ChainConfig,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	txpool *TxPool,
) (*types.Receipt, *big.Int, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if msg.To() != nil {
		//load address flag associated with this address
		shardingflag, _ := statedb.GetFlag(*msg.To())
		msg.SetShardingFlag(shardingflag)
	} else {
		//creation of newcontract will be done in synchronous way
		msg.SetShardingFlag(0)
	}

	msg.GasLimit()
	if err != nil {
		return nil, nil, err
	}
	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc, author, txpool)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, statedb, config, cfg, nil)
	// Apply the transaction to the current state (included in the env)

	_, gas, failed, err := ApplyMessage(vmenv, msg, gp)
	log.Debugf("Apply message %s with gas used: %d, failed: %v, err: %v", common.Bytes2Hex(msg.GetMsgHash().Bytes()), gas, failed, err)
	if err != nil {
		return nil, nil, err
	}

	// Update the state with pending changes
	var root []byte

	//Use the latest
	statedb.Finalise(true)
	usedGas.Add(usedGas, gas)

	// Create a new receipt for the transaction, storing the intermediate root and gasRemaining used by the tx
	// based on the eip phase, we're passing wether the root touch-delete accounts.
	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)

	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
		statedb.SetFlag(receipt.ContractAddress, 0)
	} else {
		receipt.ContractAddress = *msg.To()
	}

	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	if bc != nil {
		evmCache := vmenv
		evmCache.StateDB = statedb.Copy()
		vm.SetEVM(evmCache)
	}

	return receipt, gas, err
}

func ApplyTransactionForCalculateGas(
	config *params.ChainConfig,
	bc *BlockChain,
	author *common.Address,
	gp *GasPool,
	statedb *state.StateDB,
	header *types.Header,
	tx *types.Transaction,
	usedGas *big.Int,
	cfg vm.Config,
	txpool *TxPool,
) (*types.Receipt, *big.Int, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if msg.To() != nil {
		//load address flag associated with this address
		shardingflag, _ := statedb.GetFlag(*msg.To())
		msg.SetShardingFlag(shardingflag)
	} else {
		//creation of newcontract will be done in synchronous way
		msg.SetShardingFlag(0)
	}

	msg.GasLimit()
	if err != nil {
		return nil, nil, err
	}
	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc, author, txpool)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, statedb, config, cfg, nil)
	// Apply the transaction to the current state (included in the env)

	_, gas, failed, err := ApplyMessageForCalculateGas(vmenv, msg, gp)
	if err != nil {
		return nil, nil, err
	}

	// Update the state with pending changes
	var root []byte

	//Use the latest
	statedb.Finalise(true)
	usedGas.Add(usedGas, gas)

	// Create a new receipt for the transaction, storing the intermediate root and gasRemaining used by the tx
	// based on the eip phase, we're passing wether the root touch-delete accounts.
	receipt := types.NewReceipt(root, failed, usedGas)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = new(big.Int).Set(gas)

	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
		statedb.SetFlag(receipt.ContractAddress, 0)
	} else {
		receipt.ContractAddress = *msg.To()
	}

	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = statedb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})
	return receipt, gas, err
}
