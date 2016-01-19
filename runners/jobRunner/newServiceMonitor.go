package jobRunner

import (
	"encoding/json"
	"goBill/models"
	"goBill/util/redis"
)

func monitorNewService() chan error {
	result := make(chan error)
	go func() {
		result <- func() error {

			client := redis.NewRedis()
			client.ConfigSet("notify-keyspace-events", "KEA")
			sub, err := client.PSubscribe("__keyspace@0__:Bill-Service")

			defer sub.Close()
			if err != nil {
				return err
			}

			for {

				_, err := sub.ReceiveMessage()
				lastUpdated, _ := lastUpdatedService()
				serviceDetails, err := client.HGet("Bill-Service", lastUpdated).Result()
				if err != nil {
					return err
				}

				var requestService models.RegisterService

				if err = json.Unmarshal([]byte(serviceDetails), &requestService); err != nil {
					return err
				}
				go loadServicesJobs(requestService)
			}
		}()
	}()
	return result
}

func lastUpdatedService() (string, error) {
	serviceName, err := redis.NewRedis().Get("Bill-Service-LastUpdated").Result()

	if err != nil {
		return "", err
	}

	return serviceName, nil
}
