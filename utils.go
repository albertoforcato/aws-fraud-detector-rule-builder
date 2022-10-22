package aws_fraud_detector_rule_builder

import "fmt"

func evalElem(elem any) string {
	if _, ok := elem.(float64); ok {
		return fmt.Sprintf("%f", elem)
	} else if _, ok := elem.(int); ok {
		return fmt.Sprintf("%d", elem)
	} else {
		return fmt.Sprintf("\"%s\"", elem)
	}
}

func getSet(elems ...any) string {
	tmpExpr := "["
	for _, elem := range elems {
		tmpExpr += fmt.Sprintf("%s, ", evalElem(elem))
	}
	// Remove the last comma
	tmpExpr = tmpExpr[:len(tmpExpr)-2]

	tmpExpr += "]"

	return tmpExpr
}

func (r *RuleBuilder) addVariable(name string, varType VariableType) {
	if r == nil {
		return
	}

	if r.variables == nil {
		r.variables = make(map[string]VariableType)
	}

	if _, ok := r.variables[name]; !ok {
		r.variables[name] = varType
	}
}
