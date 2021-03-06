// Copyright 2016 The go-elementrem Authors.
// This file is part of the go-elementrem library.
//
// The go-elementrem library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-elementrem library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-elementrem library. If not, see <http://www.gnu.org/licenses/>.

package logger

import (
	"math/big"
	"time"
)

type utctime8601 struct{}

func (utctime8601) MarshalJSON() ([]byte, error) {
	timestr := time.Now().UTC().Format(time.RFC3339Nano)
	// Bounds check
	if len(timestr) > 26 {
		timestr = timestr[:26]
	}
	return []byte(`"` + timestr + `Z"`), nil
}

type JsonLog interface {
	EventName() string
}

type LogEvent struct {
	// Guid string      `json:"guid"`
	Ts utctime8601 `json:"ts"`
	// Level string      `json:"level"`
}

type LogStarting struct {
	ClientString    string `json:"client_impl"`
	ProtocolVersion int    `json:"ele_version"`
	LogEvent
}

func (l *LogStarting) EventName() string {
	return "starting"
}

type P2PConnected struct {
	RemoteId            string `json:"remote_id"`
	RemoteAddress       string `json:"remote_addr"`
	RemoteVersionString string `json:"remote_version_string"`
	NumConnections      int    `json:"num_connections"`
	LogEvent
}

func (l *P2PConnected) EventName() string {
	return "p2p.connected"
}

type P2PDisconnected struct {
	NumConnections int    `json:"num_connections"`
	RemoteId       string `json:"remote_id"`
	LogEvent
}

func (l *P2PDisconnected) EventName() string {
	return "p2p.disconnected"
}

type EleMinerNewBlock struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	LogEvent
}

func (l *EleMinerNewBlock) EventName() string {
	return "ele.miner.new_block"
}

type EleChainReceivedNewBlock struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	RemoteId      string   `json:"remote_id"`
	LogEvent
}

func (l *EleChainReceivedNewBlock) EventName() string {
	return "ele.chain.received.new_block"
}

type EleChainNewHead struct {
	BlockHash     string   `json:"block_hash"`
	BlockNumber   *big.Int `json:"block_number"`
	ChainHeadHash string   `json:"chain_head_hash"`
	BlockPrevHash string   `json:"block_prev_hash"`
	LogEvent
}

func (l *EleChainNewHead) EventName() string {
	return "ele.chain.new_head"
}

type EleTxReceived struct {
	TxHash   string `json:"tx_hash"`
	RemoteId string `json:"remote_id"`
	LogEvent
}

func (l *EleTxReceived) EventName() string {
	return "ele.tx.received"
}

//
//
// The types below are legacy and need to be converted to new format or deleted
//
//

// type P2PConnecting struct {
// 	RemoteId       string `json:"remote_id"`
// 	RemoteEndpoint string `json:"remote_endpoint"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PConnecting) EventName() string {
// 	return "p2p.connecting"
// }

// type P2PHandshaked struct {
// 	RemoteCapabilities []string `json:"remote_capabilities"`
// 	RemoteId           string   `json:"remote_id"`
// 	NumConnections     int      `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PHandshaked) EventName() string {
// 	return "p2p.handshaked"
// }

// type P2PDisconnecting struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnecting) EventName() string {
// 	return "p2p.disconnecting"
// }

// type P2PDisconnectingBadHandshake struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnectingBadHandshake) EventName() string {
// 	return "p2p.disconnecting.bad_handshake"
// }

// type P2PDisconnectingBadProtocol struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnectingBadProtocol) EventName() string {
// 	return "p2p.disconnecting.bad_protocol"
// }

// type P2PDisconnectingReputation struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnectingReputation) EventName() string {
// 	return "p2p.disconnecting.reputation"
// }

// type P2PDisconnectingDHT struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PDisconnectingDHT) EventName() string {
// 	return "p2p.disconnecting.dht"
// }

// type P2PEleDisconnectingBadBlock struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PEleDisconnectingBadBlock) EventName() string {
// 	return "p2p.ele.disconnecting.bad_block"
// }

// type P2PEleDisconnectingBadTx struct {
// 	Reason         string `json:"reason"`
// 	RemoteId       string `json:"remote_id"`
// 	NumConnections int    `json:"num_connections"`
// 	LogEvent
// }

// func (l *P2PEleDisconnectingBadTx) EventName() string {
// 	return "p2p.ele.disconnecting.bad_tx"
// }

// type EleNewBlockBroadcasted struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EleNewBlockBroadcasted) EventName() string {
// 	return "ele.newblock.broadcasted"
// }

// type EleNewBlockIsKnown struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EleNewBlockIsKnown) EventName() string {
// 	return "ele.newblock.is_known"
// }

// type EleNewBlockIsNew struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EleNewBlockIsNew) EventName() string {
// 	return "ele.newblock.is_new"
// }

// type EleNewBlockMissingParent struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EleNewBlockMissingParent) EventName() string {
// 	return "ele.newblock.missing_parent"
// }

// type EleNewBlockIsInvalid struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EleNewBlockIsInvalid) EventName() string {
// 	return "ele.newblock.is_invalid"
// }

// type EleNewBlockChainIsOlder struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EleNewBlockChainIsOlder) EventName() string {
// 	return "ele.newblock.chain.is_older"
// }

// type EleNewBlockChainIsCanonical struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EleNewBlockChainIsCanonical) EventName() string {
// 	return "ele.newblock.chain.is_cannonical"
// }

// type EleNewBlockChainNotCanonical struct {
// 	BlockNumber     int    `json:"block_number"`
// 	HeadHash        string `json:"head_hash"`
// 	BlockHash       string `json:"block_hash"`
// 	BlockDifficulty int    `json:"block_difficulty"`
// 	BlockPrevHash   string `json:"block_prev_hash"`
// 	LogEvent
// }

// func (l *EleNewBlockChainNotCanonical) EventName() string {
// 	return "ele.newblock.chain.not_cannonical"
// }

// type EleTxCreated struct {
// 	TxHash    string `json:"tx_hash"`
// 	TxSender  string `json:"tx_sender"`
// 	TxAddress string `json:"tx_address"`
// 	TxHexRLP  string `json:"tx_hexrlp"`
// 	TxNonce   int    `json:"tx_nonce"`
// 	LogEvent
// }

// func (l *EleTxCreated) EventName() string {
// 	return "ele.tx.created"
// }

// type EleTxBroadcasted struct {
// 	TxHash    string `json:"tx_hash"`
// 	TxSender  string `json:"tx_sender"`
// 	TxAddress string `json:"tx_address"`
// 	TxNonce   int    `json:"tx_nonce"`
// 	LogEvent
// }

// func (l *EleTxBroadcasted) EventName() string {
// 	return "ele.tx.broadcasted"
// }

// type EleTxValidated struct {
// 	TxHash    string `json:"tx_hash"`
// 	TxSender  string `json:"tx_sender"`
// 	TxAddress string `json:"tx_address"`
// 	TxNonce   int    `json:"tx_nonce"`
// 	LogEvent
// }

// func (l *EleTxValidated) EventName() string {
// 	return "ele.tx.validated"
// }

// type EleTxIsInvalid struct {
// 	TxHash    string `json:"tx_hash"`
// 	TxSender  string `json:"tx_sender"`
// 	TxAddress string `json:"tx_address"`
// 	Reason    string `json:"reason"`
// 	TxNonce   int    `json:"tx_nonce"`
// 	LogEvent
// }

// func (l *EleTxIsInvalid) EventName() string {
// 	return "ele.tx.is_invalid"
// }
