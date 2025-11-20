package main

import (
	"github.com/google/uuid"
	"sync"
)

var singletonMetricsServiceMu sync.Mutex
var singletonMetricsService *MetricsService

func GetMetricsService() *MetricsService {
	if singletonMetricsService == nil {
		singletonMetricsServiceMu.Lock()

		if singletonMetricsService == nil {
			singletonMetricsService = newMetricsService()
		}

		singletonMetricsServiceMu.Unlock()
	}

	return singletonMetricsService
}

type MetricsService struct {
	id string
}

func newMetricsService() *MetricsService {
	return &MetricsService{
		id: uuid.New().String(),
	}
}

func (m MetricsService) Id() string {
	return m.id
}
