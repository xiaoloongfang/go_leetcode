package main

import "sort"

type Account struct {
	name  string
	email string
}

func accountsMerge(accounts [][]string) [][]string {
	acnts := make([][]Account, len(accounts))
	fatherAccount := make(map[Account]Account)
	for i, account := range accounts {
		acnts[i] = make([]Account, 0)
		for j, mail := range account {
			if j != 0 {
				acnt := Account{email: mail, name: account[0]}
				acnts[i] = append(acnts[i], acnt)
			}
		}
	}
	for i, account := range acnts {
		father := acnts[i][0]
		if _, ok := fatherAccount[father]; ok {
			father = findFather(fatherAccount, father)
		} else {
			fatherAccount[father] = father
		}
		for j := 1; j < len(account); j++ {
			accountj := account[j]
			if _, ok := fatherAccount[accountj]; ok {
				fatherj := findFather(fatherAccount, accountj)
				union(fatherAccount, fatherj, father)
				father = findFather(fatherAccount, father)
			} else {
				fatherAccount[accountj] = father
			}
		}
	}

	resolve := make(map[Account][]Account)
	for _, account := range acnts {
		for _, item := range account {
			father := findFather(fatherAccount, item)
			println(father.email, father.name)
			if _, ok := resolve[father]; ok {
				resolve[father] = append(resolve[father], item)
			} else {
				resolve[father] = []Account{item}
			}
		}
	}

	res := make([][]string, 0)
	for _, acs := range resolve {
		sort.Slice(acs, func(i, j int) bool { return acs[i].email < acs[j].email })
		item := make([]string, 1)
		item[0] = acs[0].name
		last := "#"
		for _, account := range acs {
			if account.email != last {
				item = append(item, account.email)
				last = account.email
			}
		}
		if len(item) > 1 {
			res = append(res, item)
		}
	}

	return res
}

func findFather(fatherAccount map[Account]Account, account Account) Account {
	for fatherAccount[account] != account {
		account = fatherAccount[account]
	}
	return account
}

func union(fatherAccount map[Account]Account, a Account, b Account) map[Account]Account {
	fatherA := findFather(fatherAccount, a)
	fatherB := findFather(fatherAccount, b)
	fatherAccount[fatherB] = fatherA
	return fatherAccount
}

func main() {
	accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}})
}
