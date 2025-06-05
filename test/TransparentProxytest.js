const { expect } = require("chai");
const { log } = require("console");
const { ethers } = require("hardhat");

describe("Transparent Proxy Upgrade Test", async function () {
    let deployer, user;
    let logicV1, logicV2, proxy;
    let proxyAsV1;
    beforeEach(async function () {
        [deployer, user] = await ethers.getSigners();
        
        // 1. 部署 V1 实现
        const LogicV1 = await ethers.getContractFactory("BeggingContractTransparentV1");
        logicV1 = await LogicV1.deploy();
        await logicV1.waitForDeployment();

        // 2. 部署 Transparent Proxy 指向 V1
        const Proxy = await ethers.getContractFactory("TransparentProxy");
        const LogicV1address = await logicV1.getAddress();
        proxy = await Proxy.deploy(LogicV1address,deployer.address);
        console.log("BeggingContractTransparentV1地址为:" + LogicV1address);
        console.log("deployer地址为:" + deployer.address);
        await proxy.waitForDeployment();

        // 3. 将 proxy 转换为 V1 接口
        proxyAsV1 = await ethers.getContractAt("BeggingContractTransparentV1", LogicV1address);
        // 可选：设置 owner 为 deployer（如果你在构造函数中设置了 owner）
        // 如果没有 initializer，可略过
    });


    it("Should allow donation and read from V1", async function () {
      await  proxyAsV1.connect(user).donate({ value: ethers.parseEther("1.0") });

        //查询捐款记录
        const donation = await proxyAsV1.getDonation(user.address);
        expect(donation).to.equal(ethers.parseEther("1"));
    });

    it("Should upgrade to V2 and retain state", async function () {
        // 1. 先进行一次捐赠
        await proxyAsV1.connect(user).donate({ value: ethers.parseEther("1") });
        const donationBeforeUpgrade = await proxyAsV1.getDonation(user.address);
        expect(donationBeforeUpgrade).to.equal(ethers.parseEther("1"));
        console.log("查询状态donationBeforeUpgrade:"+donationBeforeUpgrade);
        

        // 2. 部署 V2 实现
        const LogicV2 = await ethers.getContractFactory("BeggingContractTransparentV2");
        logicV2 = await LogicV2.deploy();
        await logicV2.waitForDeployment();
        const   logicV2Address= await logicV2.getAddress()
        // // 3. 通过 proxy 调用 upgradeTo
        const tx = await proxy.upgradeTo(logicV2Address);
        await tx.wait();

        // // 4. 将 proxy 转换为 V2 接口
        const proxyAsV2 = await ethers.getContractAt("BeggingContractTransparentV2", logicV2Address);

        // // 5. 查询状态是否保留
        const donationAfterUpgrade = await proxyAsV2.getDonation(user.address);
        console.log("查询状态是否保留donationAfterUpgrade:"+donationAfterUpgrade);


    });


});