package parameters

type Service struct {
	Http     Http     `yaml:"http" mapstructure:"http" json:"http"`
	Grpc     Grpc     `yaml:"grpc" mapstructure:"grpc" json:"grpc"`
	Database Database `yaml:"db" mapstructure:"db" json:"db"`
}

type Http struct {
	Port int    `yaml:"port" mapstructure:"port" json:"port"`
	Host string `yaml:"host" mapstructure:"host" json:"host"`
}

type Grpc struct {
	Port int    `yaml:"port" mapstructure:"port" json:"port"`
	Host string `yaml:"host" mapstructure:"host" json:"host"`
}

type Database struct {
	Name     string `yaml:"name" mapstructure:"name" json:"name"`
	Host     string `yaml:"host" mapstructure:"host" json:"host"`
	Port     int    `yaml:"port" mapstructure:"port" json:"port"`
	User     string `yaml:"user" mapstructure:"user" json:"user"`
	Password string `yaml:"password" mapstructure:"password" json:"password"`
}
