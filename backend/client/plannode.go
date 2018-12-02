package client

import (
	"encoding/json"
	"github.com/pingcap/tidb/planCollector"
	"github.com/pingcap/tidb/planner/core"
	"log"
)

const DEFAULT_VALUE = 100

type PlanNode struct {
	ChildrenNode []*PlanNode `json:"children"`
	Info         string      `json:"name"`
	Value        int         `json:"value"`
}

func savePrePlanTree(plan *core.PhysicalPlan) {
	head_pre := PlanNode{}
	GetOptimizetree(plan, &head_pre)
	spre := TestGetPlanTree(head_pre, "pre")
	planCollector.P_Match.PrePlan = spre
}

func saveFinalPlanTree(plan *core.PhysicalPlan) error {
	head_pre := PlanNode{}
	GetOptimizetree(plan, &head_pre)
	spre := TestGetPlanTree(head_pre, "final")
	planCollector.P_Match.FinalPlan = spre
	return planCollector.P_Match.Send()
}

func GetOptimizetree(phycailPlan *core.PhysicalPlan, currentNode *PlanNode) {
	if phycailPlan == nil {
		return
	}
	currentNode.Info = (*phycailPlan).ExplainID()
	currentNode.Value = DEFAULT_VALUE
	currentNode.ChildrenNode = []*PlanNode{}
	//log.Printf("current explain id is %s", currentNode.Info)
	for _, node := range (*phycailPlan).Children() {
		childrenNode := PlanNode{Info: node.ExplainID()}
		// append 添加的是副本，如果要改动添加的数据需要使用指针
		currentNode.ChildrenNode = append(currentNode.ChildrenNode, &childrenNode)
		GetOptimizetree(&node, &childrenNode)
	}
}

func TestGetPlanTree(head PlanNode, stage string) string {
	res := map[string][]PlanNode{"children": {head}}
	bres, err := json.Marshal(res)
	if err != nil {
		log.Printf("=============marshal head data failed, err is %v=========", err)
		return ""
	}

	//log.Printf("=============>stage is %s, plan data is %v<============", stage, string(bres))
	return string(bres)
}
