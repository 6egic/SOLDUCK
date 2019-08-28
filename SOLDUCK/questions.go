package main

import (
	"github.com/AlecAivazis/survey"
)

// the general questions to ask
var qGeneral = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "What is the name of the requesting client?"},
		Validate: survey.Required,
	},
	{
		Name:     "linkedin",
		Prompt:   &survey.Input{Message: "What is the LinkedIn URL of the requesting client?"},
		Validate: survey.Required,
	},
	{
		Name:     "budget",
		Prompt:   &survey.Input{Message: "What is the budget for the smart contract development in USD?"},
		Validate: survey.Required,
	},
	{
		Name:     "deadline",
		Prompt:   &survey.Input{Message: "What is the deadline for the smart contract deliverable (DD/MM/YY)?"},
		Validate: survey.Required,
	},
	{
		Name: "language",
		Prompt: &survey.Select{
			Message: "What is your preferred programming language?",
			Options: []string{"Solidity", "Vyper"},
			Default: "Solidity",
		},
		Validate: survey.Required,
	},
	{
		Name: "type",
		Prompt: &survey.Select{
			Message: "What type of smart contract development are you looking for?",
			Options: []string{"Token", "Crowdsale"},
		},
		Validate: survey.Required,
	},
	{
		Name:     "outFilePath",
		Prompt:   &survey.Input{Message: "Please enter the path for output file."},
		Validate: survey.Required,
	},
}

// the crowdsale questions to ask
var qCrowdsale map[string]*survey.Question

// the token questions to ask
var qToken map[string]*survey.Question

//Populates the Map having Token Questions
func populateQuestionsToken() {
	qToken = make(map[string]*survey.Question)
	qToken["Standard"] = &survey.Question{
		Name: "standard",
		Prompt: &survey.Confirm{
			Message: "Do you wish to comply with a particular standard?"},
		Validate: survey.Required,
	}
	qToken["StandardTwo"] = &survey.Question{
		Name: "standardTwo",
		Prompt: &survey.Select{
			Message: "Please choose the standard that you wish to comply with:",
			Options: []string{"ERC20"},
			Default: "ERC20",
		},
	}
	qToken["Name"] = &survey.Question{
		Name:     "name",
		Prompt:   &survey.Input{Message: "What is the name of the token?"},
		Validate: survey.Required,
	}
	qToken["Symbol"] = &survey.Question{
		Name:     "symbol",
		Prompt:   &survey.Input{Message: "What is the symbol of the token?"},
		Validate: survey.Required,
	}
	qToken["Decimals"] = &survey.Question{
		Name:     "decimals",
		Prompt:   &survey.Input{Message: "What is the decimals of the token?"},
		Validate: survey.Required,
	}
	qToken["InitialSupply"] = &survey.Question{
		Name:     "initialSupply",
		Prompt:   &survey.Input{Message: "How many tokens do you wish to initially create?"},
		Validate: survey.Required,
	}
	qToken["TokenState"] = &survey.Question{
		Name: "tokenState",
		Prompt: &survey.Select{
			Message: "Will the token supply be dynamic or constant?",
			Options: []string{"Dynamic", "Constant"},
			Default: "Constant",
		},
		Validate: survey.Required,
	}
	qToken["MintOwnerPerm"] = &survey.Question{
		Name: "MintOwnerPerm",
		Prompt: &survey.Confirm{
			Message: "Will the creator be able to create new tokens to the initial supply?",
		},
		Validate: survey.Required,
	}
	qToken["MintPeoplePerm"] = &survey.Question{
		Name: "MintPeoplePerm",
		Prompt: &survey.Confirm{
			Message: "Will other people be able to create new tokens?",
		},
		Validate: survey.Required,
	}
	qToken["MaximumSupply"] = &survey.Question{
		Name: "maximumSupply",
		Prompt: &survey.Confirm{
			Message: "Will there be a maximum supply of tokens?:",
		},
		Validate: survey.Required,
	}
	qToken["MaximumSupplySize"] = &survey.Question{
		Name:     "maximumSupplySize",
		Prompt:   &survey.Input{Message: "What is the maximum supply of tokens?"},
		Validate: survey.Required,
	}
	qToken["Burnable"] = &survey.Question{
		Name: "Burnable",
		Prompt: &survey.Confirm{
			Message: "Will token holders be able to burn their tokens?",
		},
		Validate: survey.Required,
	}
	qToken["Pausable"] = &survey.Question{
		Name: "Pausable",
		Prompt: &survey.Confirm{
			Message: "Will the creator be able to pause the token state?",
		},
		Validate: survey.Required,
	}
	qToken["PausableTwo"] = &survey.Question{
		Name: "PausableTwo",
		Prompt: &survey.Confirm{
			Message: "Will other people be able to pause the token state?",
		},
		Validate: survey.Required,
	}
}

