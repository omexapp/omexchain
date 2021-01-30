#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-omexchaind.log}

##
## Run binary with all parameters
##
export OMEXCHAINDHOME="/omexchaind/node${ID}/omexchaind"

if [ -d "$(dirname "${OMEXCHAINDHOME}"/"${LOG}")" ]; then
  omexchaind --chain-id omexchain-1 --home "${OMEXCHAINDHOME}" "$@" | tee "${OMExCHAINDHOME}/${LOG}"
else
  omexchaind --chain-id omexchain-1 --home "${OMEXCHAINDHOME}" "$@"
fi

