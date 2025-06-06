// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.20;

import "./BeggingContractBeacon.sol";

contract Factory {
    address[] public allImpls;
    
    function createImpl() public {
        BeggingContractBeacon impl = new BeggingContractBeacon();
        allImpls.push(address(impl));
    }
}