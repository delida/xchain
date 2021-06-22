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
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"gopkg.in/fatih/set.v0"

	"github.com/MOACChain/MoacLib/common"
	"github.com/MOACChain/MoacLib/log"
	"github.com/MOACChain/MoacLib/mcdb"
	"github.com/MOACChain/MoacLib/rlp"
	"github.com/MOACChain/xchain/accounts/abi/bind"
	"github.com/MOACChain/xchain/accounts/keystore"
	"github.com/MOACChain/xchain/core"
	"github.com/MOACChain/xchain/dkg"
	"github.com/MOACChain/xchain/event"
	"github.com/MOACChain/xchain/mcclient"
	"github.com/MOACChain/xchain/vnode/config"
	"github.com/MOACChain/xchain/xdefi/vaultx"
	"github.com/MOACChain/xchain/xdefi/vaulty"
	"github.com/MOACChain/xchain/xdefi/xevents"
)

const (
	VaultCheckInterval                   = 3
	VssCheckInterval                     = 3
	PersistSeenVaultEventWithSigChanSize = 8192
	VaultEventsBatchChanSize             = 0
	ScanStep                             = uint64(20)
	MaxBlockNumber                       = uint64(1000000000000)
	Gwei                                 = int64(1000000000)
	DefaultGasPrice                      = int64(5) * Gwei
	DefaultGasLimit                      = int64(300000)
	BlockDelay                           = 12
	MaxEmptyBatchBlocks                  = 200
	MintBatchSize                        = 30
	MintIntervalBlocks                   = 50
	MyTurnSeed                           = 10000

	// event type
	DEPOSIT = 0
	BURN    = 1
)

var (
	XeventsXYAddr = common.HexToAddress("0x0000000000000000000000000000000000010000")
	XeventsYXAddr = common.HexToAddress("0x0000000000000000000000000000000000010001")
)

type Sentinel struct {
	db                    mcdb.Database
	dkg                   *dkg.DKG
	key                   *keystore.Key
	scope                 event.SubscriptionScope
	chainHeadSub          event.Subscription
	config                *config.Configuration
	Rpc                   string
	vaultsConfig          *VaultPairListConfig
	batchNumber           uint64
	batchEndNumber        uint64
	vaultEventWithSigFeed event.Feed

	// vault events
	VaultEventsReceived map[common.Hash]*set.Set
	VaultEvents         map[common.Hash]*core.VaultEventWithSig
	VaultEventProcessMu sync.RWMutex

	// persist
	PersistSeenVaultEventWithSigChan chan PersistSeenVaultEventWithSig
	VaultEventsBatchChan             chan VaultEventsBatch

	// [0x123][123, 456] => true
	VaultScanBlocks map[common.Address]map[string]bool
}

type EventData struct {
	EventData   []byte
	Sig         []byte
	BlockNumber *big.Int
}

type PersistSeenVaultEventWithSig struct {
	batchNumber       uint64
	vaultEventWithSig *core.VaultEventWithSig
}

type VaultEventsBatch struct {
	Batch            map[common.Hash]bool
	Vault            common.Address
	StartBlockNumber uint64
	EndBlockNumber   uint64
	StoreCounter     uint64
}

type XdefiContextConfig struct {
	XChainId         uint64
	XChainFuncPrefix string
	XChainRPC        string
	YChainId         uint64
	YChainFuncPrefix string
	YChainRPC        string
	VaultXAddr       common.Address
	VaultYAddr       common.Address
}

func (config *XdefiContextConfig) String() string {
	return fmt.Sprintf(
		"%d, %s, %s, %d, %s, %s, %x, %x",
		config.XChainId,
		config.XChainFuncPrefix,
		config.XChainRPC,
		config.YChainId,
		config.YChainFuncPrefix,
		config.YChainRPC,
		config.VaultXAddr,
		config.VaultYAddr,
	)
}

/////////////////////////////////////////////////////////////
///////////////    XdefiContext       ///////////////////////
/////////////////////////////////////////////////////////////

type XdefiContext struct {
	EventType     uint64
	Config        *XdefiContextConfig
	Sentinel      *Sentinel
	ClientX       *mcclient.Client
	ClientY       *mcclient.Client
	ClientXevents *mcclient.Client
	VaultX        *vaultx.VaultX
	VaultY        *vaulty.VaultY
	XeventsXY     *xevents.XEvents
	XeventsYX     *xevents.XEvents
}

func logX2Y(format string, args ...interface{}) {
	log.Infof("[Vault X -> Y] "+format, args...)
}

func logY2X(format string, args ...interface{}) {
	log.Infof("[Vault Y -> X] "+format, args...)
}

func (xdefiContext *XdefiContext) IsDeposit() bool {
	return xdefiContext.EventType == DEPOSIT
}

func (xdefiContext *XdefiContext) LogFunc() func(format string, args ...interface{}) {
	if xdefiContext.IsDeposit() {
		return logX2Y
	} else {
		return logY2X
	}
}

func (xdefiContext *XdefiContext) Xevents() *xevents.XEvents {
	if xdefiContext.IsDeposit() {
		return xdefiContext.XeventsXY
	} else {
		return xdefiContext.XeventsYX
	}
}

