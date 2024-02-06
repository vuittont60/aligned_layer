# Aligned Layer

<b> To be used in testnet only. </b>

Basic repo demoing a Stark/Snark verifier AVS middleware with full eigenlayer integration. 

## The Project 

Aligned Layer works with Eigen Layer to leverage ethereum consensus mechanism for ZK proof verification. Working outside the EVM, this allows for cheap verification of any proving system. This enables the usage of cutting edge algorithms, that may use new techniques to prove even faster. Even more, proving systems that reduces the proving overhead and adds verifier overhead, now become economically feasable to verify thanks to Aligned Layer. 

Full documentation and examples will be added soon

## Dependencies

You will need [foundry](https://book.getfoundry.sh/getting-started/installation) and [zap-pretty](https://github.com/maoueh/zap-pretty) to run the examples below.
```
curl -L https://foundry.paradigm.xyz | bash
foundryup
go install github.com/maoueh/zap-pretty@latest
```

## Running via make

Start anvil in a separate terminal:

```bash
make start-anvil-chain-with-el-and-avs-deployed
```

The above command starts a local anvil chain from a [saved state](./tests/integration/eigenlayer-and-shared-avs-contracts-deployed-anvil-state.json) with eigenlayer and Aligned Layer contracts already deployed (but no operator registered).

Start the aggregator:

```bash
make start-aggregator
```

Register the operator with eigenlayer and incredible-squaring, and then start the process:

```bash
make cli-setup-operator
make start-operator
```

## Workflow

To re compile contracts in case of changes use:

```bash
make deploy-all-to-anvil-and-save-state
```

To re generate the bindings in go:

```bash
make bindings
```

## Acknowledgment

Layr-Labs for creating [Incredible Squaring AVS](https://github.com/Layr-Labs/incredible-squaring-avs), which was used to bootstrap the initial version of Aligned Layer
