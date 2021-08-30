package model

type HistoryConfig struct {
	Name          string             `json:"name"`
	ConfigVersion []*AllocatedConfig `json:"configVersions"`
}

var histories = make(map[string]*HistoryConfig)

func putHistory(config *AllocatedConfig) *HistoryConfig {
	v, ok := histories[config.Name]
	if !ok {
		v = &HistoryConfig{
			Name: config.Name,
		}
	}
	v.ConfigVersion = append(v.ConfigVersion, config)
	histories[config.Name] = v
	return v
}

func GetHistory(name string) *HistoryConfig {
	v, _ := histories[name]
	return v
}
