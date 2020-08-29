package main

import (
	"fmt"
	"server-demo/utils"
	"testing"
)

func TestMyAdds(t *testing.T) {
	succ := utils.LdapAuth("liudf", "xxx")
	if succ {
		fmt.Println(succ)
	}
}
