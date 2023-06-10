package main

import (
	"fmt"
	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

const rule = `
rule "deposit" "deposit"  salience 0
begin
		if deposit.GetDepositTime()==1 && deposit.GetDepositAmout()>100{
			return true
		}
		else{
			return false
		}
end
`

type Deposit struct {
	DepositTime  int64
	DepositAmout int64
}

func (f *Deposit) GetDepositTime() int64 {
	return f.DepositTime
}

func (f *Deposit) GetDepositAmout() int64 {
	return f.DepositAmout
}

func (f *Deposit) Print(s string) {
	fmt.Println(s)
}

func main() {
	//初始化数据环境变量
	deposit := &Deposit{
		DepositTime:  1,
		DepositAmout: 150,
	}
	dataContext := context.NewDataContext()
	//注入println函数
	dataContext.Add("deposit", deposit)
	//初始化规则
	ruleBuilder := builder.NewRuleBuilder(dataContext)
	//读取规则
	errBuild := ruleBuilder.BuildRuleFromString(rule)
	if errBuild != nil {
		fmt.Println(errBuild)
	}
	//初始化规则引擎
	eng := engine.NewGengine()
	//执行规则引擎
	errExex := eng.Execute(ruleBuilder, true)
	if errExex != nil {
		fmt.Println(errExex)
	}
	resultMap, _ := eng.GetRulesResultMap()
	r := resultMap["deposit"]
	println("return--->", r.(bool))

}
