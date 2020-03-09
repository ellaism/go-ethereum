// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.
package params

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// Genesis hashes to enforce below configs on.
	MordorGenesisHash = common.HexToHash("0xa68ebde7932eccb177d38d55dcc6461a019dd795a681e59b5a3e4f3a7259a3f1")

	// MordorChainConfig is the chain parameters to run a node on the Ethereum Classic Mordor test network (PoW).
	MordorChainConfig = &ChainConfig{
		ChainID:             big.NewInt(63),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(301243),
		PetersburgBlock:     big.NewInt(301243),
		IstanbulBlock:       big.NewInt(999983),
		DisposalBlock:       big.NewInt(0),
		ECIP1017EraRounds:   big.NewInt(2000000),
		EIP160Block:         big.NewInt(0),
		ECIP1010PauseBlock:  big.NewInt(0),
		ECIP1010Length:      big.NewInt(2000000),
		Ethash:              new(EthashConfig),
	}
)
