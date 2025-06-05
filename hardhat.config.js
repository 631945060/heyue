require("@nomicfoundation/hardhat-toolbox");
require('hardhat-deploy');
require('@openzeppelin/hardhat-upgrades');

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.28",
    networks: {
    // localhost: {
    //   url: "http://127.0.0.1:8545",
    //   accounts: ["0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"]
    // },
    sepolia: {
      url: "https://sepolia.infura.io/v3/45dc4d4d5e1b4cf992f3532b186fc58f",
      accounts: ["e642666a050baebccd0f108edf718a79038c895164b6c40fa732cdf367369159"]
    }
  },
  namedAccounts:{
    deployer:0,
    user1:1,
    user2:2,
  }
};
