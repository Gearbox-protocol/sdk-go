#!/bin/zsh


function downloadContracts {
    mkdir -p  $HOME/contracts
    cd $HOME/contracts
    #
    if [ ! -d gearbox-contracts ]; then
        git clone "git@github.com:Gearbox-protocol/gearbox-contracts.git"
    fi
    cd gearbox-contracts
    git pull
    yarn && yarn build
    cd ..
    # #
    if [ ! -d contracts-v2 ]; then
        git clone "git@github.com:Gearbox-protocol/contracts-v2.git"
    fi
    cd "contracts-v2"
    git pull
    yarn && yarn build
    cd ..
    #
}

downloadContracts