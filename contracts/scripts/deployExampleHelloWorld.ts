import { ethers } from "hardhat"
import { ExampleHelloWorld } from "typechain-types"

const main = async (): Promise<any> => {
  const contract: ExampleHelloWorld = await ethers.deployContract("ExampleHelloWorld")
  await contract.waitForDeployment()
  console.log(`Contract deployed to: ${contract.target}`)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
