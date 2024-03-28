import { task } from "hardhat/config"
import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers"
import { BigNumber } from "ethers"

const HELLO_WORLD_ADDRESS = "0x0300000000000000000000000000000000000000"

const ROLES = {
  0: "None",
  1: "Enabled",
  2: "Admin",
}

const getRole = async (allowList, address) => {
  const role = await allowList.readAllowList(address)
  console.log(`${address} has role: ${ROLES[role.toNumber()]}`)
}

// npx hardhat accounts --network local
task("accounts", "Prints the list of accounts", async (args, hre): Promise<void> => {
  const accounts: SignerWithAddress[] = await hre.ethers.getSigners()
  accounts.forEach((account: SignerWithAddress): void => {
    console.log(account.address)
  })
})

// npx hardhat balances --network local
task("balances", "Prints the list of account balances", async (args, hre): Promise<void> => {
  const accounts: SignerWithAddress[] = await hre.ethers.getSigners()
  for (const account of accounts) {
    const balance: BigNumber = await hre.ethers.provider.getBalance(
      account.address
    )
    console.log(`${account.address} has balance ${balance.toString()}`)
  }
})

// npx hardhat balance --network local --address [address]
task("balance", "get the balance")
  .addParam("address", "the address you want to know balance of")
  .setAction(async (args, hre) => {
    const balance = await hre.ethers.provider.getBalance(args.address)
    const balanceInCoin = hre.ethers.utils.formatEther(balance)
    console.log(`balance: ${balanceInCoin} Coin`)
  })

// npx hardhat helloWorld:readRole --network local --address [address]
task("helloWorld:readRole", "Gets the network enabled allow list")
  .addParam("address", "the address you want to know the allowlist role for")
  .setAction(async (args, hre) => {
    const allowList = await hre.ethers.getContractAt("IHelloWorld", HELLO_WORLD_ADDRESS)
    await getRole(allowList, args.address)
  })

// npx hardhat helloWorld:addEnabled --network local --address [address]
task("helloWorld:addEnabled", "Adds the enabled on the allow list")
  .addParam("address", "the address you want to add as a enabled")
  .setAction(async (args, hre) => {
    const allowList = await hre.ethers.getContractAt("IHelloWorld", HELLO_WORLD_ADDRESS)
    // ADD CODE BELOW
    await allowList.setEnabled(args.address)
    await getRole(allowList, args.address)
  })

// npx hardhat helloWorld:addAdmin --network local --address [address]
task("helloWorld:addAdmin", "Adds an admin on the allowlist")
  .addParam("address", "the address you want to add as a admin")
  .setAction(async (args, hre) => {
    const allowList = await hre.ethers.getContractAt("IHelloWorld", HELLO_WORLD_ADDRESS)
    await allowList.setAdmin(args.address)
    await getRole(allowList, args.address)
  })

// npx hardhat helloWorld:sayHello --network local
task("helloWorld:sayHello", "Says hello")
  .setAction(async (args, hre) => {
    const helloWorld = await hre.ethers.getContractAt("IHelloWorld", HELLO_WORLD_ADDRESS)
    const result = await helloWorld.sayHello()
    console.log(result)
  })

// npx hardhat helloWorld:setGreeting --network local --greeting [greeting]
task("helloWorld:setGreeting", "Says hello")
  .addParam("greeting", "the greeting string you want to set")
  .setAction(async (args, hre) => {
    const helloWorld = await hre.ethers.getContractAt("IHelloWorld", HELLO_WORLD_ADDRESS)
    const result = await helloWorld.setGreeting(args.greeting)
    console.log(result)
  })



