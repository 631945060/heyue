const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("BeggingContractV1", function () {
    let BeggingContract;
    let beggingContract;
    let owner;
    let donor;

    beforeEach(async function () {
        // 获取账户
        [owner, donor] = await ethers.getSigners();

        // 部署合约
        BeggingContract = await ethers.getContractFactory("BeggingContractV1");
        beggingContract = await BeggingContract.deploy();
        await beggingContract.waitForDeployment();
    });

    describe("Donate", function () {
        it("Should record donation correctly", async function () {
            const donationAmount = ethers.parseEther("1.0"); // 1 ETH

            // 捐赠
            await beggingContract.connect(donor).donate({ value: donationAmount });

            // 查询捐赠记录
            const recordedDonation = await beggingContract.getDonation(donor.address);

            // 验证是否正确记录
            expect(recordedDonation).to.equal(donationAmount);
        });

        it("Should revert if donation amount is zero", async function () {
            await expect(
                beggingContract.connect(donor).donate({ value: 0 })
            ).to.be.revertedWith("Donation amount must be greater than 0.");
        });
    });

    describe("Withdraw", function () {
        it("Should allow owner to withdraw funds", async function () {
            const donationAmount = ethers.parseEther("1.0");

            // 捐赠
            await beggingContract.connect(donor).donate({ value: donationAmount });

            // 提取资金前的 owner 余额
            const initialBalance = await ethers.provider.getBalance(owner)
            console.log("提取前的资金为:"+initialBalance);

            // 提取资金
            await beggingContract.connect(owner).withdraw();

            // 提取资金后的 owner 余额
            const finalBalance = await ethers.provider.getBalance(owner)
            console.log("提取后的资金为:"+finalBalance);
            
            // 验证余额增加
            expect(finalBalance).to.be.gt(initialBalance);
        });

        it("Should revert if non-owner tries to withdraw", async function () {
            await expect(
                beggingContract.connect(donor).withdraw()
            ).to.be.revertedWith("Only the contract owner can withdraw funds.");
        });

        it("Should revert if no funds to withdraw", async function () {
            // 直接尝试提取资金，此时合约中没有资金
            await expect(beggingContract.connect(owner).withdraw()).to.be.revertedWith(
                "No funds to withdraw."
            );
        });
    });
});