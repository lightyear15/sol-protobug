// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.2;
pragma experimental ABIEncoderV2;
import "./test.sol";


contract SolTest{
    function decodeMainMessag(bytes memory message) public pure returns(TestMainMessage.Data memory) {
        return TestMainMessage.decode(message);
    }
}
