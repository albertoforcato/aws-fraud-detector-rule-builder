package aws_fraud_detector_rule_builder

import (
	"fmt"
)

type RuleBuilder struct {
	expression string
	variables  map[string]VariableType
}

func NewRuleBuilder() *RuleBuilder {
	return &RuleBuilder{}
}

func (r *RuleBuilder) OpenParenthesis() *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += "("
	return r
}

func (r *RuleBuilder) CloseParenthesis() *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += ")"
	return r
}

func (r *RuleBuilder) And() *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += " and "
	return r
}

func (r *RuleBuilder) Or() *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += " or "
	return r
}

func (r *RuleBuilder) Not() *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += fmt.Sprintf(" not")
	return r
}

func (r *RuleBuilder) Equal(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += fmt.Sprintf(" == %s", evalElem(elem))
	return r
}

func (r *RuleBuilder) NotEqual(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += fmt.Sprintf(" != %s", evalElem(elem))
	return r
}

func (r *RuleBuilder) GreaterThan(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += fmt.Sprintf(" > %s", evalElem(elem))
	return r
}

func (r *RuleBuilder) GreaterThanOrEqual(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += fmt.Sprintf(" >= %s", evalElem(elem))
	return r
}

func (r *RuleBuilder) LessThan(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += fmt.Sprintf(" < %s", evalElem(elem))
	return r
}

func (r *RuleBuilder) LessThanOrEqual(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}
	r.expression += fmt.Sprintf(" <= %s", evalElem(elem))
	return r
}

func (r *RuleBuilder) In(elems ...any) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += fmt.Sprintf(" in %s", getSet(elems...))

	return r
}

func (r *RuleBuilder) NotIn(elems ...any) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += fmt.Sprintf(" not in %s", getSet(elems...))

	return r
}

func (r *RuleBuilder) Sum(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += fmt.Sprintf(" + %s", evalElem(elem))

	return r
}

func (r *RuleBuilder) Sub(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += fmt.Sprintf(" - %s", evalElem(elem))

	return r
}

func (r *RuleBuilder) Mul(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += fmt.Sprintf(" * %s", evalElem(elem))

	return r
}

func (r *RuleBuilder) Div(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += fmt.Sprintf(" / %s", evalElem(elem))

	return r
}

func (r *RuleBuilder) Mod(elem any) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += fmt.Sprintf(" %% %s", evalElem(elem))

	return r
}

func (r *RuleBuilder) AddRegexMatch(regex string, varName string) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += fmt.Sprintf("regex_match(\"%s\", $%s)", regex, varName)

	return r
}

func (r *RuleBuilder) Var(varName string, varType VariableType) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.addVariable(varName, varType)

	r.expression += fmt.Sprintf("$%s", varName)

	return r
}

func (r *RuleBuilder) VarUppercase(varName string, varType VariableType) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.addVariable(varName, varType)

	r.expression += fmt.Sprintf("uppercase($%s)", varName)

	return r
}

func (r *RuleBuilder) VarLowercase(varName string, varType VariableType) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.addVariable(varName, varType)

	r.expression += fmt.Sprintf("lowercase($%s)", varName)

	return r
}

func (r *RuleBuilder) CustomString(str string) *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression += str

	return r
}

func (r *RuleBuilder) Erase() *RuleBuilder {
	if r == nil {
		return nil
	}

	r.expression = ""

	return r
}

func (r *RuleBuilder) String() string {
	if r == nil {
		return ""
	}

	return r.expression
}

func (r *RuleBuilder) GetVariables() map[string]VariableType {
	if r == nil {
		return nil
	}

	return r.variables
}

func (r *RuleBuilder) GetVariablesByTypes() map[VariableType][]string {
	if r == nil {
		return nil
	}

	variablesByTypes := make(map[VariableType][]string)

	for varName, varType := range r.variables {
		if _, ok := variablesByTypes[varType]; !ok {
			variablesByTypes[varType] = []string{}
		}

		variablesByTypes[varType] = append(variablesByTypes[varType], varName)
	}

	return variablesByTypes
}
