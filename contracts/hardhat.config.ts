import "@nomicfoundation/hardhat-toolbox";
import "./tasks"

// HardHat users must populate these environment variables in order to connect to their subnet-evm instance
// Since the blockchainID is not known in advance, there's no good default to use and we use the C-Chain here.
var local_rpc_uri = process.env.RPC_URI || "http://127.0.0.1:9650/ext/bc/C/rpc"
var local_chain_id = parseInt(process.env.CHAIN_ID, 10) || 99999

export default {
  solidity: {
    compilers: [
      {
        version: "0.8.24",
        settings: {
          evmVersion: "shanghai",
        },
      },
    ]
  },
  networks: {
    local: {
      //"http://{ip}:{port}/ext/bc/{chainID}/rpc
      // expected to be populated by the environment variables above
      url: local_rpc_uri,
      chainId: local_chain_id,
      // accounts contains private keys of following addresses(in order):
      // 0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC
      // 0x9632a79656af553F58738B0FB750320158495942
      // 0x55ee05dF718f1a5C1441e76190EB1a19eE2C9430
      // 0x4Cf2eD3665F6bFA95cE6A11CFDb7A2EF5FC1C7E4
      // 0x0B891dB1901D4875056896f28B6665083935C7A8
      // 0x01F253bE2EBF0bd64649FA468bF7b95ca933BDe2
      // 0x78A23300E04FB5d5D2820E23cc679738982e1fd5
      // 0x3C7daE394BBf8e9EE1359ad14C1C47003bD06293
      // 0x61e0B3CD93F36847Abbd5d40d6F00a8eC6f3cfFB
      // 0x0Fa8EA536Be85F32724D57A37758761B86416123
      accounts: [
        "0x56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027",
        "0x7b4198529994b0dc604278c99d153cfd069d594753d471171a1d102a10438e07",
        "0x15614556be13730e9e8d6eacc1603143e7b96987429df8726384c2ec4502ef6e",
        "0x31b571bf6894a248831ff937bb49f7754509fe93bbd2517c9c73c4144c0e97dc",
        "0x6934bef917e01692b789da754a0eae31a8536eb465e7bff752ea291dad88c675",
        "0xe700bdbdbc279b808b1ec45f8c2370e4616d3a02c336e68d85d4668e08f53cff",
        "0xbbc2865b76ba28016bc2255c7504d000e046ae01934b04c694592a6276988630",
        "0xcdbfd34f687ced8c6968854f8a99ae47712c4f4183b78dcc4a903d1bfe8cbf60",
        "0x86f78c5416151fe3546dece84fda4b4b1e36089f2dbc48496faf3a950f16157c",
        "0x750839e9dbbd2a0910efe40f50b2f3b2f2f59f5580bb4b83bd8c1201cf9a010a"
      ],
      pollingInterval: "1s"
    },
  },
  mocha: {
    timeout: 30000
  }
}
