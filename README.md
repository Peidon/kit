# kit
框架和工具
1. valuate
2. code generator

## valuate

以表达式的形式、可自定义算子的、高性能、可扩展的值计算框架

### 特点

* 支持自定义函数(算子)，且函数中自动传入context，无须在表达式中显式传递;
* 支持Accessor功能，可用数组下标取值;
* 支持Catch Error，可自定义错误处理;
* 支持操作符重载，可自定义操作符含义;

How do I use it?
--

You create a new EvaluableExpression, then call "Evaluate" on it.

```go
    expression, err := valuate.Expression("10 > 0");
    result, err := expression.Evaluate(nil);
    // result is now set to "true", the bool value.
```

Cool, but how about with parameters?

```go
    expression, err := valuate.Expression("foo > 0");
    
    parameters := make(map[string]interface{})
    parameters["foo"] = -1;

    result, err := expression.Evaluate(parameters);
    // result is now set to "false", the bool value.
```

That's cool, but we can almost certainly have done all that in code. What about a complex use case that involves some math?

```go
    expression, err := valuate.Expression("(requests_made * requests_succeeded / 100) >= 90");

    parameters := make(map[string]interface{})
    parameters["requests_made"] = 100;
    parameters["requests_succeeded"] = 80;

    result, err := expression.Evaluate(parameters);
    // result is now set to "false", the bool value.
```

Or maybe you want to check the status of an alive check ("smoketest") page, which will be a string?

```go
    expression, err := valuate.Expression("http_response_body == \"service is ok\"");

    parameters := make(map[string]interface{})
    parameters["http_response_body"] = "service is ok";

    result, err := expression.Evaluate(parameters);
    // result is now set to "true", the bool value.
```

These examples have all returned boolean values, but it's equally possible to return numeric ones.

```go
    expression, err := valuate.Expression("100 * (mem_used / total_mem)");

    parameters := make(map[string]interface{})
    parameters["total_mem"] = 1024;
    parameters["mem_used"] = 512;

    result, err := expression.Evaluate(parameters);
    // result is now set to "50.0", the float64 value.
```

What operators and types does this support?
--

* Modifiers: `+` `-` `/` `*`
* Comparators: `>` `>=` `<` `<=` `==` `!=`
* Logical ops: `||` `&&`
* Numeric constants, as 64-bit floating point (`12345.678`), as 64-bit int (`123`)
* String constants (double quotes: `"foobar"`)
* Boolean constants: `true` `false`
* Nil constants: `nil`
* Parenthesis to control order of evaluation `(` `)`
* Arrays (anything separated by `,` within parenthesis: `[1, 2, 3]`)
* Prefixes: `!` `-`

Functions
--

You may have cases where you want to call a function on a parameter during execution of the expression. Perhaps you want to aggregate some set of data, but don't know the exact aggregation you want to use until you're writing the expression itself. Or maybe you have a mathematical operation you want to perform, for which there is no operator; like `log` or `tan` or `sqrt`. For cases like this, you can provide a map of functions to `ExpressionWithFunctions`, which will then be able to use them during execution. For instance;

```go
	functions := map[string]valuate.ExpressionFunction {
		"strlen": func(ctx context.Context, args ...interface{}) (interface{}, error) {
			length := len(args[0].(string))
			return (float64)(length), nil
		},
	}

	expString := "strlen('someReallyLongInputString') <= 16"
	expression, _ := valuate.ExpressionWithFunctions(expString, functions)

	result, _ := expression.Evaluate(nil)
	// result is now "false", the boolean value
```

To use context, you need transfer the context like following:

```go
    expression.WithContext(ctx).Evaluate(params)
```

