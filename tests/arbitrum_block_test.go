// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// TestExecutionSpecArbitrumBlocktests runs the test fixtures from execution-spec-tests.
func TestExecutionSpecArbitrumBlocktests(t *testing.T) {
	if !common.FileExist(executionSpecBlockchainTestDir) {
		t.Skipf("directory %s does not exist", executionSpecBlockchainTestDir)
	}
	bt := new(testMatcher)

	// These tests fail as of https://github.com/ethereum/go-ethereum/pull/28666, since we
	// no longer delete "leftover storage" when deploying a contract.
	bt.skipLoad(`^cancun/eip6780_selfdestruct/selfdestruct/self_destructing_initcode_create_tx.json`)
	bt.skipLoad(`^cancun/eip6780_selfdestruct/selfdestruct/self_destructing_initcode.json`)
	bt.skipLoad(`^cancun/eip6780_selfdestruct/selfdestruct/recreate_self_destructed_contract_different_txs.json`)
	bt.skipLoad(`^cancun/eip6780_selfdestruct/selfdestruct/delegatecall_from_new_contract_to_pre_existing_contract.json`)
	bt.skipLoad(`^cancun/eip6780_selfdestruct/selfdestruct/create_selfdestruct_same_tx.json`)

	bt.walk(t, executionSpecBlockchainTestDir, func(t *testing.T, name string, test *BlockTest) {
		execBlockTest(t, bt, test, true)
	})
}