func (xdefiContext *XdefiContext) VaultAddrFrom() common.Address {
	if xdefiContext.IsDeposit() {
		return xdefiContext.Config.VaultXAddr
	} else {
		return xdefiContext.Config.VaultYAddr
	}
}

func (xdefiContext *XdefiContext) VaultAddrTo() common.Address {
	if xdefiContext.IsDeposit() {
		return xdefiContext.Config.VaultYAddr
	} else {
		return xdefiContext.Config.VaultXAddr
	}
}

func (xdefiContext *XdefiContext) ClientFrom() *mcclient.Client {
	if xdefiContext.IsDeposit() {
		return xdefiContext.ClientX
	} else {
		return xdefiContext.ClientY
	}
}

func (xdefiContext *XdefiContext) ClientTo() *mcclient.Client {
	if xdefiContext.IsDeposit() {
		return xdefiContext.ClientY
	} else {
		return xdefiContext.ClientX
	}
}

func (xdefiContext *XdefiContext) PrepareClients() {
	var xclient *mcclient.Client
	var yclient *mcclient.Client
	var err error

	logFunc := xdefiContext.LogFunc()
	xChainId := xdefiContext.Config.XChainId
	xChainFuncPrefix := xdefiContext.Config.XChainFuncPrefix
	xChainRPC := xdefiContext.Config.XChainRPC
	yChainId := xdefiContext.Config.YChainId
	yChainFuncPrefix := xdefiContext.Config.YChainFuncPrefix
	yChainRPC := xdefiContext.Config.YChainRPC

	xclient, err = mcclient.Dial(xChainRPC)
	if err != nil {
		logFunc("Unable to connect to network:%v\n", err)
		return
	}
	xclient.SetFuncPrefix(xChainFuncPrefix)
	// sanity check chain id
	xchainId_, err := xclient.ChainID(context.Background())
	if err != nil {
		log.Errorf("------------client chain id err: %v -----------------", err)
		return
	}
	if xChainId != xchainId_.Uint64() {
		log.Errorf(
			"Chain ID does not match, have: %d, want: %d, check vaultx.json and restart",
			xchainId_, xChainId,
		)
		return
	}

	yclient, err = mcclient.Dial(yChainRPC)
	if err != nil {
		logFunc("Unable to connect to network:%v\n", err)
		return
	}
	yclient.SetFuncPrefix(yChainFuncPrefix)
	// sanity check chain id
	ychainId_, err := yclient.ChainID(context.Background())
	if err != nil {
		log.Errorf("------------client chain id err: %v -----------------", err)
		return
	}
	if yChainId != ychainId_.Uint64() {
		log.Errorf(
			"Chain ID does not match, have: %d, want: %d, check vaultx.json and restart",
			ychainId_, yChainId,
		)
		return
	}

	clientXevents, err := mcclient.Dial(xdefiContext.Sentinel.Rpc)
	if err != nil {
		log.Errorf("------------client x err: %v %s-----------------", err, xdefiContext.Sentinel.Rpc)
		return
	}

	xdefiContext.ClientX = xclient
	xdefiContext.ClientY = yclient
	xdefiContext.ClientXevents = clientXevents

	return
}

func (xdefiContext *XdefiContext) PrepareContracts() {
	vaultXAddr := xdefiContext.Config.VaultXAddr
	vaultYAddr := xdefiContext.Config.VaultYAddr
	clientX := xdefiContext.ClientX
	clientY := xdefiContext.ClientY
	clientXevents := xdefiContext.ClientXevents

	vaultxContract, err := vaultx.NewVaultX(
		vaultXAddr,
		clientX,
	)
	if err != nil {
		log.Errorf("------------client new vault X contract err: %v -----------------", err)
		return
	}

	vaultyContract, err := vaulty.NewVaultY(
		vaultYAddr,
		clientY,
	)
	if err != nil {
		log.Errorf("------------client new vault Y contract err: %v -----------------", err)
		return
	}

	// xevents contract
	xeventsXYContract, err := xevents.NewXEvents(
		XeventsXYAddr,
		clientXevents,
	)
	if err != nil {
		log.Errorf("------------client new xevents XY contract err: %v -----------------", err)
		return
	}

	xeventsYXContract, err := xevents.NewXEvents(
		XeventsYXAddr,
		clientXevents,
	)
	if err != nil {
		log.Errorf("------------client new xevents YX contract err: %v -----------------", err)
		return
	}

	xdefiContext.VaultX = vaultxContract
	xdefiContext.VaultY = vaultyContract
	xdefiContext.XeventsXY = xeventsXYContract
	xdefiContext.XeventsYX = xeventsYXContract
	return
}

func (xdefiContext *XdefiContext) IsReady() bool {
	oneIsNil := xdefiContext.ClientX == nil
	oneIsNil = oneIsNil || xdefiContext.ClientY == nil
	oneIsNil = oneIsNil || xdefiContext.ClientXevents == nil
	oneIsNil = oneIsNil || xdefiContext.VaultX == nil
	oneIsNil = oneIsNil || xdefiContext.VaultY == nil
	oneIsNil = oneIsNil || xdefiContext.XeventsXY == nil
	oneIsNil = oneIsNil || xdefiContext.XeventsYX == nil

	if oneIsNil {
		log.Errorf(
			"---------Prepare common context: one of clients/contracts is nil: %v, # %v, # %v, # %v, # %v, # %v # %v ------------",
			xdefiContext.ClientX, xdefiContext.ClientY, xdefiContext.ClientXevents,
			xdefiContext.VaultX, xdefiContext.VaultY,
			xdefiContext.XeventsXY, xdefiContext.XeventsYX,
		)
	}

	return !oneIsNil
}

