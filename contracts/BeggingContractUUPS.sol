// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
contract BeggingContractUUPS is
    Initializable,
    UUPSUpgradeable,
    OwnableUpgradeable
{
    //记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );

    function initialize() public initializer {
        __Ownable_init();
    }

    //允许用户向合约发送以太币，并记录捐赠信息。
    function donate() public payable {
        require(msg.value > 0, "Donation amount must be greater than 0.");
        donations[msg.sender] += msg.value;
    }

    // 允许合约所有者提取所有资金
    function withdraw() public onlyOwner {
        require(
            msg.sender == owner(),
            "Only the contract owner can withdraw funds."
        );
        uint256 contractBalance = address(this).balance;
        require(contractBalance > 0, "No funds to withdraw.");
        payable(owner()).transfer(contractBalance);
    }

    //允许查询某个地址的捐赠金额
    function getDonation(address donor) public view returns (uint256) {
        return donations[donor];
    }
    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        // 只有管理员可以升级合约
        require(msg.sender == owner(), "Only admin can upgrade");
    }
}
