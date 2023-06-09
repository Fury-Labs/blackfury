<!--
parent:
  order: false
-->

<div align="center">
  <h1> Gridiron </h1>
</div>

<div align="center">
  <a href="https://github.com/fury-labs/blackfury/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/tharsis/blackfury.svg" />
  </a>
  <a href="https://github.com/fury-labs/blackfury/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/tharsis/blackfury.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/fury-labs/blackfury">
    <img alt="GoDoc" src="https://godoc.org/github.com/fury-labs/blackfury?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/fury-labs/blackfury">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/fury-labs/blackfury"/>
  </a>
  <a href="https://bestpractices.coreinfrastructure.org/projects/5018">
    <img alt="Lines of code" src="https://img.shields.io/tokei/lines/github/tharsis/blackfury">
  </a>
</div>
<div align="center">
  <a href="https://discord.gg/blackfury">
    <img alt="Discord" src="https://img.shields.io/discord/809048090249134080.svg" />
  </a>
  <a href="https://github.com/fury-labs/blackfury/actions?query=branch%3Amain+workflow%3ALint">
    <img alt="Lint Status" src="https://github.com/fury-labs/blackfury/actions/workflows/lint.yml/badge.svg?branch=main" />
  </a>
  <a href="https://codecov.io/gh/fury-labs/blackfury">
    <img alt="Code Coverage" src="https://codecov.io/gh/fury-labs/blackfury/branch/main/graph/badge.svg" />
  </a>
  <a href="https://twitter.com/GridironOrg">
    <img alt="Twitter Follow Gridiron" src="https://img.shields.io/twitter/follow/GridironOrg"/>
  </a>
</div>

Gridiron is a scalable, high-throughput Proof-of-Stake blockchain
that is fully compatible and interoperable with Ethereum.
It's built using the [Cosmos SDK](https://github.com/cosmos/cosmos-sdk/)
which runs on top of the [Tendermint Core](https://github.com/tendermint/tendermint) consensus engine.

## Quick Start

To learn how Gridiron works from a high-level perspective,
go to the [Protocol Overview](https://docs.blackfury.org/protocol) section of the documentation.
You can also check the instructions to [Run a Node](https://docs.blackfury.org/protocol/blackfury-cli#run-an-blackfury-node).

## Documentation

Our documentation is hosted in a [separate repository](https://github.com/blackfury/docs) and can be found at [docs.blackfury.org](https://docs.blackfury.org).
Head over there and check it out.

## Installation

For prerequisites and detailed build instructions
please read the [Installation](https://docs.blackfury.org/protocol/blackfury-cli) instructions.
Once the dependencies are installed, run:

```bash
make install
```

Or check out the latest [release](https://github.com/fury-labs/blackfury/releases).

## Community

The following chat channels and forums are great spots to ask questions about Gridiron:

- [Gridiron Twitter](https://twitter.com/GridironOrg)
- [Gridiron Discord](https://discord.gg/blackfury)
- [Gridiron Forum](https://commonwealth.im/blackfury)

## Contributing

Looking for a good place to start contributing?
Check out some
[`good first issues`](https://github.com/fury-labs/blackfury/issues?q=is%3Aopen+is%3Aissue+label%3A%22good+first+issue%22).

For additional instructions, standards and style guides, please refer to the [Contributing](./CONTRIBUTING.md) document.

## Careers

See our open positions on [Greenhouse](https://boards.eu.greenhouse.io/blackfury).

## Licensing

Starting from April 21st, 2023, the Gridiron repository will update its License
from GNU Lesser General Public License v3.0 (LGPLv3) to Gridiron Non-Commercial
License 1.0 (ENCL-1.0). This license applies to all software released from Gridiron
version 13 or later, except for specific files, as follows, which will continue
to be licensed under LGPLv3:

- `x/revenue/v1/` (all files in this folder)
- `x/claims/genesis.go`
- `x/erc20/keeper/proposals.go`
- `x/erc20/types/utils.go`

LGPLv3 will continue to apply to older versions (<v13.0.0) of the Gridiron
repository. For more information see LICENSE.

### SPDX Identifier

The following header including a license identifier in SPDX short form has been added to all ENCL-1.0 files:

```go
// Copyright Tharsis Labs Ltd.(Gridiron)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/fury-labs/blackfury/blob/main/LICENSE)
```

Exempted files contain the following SPDX ID:

```go
// Copyright Tharsis Labs Ltd.(Gridiron)
// SPDX-License-Identifier:LGPL-3.0-only
```

### License FAQ

Find below an overview of the Permissions and Limitations of the Gridiron Non-Commercial License 1.0.
For more information, check out the full ENCL-1.0 FAQ [here](/LICENSE_FAQ.md).

| Permissions                                                                                                                                                                  | Prohibited                                                                 |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| - Private Use, including distribution and modification<br />- Commercial use on designated blockchains<br />- Commercial use with Gridiron permit (to be separately negotiated) | - Commercial use, other than on designated blockchains, without Gridiron permit |
