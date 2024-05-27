package config

import "errors"

var (
	// ErrInvalidHost invalid host
	ErrInvalidHost = errors.New("db host is required")

	// ErrInvalidPort invalid port
	ErrInvalidPort = errors.New("db port is required")

	// ErrInvalidUser invalid user
	ErrInvalidUser = errors.New("db user is required")

	// ErrInvalidPassword invalid password
	ErrInvalidPassword = errors.New("db password is required")

	// ErrInvalidDbName invalid db name
	ErrInvalidDbName = errors.New("db name is required")

	// ErrInvalidDBPath invalid db path
	ErrInvalidDBPath = errors.New("sqlite db path is required")

	// ErrInvalidMaxIdleConns invalid max idle conns
	ErrInvalidMaxIdleConns = errors.New("max idle conns is required")

	// ErrInvalidMaxOpenConns invalid max open conns
	ErrInvalidMaxOpenConns = errors.New("max open conns is required")

	// ErrInvalidLogLevel invalid log level
	ErrInvalidLogLevel = errors.New("log level is required")

	// ErrInvalidLogFormat invalid log format
	ErrInvalidLogFormat = errors.New("log format is required")

	// ErrInvalidHttpPort invalid http port
	ErrInvalidHttpPort = errors.New("http port is required")

	// ErrInvalidGrpcPort invalid grpc port
	ErrInvalidGrpcPort = errors.New("grpc port is required")

	// ErrInvalidGraphqlPort invalid graphql port
	ErrInvalidGraphqlPort = errors.New("graphql port is required")

	// ErrInvalidHttpHost invalid http host
	ErrInvalidHttpHost = errors.New("http host is required")

	// ErrInvalidGrpcHost invalid grpc host
	ErrInvalidGrpcHost = errors.New("grpc host is required")

	// ErrInvalidGraphqlHost invalid graphql host
	ErrInvalidGraphqlHost = errors.New("graphql host is required")

	// ErrInvalidSslMode invalid ssl mode
	ErrInvalidSslMode = errors.New("ssl mode is required")

	// ErrInvalidDriverName invalid driver name
	ErrInvalidDriverName = errors.New("driver name is required")

	// ErrInvalidDbSid invalid db sid
	ErrInvalidDbSid = errors.New("oracle db sid is required")
)