/////////////////////////////////////////////////////////////
///////////////       Sentinel        ///////////////////////
/////////////////////////////////////////////////////////////

func New(
	bc *core.BlockChain,
	vaultsConfig *VaultPairListConfig,
	db mcdb.Database,
	dkg *dkg.DKG,
	rpc string,
	key *keystore.Key,
) *Sentinel {
	sentinel := &Sentinel{
		db:                   db,
		VaultEvents:          make(map[common.Hash]*core.VaultEventWithSig),
		VaultEventsReceived:  make(map[common.Hash]*set.Set),
		config:               bc.VnodeConfig(),
		vaultsConfig:         vaultsConfig,
		dkg:                  dkg,
		key:                  key,
		Rpc:                  "http://" + rpc,
		VaultScanBlocks:      make(map[common.Address]map[string]bool),
		VaultEventsBatchChan: make(chan VaultEventsBatch, VaultEventsBatchChanSize),
		PersistSeenVaultEventWithSigChan: make(
			chan PersistSeenVaultEventWithSig, PersistSeenVaultEventWithSigChanSize),
	}
	sentinel.scope.Open()
	log.Infof("sentinel start with config: %v", sentinel.vaultsConfig)
	go sentinel.start()

	return sentinel
}

func (sentinel *Sentinel) threshold() int {
	if sentinel.dkg.IsVSSReady() {
		return sentinel.dkg.Bls.Threshold
	} else {
		return 0
	}
}

func (sentinel *Sentinel) SubscribeVaultEventWithSig(ch chan<- core.VaultEventWithSig) event.Subscription {
	return sentinel.scope.Track(sentinel.vaultEventWithSigFeed.Subscribe(ch))
}

func (sentinel *Sentinel) PrintVaultEventsReceived() {
	log.Debugf("Vault events received:")
	for hash, received := range sentinel.VaultEventsReceived {
		log.Debugf("\t%x, %d", hash.Bytes()[:8], received.Size())
	}
}

func (sentinel *Sentinel) getTransactor(
	client *mcclient.Client,
	gasPrice int64,
	gasLimit int64,
) *bind.TransactOpts {
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		return nil
	}

	// build transactor
	transactor, _ := bind.NewKeyedTransactorWithChainID(
		sentinel.key.PrivateKey,
		chainId,
	)
	transactor.GasPrice = big.NewInt(gasPrice)
	transactor.GasLimit = uint64(gasLimit)

	nonceAt, err := client.NonceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return nil
	}
	transactor.Nonce = big.NewInt(int64(nonceAt))

	return transactor
}

func (sentinel *Sentinel) IsVaultEventsBatchAllReceived(batch VaultEventsBatch, threshold int) bool {
	for vaultEventHash, _ := range batch.Batch {
		if sentinel.VaultEventsReceived[vaultEventHash].Size() < threshold {
			return false
		}
	}
	return true
}

func (sentinel *Sentinel) RecordVaultEventWithSig(
	vaultEventWithSig *core.VaultEventWithSig,
) common.Hash {
	// both mc/handler and this process call this function
	sentinel.VaultEventProcessMu.Lock()
	defer sentinel.VaultEventProcessMu.Unlock()

	vaultEvent := vaultEventWithSig.Event
	vaultEventHash := vaultEvent.Hash()
	// keep count of the event
	if received, found := sentinel.VaultEventsReceived[vaultEventHash]; found {
		received.Add(string(vaultEventWithSig.Blssig))
	} else {
		sentinel.VaultEventsReceived[vaultEventHash] = set.New()
		sentinel.VaultEventsReceived[vaultEventHash].Add(string(vaultEventWithSig.Blssig))
	}
	// keep the event
	sentinel.VaultEvents[vaultEventHash] = vaultEventWithSig

	return vaultEventHash
}

func (sentinel *Sentinel) myTurn(number uint64) bool {
	// use node index to determine if this node should update xevents this round
	// batchNumber's value should be consistent between validators
	nodeIndex, nodeCount := sentinel.dkg.FindAddressInNodelist(sentinel.dkg.Vssid)
	myTurn := false
	if nodeIndex == -1 {
		log.Errorf(
			"----------- My turn err: index: %d, count: %d, last block: %d-----------------",
			nodeIndex,
			nodeCount,
			number,
		)
		return false
	} else {
		myTurn = number%uint64(nodeCount) == uint64(nodeIndex)
	}

	return myTurn
}

func (sentinel *Sentinel) scanCacheCheck(vault common.Address, start, end uint64) bool {
	if vaultBlocks, found := sentinel.VaultScanBlocks[vault]; found {
		key := fmt.Sprintf("%d,%d", start, end)
		_, found := vaultBlocks[key]
		return found
	}

	return false
}

