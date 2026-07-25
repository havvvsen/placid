package config

import (
	"fmt"
	"os"
	"placid-backend/internal/placid"
	placiderror "placid-backend/internal/placid_error"
	"strings"
)

func FillEnv() (*placid.Config, error) {
	// TODO: Understand why creating a pointer to an instance of placid.Config won't
	// automatically create pointers to nested structs

	cfg := &placid.Config{
		ApiConfig:      &placid.ApiConfig{},
		Endpoints:      &placid.EndPoint{},
		RedisConfig:    &placid.RedisConfig{},
		PostgresConfig: &placid.PostgresConfig{},
	}

	var isSet bool

	// ALLOWED CORS ORIGINS
	allowedCORSOrigins, isSet := os.LookupEnv("ALLOWED_CORS_ORIGINS")

	if !isSet {
		return nil, placiderror.ErrApiPortRequired
	}

	cfg.AllowedCORSOrigins = strings.Split(allowedCORSOrigins, " ")

	// API CONFIG
	cfg.ApiConfig.ApiHost, isSet = os.LookupEnv("API_HOST")
	if !isSet {
		return nil, placiderror.ErrApiHostRequired
	}
	cfg.ApiConfig.ApiPort, isSet = os.LookupEnv("API_PORT")
	if !isSet {
		return nil, placiderror.ErrApiPortRequired
	}
	cfg.ApiConfig.ApiVersion, isSet = os.LookupEnv("API_VERSION")
	if !isSet {
		return nil, placiderror.ErrApiVersionRequired
	}

	// ENDPOINTS
	cfg.Endpoints.Register = fmt.Sprintf("/api/%s/register", cfg.ApiConfig.ApiVersion)
	cfg.Endpoints.Login = fmt.Sprintf("/api/%s/login", cfg.ApiConfig.ApiVersion)
	cfg.Endpoints.User = fmt.Sprintf("/api/%s/user", cfg.ApiConfig.ApiVersion)
	cfg.Endpoints.DeleteAccount = fmt.Sprintf("/api/%s/delete-account", cfg.ApiConfig.ApiVersion)
	cfg.Endpoints.Tracks = fmt.Sprintf("/api/%s/tracks", cfg.ApiConfig.ApiVersion)
	cfg.Endpoints.AddNewsletterSubscriber = fmt.Sprintf("/api/%s/join-newsletter", cfg.ApiConfig.ApiVersion)
	cfg.Endpoints.DeleteNewsletterSubscriber = fmt.Sprintf("/api/%s/exit-newsletter", cfg.ApiConfig.ApiVersion)
	cfg.Endpoints.AdminUploadTrack = fmt.Sprintf("/api/%s/admin/upload-tracks", cfg.ApiConfig.ApiVersion)
	cfg.Endpoints.AdminDeleteTrack = fmt.Sprintf("/api/%s/admin/delete-tracks", cfg.ApiConfig.ApiVersion)

	// REDIS CONFIG
	cfg.RedisConfig.RedisHost, isSet = os.LookupEnv("REDIS_HOST")
	if !isSet {
		return nil, placiderror.ErrRedisHostRequired
	}
	cfg.RedisConfig.RedisPort, isSet = os.LookupEnv("REDIS_PORT")
	if !isSet {
		return nil, placiderror.ErrRedisPortRequired
	}
	cfg.RedisConfig.RedisUsername, isSet = os.LookupEnv("REDIS_USERNAME")
	if !isSet {
		return nil, placiderror.ErrRedisUsernameRequired
	}
	cfg.RedisConfig.RedisPassword, isSet = os.LookupEnv("REDIS_PASSWORD")
	if !isSet {
		return nil, placiderror.ErrRedisPasswordRequired
	}

	// POSTGRES CONFIG
	cfg.PostgresConfig.PostgresHost, isSet = os.LookupEnv("POSTGRES_HOST")
	if !isSet {
		return nil, placiderror.ErrPostgresHostRequired
	}
	cfg.PostgresConfig.PostgresPort, isSet = os.LookupEnv("POSTGRES_PORT")
	if !isSet {
		return nil, placiderror.ErrPostgresPortRequired
	}
	cfg.PostgresConfig.PostgresUsername, isSet = os.LookupEnv("POSTGRES_USER")
	if !isSet {
		return nil, placiderror.ErrPostgresUsernameRequired
	}
	cfg.PostgresConfig.PostgresPassword, isSet = os.LookupEnv("POSTGRES_PASSWORD")
	if !isSet {
		return nil, placiderror.ErrPostgresPasswordRequired
	}
	cfg.PostgresConfig.PostgresDb, isSet = os.LookupEnv("POSTGRES_DB")
	if !isSet {
		return nil, placiderror.ErrPostgresDbRequired
	}

	return cfg, nil
}
