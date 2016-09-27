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

package core

import (
	"errors"
	"math/big"

	"github.com/elementrem/go-elementrem/core/vm"
)

var ChainConfigNotFoundErr = errors.New("ChainConfig not found") // general config not found error

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	HomesteadBlock *big.Int `json:"homesteadBlock"` // Homestead switch block (nil = no fork, 0 = already homestead)
	INTERSTELLARleapBlock   *big.Int `json:"interstellarLeapBlock"`   // TheINTERSTELLAR hyperz-leap switch block (nil = no fork)
	INTERSTELLARleapSupport bool     `json:"interstellarLeapSupport"` // Whether the nodes supports or opposes the INTERSTELLAR hyperz-leap

	VmConfig vm.Config `json:"-"`
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	if c.HomesteadBlock == nil || num == nil {
		return false
	}
	return num.Cmp(c.HomesteadBlock) >= 0
}
