// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./BeggingContractNormalV1.sol";


contract BeggingContractNormalV2 is BeggingContractNormalV1 {
    uint256 public newValue; // 新增状态变量
     function setValue(uint256 _newValue) public virtual override onlyOwner {
        newValue = _newValue;
    }

    function getValue() public view virtual override returns (uint256) {
        return value + newValue;
    }

    // function _authorizeUpgrade(
    //     address newImplementation
    // ) internal view override {
    //     // 只有管理员可以升级合约
    //     require(msg.sender == owner, "Only admin can upgrade");
    // }
}
