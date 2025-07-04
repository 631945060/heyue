const { deployments, upgrades, ethers } = require("hardhat");
const fs = require("fs");
const path = require("path");
module.exports = async ({ getNamedAccounts, deployments }) => {
    const { save } = deployments;
    const { deployer } = await getNamedAccounts();
    console.log("部署用户地址:", deployer)
    const BeggingContract = await ethers.getContractFactory('BeggingContractV1');
    const beggingContractProxy = await upgrades.deployProxy(BeggingContract, [
        // "0x00000000000000000000000000000000",
        // 1000 * 100,
        // ethers.parseEther("0.000000000001"),
        // ethers.ZeroAddress,
        // 1
    ], {
        initializer: "initialize",
    })
    await beggingContractProxy.waitForDeployment();
    const proxyAddress = await beggingContractProxy.getAddress()
    console.log("代理合约地址", proxyAddress);
    const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress)
    console.log("实现合约地址", implAddress);
    const storePath = path.resolve(__dirname, "./.cache/proxyBegging.json");

    fs.writeFileSync(
        storePath,
        JSON.stringify({
            proxyAddress,
            implAddress,
            abi: BeggingContract.interface.format("json"),
        })
    );

    await save("BeggingProxy", {
        abi: BeggingContract.interface.format("json"),
        address: proxyAddress,
    })
};
module.exports.tags = ['deployBegging'];