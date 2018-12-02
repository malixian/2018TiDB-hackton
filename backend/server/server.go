package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

type PlanMatch struct {
	SQL       string `json:"sql"`
	PrePlan   string `json:"pre"`
	FinalPlan string `json:"final"`
}

var conn *redis.Client

func SavePlan(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	planData := r.PostForm.Get("query")
	pMatch := PlanMatch{}
	err := json.Unmarshal([]byte(planData), &pMatch)
	if err != nil {
		log.Printf("unmarshal failed")
	}
	log.Printf("query is %v", planData)
	value := map[string]interface{}{
		"pre_plan":   pMatch.PrePlan,
		"final_plan": pMatch.FinalPlan,
	}
	_, err = conn.HMSet(pMatch.SQL, value).Result()
	if err != nil {
		log.Printf("hash set value failed, error is %v", err)
	}
	log.Printf("save success")
	fmt.Fprintf(w, "success!") //这个写入到w的是输出到客户端的

}

func GetPlan(w http.ResponseWriter, r *http.Request) {
	key := r.Form.Get("key")
	if key == "" {
		fmt.Fprintf(w, "key is invaild")
		return
	}
	res, err := conn.HMGet(key).Result()
	if err != nil {
		log.Printf("get value failed, error is %v", err)
		return
	}
	prePlan := res[0].(string)
	finalPlan := res[1].(string)
	pm := PlanMatch{
		SQL:       key,
		PrePlan:   prePlan,
		FinalPlan: finalPlan,
	}
	bpm, err := json.Marshal(pm)
	if err != nil {
		log.Printf("marshal data failed, error is %v", err)
		return
	}
	fmt.Fprintf(w, string(bpm))
}

func NewRedisClinet() *redis.Client {
	conn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return conn
}
