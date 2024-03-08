# kit
Golang Coder 的框架和工具
1. valuate
2. code generator

## valuate

基于ANTLR4的表达式计算框架

特点
--

* 支持自定义函数;
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

	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = -1;

	result, err := expression.Evaluate(parameters);
	// result is now set to "false", the bool value.
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

测试环境
--

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
```

性能对比
--

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

分析
--

从BenchmarkSingleParse和BenchmarkSimpleParse的对比中可以看出，taskflow/common/valuate的操作耗时、内存占用、内存申请次数分别是github.com/govaluate的8到10倍，前者性能远远低于后者。当Parse的内容越长，性能差距越明显，BenchmarkFullParse将这个差距扩大到了200多倍，BenchmarkComplexExpression将这差距扩大到60多倍。然而，当Parse和Evaluate操作同时进行时，性能差距又明显缩小了。这表明，性能差距主要在Parse操作上。

以上数据体现了ANTLR4的性能短板，语法树创建的对象很多，申请的内存很大，整个解析过程是耗时的。

优化思路
--

尽量少调用Parse，复用Expression对象，用对象池复用已经申请过的内存。
## code generator
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
