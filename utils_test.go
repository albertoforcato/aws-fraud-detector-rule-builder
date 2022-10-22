package rlbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalElem(t *testing.T) {
	el := evalElem(1)
	assert.Equal(t, "1", el)

	el = evalElem(1.1)
	assert.Equal(t, "1.100000", el)

	el = evalElem("test")
	assert.Equal(t, "\"test\"", el)
}

func Test_getSet(t *testing.T) {
	el := getSet(1, 2, 3)
	assert.Equal(t, "[1, 2, 3]", el)

	el = getSet(1.1, 2.2, 3.3)
	assert.Equal(t, "[1.100000, 2.200000, 3.300000]", el)

	el = getSet("test1", "test2", "test3")
	assert.Equal(t, "[\"test1\", \"test2\", \"test3\"]", el)
}

func TestRuleBuilder_addVariable(t *testing.T) {
	r := NewRuleBuilder()
	r.addVariable("test", VariableTypeEmail)
	assert.Equal(t, map[string]VariableType{"test": VariableTypeEmail}, r.variables)
}

func TestRuleBuilder_addVariable_nil(t *testing.T) {
	var r *RuleBuilder
	r.addVariable("test", VariableTypeEmail)
	assert.Nil(t, r)
}