func (sentinel *Sentinel) scanCacheStore(vault common.Address, start, end uint64) {
	if _, found := sentinel.VaultScanBlocks[vault]; !found {
		sentinel.VaultScanBlocks[vault] = make(map[string]bool)
	}
	key := fmt.Sprintf("%d,%d", start, end)
	sentinel.VaultScanBlocks[vault][key] = true
}

func (sentinel *Sentinel) prepareCommonContext(
	xChainId uint64,
	xChainFuncPrefix string,
	xChainRPC string,
	xVaultAddr common.Address,
	yChainId uint64,
	yChainFuncPrefix string,
	yChainRPC string,
	yVaultAddr common.Address,
) *XdefiContext {
	xdefiContextConfig := &XdefiContextConfig{
		xChainId,
		xChainFuncPrefix,
		xChainRPC,
		yChainId,
		yChainFuncPrefix,
		yChainRPC,
		xVaultAddr,
		yVaultAddr,
	}
	xdefiContext := &XdefiContext{}
	xdefiContext.EventType = DEPOSIT
	xdefiContext.Config = xdefiContextConfig
	xdefiContext.Sentinel = sentinel

	log.Infof("Xdefi context config: %s", xdefiContext.Config)

	// clients are contracts are set inline with context
	xdefiContext.PrepareClients()
	xdefiContext.PrepareContracts()

	if xdefiContext.IsReady() {
		return xdefiContext
	} else {
		return nil
	}
}

func (sentinel *Sentinel) initVaultParams(xdefiContext *XdefiContext) (uint64, uint64) {
	xevents := xdefiContext.Xevents()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	callOpts := &bind.CallOpts{}
	vaultWatermark, err := xevents.VaultWatermark(callOpts, vaultAddrFrom)
	if err != nil {
		log.Errorf("------------client watermark block err: %v -----------------", err)
		return 0, 0
	}

	batchNumber := uint64(0)
	storeCounter := uint64(0)
	lastBlock := uint64(0)
	var createdAt *big.Int
	// fallback to use createAt
	if vaultWatermark.Uint64() == 0 {
		if xdefiContext.IsDeposit() {
			createdAt, err = xdefiContext.VaultX.CreatedAt(callOpts)
		} else {
			createdAt, err = xdefiContext.VaultY.CreatedAt(callOpts)
		}
		if err != nil {
			log.Errorf(
				"------------client CreateAt err: %v, set lastBlock to %d-----------------",
				err, lastBlock,
			)
			return 0, 0
		} else {
			// first time
			lastBlock = createdAt.Uint64()
			storeCounter = uint64(0)
		}
	} else {
		// normal case
		// we've scanned watermark block, proceed to next one
		batchNumber = vaultWatermark.Uint64()
		lastBlock = vaultWatermark.Uint64() + 1
		counter, err := xevents.VaultStoreCounter(
			callOpts, vaultAddrFrom, big.NewInt(int64(batchNumber)),
		)
		if err != nil {
			log.Errorf("---------------client vault store counter error %v------------", err)
			return 0, 0
		}
		storeCounter = counter.Uint64()
	}

	// share with other go routine
	sentinel.batchNumber = lastBlock

	return lastBlock, storeCounter
}