//populates the map of CrowdSale questions
func populateQuestionsCrowdsale() {
	qCrowdsale = make(map[string]*survey.Question)
	qCrowdsale["RefundableSale"] = &survey.Question{
		Name: "refundableSale",
		Prompt: &survey.Confirm{
			Message: "Will the crowdsale be refundable?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["FinalizableSale"] = &survey.Question{
		Name: "finalizableSale",
		Prompt: &survey.Confirm{
			Message: "Will the crowdsale be finalizable?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["InstantDistribution"] = &survey.Question{
		Name: "instantDistribution",
		Prompt: &survey.Confirm{
			Message: "Will the crowdsale tokens be delivered instantly as per contribution?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["PostDistribution"] = &survey.Question{
		Name: "postDistribution",
		Prompt: &survey.Confirm{
			Message: "Will the crowdsale tokens be delivered post crowdsale?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["AllowanceEmission"] = &survey.Question{
		Name: "allowanceEmission",
		Prompt: &survey.Confirm{
			Message: "Will the emission of tokens be allowance based?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["MintEmission"] = &survey.Question{
		Name: "mintEmission",
		Prompt: &survey.Confirm{
			Message: "Will the emission be minting based?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["Whitelisting"] = &survey.Question{
		Name: "whitelisting",
		Prompt: &survey.Confirm{
			Message: "Will participation in the crowdsale require whitelisting?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["TimedSale"] = &survey.Question{
		Name: "timedSale",
		Prompt: &survey.Confirm{
			Message: "Will the crowdsale be timed?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["StartDate"] = &survey.Question{
		Name:     "startDate",
		Prompt:   &survey.Input{Message: "What is the start date?"},
		Validate: survey.Required,
	}
	qCrowdsale["EndDate"] = &survey.Question{
		Name:     "endDate",
		Prompt:   &survey.Input{Message: "What is the end date?:"},
		Validate: survey.Required,
	}
	qCrowdsale["Pausable"] = &survey.Question{
		Name: "pausable",
		Prompt: &survey.Confirm{
			Message: "Will the crowdsale be pausable?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["PausePerm"] = &survey.Question{
		Name: "pausePerm",
		Prompt: &survey.Select{
			Message: "Who will be able to pause the crowdsale?",
			Options: []string{"The Owner", "The Owner and Permissioned Individuals"},
		},
		Validate: survey.Required,
	}
	qCrowdsale["CappedSale"] = &survey.Question{
		Name: "cappedSale",
		Prompt: &survey.Confirm{
			Message: "Will the crowdsale be capped?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["SoftCap"] = &survey.Question{
		Name: "softCap",
		Prompt: &survey.Confirm{
			Message: "Will there be a soft cap?:",
		},
		Validate: survey.Required,
	}
	qCrowdsale["SoftCapSize"] = &survey.Question{
		Name:     "softCapSize",
		Prompt:   &survey.Input{Message: "What is the soft cap?"},
		Validate: survey.Required,
	}
	qCrowdsale["MaximumCap"] = &survey.Question{
		Name: "maximumCap",
		Prompt: &survey.Confirm{
			Message: "Will there be a maximum cap?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["MaximumCapSize"] = &survey.Question{
		Name:     "maximumCapSize",
		Prompt:   &survey.Input{Message: "What is the maximum cap?"},
		Validate: survey.Required,
	}
	qCrowdsale["IndividuallyCappedSale"] = &survey.Question{
		Name: "individuallyCappedSale",
		Prompt: &survey.Confirm{
			Message: "Will the crowdsale be individually capped?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["IndividualMinSize"] = &survey.Question{
		Name:     "individualMinSize",
		Prompt:   &survey.Input{Message: "What is the minimum individual contribution?"},
		Validate: survey.Required,
	}
	qCrowdsale["IndividualMax"] = &survey.Question{
		Name: "individualMax",
		Prompt: &survey.Confirm{
			Message: "Will there be a maximum individual cap?"},
	}
	qCrowdsale["IndividualMaxSize"] = &survey.Question{
		Name:     "individualMaxSize",
		Prompt:   &survey.Input{Message: "What is the maximum individual cap?"},
		Validate: survey.Required,
	}
	qCrowdsale["Token"] = &survey.Question{
		Name:     "token",
		Prompt:   &survey.Input{Message: "Which token will be offered during the crowdsale?"},
		Validate: survey.Required,
	}
	qCrowdsale["Rate"] = &survey.Question{
		Name:     "rate",
		Prompt:   &survey.Input{Message: "What will the token rate be per contributed ether?"},
		Validate: survey.Required,
	}
	qCrowdsale["Managing"] = &survey.Question{
		Name: "managing",
		Prompt: &survey.Confirm{
			Message: "Will multiple people be able to manage the crowdsale?:",
		},
		Validate: survey.Required,
	}
	qCrowdsale["Bonus"] = &survey.Question{
		Name: "bonus",
		Prompt: &survey.Confirm{
			Message: "Will there be contribution bonuses in the crowdsale?",
		},
		Validate: survey.Required,
	}
	qCrowdsale["BonusTrigger"] = &survey.Question{
		Name: "bonusTrigger",
		Prompt: &survey.Select{
			Message: "What will the bonus be granted from?",
			Options: []string{"Contribution Size", "Contribution Earliness"},
		},
		Validate: survey.Required,
	}
}
