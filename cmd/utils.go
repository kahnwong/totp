package cmd

func getOrgs() []string {
	orgs := make([]string, 0)
	for _, v := range config.Totp {
		orgs = append(orgs, v.Org)
	}

	return orgs
}

func getAccounts(org string) []string {
	accounts := make([]string, 0)

out:
	for _, v := range config.Totp {
		if v.Org == org {
			for _, v := range v.Accounts {
				accounts = append(accounts, v.Name)
			}
			break out
		}
	}

	return accounts
}
