package accountsmerge

import (
	"fmt"
	"sort"
)

/*
Doesn't work for case where same emails happened after

	accounts := [][]string{
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"Mary", "mary@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "john00@mail.com", "johnnybravo@mail.com"},
	}

Expected:

	[
		["Mary","mary@mail.com"],
		["John","john00@mail.com","john_newyork@mail.com","johnnybravo@mail.com","johnsmith@mail.com"]
	]
*/
func AccountsMergeNotWorking(accounts [][]string) [][]string {
	res := [][]string{}

	acMap := make(map[string]map[string]byte)

	for _, account := range accounts {
		name := account[0]
		emails := account[1:]
		// fmt.Printf("name = %v, emails = %v", name, emails)
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
					delete(acMap[name], email)
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

/*

create a map of email string to int
loop through the input table, every found email, added to map and increment its val by 1

for emails with val > 1, meaning shared account found

* This way not work for different 2 Johns

		accounts := [][]string{
			{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
			{"John", "johnsmith@mail.com", "john00@mail.com"},
			{"Mary", "mary@mail.com"},
			{"John", "johnnybravo@mail.com"},
			{"John", "johnnybravo@mail.com", "johnnybravo1@mail.com"},
		}
----------

create map[dup-email]obj where
obj
{
	name
	count
	map[email]int
}

*/

type AccountEmailCount struct {
	Name    string
	Visited bool
	Emails  []string
}

func AccountsMergeBFS(accounts [][]string) [][]string {
	res := [][]string{}

	// create graph
	eMap := make(map[string]*AccountEmailCount) // the val is a pointer instead of a copy, this way we can directly modify its porperties
	for _, account := range accounts {
		name := account[0]
		emails := account[1:]
		for i, email := range emails {
			otherEmails := make([]string, len(emails)-1)
			otherEmails = append(otherEmails, emails[:i]...)
			otherEmails = append(otherEmails, emails[i+1:]...)
			if entry, ok := eMap[email]; ok { // found
				entry.Emails = append(entry.Emails, otherEmails...)
				// if not *AccountEmailCount, we have to do ' eMap[email] = entry
			} else { // first time
				eMap[email] = &AccountEmailCount{
					Name:   name,
					Emails: otherEmails,
				}
			}
		}
	}

	for k, v := range eMap {
		fmt.Printf("email = %v, emails = %v", k, v.Emails)
		fmt.Println()
	}

	// BFS begins
	for email, property := range eMap {
		if property.Visited {
			continue
		}

		account := []string{}
		queue := []string{email}

		for len(queue) != 0 {
			accEmail := queue[0]
			queue = queue[1:]
			account = append(account, accEmail)
			eMap[accEmail].Visited = true

			for _, childEmail := range eMap[accEmail].Emails {
				if _, ok := eMap[childEmail]; ok && !eMap[childEmail].Visited {
					queue = append(queue, childEmail)
					eMap[childEmail].Visited = true // avoid appending the same email twice
				}
			}
		}
		sort.SliceStable(account, func(i, j int) bool {
			return account[i] < account[j]
		})
		account = append([]string{property.Name}, account...)
		res = append(res, account)
	}

	return res
}

func AccountsMergeDFS(accounts [][]string) [][]string {
	res := [][]string{}

	// create graph
	eMap := make(map[string]*AccountEmailCount) // the val is a pointer instead of a copy, this way we can directly modify its porperties
	for _, account := range accounts {
		name := account[0]
		emails := account[1:]
		for i, email := range emails {
			otherEmails := make([]string, len(emails)-1)
			otherEmails = append(otherEmails, emails[:i]...)
			otherEmails = append(otherEmails, emails[i+1:]...)
			if entry, ok := eMap[email]; ok { // found
				entry.Emails = append(entry.Emails, otherEmails...)
				// if not *AccountEmailCount, we have to do ' eMap[email] = entry
			} else { // first time
				eMap[email] = &AccountEmailCount{
					Name:   name,
					Emails: otherEmails,
				}
			}
		}
	}

	for k, v := range eMap {
		fmt.Printf("email = %v, emails = %v", k, v.Emails)
		fmt.Println()
	}

	// DFS begins
	for email, property := range eMap {
		if property.Visited {
			continue
		}

		account := []string{}
		property.Visited = true
		dfs(&account, email, eMap)

		sort.SliceStable(account, func(i, j int) bool {
			return account[i] < account[j]
		})
		account = append([]string{property.Name}, account...)
		res = append(res, account)
	}

	return res
}

func dfs(account *[]string, email string, eMap map[string]*AccountEmailCount) {
	if eMap[email].Visited {
		*account = append(*account, email)
	}

	for _, childEmail := range eMap[email].Emails {
		if _, ok := eMap[childEmail]; ok && !eMap[childEmail].Visited {
			eMap[childEmail].Visited = true
			dfs(account, childEmail, eMap)
		}
	}
}
