import { task } from "hardhat/config"
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers"
import { BigNumber } from "ethers"

task("accounts", "Prints the list of accounts", async (args, hre): Promise<void> => {
  const accounts: SignerWithAddress[] = await hre.ethers.getSigners()
  accounts.forEach((account: SignerWithAddress): void => {
    console.log(account.address)
  })
})

task("balances", "Prints the list of account balances", async (args, hre): Promise<void> => {
  const accounts: SignerWithAddress[] = await hre.ethers.getSigners()
  for (const account of accounts) {
    const balance: BigNumber = await hre.ethers.provider.getBalance(
      account.address
    )
    console.log(`${account.address} has balance ${balance.toString()}`)
  }
})


task("balance", "get the balance")
  .addParam("address", "the address you want to know balance of")
  .setAction(async (args, hre) => {
    const balance = await hre.ethers.provider.getBalance(args.address)
    const balanceInCoin = hre.ethers.utils.formatEther(balance)
    console.log(`balance: ${balanceInCoin} Coin`)
  })



