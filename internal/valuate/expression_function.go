/*
 * @Copyright @ Shopee
 * @Author: Peidong Xu
 * @Email: peidong.xu@shopee.com
 * @Date: 2023/11/10 5:29 PM
 * @Version: kit
 */

package valuate

type ExpressionFunction func(arguments ...interface{}) (interface{}, error)
