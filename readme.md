<br />
<p align="center">
  <img src="https://raw.githubusercontent.com/crypto-com/chain/master/assets/logo.svg" alt="Crypto.com Chain" width="400">
</p>
<br />

<p align="center">
  <a href="https://github.com/crypto-com/chain-main/workflows"><img label="Build Status" src="https://github.com/crypto-com/chain-main/workflows/Build/badge.svg" /></a>
  <a href="https://codecov.io/gh/crypto-com/chain-main"><img label="Code Coverage" src="https://codecov.io/gh/crypto-com/chain-main/branch/master/graph/badge.svg" /></a>
  <a href="https://gitter.im/crypto-com/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge"><img label="Gitter" src="https://badges.gitter.im/crypto-com/community.svg" /></a>
</p>


## Table of Contents

1. [Description](#description)
2. [Contributing](#contributing)
3. [License](#license)
4. [Documentation](#documentation)
5. [Build](#build)
   1. [Nix](#nix)
6. [Start a Local Full Node](#start-local-full-node)
7. [Send your First Transaction](#send-first-transaction)
8. [Testing](#testing)
9. [Useful Links](#useful-links)

<a id="description" />

## 1. Description

**Crypto.com Chain** is a blockchain application built using Cosmos SDK and Tendermint,
intended as a backbone for some of the existing and future Crypto.com ecosystem.

<a id="contributing" />

## 2. Contributing
Please abide by the [Code of Conduct](CODE_OF_CONDUCT.md) in all interactions,
and the [contributing guidelines](CONTRIBUTING.md) when submitting code.

<a id="license" />

## 3. License

[Apache 2.0](./LICENSE)

<a id="documentation" />

## 4. Documentation

Technical documentation can be found in this [Github repository](https://github.com/crypto-com/chain-docs) (you can read it in [this hosted version](https://chain.crypto.com/docs)).

<a id="build" />

## 5. Build full node

<a id="nix" />

#### 1. Nix
Nix is a (cross-language) package manager for reproducible builds.
On Linux and macOS, you can [install it as follows](https://nixos.org/download.html) (on Windows 10, you can possibly use the Windows Subsystem for Linux):

```
$ curl -L https://nixos.org/nix/install | sh
```

You can then run:

```
$ make chaindImage
```

Which will build a docker image that contains the full node binary.

Optionally, you can also use a binary cache to speed up the build process:

```
$ nix-env -iA cachix -f https://cachix.org/api/v1/install
$ cachix use crypto-com
```
<a id="start-local-full-node" />
## 6. Start a Local Full Node

TODO

<a id="send-first-transaction" />

## 7. Send Your First Transaction

TODO

<a id="testing" />

## 8. Testing

There are different tests that can be executed in the following ways:

- unit tests: `make test`
- simulations: `make test-sim-*` (e.g. `make test-sim-nondeterminism`)
- integrations tests: `make nix-integration-test` (see more details in [their documentation](integration_tests/README.md))

---

<a id="useful-links" />

## Useful links

- [Project Website](http://chain.crypto.com/)
- [Technical Documentation](http://chain.crypto.com/)
- Community chatrooms (non-technical): [Discord](https://discord.gg/nsp9JTC) [Telegram](https://t.me/CryptoComOfficial)
- Developer community chatroom (technical): [![Gitter](https://badges.gitter.im/crypto-com/community.svg)](https://gitter.im/crypto-com/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Discord](https://discord.gg/W8trcGV)
- [pystarport](./pystarport/README.md)
