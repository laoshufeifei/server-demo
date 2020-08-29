package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAdd(t *testing.T) {
	test := assert.New(t)
	test.Equal(MyAdd(1, 2), 3)
	test.NotEqual(MyAdd(1, 2), 4)
}
