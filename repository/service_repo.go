package repository

import "github.com/costap/healthcheckapp/model"

type ServiceInfoRepository struct {
	services map[string]model.ServiceInfo
}

func NewServiceInfoRepository() *ServiceInfoRepository {
	return &ServiceInfoRepository{services: make(map[string]model.ServiceInfo)}
}

func (r *ServiceInfoRepository) Save(si model.ServiceInfo) {
	r.services[si.Name] = si
}

func (r *ServiceInfoRepository) List() []model.ServiceInfo {
	l := make([]model.ServiceInfo, 0)
	for _, v := range r.services {
		l = append(l, v)
	}
	return l
}
