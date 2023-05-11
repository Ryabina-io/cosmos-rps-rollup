# RPS Rollup
This repository is the simplest implementation of the Rock-Paper-Scissors game using Cosmos-SDK and Rollkit.
It is based on the official Rollkit [documentation](https://rollkit.dev/docs/tutorials/gm-world/). Familiarity with Rollkit may be necessary to build and run everything.
### Building binary

1. [Install go](https://go.dev/doc/install)
1. [Install Ignite cli](https://docs.ignite.com/welcome/install)
1. Build chain using `ignite chain build` in project folder. (TODO: add go building instructions without ignite)

### Connecting to running rollup
There is a version of rollup that we will aggregate for some time. However, there is no faucet available, so syncing with it would be purely for entertainment purposes. When our aggregation node fails to send a blob on a certain block, the syncing of your node will stuck. There is an issue with Rollkit that we are unable to fix on our own at the moment.
1. Setup blockspacerace DA node (light or full). 
    - `celestia light start --core.ip https://rpc-blockspacerace.pops.one --gateway --gateway.addr 127.0.0.1 --gateway.port 26659 --p2p.network blockspacerace`
    - Or use any public DA node.
1. Init chain
  `rps init <moniker>`
1. Swap genesis file in chain home directory (default: ~/.rpsd) <HOME>/.rps/config/genesis.json with [genesis.json](https://github.com/Ryabina-io/cosmos-rps-rollup/raw/master/genesis.json)
  `cp genesis.json ~/.rps/config/genesis.json`
1. Use provided parameters of rollup to sync for
`export DA_BLOCK_HEIGHT=455368`
`export NAMESPACE_ID=4e114dd721e93402`
1. Start rollup with command:
`rpsd start --rollkit.block_time 10s --rollkit.da_block_time 10s --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT`
On some block, on which our aggregation misses sending blob, syncing of you node will stop.
### Starting your own rollup
You can use local devnet DA:
`docker run --platform linux/amd64 -p 26650:26657 -p 26659:26659 ghcr.io/rollkit/local-celestia-devnet:v0.9.1`
Or any other network
[https://docs.celestia.org/nodes/quick-start/](https://docs.celestia.org/nodes/quick-start/)
1. Run any DA node with some balance (any network should be fine). Local devnet should be ok (`docker run --platform linux/amd64 -p 26650:26657 -p 26659:26659 ghcr.io/rollkit/local-celestia-devnet:v0.9.3`)
1. Run [init-local.sh](./init-local.sh) script 

### Playing the game
Two players initiate the game by first submitting their moves, each hashed with a unique salt. After both hashes have been submitted, each player then reveals their moves and salts. Once the last player has submitted their move, the winner is determined, and the rewards are sent to the corresponding player.
Currently, there is no time constraint for players to reveal their moves. Consequently, the second player could potentially never disclose their move. We may consider adding this functionality at a later date.

To create game start with command. You can choose any move (rock, scissors, paper).
```bash
$ rpsd tx rps create-game 1stake rock --from bob --keyring-backend test --chain-id=rps
```
In first string you will recieve generated salt for you
`Remember your Salt:  vcuOAlYgvH`
Remeber that - it will help you to finish game later.

Then you can view created games
```bash
$ rpsd q rps list-games --chain-id=rps 
games:
- betAmount:
    amount: "1"
    denom: stake
  index: "1"
  player1: rps1z8p3n5hap9vz7w6420plkpfzxsx9y8ulg5x7r7
  player2: ""
  turn1: ""
  turn2: ""
  turnHash1: ebfb74b700f5a504e673a228129a47b9e645607305f818cd80dc3277
  turnHash2: ""
pagination:
  next_key: null
  total: "0"
```

You can remove your game if no other players have joined it, by using the following command:
```bash
$ rpsd tx rps remove-game 1 --from bob -y
```

To join game use command
```bash
$ rpsd tx rps join-game 1 rock --from alice --keyring-backend test           
Remember your Salt:  BKwoYdzNKM
```

Next, both players should reveal their moves:
```bash
$ rpsd tx rps reveal-game 1 rock vcuOAlYgvH --from bob --keyring-backend test
$ rpsd tx rps reveal-game 1 rock BKwoYdzNKM --from alice --keyring-backend test
```

You can then verify the balances of your players using the `bank` module, and confirm that the game has disappeared `rpsd q rps list-games`