func (sentinel *Sentinel) scanVaultEvents(
	xdefiContext *XdefiContext,
	lastBlock uint64,
	storeCounter uint64,
) *VaultEventsBatch {
	eventCount := uint64(0)
	vaultX := xdefiContext.VaultX
	vaultY := xdefiContext.VaultY
	logFunc := xdefiContext.LogFunc()
	clientFrom := xdefiContext.ClientFrom()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	startBlock := lastBlock

	// the outer for loop is to skip void blocks with no events
	for {
		// throttling for sending query
		time.Sleep(1000 * time.Millisecond)
		// once the filter return some events, we're done for this round.
		if eventCount > 0 {
			return nil
		}
		currentBlock, err := clientFrom.BlockNumber(context.Background())
		if err != nil {
			log.Errorf(
				"----------------- sentinel: unable to get current block number from chain: %v ---------------", err)
		}
		if lastBlock > currentBlock {
			return nil
		}

		// filter events from vaultx, [start, end] are inclusive
		endBlock := lastBlock + ScanStep - 1
		if endBlock > currentBlock-BlockDelay {
			endBlock = currentBlock - BlockDelay
		}
		if endBlock < lastBlock {
			endBlock = lastBlock + 1
		}

		filterOpts := &bind.FilterOpts{
			Context: context.Background(),
			Start:   lastBlock,
			End:     &endBlock,
		}

		logFunc("--------------- filter opts:[batch=%d], start: %d, end: %d-----------",
			startBlock, lastBlock, endBlock,
		)

		//var itrY *vaulty.VaultYTokenDepositIterator
		batch := make(map[common.Hash]bool)
		if xdefiContext.IsDeposit() {
			itrX, err := vaultX.FilterTokenDeposit(
				filterOpts,
				[]common.Address{},
				[]common.Address{},
				[]*big.Int{},
			)

			if err != nil {
				log.Errorf(
					"--------------- sentinel: unable to get token deposit iterator: %v ------------", err)
				return nil
			}

			if !sentinel.dkg.IsVSSReady() {
				log.Errorf("------------------ sentinel: unable to get bls signer: %v ----------", err)
				return nil
			}

			for itrX.Next() {
				eventCount += 1
				event := itrX.Event
				vaultEvent := core.VaultEvent{
					vaultAddrFrom,
					event.SourceChainid,
					event.SourceToken,
					event.MappedChainid,
					event.MappedToken,
					event.From,
					event.Amount,
					event.Nonce,
					event.BlockNumber,
					event.Tip,
				}
				blssig := sentinel.dkg.Bls.SignBytes(vaultEvent.Hash().Bytes())
				vaultEventWithSig := core.VaultEventWithSig{
					vaultEvent,
					blssig,
				}
				logFunc(
					"-------Scanned vault event: %x, nonce: %x, block: %d, source: %x, mapped: %x--------",
					vaultAddrFrom,
					event.Nonce,
					event.BlockNumber,
					event.SourceToken,
					event.MappedToken,
				)
				// process the event in this sentinel, include this node itself.
				vaultEventHash := sentinel.RecordVaultEventWithSig(&vaultEventWithSig)
				batch[vaultEventHash] = true

				// broad cast to other nodes
				sentinel.vaultEventWithSigFeed.Send(vaultEventWithSig)
			}
			logFunc("---------------- raw filter result: batch size = %d--------------", len(batch))
		} else {
			itrY, err := vaultY.FilterTokenBurn(
				filterOpts,
				[]common.Address{},
				[]common.Address{},
				[]*big.Int{},
			)

			if err != nil {
				log.Errorf(
					"--------------- sentinel: unable to get token deposit iterator: %v ------------", err)
				return nil
			}

			if !sentinel.dkg.IsVSSReady() {
				log.Errorf("------------------ sentinel: unable to get bls signer: %v ----------", err)
				return nil
			}

			for itrY.Next() {
				eventCount += 1
				event := itrY.Event
				vaultEvent := core.VaultEvent{
					vaultAddrFrom,
					event.SourceChainid,
					event.SourceToken,
					event.MappedChainid,
					event.MappedToken,
					event.From,
					event.Amount,
					event.Nonce,
					event.BlockNumber,
					event.Tip,
				}
				blssig := sentinel.dkg.Bls.SignBytes(vaultEvent.Hash().Bytes())
				vaultEventWithSig := core.VaultEventWithSig{
					vaultEvent,
					blssig,
				}
				logFunc(
					"-------Scanned vault event: %x, nonce: %x, block: %d, source: %x, mapped: %x--------",
					vaultAddrFrom,
					event.Nonce,
					event.BlockNumber,
					event.SourceToken,
					event.MappedToken,
				)
				// process the event in this sentinel, include this node itself.
				vaultEventHash := sentinel.RecordVaultEventWithSig(&vaultEventWithSig)
				batch[vaultEventHash] = true

				// broad cast to other nodes
				sentinel.vaultEventWithSigFeed.Send(vaultEventWithSig)
			}
			logFunc("---------------- raw filter result: batch size = %d--------------", len(batch))
		}

		// once we receive the batch with events,
		// scan for this round is done, just return
		if sentinel.batchNumber != 0 && len(batch) > 0 {
			logFunc(
				"----------- New batch mined (commit = %t), number: %d, start: %d, end: %d, events: %d -------------",
				sentinel.myTurn(MyTurnSeed),
				sentinel.batchNumber,
				startBlock,
				endBlock,
				len(batch),
			)
			return &VaultEventsBatch{
				batch,
				vaultAddrFrom,
				startBlock,
				endBlock,
				storeCounter,
			}
		}

		// corner case where we scan many blocks but getting no events
		if endBlock-startBlock >= MaxEmptyBatchBlocks && len(batch) == 0 {
			return &VaultEventsBatch{
				batch,
				vaultAddrFrom,
				startBlock,
				endBlock,
				storeCounter,
			}
		}

		lastBlock += ScanStep
	}
}

