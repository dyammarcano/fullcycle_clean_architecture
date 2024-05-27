package config

import (
	"fmt"
	"strings"
)

type NamedDriver string

const (
	PostgresNamedDriver NamedDriver = "postgres"
	MySQLNamedDriver    NamedDriver = "mysql"
	SQLiteNamedDriver   NamedDriver = "sqlite"
	OracleNamedDriver   NamedDriver = "oracle"
	BoltNamedDriver     NamedDriver = "bolt"
)

type LogFormat string

const (
	JSONLogFormat LogFormat = "json"
	TextLogFormat LogFormat = "text"
)

type LevelName string

const (
	LevelDebug LevelName = "debug"
	LevelInfo  LevelName = "info"
	LevelError LevelName = "error"
)

var G *Service

func init() {
	G = &Service{
		Http: Http{
			Port: 8080,
			Host: "localhost",
		},
		Grpc: Grpc{
			Port: 50051,
			Host: "localhost",
		},
		Logger: Logger{
			LogLevel:  LevelInfo,
			LogFormat: JSONLogFormat,
		},
		Db: Db{
			Driver:       SQLiteNamedDriver,
			Sslmode:      "disable",
			Sid:          "xe",
			MaxIdleConns: 2,
			MaxOpenConns: 10,
		},
	}
}

type Service struct {
	Http   Http   `yaml:"http" mapstructure:"http" json:"http"`
	Grpc   Grpc   `yaml:"grpc" mapstructure:"grpc" json:"grpc"`
	Logger Logger `yaml:"logger" mapstructure:"logger" json:"logger"`
	Db     Db     `yaml:"db" mapstructure:"db" json:"db"`
}

type Http struct {
	Port int    `yaml:"port" mapstructure:"port" json:"port"`
	Host string `yaml:"host" mapstructure:"host" json:"host"`
}

type Grpc struct {
	Port int    `yaml:"port" mapstructure:"port" json:"port"`
	Host string `yaml:"host" mapstructure:"host" json:"host"`
}

type Logger struct {
	LogLevel  LevelName `yaml:"logLevel" mapstructure:"logLevel" json:"logLevel"`
	LogFormat LogFormat `yaml:"logFormat" mapstructure:"logFormat" json:"logFormat"`
}

