package utils

import (
	"fmt"
	"log"
	"net"

	"github.com/go-ldap/ldap/v3"
)

// Index is for test
var Index = 0

// MyAdd is function to show how to write test cases
func MyAdd(a, b int) int {
	return a + b
}

// LocalIPs is get the local ips([]string)
func LocalIPs() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	// ips = strings.TrimSpace(ips)
	return
}

// LdapAuth is auth by ldap
// https://github.com/go-ldap/ldap/blob/master/examples_test.go#L164
// https://www.jianshu.com/p/f091c3f22806
func LdapAuth(userName, password string) (succ bool) {
	succ = false

	// bind user for search
	bindUserName := fmt.Sprintf("uid=%s,ou=user,dc=mapbar,dc=com", userName)
	bindPassword := password

	conn, err := ldap.DialURL("ldap://ldap.mapbar.com")
	if err != nil {
		log.Printf("ldap.DialURL had error: %v\n", err)
		return
	}
	defer conn.Close()

	// First bind with a read only user
	err = conn.Bind(bindUserName, bindPassword)
	if err != nil {
		log.Printf("conn.Bind had error: %v\n", err)
		return
	}

	// Search for the given userName
	searchRequest := ldap.NewSearchRequest(
		"dc=mapbar,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(uid=%s)", userName),
		[]string{"dn"},
		nil,
	)

	result, err := conn.Search(searchRequest)
	if err != nil {
		log.Printf("conn.Search had error: %v\n", err)
		return
	}

	if len(result.Entries) != 1 {
		log.Printf("entries count is %d\n", len(result.Entries))
		return
	}

	// userdn := result.Entries[0].DN
	// // Bind as the user to verify their password
	// err = conn.Bind(userdn, password)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Rebind as the read only user for any further queries
	// err = conn.Bind(bindUserName, bindPassword)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	succ = true
	return
}
