package jobRunner

import "log"

func removeExistingService(serviceName string) {
	if existingService, i := serviceExists(serviceName); existingService != nil {

		log.Println("Removing " + serviceName)
		for _, e := range existingService.JobIds {
			c.Remove(e)
		}
		serviceList = append(serviceList[:i], serviceList[i+1:]...) // remove Old element

	}
}

func serviceExists(name string) (*ServiceJobs, int) {
	for i, service := range serviceList {
		if service.Name == name {
			return &service, i
		}
	}
	return nil, -1
}