Functions can accept any number of arguments, correctly handles nested functions, and arguments can be of any type (even if none of this library's operators support evaluation of that type). For instance, each of these usages of functions in an expression are valid (assuming that the appropriate functions and parameters are given):

```go
"sqrt(x1 * y1, x2 + y2)"
"max(someValue, abs(anotherValue), 10 * lastValue)"
```

Builtin functions:

```go
"json_marshal(someStruct)" // return a string
"json_unmarshal(jsonStr, b)" // jsonStr := `{"name": "foo"}`, b is a struct inference
"unix_timestamp(t)" // the type of `t` must be time.Time or *time.Time
"len(abc)" // the type of `abc` must be string, slice or array
```

Accessors
--

If you have structs in your parameters, you can access their fields in the usual way. For instance, given a struct that has a field "bar", present in the parameters as `foo`, the following is valid:

	"foo.bar"

Assuming `foo` has a field called "Size":

	"foo.Size > 9000"

Accessors can be nested to any depth, like the following:

	"foo.Bar.Baz.Length"

Assuming `foo.bar.Baz` is an `array` or a `slice`:

	"foo.bar.Baz[0]"

You can access the element with an index parameter, Assuming `idx` is an integer:

    "foo.bar.Baz[idx]"

It will return error when the field not exists or index out of range.

Catch Error
--

Sometimes you want to handle the error by yourself.

```go
    // this is a simplest strategy, just omit, and return a mock value.
    func omitError(_ error) (interface{}, error) {
	    return 0, nil
    }

    expression, err := valuate.Expression("a + b");
    parameters := make(map[string]interface{})
    parameters["a"] = 1
    // missing vairable `b`
    
    
    result, err := expression.Evaluate(parameters)
    // return ParamNotFound error, result is nil
    
    result, _ = expression.CatchError(omitError).Evaluate(parameters)
    // result is now 1, the int value.
```

Escaping characters
--

Sometimes you'll have parameters that have spaces, slashes, pluses, ampersands or some other character
that this library interprets as something special. For example, the following expression will not
act as one might expect:

	"response.time < 100"

As written, the library will parse it as "[response] dot [time] is less than 100". In reality,
"response.time" is meant to be one variable that just happens to have a dot in it.

There are two ways to work around this. First, you can escape the entire parameter name:

 	"${response.time} < 100"

Or you can give one more parameter named "response" which is a struct has field "time".

Modifier Overwrite
--

Consider the following expression:

    "a + b"

`a` and `b` must be number.

But sometimes you want to overwrite the `+` operator, you can implement the interface `valuate.Modifier` like following:

```go
type emptyValue struct{}

func (empty emptyValue) Modify(_ context.Context, op string, other interface{}) (interface{}, error) {
	if op == "+" {
		return 0, nil
	}
	return other, nil
}

expression, err := valuate.Expression("a + b");
parameters := make(map[string]interface{})
parameters["a"] = 1
parameters["b"] = emptyValue{}

result, _ = expression.Evaluate(parameters)
// result is now 0
```

Parallelization
--

Sometimes operator or function are time-costly, 

    "get_result(call_api(api_a, request_of_a), query_db())"

`call_api` and `query_db` could execute at the same time.

You can use `parallelization` mode like following:

```go
expression, err := valuate.Expression("get_result(call_api(api_a, request_of_a), query_db())")
expr.Parallel()
```

### 测试环境

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
```

### 性能对比

pkg: github.com/govaluate

```
BenchmarkSingleParse-16                          1893027               619.2 ns/op           344 B/op          9 allocs/op
BenchmarkSimpleParse-16                           230068              4630 ns/op            2776 B/op         52 allocs/op
BenchmarkFullParse-16                              51336             23321 ns/op           13440 B/op        284 allocs/op
BenchmarkEvaluationSingle-16                    62566802                18.62 ns/op            0 B/op          0 allocs/op
BenchmarkEvaluationNumericLiteral-16            21492518                54.18 ns/op            0 B/op          0 allocs/op
BenchmarkEvaluationLiteralModifiers-16           9564715               127.1 ns/op             8 B/op          1 allocs/op
BenchmarkEvaluationParameter-16                 19523344                60.69 ns/op           16 B/op          1 allocs/op
BenchmarkEvaluationParameters-16                12093152                98.86 ns/op           16 B/op          1 allocs/op
BenchmarkEvaluationParametersModifiers-16        5647666               197.6 ns/op            32 B/op          3 allocs/op
BenchmarkComplexExpression-16                   23106603                50.42 ns/op           16 B/op          1 allocs/op
BenchmarkRegexExpression-16                       795068              1427 ns/op            1437 B/op         20 allocs/op
BenchmarkConstantRegexExpression-16              2723265               457.6 ns/op            32 B/op          3 allocs/op
BenchmarkAccessors-16                            6701733               170.1 ns/op            32 B/op          3 allocs/op
BenchmarkNestedAccessors-16                      4488036               243.7 ns/op            32 B/op          3 allocs/op
PASS
ok      github.com/govaluate    41.946s
```

pkg: kit/pkg/valuate

```
BenchmarkSingleParse-16                           220531              4732 ns/op            3688 B/op         64 allocs/op
BenchmarkSimpleParse-16                            35017             34767 ns/op           21728 B/op        344 allocs/op
BenchmarkFullParse-16                                319           3843065 ns/op         2578144 B/op      37762 allocs/op
BenchmarkEvaluationSingle-16                    51548596                22.80 ns/op            0 B/op          0 allocs/op
BenchmarkEvaluationNumericLiteral-16             4159724               286.7 ns/op            96 B/op          4 allocs/op
BenchmarkEvaluationLiteralModifiers-16           2155542               565.6 ns/op           192 B/op          8 allocs/op
BenchmarkEvaluationParameter-16                 58607061                20.14 ns/op            0 B/op          0 allocs/op
BenchmarkEvaluationParameters-16                 4342616               270.5 ns/op            96 B/op          4 allocs/op
BenchmarkEvaluationParametersModifiers-16        1395748               824.2 ns/op           304 B/op         14 allocs/op
BenchmarkComplexExpression-16                     311020              3721 ns/op            1344 B/op         60 allocs/op
BenchmarkRegexExpression-16                      3248205               380.5 ns/op           128 B/op          6 allocs/op
BenchmarkConstantRegexExpression-16              1338622               900.7 ns/op           320 B/op         14 allocs/op
BenchmarkAccessors-16                             772376              1617 ns/op            1056 B/op         26 allocs/op
BenchmarkNestedAccessors-16                       653570              1803 ns/op            1136 B/op         29 allocs/op
PASS
ok      kit/pkg/valuate 21.734s
```

### 分析

从BenchmarkSingleParse和BenchmarkSimpleParse的对比中可以看出，taskflow/common/valuate的操作耗时、内存占用、内存申请次数分别是github.com/govaluate的8到10倍，前者性能远远低于后者。当Parse的内容越长，性能差距越明显，BenchmarkFullParse将这个差距扩大到了200多倍，BenchmarkComplexExpression将这差距扩大到60多倍。然而，当Parse和Evaluate操作同时进行时，性能差距又明显缩小了。这表明，性能差距主要在Parse操作上。

以上数据体现了ANTLR4的性能短板，语法树创建的对象很多，申请的内存很大，整个解析过程是耗时的。

### 优化思路

尽量少调用Parse，复用Expression对象，用对象池复用已经申请过的内存。

## code generator
基于ddl生成对应的golang struct

### usage

1. go install
2. kit --help

Usage of generating code depending on ddl:
```cmd
  -ddl-file string
        ddl file path, is Required
  -entity-output string
        entity struct output dir
  -model-output string
        model struct output dir
  -repo-output string
        repository file output dir
  -repo-tpl string
        repository template file
```

example command:  
```cmd
kit -ddl-file /Users/peidongxu/kit/tmp/sql \  
    -entity-output /Users/peidongxu/home/kit/entity \  
    -model-output /Users/peidongxu/home/kit/model/tm \  
    -repo-tpl /Users/peidongxu/kit/tmp/tpl \  
    -repo-output /Users/peidongxu/home/
```
