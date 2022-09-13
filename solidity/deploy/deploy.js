// pull contract deployments from hardhat deploy
const { network } = require("hardhat")
const { verify } = require("../utils/verify")

module.exports = async ({ getNamedAccounts, deployments }) => {
    const { deploy, log } = deployments
    const { deployer } = await getNamedAccounts()

    log("----------------------------------------------------")
    log("Deploying contract and waiting for confirmations...")
    const PriceSetter = await deploy("PriceSetter", {
        from: deployer,
        args: [],
        log: true,
        waitConfirmations: network.config.blockConfirmations || 1,
    })
    log(`Contract deployed at ${PriceSetter.address}`)

    // verify contract on etherscan
    await verify(PriceSetter.address, [])

}