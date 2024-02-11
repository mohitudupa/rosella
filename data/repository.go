package data

type Config map[string]any

type Repository interface {
	Connect() error
	Close() error

	ListFlags(groups string) ([]string, error)
	GetFlag(groups string, key string) (bool, error)
	SetFlag(groups string, key string, value bool) error
	DeleteFlag(groups string, key string) error

	ListLimits(groups string) ([]string, error)
	GetLimit(groups string, key string) (float64, error)
	SetLimit(groups string, key string, value float64) error
	DeleteLimit(groups string, key string) error

	ListValues(groups string) ([]string, error)
	GetValue(groups string, key string) (string, error)
	SetValue(groups string, key string, value string) error
	DeleteValue(groups string, key string) error

	ListConfigs(groups string) ([]string, error)
	GetConfig(groups string, key string) (Config, error)
	SetConfig(groups string, key string, value Config) error
	DeleteConfig(groups string, key string) error
}
