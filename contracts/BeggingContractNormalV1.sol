// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
// import "https://github.com/smartcontractkit/chainlink/blob/v0.219.1/src/v0.8/interfaces/AggregatorV3Interface.sol";

contract BeggingContractNormalV1 is Initializable, UUPSUpgradeable,OwnableUpgradeable {
    //记录每个捐赠者的捐赠金额
    mapping(address => uint256) public donations;
    uint256 public value;

    function initialize(address initialOwner) public initializer {
        __Ownable_init(initialOwner);
    }
    // AggregatorV3Interface public priceFeed; // ETH/USD 预言机接口

    mapping(address => AggregatorV3Interface) public priceFeeds;
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );


    // 设置新的预言机地址（仅限 owner）
    function setPriceFeed(
        address tokenAddress,
        address _priceFeedAddress
    ) public {
        require(msg.sender == owner(), "Only owner can set price feed");
        priceFeeds[tokenAddress] = AggregatorV3Interface(_priceFeedAddress);
    }
    function getChainlinkDataFeedLatestAnswer(
        address tokenAddress
    ) public view returns (int) {
        AggregatorV3Interface priceFeed = priceFeeds[tokenAddress];
        require(address(priceFeed) != address(0), "Price feed not set");
        // prettier-ignore
        (
            /* uint80 roundId */,
            int256 answer,
            /*uint256 startedAt*/,
            /*uint256 updatedAt*/,
            /*uint80 answeredInRound*/
        ) = priceFeed.latestRoundData();
        return answer;
    }

    //允许用户向合约发送以太币，并记录捐赠信息。
    function donate() public payable // address _tokenAddress

    {
        require(msg.value > 0, "Donation amount must be greater than 0.");
        donations[msg.sender] += msg.value;
    }

    // 允许合约所有者提取所有资金
    function withdraw() public {
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
    function setValue(uint256 _value) public virtual {
        require(msg.sender == owner(), "Only owner can set price feed");
        value = _value;
    }

    function getValue() public view virtual returns (uint256) {
        return value;
    }
    function _authorizeUpgrade(
        address newImplementation
    ) internal view override onlyOwner {
        // 只有管理员可以升级合约
        require(msg.sender == owner(), "Only admin can upgrade");
    }
    // function _authorizeUpgrade(
    //     address newImplementation
    // ) internal view override {
    //     // 只有管理员可以升级合约
    //     require(msg.sender == owner, "Only admin can upgrade");
    // }
}
