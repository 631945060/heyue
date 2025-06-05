const { ethers, upgrades } = require("hardhat")
const fs = require("fs")
const path = require("path")

module.exports = async function ({ getNamedAccounts, deployments }) {
  const { save } = deployments
  const { deployer } = await getNamedAccounts()
  console.log("部署用户地址：", deployer)

  // 读取 .cache/proxyBegging.json文件
  const storePath = path.resolve(__dirname, "./.cache/proxyBegging.json");
  const storeData = fs.readFileSync(storePath, "utf-8");
  const { proxyAddress, implAddress, abi } = JSON.parse(storeData);

  // 升级版的业务合约
  const BeggingContractV3 = await ethers.getContractFactory("BeggingContractV3")

  // 升级代理合约
  const beggingProxyV3 = await upgrades.upgradeProxy(proxyAddress, BeggingContractV3)
  await beggingProxyV3.waitForDeployment()
  const proxyAddressV3 = await beggingProxyV3.getAddress()

  //   // 保存代理合约地址
  //   fs.writeFileSync(
  //     storePath,
  //     JSON.stringify({
  //       proxyAddress: proxyAddressV2,
  //       implAddress,
  //       abi,
  //     })
  //   );

  await save("BeggingContractV3", {
    abi,
    address: proxyAddressV3,
  })
}


module.exports.tags = ["upgradeBegging"]