func (sentinel *Sentinel) commitBatch(
	xdefiContext *XdefiContext,
	batch *VaultEventsBatch,
) (uint64, uint64, uint64, *bind.TransactOpts) {
	clientXevents := xdefiContext.ClientXevents
	xevents := xdefiContext.Xevents()
	logFunc := xdefiContext.LogFunc()

	logFunc(
		"-----------Vault FROM commit batch BEGIN from %d to %d ]-----------------",
		batch.StartBlockNumber,
		batch.EndBlockNumber,
	)
	defer logFunc(
		"-----------Vault FROM commit batch END from %d to %d-----------------",
		batch.StartBlockNumber,
		batch.EndBlockNumber,
	)

	// sanity check chain id
	chainId, err := clientXevents.ChainID(context.Background())
	if err != nil {
		log.Errorf("------------client xevents chain id err: %v -----------------", err)
		return 0, 0, 0, nil
	}

	balance, err := clientXevents.BalanceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return 0, 0, 0, nil
	}
	nonceAt, err := clientXevents.NonceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return 0, 0, 0, nil
	}
	logFunc(
		"---------  %x, balance: %d, nonce: %d, chainId: %d ----------",
		sentinel.key.Address.Bytes(), balance, nonceAt, chainId,
	)
	//  if we don't have money, return
	if balance.Uint64() == 0 {
		return 0, 0, 0, nil
	}
	callOpts := &bind.CallOpts{}
	vaultWatermark, err := xevents.VaultWatermark(
		callOpts, XeventsXYAddr,
	)
	if err != nil {
		log.Errorf("------------client watermark block err: %v -----------------", err)
		return 0, 0, 0, nil
	}
	// prevent re-entry
	if vaultWatermark.Uint64() >= batch.StartBlockNumber {
		log.Errorf(
			"---------------Vault FROM skip batch commit: vault: %d, batch number: %d ----------",
			vaultWatermark.Uint64(),
			batch.StartBlockNumber,
		)
		return 0, 0, 0, nil
	}

	// setup transactor
	transactor := sentinel.getTransactor(
		clientXevents,
		DefaultGasPrice,
		DefaultGasLimit,
	)

	// order the commit sequence by nonce
	nonce := 0
	nonces := make([]int, 0)
	noncesMap := make(map[int]common.Hash)
	for vaultEventHash, _ := range batch.Batch {
		vaultEventWithSig := sentinel.VaultEvents[vaultEventHash]
		vaultEvent := vaultEventWithSig.Event
		nonce = int(vaultEvent.Nonce.Int64())
		nonces = append(nonces, int(vaultEvent.Nonce.Int64()))
		noncesMap[nonce] = vaultEventHash
		logFunc(
			"------------- Vault event: vault: %x, tokenMapping: %s, sha256: %x, nonce: %d ---------------",
			vaultEvent.Vault,
			vaultEvent.TokenMapping(),
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
		)
	}
	sort.Ints(nonces)

	var vaultEventHash common.Hash
	logFunc("-------------- Vault FROM nonces: %v -------------", nonces)

	committed := uint64(0)
	errors := uint64(0)
	omitted := uint64(0)
	// commit all events
	for _, nonce := range nonces {
		vaultEventHash = noncesMap[nonce]
		vaultEventWithSig := sentinel.VaultEvents[vaultEventHash]
		vaultEvent := vaultEventWithSig.Event

		callOpts := &bind.CallOpts{}
		watermark, err := xevents.VaultEventWatermark(
			callOpts, vaultEvent.Vault, vaultEvent.TokenMappingSha256(),
		)
		if err != nil {
			log.Errorf("-----------------Vault FROM watermark err: %v -------------", err)
		} else {
			logFunc(
				"---------xevents tokenmapping watermark %d, this tokenmapping nonce: %d ----",
				watermark, vaultEvent.Nonce,
			)
		}

		if watermark.Int64() > vaultEvent.Nonce.Int64() {
			omitted += 1
			logFunc(
				"---------------- xevents tokenmapping nonce low, omitted: %d------------------",
				vaultEvent.Nonce.Int64(),
			)
			continue
		}

		logFunc(
			"--------- xevents Store tx: vault: %x, tokenmapping: %x, nonce: %d, block: %d, %d, %d -------",
			vaultEvent.Vault.Bytes()[:8],
			vaultEvent.TokenMappingSha256(),
			vaultEvent.Nonce,
			vaultEvent.BlockNumber,
			transactor.GasPrice,
			transactor.GasLimit,
		)
		tx, err := xevents.Store(
			transactor,
			vaultEventWithSig.Blssig,
			vaultEvent.Vault,
			vaultEvent.Nonce,
			vaultEvent.TokenMappingSha256(),
			vaultEvent.BlockNumber,
			vaultEvent.Bytes(),
		)
		transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)

		if err != nil {
			errors += 1
			log.Errorf("------------Vault FROM sentinel xevents store err: %v ------------------", err)
		} else {
			logFunc(
				"---------Vault FROM sentinel xevents store tx: %x, from: %x, nonce: %d ---------",
				tx.Hash(), sentinel.key.Address.Bytes(), tx.Nonce(),
			)
			committed += 1
		}
	}

	return omitted, committed, errors, transactor
}

func (sentinel *Sentinel) watchVault(xdefiContext *XdefiContext) {
	logFunc := xdefiContext.LogFunc()
	defer logFunc("*********************END WATCH LOOP***********************")
	for {
		// sleep for interval
		time.Sleep(VaultCheckInterval * time.Second)

		// # 0
		xevents := xdefiContext.Xevents()
		lastBlock, storeCounter := sentinel.initVaultParams(xdefiContext)

		// # 1
		batch := sentinel.scanVaultEvents(
			xdefiContext,
			lastBlock,
			storeCounter,
		)

		if batch == nil {
			continue
		}

		if sentinel.myTurn(MyTurnSeed) {
			// # 2
			omitted, committed, errors, transactor := sentinel.commitBatch(
				xdefiContext,
				batch,
			)

			if transactor == nil {
				continue
			}

			// # 3
			wait := 0
			for {
				time.Sleep(3 * time.Second)
				callOpts := &bind.CallOpts{}
				newStoreCounter, err := xevents.StoreCounter(callOpts)
				if err == nil {
					logFunc(
						"------- Before vault watermark commit, store counter %d, prev %d, batch size: %d, omitted: %d, commit %d, errors: %d -----------",
						newStoreCounter,
						batch.StoreCounter,
						len(batch.Batch),
						omitted,
						committed,
						errors,
					)

					// wait for the new store counter to increase and then update vault water mark
					//oldBatchPass := newStoreCounter.Uint64() == batch.StoreCounter && uint64(len(batch.Batch)) == omitted
					oldBatchPass := false
					newBatchPass := newStoreCounter.Uint64() == batch.StoreCounter+uint64(len(batch.Batch))
					if oldBatchPass || newBatchPass {
						logFunc(
							"-------------------MOVE vault watermark HIGHER to %d---------------",
							batch.EndBlockNumber,
						)
						xevents.UpdateVaultWatermark(
							transactor,
							batch.Vault,
							big.NewInt(int64(batch.EndBlockNumber)),
						)
						break
					}
				}
				wait += 1
				if wait > 20 {
					logFunc("--------------wait time out for storeCounter------------")
					break
				}
			}
		}
	}
}

