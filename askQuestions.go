package main

import (
	"strings"

	"github.com/AlecAivazis/survey"
)

//asks the questions related to Crowd Sale
func askCrowdSaleQuestions() error {
	err := survey.AskOne(qCrowdsale["RefundableSale"].Prompt, &ansCrowdSale.RefundableSale)
	if err != nil {
		return err
	}
	err = survey.AskOne(qCrowdsale["FinalizableSale"].Prompt, &ansCrowdSale.FinalizableSale)
	if err != nil {
		return err
	}
	err = survey.AskOne(qCrowdsale["InstantDistribution"].Prompt, &ansCrowdSale.InstantDistribution)
	if err != nil {
		return err
	}
	if !ansCrowdSale.InstantDistribution {
		err = survey.AskOne(qCrowdsale["PostDistribution"].Prompt, &ansCrowdSale.PostDistribution)
		if err != nil {
			return err
		}
	}
	err = survey.AskOne(qCrowdsale["AllowanceEmission"].Prompt, &ansCrowdSale.AllowanceEmission)
	if err != nil {
		return err
	}

	if !ansCrowdSale.AllowanceEmission {
		err = survey.AskOne(qCrowdsale["MintEmission"].Prompt, &ansCrowdSale.MintEmission)
		if err != nil {
			return err
		}
	}
	err = survey.AskOne(qCrowdsale["Whitelisting"].Prompt, &ansCrowdSale.Whitelisting)
	if err != nil {
		return err
	}
	err = survey.AskOne(qCrowdsale["TimedSale"].Prompt, &ansCrowdSale.TimedSale)
	if err != nil {
		return err
	}
	if ansCrowdSale.TimedSale {
		err = survey.AskOne(qCrowdsale["StartDate"].Prompt, &ansCrowdSale.StartDate)
		if err != nil {
			return err
		}
		err = survey.AskOne(qCrowdsale["EndDate"].Prompt, &ansCrowdSale.EndDate)
		if err != nil {
			return err
		}
	}
	err = survey.AskOne(qCrowdsale["Pausable"].Prompt, &ansCrowdSale.Pausable)
	if err != nil {
		return err
	}

	if ansCrowdSale.Pausable {
		err = survey.AskOne(qCrowdsale["PausePerm"].Prompt, &ansCrowdSale.PausePerm)
		if err != nil {
			return err
		}
	}

	err = survey.AskOne(qCrowdsale["CappedSale"].Prompt, &ansCrowdSale.CappedSale)
	if err != nil {
		return err
	}

	if ansCrowdSale.CappedSale {
		err = survey.AskOne(qCrowdsale["SoftCap"].Prompt, &ansCrowdSale.SoftCap)
		if err != nil {
			return err
		}
	}

	if ansCrowdSale.SoftCap {
		err = survey.AskOne(qCrowdsale["SoftCapSize"].Prompt, &ansCrowdSale.SoftCapSize)
		if err != nil {
			return err
		}
	}

	err = survey.AskOne(qCrowdsale["MaximumCap"].Prompt, &ansCrowdSale.MaximumCap)
	if err != nil {
		return err
	}

	if ansCrowdSale.MaximumCap {
		err = survey.AskOne(qCrowdsale["MaximumCapSize"].Prompt, &ansCrowdSale.MaximumCapSize)
		if err != nil {
			return err
		}
	}

	err = survey.AskOne(qCrowdsale["IndividuallyCappedSale"].Prompt, &ansCrowdSale.IndividuallyCappedSale)
	if err != nil {
		return err
	}

	if ansCrowdSale.IndividuallyCappedSale {
		err = survey.AskOne(qCrowdsale["IndividualMinSize"].Prompt, &ansCrowdSale.IndividualMinSize)
		if err != nil {
			return err
		}

		err = survey.AskOne(qCrowdsale["IndividualMax"].Prompt, &ansCrowdSale.IndividualMax)
		if err != nil {
			return err
		}
	}
	err = survey.AskOne(qCrowdsale["IndividualMaxSize"].Prompt, &ansCrowdSale.IndividualMaxSize)
	if err != nil {
		return err
	}
	err = survey.AskOne(qCrowdsale["Token"].Prompt, &ansCrowdSale.Token)
	if err != nil {
		return err
	}
	err = survey.AskOne(qCrowdsale["Rate"].Prompt, &ansCrowdSale.Rate)
	if err != nil {
		return err
	}
	err = survey.AskOne(qCrowdsale["Managing"].Prompt, &ansCrowdSale.Managing)
	if err != nil {
		return err
	}
	err = survey.AskOne(qCrowdsale["Bonus"].Prompt, &ansCrowdSale.Bonus)
	if err != nil {
		return err
	}

	if ansCrowdSale.Bonus {
		err = survey.AskOne(qCrowdsale["BonusTrigger"].Prompt, &ansCrowdSale.Bonus)
		if err != nil {
			return err
		}
	}

	return nil
}

//asks questions related to token
func askTokenQuestions() error {
	err := survey.AskOne(qToken["Standard"].Prompt, &ansToken.Standard)
	if err != nil {
		return err
	}

	err = survey.AskOne(qToken["StandardTwo"].Prompt, &ansToken.StandardTwo)
	if err != nil {
		return err
	}

	err = survey.AskOne(qToken["Name"].Prompt, &ansToken.Name)
	if err != nil {
		return err
	}

	err = survey.AskOne(qToken["Symbol"].Prompt, &ansToken.Symbol)
	if err != nil {
		return err
	}

	err = survey.AskOne(qToken["Decimals"].Prompt, &ansToken.Decimals)
	if err != nil {
		return err
	}

	err = survey.AskOne(qToken["InitialSupply"].Prompt, &ansToken.InitialSupply)
	if err != nil {
		return err
	}

	err = survey.AskOne(qToken["TokenState"].Prompt, &ansToken.TokenState)
	if err != nil {
		return err
	}

	if strings.ToUpper(ansToken.TokenState) == "DYNAMIC" {
		err = survey.AskOne(qToken["MintOwnerPerm"].Prompt, &ansToken.MintOwnerPerm)
		if err != nil {
			return err
		}
	}

	err = survey.AskOne(qToken["MintPeoplePerm"].Prompt, &ansToken.MintPeoplePerm)
	if err != nil {
		return err
	}

	err = survey.AskOne(qToken["MaximumSupply"].Prompt, &ansToken.MaximumSupply)
	if err != nil {
		return err
	}

	if ansToken.MaximumSupply {
		err = survey.AskOne(qToken["MaximumSupplySize"].Prompt, &ansToken.MaximumSupplySize)
		if err != nil {
			return err
		}
	}

	err = survey.AskOne(qToken["Burnable"].Prompt, &ansToken.Burnable)
	if err != nil {
		return err
	}

	err = survey.AskOne(qToken["Pausable"].Prompt, &ansToken.Pausable)
	if err != nil {
		return err
	}

	if ansToken.Pausable {
		err = survey.AskOne(qToken["PausableTwo"].Prompt, &ansToken.PausableTwo)
		if err != nil {
			return err
		}
	}
	return nil
}
