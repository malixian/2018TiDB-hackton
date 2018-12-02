package client

import (
	"fmt"
	"encoding/json"
    "github.com/pkg/errors"
	"net/http"
	"net/url"
)

const SAVE_URL = "http://127.0.0.1:8001/savePlan"

var P_Match PlanMatch

type PlanMatch struct {
	SQL       string `json:"sql"`
	PrePlan   string `json:"pre"`
	FinalPlan string `json:"final"`
}

func NewPlanMatch(Sql string) *PlanMatch {
	P_Match = PlanMatch{
		SQL: Sql,
	}
	return &P_Match
}

func (match *PlanMatch) SetSql(sql string) {
	match.SQL = sql

}

func (match *PlanMatch) SetPrePlanData(planData string) {
	match.PrePlan = planData
}

func (match *PlanMatch) SetFinalPlanData(planData string) {
	match.FinalPlan = planData
}

func (match *PlanMatch) Send() error {
	client := http.Client{}
	bmatch, err := json.Marshal(match)
	if err != nil {
		return errors.New(fmt.Sprintf("send data failed, error is %v", err))
	}
	resp, err := client.PostForm(SAVE_URL, url.Values{"query": {string(bmatch)}})
	if err != nil || resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("response failed, err is %v", err))
	}
	return nil
}