func (sentinel *Sentinel) scanAndForwardVaultEvents(
	xdefiContext *XdefiContext,
	tokenMapping TokenMapping,
) {
	clientTo := xdefiContext.ClientTo()
	xevents := xdefiContext.Xevents()
	vaultAddrFrom := xdefiContext.VaultAddrFrom()
	logFunc := xdefiContext.LogFunc()

	callOpts := &bind.CallOpts{}
	var waterMark *big.Int
	var err error
	if xdefiContext.IsDeposit() {
		waterMark, err = xdefiContext.VaultY.TokenMappingWatermark(
			callOpts,
			common.HexToAddress(tokenMapping.SourceToken),
			common.HexToAddress(tokenMapping.MappedToken),
		)
		if err != nil {
			log.Errorf("-------------- Vault TO watermark err: %v--------------", err)
			return
		}
		logFunc("-----------Vault TO watermark: %d, %s, %s------------------------------",
			waterMark,
			tokenMapping.SourceToken,
			tokenMapping.MappedToken,
		)
	} else {
		waterMark, err = xdefiContext.VaultX.TokenMappingWatermark(
			callOpts,
			common.HexToAddress(tokenMapping.SourceToken),
			common.HexToAddress(tokenMapping.MappedToken),
		)
		if err != nil {
			log.Errorf("-------------- Vault TO watermark err: %v--------------", err)
			return
		}
		logFunc("-----------Vault TO watermark: %d, %s, %s------------------------------",
			waterMark,
			tokenMapping.SourceToken,
			tokenMapping.MappedToken,
		)

	}

	// each node will mint for n blocks, then hand it over to next node
	if !sentinel.myTurn(MyTurnSeed) {
		logFunc("---------Vault TO Not My Turn to Mint ------------------")
		return
	} else {
		logFunc("---------Vault TO My Turn to Mint ------------------")
	}

	// setup transactor
	transactor := sentinel.getTransactor(
		clientTo,
		DefaultGasPrice,
		DefaultGasLimit,
	)
	if transactor == nil {
		return
	}

	// check balance
	balance, err := clientTo.BalanceAt(context.Background(), sentinel.key.Address, nil)
	if err != nil {
		return
	}
	//  if we don't have money, return
	if balance.Uint64() == 0 {
		log.Errorf(
			"---------------No balance for account on Vault TO chain: %x--------------------",
			sentinel.key.Address,
		)
		return
	}

	committed := uint64(0)
	errors := uint64(0)
	for i := int64(0); i < int64(MintBatchSize); i++ {
		var vaultEvent core.VaultEvent
		vaultEventData, err := xevents.VaultEvents(
			callOpts, vaultAddrFrom,
			tokenMapping.Sha256(),
			big.NewInt(waterMark.Int64()+i),
		)
		if err != nil {
			log.Errorf("------------ Xevents: retrieve vault events err: %v---------------", err)
		}
		rlp.DecodeBytes(vaultEventData.EventData, &vaultEvent)
		logFunc(
			"-----------Vault [TO] TO MINT: nonce: %d amount: %d, tip: %d, to: %x, mappedToken: %x--------",
			big.NewInt(waterMark.Int64()+i),
			vaultEvent.Amount,
			vaultEvent.Tip,
			vaultEvent.To,
			vaultEvent.MappedToken,
		)
		// if no more vault event
		if vaultEvent.Amount == nil || vaultEvent.Amount.Int64() == 0 {
			break
		} else {
			if xdefiContext.IsDeposit() {
				tx, err := xdefiContext.VaultY.Mint(
					transactor,
					vaultEvent.SourceToken,
					vaultEvent.MappedToken,
					vaultEvent.To,
					vaultEvent.Amount,
					vaultEvent.Tip,
					vaultEvent.Nonce,
				)
				transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)

				if err != nil {
					errors += 1
					log.Errorf("---------------sentinel vault TO 'mint' tx err: %v ------------------", err)
				} else {
					logFunc(
						"---------sentinel vault [TO] mint() tx: %x, from: %x, nonce: %d, %d, %d ---------",
						tx.Hash(),
						sentinel.key.Address.Bytes(),
						tx.Nonce(),
						transactor.GasPrice,
						transactor.GasLimit,
					)
					committed += 1
				}
			} else {
				tx, err := xdefiContext.VaultX.Withdraw(
					transactor,
					vaultEvent.SourceToken,
					vaultEvent.MappedToken,
					vaultEvent.To,
					vaultEvent.Amount,
					vaultEvent.Tip,
					vaultEvent.Nonce,
				)
				transactor.Nonce = big.NewInt(transactor.Nonce.Int64() + 1)

				if err != nil {
					errors += 1
					log.Errorf("---------------sentinel vault TO 'withdraw' tx err: %v ------------------", err)
				} else {
					logFunc(
						"---------sentinel vault TO withdraw() tx: %x, from: %x, nonce: %d, %d, %d ---------",
						tx.Hash(),
						sentinel.key.Address.Bytes(),
						tx.Nonce(),
						transactor.GasPrice,
						transactor.GasLimit,
					)
					committed += 1
				}
			}
		}
	}

	waterMarkBefore := waterMark.Uint64()
	// wait for watermark to update
	wait := 0
	for {
		if committed == 0 {
			logFunc("--------------No vault TO mint/withdraw tx called-------------------")
			break
		}
		var waterMark *big.Int
		if xdefiContext.IsDeposit() {
			waterMark, err = xdefiContext.VaultY.TokenMappingWatermark(
				callOpts,
				common.HexToAddress(tokenMapping.SourceToken),
				common.HexToAddress(tokenMapping.MappedToken),
			)
			logFunc(
				"---------Vault TO watermark: %d, before: %d, committed: %d ---------",
				waterMark, waterMarkBefore, committed,
			)
		} else {
			waterMark, err = xdefiContext.VaultX.TokenMappingWatermark(
				callOpts,
				common.HexToAddress(tokenMapping.SourceToken),
				common.HexToAddress(tokenMapping.MappedToken),
			)
			logFunc(
				"---------Vault TO watermark: %d, before: %d, committed: %d ---------",
				waterMark, waterMarkBefore, committed,
			)
		}
		if err != nil {
			log.Errorf("---------Err with calling vault Y watermark: %v-----------", err)
			continue
		}
		if waterMark.Uint64() == waterMarkBefore+committed {
			break
		}
		time.Sleep(10 * time.Second)
		wait += 1
		if wait > 12 {
			logFunc(
				"------Vault Y watermark wait time out, watermark: %d, before: %d, committed: %d ------",
				waterMark, waterMarkBefore, committed,
			)
			break
		}
	}
}

