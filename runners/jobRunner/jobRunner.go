package jobRunner

import (
	"encoding/json"
	"goBill/models"
	"goBill/util/redis"
	"log"
	"strconv"
	"time"

	"github.com/parnurzeal/gorequest"
)

var redisClient = redis.NewRedis()

func checkJob(job models.Job, serviceName string) {

	r := getRequest(job)

	log.Println("Checking " + job.Name)
	var now = time.Now()
	resp, _, err := r.End()
	responseTime := time.Since(now).Nanoseconds()
	errorMessage := ""
	status := ""
	if err != nil {
		for _, e := range err {
			errorMessage += e.Error()
		}
	}

	if resp == nil {
		status = "Service Unavailable"
	} else {
		status = resp.Status
	}
	result := models.JobResult{StatusCode: status, Started: now, ResponseTime: responseTime, Error: errorMessage}
	marshelledService, _ := json.Marshal(result)
	redisClient.LPush("billResult-"+serviceName+"-"+job.Name, string(marshelledService))
}

func interfaceToString(value interface{}) string {
	switch str := value.(type) {
	case string:
		return str
	case int:
		return strconv.Itoa(str)
	case float64:
		return strconv.FormatFloat(str, 'f', -1, 64)
	}
	return ""
}

func getRequest(job models.Job) *gorequest.SuperAgent {
	var request = gorequest.New()
	var r *gorequest.SuperAgent

	if job.HTTPVerb == gorequest.GET {
		r = request.Get(job.URI)
	}

	if job.HTTPVerb == gorequest.POST {
		r = request.Post(job.URI)
	}
	r.Timeout(time.Second * 20)

	for key, header := range job.Headers {
		r.Set(key, interfaceToString(header))
	}
	return r
}
