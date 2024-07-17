/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2024/7/17 2:59 PM
 * @Version: kit
 */

package taskflow

import (
	"context"
	"runtime"
)

func LogStackBuf(errMsg string, r interface{}) {
	const size = 64 << 10
	stackBuf := make([]byte, size)
	stackBuf = stackBuf[:runtime.Stack(stackBuf, false)]
	println(errMsg, r, "\n", string(stackBuf))
}

func withRecover(ctx context.Context, fn func(ctx context.Context)) {
	subCtx, cancel := context.WithCancel(ctx)
	defer func() {
		cancel()
		if r := recover(); r != nil {
			LogStackBuf("[flow running panic]", r)
		}
	}()
	fn(subCtx)
}