type Db struct {
	Driver       NamedDriver `yaml:"driver" mapstructure:"driver" json:"driver"`
	Sid          string      `yaml:"sid" mapstructure:"sid" json:"sid"`
	Name         string      `yaml:"name" mapstructure:"name" json:"name"`
	Host         string      `yaml:"host" mapstructure:"host" json:"host"`
	Port         int         `yaml:"port" mapstructure:"port" json:"port"`
	User         string      `yaml:"user" mapstructure:"user" json:"user"`
	Password     string      `yaml:"password" mapstructure:"password" json:"password"`
	Sslmode      string      `yaml:"sslmode" mapstructure:"sslmode" json:"sslmode"`
	Dbname       string      `yaml:"dbName" mapstructure:"dbName" json:"dbName"`
	MaxIdleConns int         `yaml:"maxIdleConns" mapstructure:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns int         `yaml:"maxOpenConns" mapstructure:"maxOpenConns" json:"maxOpenConns"`
	DBPath       string      `yaml:"dbPath" mapstructure:"dbPath" json:"dbPath"`
}

type OptsFunc func(*Service)

// WithDriver sets database driver, e.g. postgres, mysql, sqlite, oracle
//
// Default is sqlite
func WithDriver(driver NamedDriver) OptsFunc {
	return func(o *Service) {
		o.Db.Driver = driver
	}
}

// WithUser sets database user for connection, default is empty
func WithUser(user string) OptsFunc {
	return func(o *Service) {
		o.Db.User = user
	}
}

// WithPassword sets database password for connection, default is empty
func WithPassword(password string) OptsFunc {
	return func(o *Service) {
		o.Db.Password = password
	}
}

// WithHost sets database host for connection, default is empty
func WithHost(host string) OptsFunc {
	return func(o *Service) {
		o.Db.Host = host
	}
}

// WithPort sets database port, default is empty
func WithPort(port int) OptsFunc {
	return func(o *Service) {
		o.Db.Port = port
	}
}

// WithDatabase sets database name for connection, default is empty
func WithDatabase(database string) OptsFunc {
	return func(o *Service) {
		o.Db.Dbname = database
	}
}

// WithSID sets oracle db sid, default is xe
func WithSID(sid string) OptsFunc {
	return func(o *Service) {
		o.Db.Sid = sid
	}
}

// WithMaxIdleConns sets max idle connections for database, default is 2
func WithMaxIdleConns(maxIdleConns int) OptsFunc {
	return func(o *Service) {
		o.Db.MaxIdleConns = maxIdleConns
	}
}

// WithMaxOpenConns sets max open connections for database, default is 10
func WithMaxOpenConns(maxOpenConns int) OptsFunc {
	return func(o *Service) {
		o.Db.MaxOpenConns = maxOpenConns
	}
}

// WithLogLevel sets log level, e.g. info, debug, error
func WithLogLevel(logLevel LevelName) OptsFunc {
	return func(o *Service) {
		o.Logger.LogLevel = logLevel
	}
}

// WithLogFormat sets log format
func WithLogFormat(logFormat LogFormat) OptsFunc {
	return func(o *Service) {
		o.Logger.LogFormat = logFormat
	}
}

// WithHttpPort sets http port
func WithHttpPort(port int) OptsFunc {
	return func(o *Service) {
		o.Http.Port = port
	}
}

// WithGrpcPort sets grpc port
func WithGrpcPort(port int) OptsFunc {
	return func(o *Service) {
		o.Grpc.Port = port
	}
}

// WithHttpHost sets http host
func WithHttpHost(host string) OptsFunc {
	return func(o *Service) {
		o.Http.Host = host
	}
}

// WithGrpcHost sets grpc host
func WithGrpcHost(host string) OptsFunc {
	return func(o *Service) {
		o.Grpc.Host = host
	}
}

// WithSslMode sets ssl mode for database connection, except for sqlite or bolt database
func WithSslMode(sslMode string) OptsFunc {
	return func(o *Service) {
		o.Db.Sslmode = sslMode
	}
}

// WithBoltDB sets bolt db name
func WithBoltDB(name, path string) OptsFunc {
	if !strings.HasSuffix(path, ".db") {
		path = fmt.Sprintf("%s.db", path)
	}
	return func(o *Service) {
		o.Db.Dbname = name
		o.Db.DBPath = path
	}
}

// WithSqliteDB sets sqlite db path name
func WithSqliteDB(name, path string) OptsFunc {
	if !strings.HasSuffix(path, ".sqlite") {
		path = fmt.Sprintf("%s.sqlite", path)
	}
	return func(o *Service) {
		o.Db.Dbname = name
		o.Db.DBPath = path
	}
}

// validateCommonDatabaseValues validates common database values
func validateCommonDatabaseValues(db *Db) error {
	if db.Host == "" {
		return ErrInvalidHost
	}

	if db.Port == 0 {
		return ErrInvalidPort
	}

	if db.User == "" {
		return ErrInvalidUser
	}

	if db.Password == "" {
		return ErrInvalidPassword
	}

	if db.Dbname == "" {
		return ErrInvalidDbName
	}

	return nil
}

// Validate validates database configuration
func (s *Service) Validate() error {
	if s.Db.Driver == "" {
		return ErrInvalidDriverName
	}

	if s.Db.Driver == SQLiteNamedDriver || s.Db.Driver == BoltNamedDriver {
		if s.Db.DBPath == "" {
			return ErrInvalidDBPath
		}
	}

	if s.Db.Driver == PostgresNamedDriver || s.Db.Driver == MySQLNamedDriver {
		if err := validateCommonDatabaseValues(&s.Db); err != nil {
			return err
		}
	}

	if s.Db.Driver == OracleNamedDriver {
		if err := validateCommonDatabaseValues(&s.Db); err != nil {
			return err
		}

		if s.Db.Sid == "" {
			return ErrInvalidDbSid
		}
	}

	if err := validateLoggerValues(&s.Logger); err != nil {
		return err
	}

	return nil
}

func validateLoggerValues(logger *Logger) error {
	if logger.LogLevel == "" {
		return ErrInvalidLogLevel
	}

	if logger.LogFormat == "" {
		return ErrInvalidLogFormat
	}

	return nil
}

// NewConfig creates a new service configuration
func NewConfig(opts ...OptsFunc) *Service {
	for _, fn := range opts {
		fn(G)
	}

	return G
}