func (sentinel *Sentinel) doForward(
	xdefiContext *XdefiContext, tokenMapping TokenMapping,
) {
	logFunc := xdefiContext.LogFunc()
	defer logFunc("*********************END LOOP Foward Vault Events***********************")
	for {
		time.Sleep(VaultCheckInterval * time.Second)

		clientY := xdefiContext.ClientY
		clientXevents := xdefiContext.Xevents()
		vaultyContract := xdefiContext
		xeventsContract := xdefiContext

		if clientY == nil || clientXevents == nil || vaultyContract == nil || xeventsContract == nil {
			log.Errorf(
				"---------Do Forward: one of clients/contracts is nil: %v, # %v, # %v, # %v -------",
				clientY, clientXevents, vaultyContract, xeventsContract,
			)
			continue
		}

		sentinel.scanAndForwardVaultEvents(
			xdefiContext,
			tokenMapping,
		)
	}
}

func (sentinel *Sentinel) start() {
	for {
		time.Sleep(VssCheckInterval * time.Second)
		if sentinel.dkg.IsVSSReady() {
			break
		}
		log.Infof("In sentinel start(): wait for vss ready")
	}
	log.Infof("In sentinel start(): vss is ready, t=%d, start watcher and forwarder",
		sentinel.dkg.Bls.Threshold,
	)
	// create all go routines for watching vault contracts on various blockchains
	for _, vaultPairConfig := range sentinel.vaultsConfig.Vaults {
		vaultx := vaultPairConfig.VaultX
		vaulty := vaultPairConfig.VaultY

		log.Infof("config: %v, %v", vaultx, vaulty)
		// deposit & mint
		xdefiContextXY := sentinel.prepareCommonContext(
			vaultx.ChainId,
			vaultx.ChainFuncPrefix,
			vaultx.ChainRPC,
			common.HexToAddress(vaultx.VaultAddress),
			vaulty.ChainId,
			vaulty.ChainFuncPrefix,
			vaulty.ChainRPC,
			common.HexToAddress(vaulty.VaultAddress),
		)
		xdefiContextXY.EventType = DEPOSIT
		go sentinel.watchVault(xdefiContextXY)
		for _, tokenMapping := range vaultPairConfig.TokenMappings {
			go sentinel.doForward(xdefiContextXY, tokenMapping)
		}

		// burn & withdarw
		xdefiContextYX := sentinel.prepareCommonContext(
			vaultx.ChainId,
			vaultx.ChainFuncPrefix,
			vaultx.ChainRPC,
			common.HexToAddress(vaultx.VaultAddress),
			vaulty.ChainId,
			vaulty.ChainFuncPrefix,
			vaulty.ChainRPC,
			common.HexToAddress(vaulty.VaultAddress),
		)
		xdefiContextXY.EventType = BURN
		go sentinel.watchVault(xdefiContextYX)
		for _, tokenMapping := range vaultPairConfig.TokenMappings {
			go sentinel.doForward(xdefiContextYX, tokenMapping)
		}
	}
}
