![solduck](https://github.com/realphant0m/SOLDUCK/blob/master/solduck.png)

The SOLDUCK CLI provides a simple and interactive shell to document Ethereum-based smart contract development. 

> SOLDUCK is a developer tool that helps organize scattered or premature documentation and clears out information gaps related to standard Ethereum-based smart contract projects i.e. tokens, crowdsales, etc.

## Getting Started

Build
1. Start the interactive shell: ./SOLDUCK.go
2. Enjoy!

### Prerequisites

To use this tool a few prerequisites are required.
1. Make sure Golang is installed.
2. Download this repository.
3. Install the required packages. 
4. Build the binaries. 

```
git clone https://github.com/realphant0m/SOLDUCK.git
go get github.com/AlecAivazis/survey
go build 
```


## Sample Documentation

By completing the interactive shell succesfully, an output should be generated similar to below.

| Sample Documentation  | 
| ------------- | 
| The name of the requesting client: John Doe.  | 
| The LinkedIn URL of the requesting client: linkedin.com/users/JohnDoe. | 
| The name of the requesting client: John Doe.  | 
| The budget for the smart contract development: 500. | 
| The deadline for the smart contract deliverable: 31.12.19.  | 
| The client’s preferred programming language: Solidity. | 
| The type of smart contract development: Token.  | 
| The name of the token: JohnsToken.| 
| The symbol of the token: JTKN. | 
| The decimals of the token: 18. | 
| The amount of tokens to be initially created: 1000000000. | 
| The token supply will be: Constant. | 
| Token holders will be able to burn their tokens from the total supply.  | 


## Built With

* [survey] (https://github.com/AlecAivazis/survey) - A golang library for building interactive prompts

## Contributing

Please feel free to open issues, help improve the tool or similar! 

## Future

In the future when I have more time, I plan on updating this tool to be capable of automatically generating smart contracts templated using user-provided shell inputs and based on the smart contracts provided by OpenZeppelin. 

