package utils

import (
	apps "k8s.io/api/apps/v1"
)

type StatefulRollout struct {
}

func (s *StatefulRollout) GetStatefulSet() *apps.StatefulSet {
	return &apps.StatefulSet{}
}
