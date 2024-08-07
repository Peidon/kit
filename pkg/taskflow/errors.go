/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2024/7/17 3:15 PM
 * @Version: kit
 */

package taskflow

import "errors"

var (
	TimeoutWithContext = errors.New("timeout with context. ")
	DeadLockError      = errors.New("task flow dead lock. ")
	UnexpectedStopped  = errors.New("flow unexpected stopped. ")
	ValueNotExists     = errors.New("value not exists")
	// GraphCircuitError  = errors.New("graph circuit error")
)
