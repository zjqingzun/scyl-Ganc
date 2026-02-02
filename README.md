# Ganc
*(update)*


## Usage
### Clone Repository & Directory 
```bash
git clone https://github.com/zjqingzun/scyl-Ganc.git
cd scyl-Ganc
```

### Check Environment & Package (Go, Ignite, Rustup, Circom 2)
```bash
# go
go version

# ignite
ignite version

# cargo (Rust)
cargo --version

# circom
circom --version
```
If you haven't installed Ignite CLI yet, please refer to the official [Ignite CLI installation guide](./cli/readme.md). <br>

If you haven't installed Circom 2, execute the following commands:
```bash
# Rustup (optional)
curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
source "$HOME/.cargo/env"
cargo --version

# Circom 2
git clone https://github.com/iden3/circom.git
cd circom
cargo build --release
cargo install --path circom
circom --version

circom --help
```

### Run Chain
```bash
cd sw/ob
ignite chain serve -r
```

#### Run Node (Ob Node)
***Mandatory:*** 
*The prerequisite is that the chain must be running first.*
```bash
# Open a new terminal
cd sw/ob
obd tx dex --help
```

#### Test Matching Orderbook
```bash
# Open a new terminal 
cd sw/sh-scyl/test/
bash order-matching@10S10B-empty-sort.sh
```


## Contribution 
*(update)*


## Security
*(update)*


## License
*(update)*


