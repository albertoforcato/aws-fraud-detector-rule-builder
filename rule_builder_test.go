package rlbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleBuilder_AddRegexMatch(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.AddRegexMatch(".*", "email")
	assert.Equal(t, "regex_match(\".*\", $email)", s.Build())
}

func TestRuleBuilder_And(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).GreaterThan(10).And().Var("b", VariableTypeNumeric).LessThan(20).Build()
	assert.Equal(t, "$a > 10 and $b < 20", s)
}

func TestRuleBuilder_Div(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).Div(10).Build()
	assert.Equal(t, "$a / 10", s)
}

func TestRuleBuilder_Equal(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).Equal(10).Build()
	assert.Equal(t, "$a == 10", s)
}

func TestRuleBuilder_GreaterThan(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).GreaterThan(10).Build()
	assert.Equal(t, "$a > 10", s)
}

func TestRuleBuilder_GreaterThanOrEqual(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).GreaterThanOrEqual(10).Build()
	assert.Equal(t, "$a >= 10", s)
}

func TestRuleBuilder_In(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("email", VariableTypeEmail).In("a", "b", "c").Build()
	assert.Equal(t, "$email in [\"a\", \"b\", \"c\"]", s)
}

func TestRuleBuilder_LessThan(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).LessThan(10).Build()
	assert.Equal(t, "$a < 10", s)
}

func TestRuleBuilder_LessThanOrEqual(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).LessThanOrEqual(10).Build()
	assert.Equal(t, "$a <= 10", s)
}

func TestRuleBuilder_Mul(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).Mul(10).Build()
	assert.Equal(t, "$a * 10", s)
}

func TestRuleBuilder_NotEqual(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).NotEqual(10).Build()
	assert.Equal(t, "$a != 10", s)
}

func TestRuleBuilder_NotIn(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("email", VariableTypeEmail).NotIn("a", "b", "c").Build()
	assert.Equal(t, "$email not in [\"a\", \"b\", \"c\"]", s)
}

func TestRuleBuilder_Or(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).GreaterThan(10).Or().Var("b", VariableTypeNumeric).LessThan(20).Build()
	assert.Equal(t, "$a > 10 or $b < 20", s)
}

func TestRuleBuilder_Sub(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).Sub(10).Build()
	assert.Equal(t, "$a - 10", s)
}

func TestRuleBuilder_Not(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("email", VariableTypeEmail).Not().In("a", "b", "c").Build()
	assert.Equal(t, "$email not in [\"a\", \"b\", \"c\"]", s)
}

func TestRuleBuilder_OpenParenthesisAndCloseParenthesis(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).GreaterThan(10).And().OpenParenthesis().Var("b", VariableTypeNumeric).LessThan(20).CloseParenthesis().Build()
	assert.Equal(t, "$a > 10 and ($b < 20)", s)
}

func TestRuleBuilder_Erase(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).GreaterThan(10).Erase().Var("b", VariableTypeNumeric).LessThan(20).Build()
	assert.Equal(t, "$b < 20", s)
}

func TestRuleBuilder_Sum(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).Sum(10).Build()
	assert.Equal(t, "$a + 10", s)
}

func TestRuleBuilder_Mod(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("a", VariableTypeNumeric).Mod(10).Build()
	assert.Equal(t, "$a % 10", s)
}

func TestRuleBuilder_GetVariables(t *testing.T) {
	rb := NewRuleBuilder()

	rb.Var("a", VariableTypeNumeric).GreaterThan(10).And().OpenParenthesis().Var("b", VariableTypeNumeric).LessThan(20).CloseParenthesis()
	assert.Equal(t, map[string]VariableType{
		"a": VariableTypeNumeric,
		"b": VariableTypeNumeric,
	}, rb.GetVariables())
}

