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
./verifier -conf ../configs/config.yaml -private-key <Your Private Key> -reward-address <Your Reward Address> -commission-rate <Your Commission Rate>

# By specifying keystore
./verifier -conf ../configs/config.yaml -keystore-path <Path to keystore file> -keystore-password <keystore password> -reward-address <Your Reward Address> -commission-rate <Your Commission Rate>
```

In order to facilitate user operation, verifier provides a tool to generate a new keystore, run
```shell
./verifier -generate-keystore -keystore-path <path to generate your keystore file>
```
After running the command, you will be prompted to enter a password for the keystore. After entering the password, the keystore file will be generated in the specified path.
And then you can run the verifier with the keystore file path and password before you delegate to the keystore address.
> Before you run the verifier, you must delegate to the keystore address. After the delegation is successful, you can run the verifier again.
> Delegated address you can get from the terminal after generating the keystore.
> You can delegate your License by below url:
> -  [Delegate-testnet](https://testnet-explorer.carv.io/verifiers)
> - [Delegate-mainnet](https://explorer.carv.io/verifiers)


#### through configuration file

Configure the plain text private key: 
1. Set `wallet.mode` in the configuration file to `1`, 
2. Write the plain text private key into `wallet.private_key`, 
3. Write your reward address into `wallet.reward_claimer_addr`, 
4. Write your commission rate into `wallet.commission_rate`.
Then run

```shell
./verifier -conf ../configs/config.yaml 
```

Configure the path and password of the keystore: 
1. Set `wallet.mode` in the configuration file to `2`, 
2. Write the path and password of the keystore file into `wallet.keystore_path` 
3. Write the keystore password into `wallet.keystore_password`. 
4. Write your reward address into `wallet.reward_claimer_addr`,
5. Write your commission rate into `wallet.commission_rate`.
Then run

```shell
./verifier -conf ../configs/config.yaml
```

### Run via docker

One of the quickest ways to run verifier is by using Docker:

#### Use Private Key (mode = 0)

#### Update config_docker.yaml
Update your private_key and run :
```yaml
......
wallet:
  # wallet mode, by which way to pass the private key
  # 0: through startup parameters
  # 1: through plain text private key in config file
  # 2: through path and password of the keystore in config file
  mode: 0
  # plain text private key, needed when mode is 1
  private_key: "99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591"
  # path of the keystore, needed when mode is 2
  keystore_path: ""
  # password of the keystore, needed when mode is 2
  keystore_password: ""
......
```

```shell
docker run -d --name verifier -v /<Path To This Repository>/verifier/configs:/data/conf carvprotocol/verifier
```

#### Use Keystore (mode = 2)
If you want to use keystore, you need to generate a keystore file first. You can use the following command to generate a keystore file.
```shell
./verifier -generate-keystore -keystore-path <path to generate your keystore file>
```
After running the command, you can run verifier by following the steps below.
##### update config_docker.yaml
Update your keystore_path and keystore_password and run :
```yaml
......

wallet:
  # wallet mode, by which way to pass the private key
  # 0: through startup parameters
  # 1: through plain text private key in config file
  # 2: through path and password of the keystore in config file
  mode: 2
  # plain text private key, needed when mode is 1
  private_key: ""
  # path of the keystore, needed when mode is 2
  keystore_path: "/data/conf/xxxx/UTC--2021-09-29T07-00-00.000000000Z--xxxx"
  # password of the keystore, needed when mode is 2
  keystore_password: "123456"
......

```

```shell
docker run -d --name verifier -v /<Path To This Repository>/verifier/configs:/data/conf -v /<Path To Keystore direction>:/data/keystore carvprotocol/verifier
```

### Configuration

```yaml
chain:
  chain_id: 421614
  chain_name: "arbitrum-sepolia"
  rpc_url: "https://sepolia-rollup.arbitrum.io/rpc"
  start_block: 0
  offset_block: 3600 # opBNB block time: 1 sec. An offset of 3600 starts fetching blocks from 1 hours ago.
contract:
  addr: "0xcb37148add8b8be58034a742495d935c78d9fd76"
  tee_addr : "0x7f57004E08ef1702b2b88b87ae01a561ae10F10e"
  topic1: "0x89a3b784b99180438f3b2027aa89e97c3c3ed66e8dc78a555d7013b39caf1a89"
  topic2: "0x455929120054502ca2ea8194b26e7bb3acb631d30177f6881ffa70581abd4a13"
settings_contract:
  addr: "0xBeFAB38c0Cf603Ca59fe064fc7B3c2BcA726A28F"
#wallet:
  ## wallet mode, by which way to pass the private key
  ## 0: through startup parameters
  ## 1: through plain text private key in config file
  ## 2: through path and password of the keystore in config file
#  mode: 1
#  private_key: "99a038e9d345d0b12130b3b1fb003bf8f2d3a5c27ce2a800bbb1608efff6c591"
#  keystore_path: ""
#  keystore_password: ""
#  reward_claimer_addr: "0x689d0b32Da0480095b7AE7b604f1b1997736B3F9"
#  commission_rate: 100
wallet:
  mode: 2
  private_key: ""
  keystore_path: "./keystore/UTC--2024-06-25T12-01-01.663731000Z--031cff11b035aa5f5189f163c1fc937bf0be235c"
  keystore_password: "150318"
  reward_claimer_addr: "0x689d0b32Da0480095b7AE7b604f1b1997736B3F9"
  commission_rate: 100
signature:
  domain_name: "ProtocolService"
  domain_version: "1.0.0"
  expired_time: 3600
gasless_service:
  url: "https://dev-interface.carv.io"
```

## Contribution
We welcome contributions of all forms! If you have any questions, suggestions, or have identified any bugs, please submit them through Issues. If you would like to contribute code, please submit it via Pull Requests.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.
