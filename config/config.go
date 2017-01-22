package config

import (
	"github.com/ferrariframework/ferrariserver/services/job"
)

//JobService Configures a new instance of a job service
func JobService() job.Service {
	return job.New()
}