func TestRuleBuilder_GetVariablesByTypes(t *testing.T) {
	rb := NewRuleBuilder()

	rb.Var("a", VariableTypeNumeric).GreaterThan(10).And().OpenParenthesis().Var("b", VariableTypeNumeric).LessThan(20).CloseParenthesis()
	assert.Equal(t, map[VariableType][]string{
		VariableTypeNumeric: {"a", "b"},
	}, rb.GetVariablesByTypes())
}

func TestRuleBuilder_VarUppercase(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.VarUppercase("a", VariableTypeNumeric).GreaterThan(10).Build()
	assert.Equal(t, "uppercase($a) > 10", s)
}

func TestRuleBuilder_VarLowercase(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.VarLowercase("a", VariableTypeNumeric).GreaterThan(10).Build()
	assert.Equal(t, "lowercase($a) > 10", s)
}

func TestRuleBuilder_CoverNilValues(t *testing.T) {
	var rb *RuleBuilder
	returnedRb := rb.OpenParenthesis()
	assert.Nil(t, returnedRb)

	returnedRb = rb.CloseParenthesis()
	assert.Nil(t, returnedRb)

	returnedRb = rb.And()
	assert.Nil(t, returnedRb)

	returnedRb = rb.Or()
	assert.Nil(t, returnedRb)

	returnedRb = rb.Not()
	assert.Nil(t, returnedRb)

	returnedRb = rb.Erase()
	assert.Nil(t, returnedRb)

	returnedRb = rb.Var("a", VariableTypeNumeric)
	assert.Nil(t, returnedRb)

	returnedRb = rb.Equal(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.NotEqual(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.GreaterThan(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.GreaterThanOrEqual(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.LessThan(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.LessThanOrEqual(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.In("a", "b", "c")
	assert.Nil(t, returnedRb)

	returnedRb = rb.NotIn("a", "b", "c")
	assert.Nil(t, returnedRb)

	returnedRb = rb.Sum(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.Sub(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.Mul(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.Div(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.Mod(10)
	assert.Nil(t, returnedRb)

	returnedRb = rb.AddRegexMatch(".*", "a")
	assert.Nil(t, returnedRb)

	returnedRb = rb.VarUppercase("a", VariableTypeNumeric)
	assert.Nil(t, returnedRb)

	returnedRb = rb.VarLowercase("a", VariableTypeNumeric)
	assert.Nil(t, returnedRb)

	returnedRb = rb.CustomString("a")
	assert.Nil(t, returnedRb)

	returnedMap := rb.GetVariables()
	assert.Nil(t, returnedMap)

	returnedMap2 := rb.GetVariablesByTypes()
	assert.Nil(t, returnedMap2)

	returnedStr := rb.Build()
	assert.Equal(t, "", returnedStr)
}

func TestRuleBuilder_CustomString(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.CustomString("a > 10").Build()
	assert.Equal(t, "a > 10", s)
}

func TestRuleBuilder(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("email", VariableTypeEmail).NotIn("a", "b", "c").Build()
	assert.Equal(t, "$email not in [\"a\", \"b\", \"c\"]", s)

	rb.Erase()

	s = rb.Var("address_zip", VariableTypeBillingZip).In("12345", "67890").Build()
	assert.Equal(t, "$address_zip in [\"12345\", \"67890\"]", s)

	rb.Erase()

	s = rb.Var("a", VariableTypeNumeric).GreaterThan(10).And().Var("b", VariableTypeNumeric).LessThan(20).Build()
	assert.Equal(t, "$a > 10 and $b < 20", s)

	rb.Erase()

	s = rb.Var("a", VariableTypeNumeric).GreaterThan(10).
		Or().
		Var("b", VariableTypeNumeric).LessThan(20).
		And().
		OpenParenthesis().
		Var("c", VariableTypeNumeric).GreaterThan(30).
		And().
		Var("d", VariableTypeNumeric).LessThan(40).
		CloseParenthesis().
		Build()
	assert.Equal(t, "$a > 10 or $b < 20 and ($c > 30 and $d < 40)", s)
}
