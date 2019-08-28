package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey"
)

// the general answers will be written to this struct
type aGeneral struct {
	Name        string
	Linkedin    string
	Budget      string
	Deadline    string
	Language    string `survey:"language"`
	Type        string `survey:"type"`
	OutFilePath string
}

// the token answers will be written to this struct
type aToken struct {
	Standard          bool
	StandardTwo       string `survey:"standardTwo"`
	Name              string
	Symbol            string
	Decimals          string
	InitialSupply     string
	TokenState        string
	MintOwnerPerm     bool
	MintPeoplePerm    bool
	MaximumSupply     bool
	MaximumSupplySize string
	Burnable          bool
	Pausable          bool
	PausableTwo       bool
}

// the crowdsale answers will be written to this struct
type aCrowdsale struct {
	RefundableSale         bool
	FinalizableSale        bool
	InstantDistribution    bool
	PostDistribution       bool
	AllowanceEmission      bool
	MintEmission           bool
	Whitelisting           bool
	TimedSale              bool
	StartDate              string
	EndDate                string
	Pausable               bool
	PausePerm              string `survey:"pausePerm"`
	CappedSale             bool
	SoftCap                bool
	SoftCapSize            string
	MaximumCap             bool
	MaximumCapSize         string
	IndividuallyCappedSale bool
	IndividualMinSize      string
	IndividualMax          bool
	IndividualMaxSize      string
	Token                  string
	Rate                   string
	Managing               bool
	Bonus                  bool
	BonusTrigger           string `survey:"bonusTrigger"`
}

var ansGeneral aGeneral
var ansToken aToken
var ansCrowdSale aCrowdsale

func writeToFile(fileName, line string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(line + "\n"); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Greetings! Together we will now document the smart contract development scope.")

	// populate question maps
	populateQuestionsCrowdsale()
	populateQuestionsToken()

	// perform the questions
	err := survey.Ask(qGeneral, &ansGeneral)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ansToken = aToken{
		Standard: false,
	}
	ansCrowdSale = aCrowdsale{}

	if ansGeneral.Type == "Token" {
		err = askTokenQuestions()
	} else if ansGeneral.Type == "Crowdsale" {
		err = askCrowdSaleQuestions()
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//General Data Logs
	fmt.Println("\n...Creating Documentation")
	writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The name of the requesting client: %s.", ansGeneral.Name))
	writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The LinkedIn URL of the requesting client: %s.", ansGeneral.Linkedin))
	writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The budget for the smart contract development: %s.", ansGeneral.Budget))
	writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The deadline for the smart contract deliverable: %s.", ansGeneral.Deadline))

	if ansGeneral.Language == "Solidity" {
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The client’s preferred programming language: %s.", ansGeneral.Language))
	} else if ansGeneral.Language == "Vyper" {
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The client’s preferred programming language: %s.", ansGeneral.Language))
	}

	if ansGeneral.Type == "Token" {
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The type of smart contract development: %s.", ansGeneral.Type))
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The name of the token: %s.", ansToken.Name))
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The symbol of the token: %s.", ansToken.Symbol))
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The decimals of the token: %s.", ansToken.Decimals))
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The amount of tokens to be initially created: %s.", ansToken.InitialSupply))
		if ansToken.TokenState == "Dynamic" {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The token supply will be: %s.", ansToken.TokenState))
			if ansToken.MintOwnerPerm == true {
				writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The creator will be able to mint new tokens to the supply."))
			}
			if ansToken.MintPeoplePerm == true {
				writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("Other people will also be able to mint new tokens."))
			}
		} else if ansToken.TokenState == "Constant" {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The token supply will be: %s.", ansToken.TokenState))
		}
		if ansToken.MaximumSupply == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("There will be a maximum supply of tokens."))
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The maximum supply of tokens: %s.", ansToken.MaximumSupplySize))
		}
		if ansToken.Burnable == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("Token holders will be able to burn their tokens from the total supply"))
		}
		if ansToken.Pausable == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The creator will be able to pause the state of the token."))
			if ansToken.PausableTwo == true {
				writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("Other people will be able to pause the state of the token."))
			}
		}

	} else if ansGeneral.Type == "Crowdsale" {
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The type of smart contract development: %s.\n", ansGeneral.Type))
		if ansCrowdSale.RefundableSale {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The crowdsale will be refundable."))
		}
		if ansCrowdSale.FinalizableSale == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The crowdsale will be finalizable."))
		}
		if ansCrowdSale.InstantDistribution == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The tokens offered in the crowdsale will be delivered instantly as per contribution."))
		}
		if ansCrowdSale.PostDistribution == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The tokens offered in the crowdsale will be delivered after the crowdsale."))
		}
		if ansCrowdSale.AllowanceEmission == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The tokens offered in the crowdsale will be ‘allowanced’ from a wallet."))
		}
		if ansCrowdSale.MintEmission == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The tokens offered in the crowdsale will be freshly minted as per contribution."))
		}
		if ansCrowdSale.Whitelisting == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("Participation will require users to be whitelisted beforehand."))
		}
		if ansCrowdSale.TimedSale == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The Crowdsale will be timed."))
		}
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The start date of the crowdsale:%s.\n", ansCrowdSale.StartDate))
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The end date of the crowdsale: %s.\n", ansCrowdSale.EndDate))

		if ansCrowdSale.Pausable == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The Crowdsale will be pausable."))
			if ansCrowdSale.PausePerm == "The Owner" {
				writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The Crowdsale will only be pausable by: %s.\n", ansCrowdSale.PausePerm))
			} else if ansCrowdSale.PausePerm == "The Owner and Permissioned Individuals\n" {
				writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The Crowdsale will be pausable by: %s.\n", ansCrowdSale.PausePerm))
			}
		}
		if ansCrowdSale.CappedSale == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The crowdsale will be capped at a certain amount of ether raised."))
		}
		if ansCrowdSale.SoftCap == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The crowdsale will have a soft cap."))
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The soft cap of the crowdsale: %s.\n", ansCrowdSale.SoftCapSize))
		}
		if ansCrowdSale.MaximumCap == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The crowdsale will have a maximum cap."))
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The maximum cap of the crowdsale: %s.\n", ansCrowdSale.MaximumCapSize))
		}
		if ansCrowdSale.IndividuallyCappedSale == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The crowdsale will be individually capped."))
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The minimum individual contribution: %s.\n", ansCrowdSale.IndividualMinSize))
			if ansCrowdSale.IndividualMax == true {
				writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The crowdsale will have a maximum individual cap."))
				writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The maximum individual cap: %s.\n", ansCrowdSale.IndividualMaxSize))
			}
		}
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The token offered during the crowdsale:  %s\n.", ansCrowdSale.Token))
		writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The token rate per contributed ether:  %s.\n", ansCrowdSale.Rate))
		if ansCrowdSale.Managing == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("Multiple accounts will be able to manage the crowdsale."))
		}
		if ansCrowdSale.Bonus == true {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("There will be bonuses associated with contributions during the crowdsale."))
		}
		if ansCrowdSale.BonusTrigger == "Contribution Size" {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The bonus will be based on: %s.\n", ansCrowdSale.BonusTrigger))
		} else if ansCrowdSale.BonusTrigger == "Contribution Earliness" {
			writeToFile(ansGeneral.OutFilePath, fmt.Sprintf("The bonus will be based on: %s.\n", ansCrowdSale.BonusTrigger))
		}
	}
}
