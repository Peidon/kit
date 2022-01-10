package sql

import (
	"fmt"
	"testing"
)

func TestColX_Enter(t *testing.T) {
	astNode, err := parse("SELECT a, b FROM t GROUP BY (a, b) HAVING a > c ORDER BY b")
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}
	fmt.Printf("%v\n", extract(astNode))
}
