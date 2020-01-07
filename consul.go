package main

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

type meta struct {
	URL       string `json:"url"`
	RouteType int    `json:"route_type"`
	Remove    bool   `json:"remove"`
}

// DisableService 服务禁用
// 判断consul中是否已经禁用此服务, 如果没有禁用，则修改服务标志为禁用
// 如果已经禁用，则不再重复赋值
func (m monImplement) DisableService(name string) (bool, error) {
	hasDisable := false

	path := fmt.Sprintf("tio/v1/gateway/services/%s", name)
	logrus.Debugf("Disable %s ", path)

	val, _, err := m.consulCli.KV().Get(path, nil)
	if err != nil {
		return hasDisable, err
	}

	if val == nil {
		return hasDisable, nil
	}

	var mta meta
	json.Unmarshal(val.Value, &mta)
	logrus.Debugf("%s value: %v", name, mta)

	if mta.Remove {
		return !hasDisable, nil
	}

	mta.Remove = true

	content, _ := json.Marshal(mta)

	_, err = m.consulCli.KV().Put(&api.KVPair{
		Key:   val.Key,
		Value: content,
	}, nil)

	return hasDisable, err
}
