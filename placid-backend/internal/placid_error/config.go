package placiderror

import "errors"

var ErrAllowedCORSOriginsRequired error = errors.New("$ALLOWED_CORS_ORIGINS is required")

// API CONFIG
var ErrApiHostRequired error = errors.New("$API_HOST is required")
var ErrApiPortRequired error = errors.New("$API_PORT is required")
var ErrApiVersionRequired error = errors.New("$API_VERSION is required")

// REDIS CONFIG
var ErrRedisHostRequired error = errors.New("$REDIS_HOST is required")
var ErrRedisPortRequired error = errors.New("$REDIS_PORT is required")
var ErrRedisUsernameRequired error = errors.New("$REDIS_USERNAME is required")
var ErrRedisPasswordRequired error = errors.New("$REDIS_PASSWORD is required")

// POSTGRES CONFIG
var ErrPostgresHostRequired error = errors.New("$POSTGRES_HOST is required")
var ErrPostgresPortRequired error = errors.New("$POSTGRES_PORT is required")
var ErrPostgresUsernameRequired error = errors.New("$POSTGRES_USERNAME is required")
var ErrPostgresPasswordRequired error = errors.New("$POSTGRES_PASSWORD is required")
var ErrPostgresDbRequired error = errors.New("$POSTGRES_DB is required")

// SECRET KEYS
var ErrJwtSecretKeyRequired error = errors.New("$JWT_SECRET_KEY is required")
