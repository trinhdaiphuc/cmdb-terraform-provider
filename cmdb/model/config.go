package model

import "time"

// AllocatedConfig holds metadata of a name allocated to a deployed resource
// within one of our theoretical deployment environments
type AllocatedConfig struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

var configs = make(map[string]*AllocatedConfig)

func PutAllocatedConfig(name, value string) *AllocatedConfig {
	v, ok := configs[name]
	if ok {
		v.Value = value
		v.UpdatedAt = time.Now().Format(time.RFC3339)
	} else {
		v = &AllocatedConfig{
			Name:      name,
			Value:     value,
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		}
	}
	putHistory(v)
	configs[name] = v
	return v
}

func DeleteAllocatedConfig(name string) {
	if _, ok := configs[name]; ok {
		delete(configs, name)
		delete(histories, name)
	}
}
func GetAllocatedConfig(name string) *AllocatedConfig {
	v, _ := configs[name]
	return v
}
