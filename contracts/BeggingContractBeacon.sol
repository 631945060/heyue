// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract BeggingContractBeacon is Initializable, OwnableUpgradeable {
    // 存储每个地址的捐赠金额（以代币数量为单位）
    mapping(address => uint256) public donations;

    // 合约所有者和 ERC20 代币地址
    IERC20 public token; // ERC20 代币合约地址

    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
    event Donation(address indexed donor, uint256 amount);

    function initialize(address _tokenAddress) public initializer {
        __Ownable_init(msg.sender);
        token = IERC20(_tokenAddress); // 初始化时设置 ERC-20 代币地址
    }

    // 设置新的 ERC-20 代币地址（仅限 owner）
    function setTokenAddress(address _tokenAddress) external onlyOwner {
        token = IERC20(_tokenAddress);
    }

    // 用户进行 ERC-20 捐赠
    function donate(uint256 amount) public {
        require(amount > 0, "Donation amount must be greater than zero");

        // 从用户钱包中转账到合约
        require(
            token.transferFrom(msg.sender, address(this), amount),
            "Transfer failed"
        );

        // 更新捐款记录
        donations[msg.sender] += amount;

        emit Donation(msg.sender, amount);
    }

    // 合约所有者提取所有代币
    function withdraw(uint256 amount) public onlyOwner {
        require(amount > 0, "Amount must be greater than zero");
        require(
            token.balanceOf(address(this)) >= amount,
            "Insufficient contract balance"
        );

        // 将指定数量的代币转给 owner
        require(token.transfer(owner(), amount), "Withdrawal failed");
    }

    // 查询某个地址的捐款记录
    function getDonation(address donor) public view returns (uint256) {
        return donations[donor];
    }

    // 获取合约当前持有的代币余额
    function getContractTokenBalance() public view returns (uint256) {
        return token.balanceOf(address(this));
    }
}
