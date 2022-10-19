package aws_fraud_detector_rule_builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleBuilder(t *testing.T) {
	rb := NewRuleBuilder()

	s := rb.Var("email", VariableTypeEmail).NotIn("a", "b", "c").String()
	assert.Equal(t, "$email not in [\"a\", \"b\", \"c\"]", s)

	rb.Erase()

	s = rb.Var("address_zip", VariableTypeBillingZip).In("12345", "67890").String()
	assert.Equal(t, "$address_zip in [\"12345\", \"67890\"]", s)

	rb.Erase()

	s = rb.Var("a", VariableTypeNumeric).GreaterThan(10).And().Var("b", VariableTypeNumeric).LessThan(20).String()
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
		String()
	assert.Equal(t, "$a > 10 or $b < 20 and ($c > 30 and $d < 40)", s)
}
