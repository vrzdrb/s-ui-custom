package cronjob

import (
	"github.com/vrzdrb/s-ui-custom/logger"
	"github.com/vrzdrb/s-ui-custom/service"
)

type StatsJob struct {
	service.StatsService
	enableTraffic bool
}

func NewStatsJob(saveTraffic bool) *StatsJob {
	return &StatsJob{
		enableTraffic: saveTraffic,
	}
}

func (s *StatsJob) Run() {
	err := s.StatsService.SaveStats(s.enableTraffic)
	if err != nil {
		logger.Warning("Get stats failed: ", err)
		return
	}
}
