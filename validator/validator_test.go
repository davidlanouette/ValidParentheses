package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyString(t *testing.T) {
	assert.Panics(t, func() { ValidateParens("") })
}
func TestValidateSingleParen(t *testing.T) {
	assert.False(t, ValidateParens("("))
	assert.False(t, ValidateParens(")"))
}

func TestValidateHappyPath(t *testing.T) {
	assert.True(t, ValidateParens("()"))
	assert.True(t, ValidateParens("()()()()()"))
	assert.True(t, ValidateParens("(()()()()())"))
	assert.True(t, ValidateParens("(()())(()()())"))
}

func TestValidateTooManyClosingParens(t *testing.T) {
	assert.False(t, ValidateParens("())"))
	assert.False(t, ValidateParens("()())"))
	assert.False(t, ValidateParens("())))))"))
}

func TestValidateTooManyOpeningParens(t *testing.T) {
	assert.False(t, ValidateParens("()("))
	assert.False(t, ValidateParens("(()"))
	assert.False(t, ValidateParens("()()("))
	assert.False(t, ValidateParens("()((((((((("))
}
