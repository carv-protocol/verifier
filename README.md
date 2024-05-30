# Verifier

Verifier nodes are lightweight nodes managed by the community. They play a crucial role in maintaining the protocol's integrity, thereby enhancing its trustworthiness and security. Currently, these nodes primarily check TEE attestations on-chain, verifying that the results are reliable and that the process maintains the privacy of user data.

For more information about verifier please read [verifier-node-explained](https://docs.carv.io/carv-protocol/verifier-node-explained)

## Building the source

Building verifier requires a Go (version 1.21 or later). You can install it using your favourite package manager. Then run

```shell
make build
```

After this, the verifier executable program has been compiled. If you just need to run the node, you can jump directly to [Run via binary program](#Run-via-binary-program).

If you want to become a contributor to this project, you can follow the following steps to compile verifier after completing the modifications.

To download dependencies during compilation, run

```shell
make init
```

These dependencies will be downloaded and compiled to `$GOPATH/bin`. This path needs to be added to the environment variable by running `export PATH=$PATH:$GOPATH/bin`. If it has already been added, ignore it.

Then, build the project by running

```shell
make all
```

## Running verifier

### Hardware Requirements

Minimum:

- CPU with 1+ cores
- 2GB RAM
- 4 MBit/sec download Internet service

Recommended:

- Fast CPU with 2+ cores
- 4GB+ RAM
- 8+ MBit/sec download Internet service

### Run via binary program

After executing `make build` or `make all`, the verifier executable file will be compiled into the `./bin` directory. You need to switch to the `./bin` directory before executing verifier. run
```shell
cd bin
```

If this is your first time running verifier, you need to specify a private key. The private key will sign the verification transaction, and the corresponding address will receive carv rewards.

The private key can be passed to the verifier through startup parameters, or written into the configuration file.

We have prepared a `config.yaml` file template in the `./configs` directory, and have written some basic default configurations into the file. If you need to use more customized configurations, you can refer to [Configuration](#Configuration)

#### through startup parameters

If you want to pass the private key through startup parameters, you need to set `wallet.mode` in the configuration file to `0`.

```shell
# Pass the private key in clear text
./verifier -conf ../configs/config.yaml -private-key <Your Private Key>

# By specifying keystore
./verifier -conf ../configs/config.yaml -keystore-path <Path to keystore file> -keystore-password <keystore's password>
```

In order to facilitate user operation, verifier provides a tool to generate a new keystore, run
```shell
./verifier -generate-keystore -keystore-path <path to generate your keystore file>
```

#### through configuration file

Configure the plain text private key: Set `wallet.mode` in the configuration file to `1`, and write the plain text private key into `wallet.private_key`. Then run

```shell
./verifier -conf ../configs/config.yaml
```

- Configure the path and password of the keystore: Set `wallet.mode` in the configuration file to `2`, and write the path and password of the keystore file into `wallet.keystore_path` and `wallet.keystore_password`. Then run

```shell
./verifier -conf ../configs/config.yaml
```

### Run via docker

One of the quickest ways to run verifier is by using Docker:

```shell
docker run -d --name verifier -v /<Path To This Repository>/verifier/configs:/data/conf carvprotocol/verifier
```

Note that you need to configure your private key in the configuration file. Similarly, you can configure both plaintext and keystore methods.

### Configuration

```yaml
chain:
  # The chain id of the chain that the verifier program interacts with
  chain_id: 5611
  # The chain name of the chain that the verifier program interacts with
  chain_name: "opBNB"
  # The rpc url of the blockchain node which verifier program interacts with
  rpc_url: "https://opbnb-testnet.nodereal.io/v1/9e210feafbec4ed9bd48f855c2bd979a"
  # The block height at which verifier starts synchronizing data
  start_block: 26623154
  # block height interval for synchronizing data at one time
  offset_block: 1
contract:
  # The contract address of carv on this chain
  addr: "0x39afd848cd4a83cddd06f340ae584c547e53873d"
  # The contract address of tee on this chain
  tee_addr : "0x19baa72643aa11b28cb6251fd7596201778ead9a"
  # The topic id of the event corresponding to the attestation published by tee
  topic: "0x99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591"
wallet:
  # wallet mode, by which way to pass the private key
  # 0: through startup parameters
  # 1: through plain text private key in config file
  # 2: through path and password of the keystore in config file
  mode: 0
  # plain text private key, needed when mode is 1
  private_key: "99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591"
  # path of the keystore, needed when mode is 2
  keystore_path: "/Users/alice/.keystore"
  # password of the keystore, needed when mode is 2
  keystore_password: "123456"
```

## Contribution
We welcome contributions of all forms! If you have any questions, suggestions, or have identified any bugs, please submit them through Issues. If you would like to contribute code, please submit it via Pull Requests.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.
