package jobRunner

import (
	"encoding/json"
	"goBill/models"
	"goBill/util/httpRequestor"
	"goBill/util/redis"
	"log"
	"strconv"

	cron "gopkg.in/robfig/cron.v2"
)

var c = cron.New()

//ServiceJobs holds the Ids for the tasks that are running for the service
type ServiceJobs struct {
	Name   string
	JobIds []cron.EntryID
}

var serviceList []ServiceJobs

//InitialiseJobs - loads up the current jobs from Redis and starts a schedule for them.
func InitialiseJobs() error {
	services, err := redis.NewRedis().HGetAllMap("Bill-Service").Result()
	if err != nil {
		return err
	}

	defer c.Start()
	go loadCurrentService(services)
	go monitorNewService()
	return nil
}

func loadCurrentService(services map[string]string) error {
	var requestService models.RegisterService
	for _, service := range services {
		err := json.Unmarshal([]byte(service), &requestService)
		if err != nil {
			return err
		}
		go loadServicesJobs(requestService)
	}
	return nil
}

func loadServicesJobs(requestService models.RegisterService) chan error {
	var jobs []models.Job
	httpRequestor.GetBody(requestService.BaseURI+"/jobs", &jobs)

	go removeExistingService(requestService.ServiceName)

	result := make(chan error)
	go func() {
		result <- func() error {

			jobList := make([]cron.EntryID, len(jobs))
			for i, job := range jobs {
				newJob := job

				log.Println("Loading up Job " + newJob.Name)
				entry, err := c.AddFunc("@every "+strconv.Itoa(newJob.Timeout)+"ms", func() { checkJob(newJob, requestService.ServiceName) })

				if err != nil {
					return err
				}
				jobList[i] = entry
			}
			serviceList = append(serviceList, ServiceJobs{Name: requestService.ServiceName, JobIds: jobList})
			return nil
		}()
	}()
	return result
}
