package accountsmerge

import (
	"fmt"
	"sort"
)

func AccountsMerge(accounts [][]string) [][]string {
	res := [][]string{}

	acMap := make(map[string]map[string]byte)

	for _, account := range accounts {
		name := account[0]
		emails := account[1:]
		fmt.Printf("name = %v, emails = %v", name, emails)
		if _, ok := acMap[name]; !ok { // name not found
			acMap[name] = make(map[string]byte) // init inner map (set)
			for _, email := range emails {
				acMap[name][email] = 0
			}
		} else { // name found, could be the same, could be diff persion
			sameAcc := false
			for _, email := range emails {
				if _, ok := acMap[name][email]; ok {
					sameAcc = true
				}
			}
			if sameAcc { // found same email
				sumEmails := []string{}
				for _, email := range emails { // add new account's emails into sum
					sumEmails = append(sumEmails, email)
				}
				for email := range acMap[name] { // add existing account's email into sum
					sumEmails = append(sumEmails, email)
				}
				sort.Slice(sumEmails, func(i, j int) bool {
					return sumEmails[i] > sumEmails[j]
				})
				sumAcc := append([]string{name}, sumEmails...)
				res = append(res, sumAcc)
			} else { // not the same account
				existingAcc := []string{}
				for email := range acMap[name] {
					existingAcc = append(existingAcc, email)
				}
				sort.Slice(existingAcc, func(i, j int) bool {
					return existingAcc[i] > existingAcc[j]
				})
				existingAcc = append([]string{name}, existingAcc...)
				res = append(res, existingAcc)

				newAcc := []string{}
				for _, email := range emails {
					newAcc = append(newAcc, email)
				}
				sort.Slice(newAcc, func(i, j int) bool {
					return newAcc[i] > newAcc[j]
				})
				newAcc = append([]string{name}, newAcc...)
				res = append(res, newAcc)
			}
			delete(acMap, name)
		}
	}
	for name, emails := range acMap {
		remainAcc := []string{}
		for email := range emails {
			remainAcc = append(remainAcc, email)
		}
		sort.Slice(remainAcc, func(i, j int) bool {
			return remainAcc[i] > remainAcc[j]
		})
		remainAcc = append([]string{name}, remainAcc...)
		res = append(res, remainAcc)
	}
	return res
}
