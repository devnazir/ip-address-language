package parser

// TODO: Implement type annotation parsing

// func (p *Parser) ParseTypeAnnotation(v node.ASTNode) string {
// 	result := ""
// 	vType := reflect.TypeOf(v)

// 	switch vType {
// 	case reflect.TypeOf(node.VariableDeclaration{}):
// 		node := v.(node.VariableDeclaration)
// 		_, valueType := p.ReflectInitVariableDeclaratorType(node)

// 		if node.TypeAnnotation != "" {
// 			if valueType != node.TypeAnnotation {
// 				oops.TypeMismatchError(p.peek(), node.TypeAnnotation, valueType)
// 			}

// 			result = node.TypeAnnotation
// 		} else {
// 			result = p.InferType(v)
// 		}
// 	default:
// 		break
// 	}

// 	return result
// }

// func (p *Parser) InferType(v node.ASTNode) string {
// 	result := ""
// 	vType := reflect.TypeOf(v)

// 	switch vType {
// 	case reflect.TypeOf(node.VariableDeclaration{}):
// 		node := v.(node.VariableDeclaration)
// 		_, valueType := p.ReflectInitVariableDeclaratorType(node)
// 		result = valueType
// 	default:
// 		result = ""
// 	}

// 	return result
// }

// func (p *Parser) ReflectInitVariableDeclaratorType(v node.VariableDeclaration) (reflect.Type, string) {
// 	initType := reflect.TypeOf(v.Declarations[0].Init)
// 	valueType := ""

// 	if initType.Kind() == reflect.Struct {
// 		switch initType {
// 		case reflect.TypeOf(node.Literal{}):
// 			valueType = reflect.TypeOf(v.Declarations[0].Init.(node.Literal).Value).String()
// 		case reflect.TypeOf(node.BinaryExpression{}):
// 			isConcat := p.IsConcatenation(v.Declarations[0].Init.(node.BinaryExpression))

// 			if isConcat {
// 				valueType = "string"
// 				break
// 			}

// 			result := interpreterEvaluateBinaryExpr(v.Declarations[0].Init.(node.BinaryExpression))
// 			valueType = reflect.TypeOf(result).String()
// 			break
// 		default:
// 			valueType = ""
// 		}
// 	}

// 	return initType, valueType
// }

// func (p *Parser) IsConcatenation(b node.ASTNode) bool {
// 	if reflect.TypeOf(b) == reflect.TypeOf(node.Literal{}) {
// 		fmt.Printf("Tokens %+v", p.tokens)
// 		return reflect.TypeOf(b.(node.Literal).Value) == reflect.TypeOf("")
// 	}

// 	if reflect.TypeOf(b) == reflect.TypeOf(node.BinaryExpression{}) {
// 		leftTypeIsString := p.IsConcatenation(b.(node.BinaryExpression).Left)
// 		rightTypeIsString := p.IsConcatenation(b.(node.BinaryExpression).Right)
// 		isPlusSign := b.(node.BinaryExpression).Operator == "+"

// 		if (leftTypeIsString || rightTypeIsString) && !isPlusSign {
// 			oops.InvalidConcatenationError(p.peek(), ""+b.(node.BinaryExpression).Operator+" operator"+" is not allowed")
// 		}

// 		return (leftTypeIsString || rightTypeIsString) && isPlusSign
// 	}

// 	return false
// }
