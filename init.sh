#!/bin/sh

VALIDATOR_NAME=validator1
CHAIN_ID=rpsd
KEY_NAME=rpsd-key
CHAINFLAG="--chain-id ${CHAIN_ID}"
TOKEN_AMOUNT="10000000000000000000000000stake"
STAKING_AMOUNT="1000000000stake"

NAMESPACE_ID=$(openssl rand -hex 8)
echo $NAMESPACE_ID
DA_BLOCK_HEIGHT=$(curl https://rpc-blockspacerace.pops.one/block | jq -r '.result.block.header.height')
echo $DA_BLOCK_HEIGHT

ignite chain build
rpsd tendermint unsafe-reset-all
rpsd init $VALIDATOR_NAME --chain-id $CHAIN_ID

rpsd keys add $KEY_NAME --keyring-backend test
rpsd keys add alice --keyring-backend test
rpsd keys add bob --keyring-backend test
rpsd add-genesis-account $KEY_NAME $TOKEN_AMOUNT --keyring-backend test
rpsd add-genesis-account alice $TOKEN_AMOUNT --keyring-backend test
rpsd add-genesis-account bob $TOKEN_AMOUNT --keyring-backend test
rpsd gentx $KEY_NAME $STAKING_AMOUNT --chain-id $CHAIN_ID --keyring-backend test
rpsd collect-gentxs
rpsd start --rollkit.aggregator true --rollkit.block_time 2s --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT --api.enable --api.enabled-unsafe-cors
# rpsd start --rollkit.aggregator false --rollkit.da_layer celestia --rollkit.da_config='{"base_url":"http://localhost:26659","timeout":60000000000,"fee":6000,"gas_limit":6000000}' --rollkit.namespace_id $NAMESPACE_ID --rollkit.da_start_height $DA_BLOCK_HEIGHT --api.enable --api.enabled-unsafe-cors

# a25255eb9d95d989
# 336047