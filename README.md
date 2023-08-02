# Temporal Zone
**Temporal Zone** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Custom Modules

`compound` Compounds the native staking token.

`record` Keeps track of when someone staked and how much.

## How to locally run the chain

Install gcc & make

Install [go](https://go.dev/doc/install)

Optional Installs:

docker - if you plan on adding or updating protos

rustc & hermes - if you want to run two local networks

On first run you might need to add the following go path to your $HOME/.profile:
```
export PATH=$PATH:/usr/local/go/bin:~/go/bin/
```

## Using Make

Build the binary:
```
make install
```

Start the chain:
```
temporald start
```

## Using Ignite

Install [Ignite](https://docs.ignite.com/welcome/install)

From the root of the temporal directory run:
```
ignite chain serve -v
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

Various parameters can be configured with `config.yml`. To learn more, see the [Ignite Config docs](https://docs.ignite.com/references/config).

## Attribution

Thank you to the following projects for inspiration, code and helping us get started:

Osmosis

Stride

Mars Hub
