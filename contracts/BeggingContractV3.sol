// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

contract BeggingContractV3 is Initializable {
    //记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;
    address public owner;
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    function initialize() initializer public {
        owner = msg.sender;
    }

    //允许用户向合约发送以太币，并记录捐赠信息。
    function donate() public payable {
        require(msg.value > 10, "Donation amount must be greater than 10.");
        donations[msg.sender] += msg.value;
    }

    // 允许合约所有者提取所有资金
    function withdraw() public {
        require(
            msg.sender == owner,
            "Only the contract owner can withdraw funds."
        );
        uint256 contractBalance = address(this).balance;
        require(contractBalance > 0, "No funds to withdraw.");
        payable(owner).transfer(contractBalance);
    }

    //允许查询某个地址的捐赠金额
    function getDonation(address donor) public view returns (uint256) {
        return donations[donor];
    }
}
