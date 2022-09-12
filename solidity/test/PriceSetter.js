const { ethers } = require("hardhat");
const { expect } = require("chai");

describe("PriceSetter", function () {

  let PriceSetter;

  before(async () => {
    const PriceSetterContract = await ethers.getContractFactory("PriceSetter");
    PriceSetter = await PriceSetterContract.deploy();
  });

  it("Set first price for ETH symbol", async function () {

    const newPrice = ethers.utils.parseEther("1");

    const setPriceTx = await PriceSetter.set("ETH", newPrice);
    await setPriceTx.wait();

    const currentPrice = await PriceSetter.priceOf("ETH");

    expect(newPrice.toString()).to.equal(currentPrice.toString());
  });

  it("Set new price for ETH symbol", async function () {

    const newPrice = ethers.utils.parseEther("1.1");

    const setPriceTx = await PriceSetter.set("ETH", newPrice);
    await setPriceTx.wait();

    const currentPrice = await PriceSetter.priceOf("ETH");

    expect(newPrice.toString()).to.equal(currentPrice.toString());
  });

  it("Try to set lower than 2% price difference", async function () {

    const setupPriceTx = await PriceSetter.set("ETH", ethers.utils.parseEther("1"));
    await setupPriceTx.wait();

    const newPrice = ethers.utils.parseEther("1.02");

    await expect(PriceSetter.set("ETH", newPrice)).to.be.revertedWith("PriceSetter__InsufficientPriceDifference()");
  });

  it("Try to set with empty symbol", async function () {
    await expect(PriceSetter.set("", ethers.utils.parseEther("1"))).to.be.revertedWith("PriceSetter__SymbolCantBeEmpty()");
  });
});
