const { ethers, upgrades } = require("hardhat");

// bounty factory也是uups，基本和bountyFactory一样
async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contracts with the account:", deployer.address);

    // 部署逻辑合约
    const Bounty = await ethers.getContractFactory("BeggingContractV1");
    const BountyDeploy = await Bounty.deploy();
    await BountyDeploy.waitForDeployment()
    addresswait=await BountyDeploy.getAddress();
    console.log("bounty implementation: ", addresswait);

}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    }); 