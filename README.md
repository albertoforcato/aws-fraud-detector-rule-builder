# AWS Fraud Detector Rule Builder

This is a simple tool to help you build AWS Fraud Detector rules. It is not a replacement for the AWS Fraud Detector
console, but it can help you build rules faster.

## Install

```shell
go get github.com/albertoforcato/aws-fraud-detector-rule-builder@0.1.0
```

## Examples

### Simple rule

```go
package main

import (
	"fmt"
	"github.com/albertoforcato/aws-fraud-detector-rule-builder"
)

func main() {
	rule := rlbuilder.NewRuleBuilder()
	.Var("email", rlbuilder.VariableTypeEmail)
	.NotIn("a", "b", "c")
	.Build()

	// Output: $email not in ["a", "b", "c"]
	fmt.Println(rule)
}
```