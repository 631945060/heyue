const { ethers, deployments, upgrades } = require("hardhat")
const { expect } = require("chai");
describe("upgrde", async function () {
  it("should be able to deploy", async function () {
    // 获取账户
    let [owner, donor] = await ethers.getSigners();
    //1:部署业务合约
    await deployments.fixture("deployBegging")
    const beggingProxy = await deployments.get("BeggingProxy");
    const BeggingContractV2 = await ethers.getContractAt(
      "BeggingContractV2",
      beggingProxy.address
    );
    const donationAmount = ethers.parseEther("1.0"); // 1 ETH
    // 捐赠
    await BeggingContractV2.connect(donor).donate({ value: donationAmount });
    const donations1 = await BeggingContractV2.donations(beggingProxy.address);
    console.log("捐赠成功：：", donations1);
    const implAddress1 = await upgrades.erc1967.getImplementationAddress(
      beggingProxy.address
    );

    // 3. 升级合约
    await deployments.fixture(["upgradeBegging"]);
    const implAddress2 = await upgrades.erc1967.getImplementationAddress(
      beggingProxy.address
    );

    const BeggingContractV3 = await ethers.getContractAt(
      "BeggingContractV3",
      beggingProxy.address
    );
    console.log("implAddress1::", implAddress1, "\nimplAddress2::", implAddress2);

    // 捐赠
    await BeggingContractV3.connect(donor).donate({ value: donationAmount });
    const donations2 = await BeggingContractV3.donations(beggingProxy.address);
    console.log("升级后捐赠成功：：", donations2);

  })

})