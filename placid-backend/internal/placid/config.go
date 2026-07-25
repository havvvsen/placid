package placid

type EndPoint struct {
	Register                   string
	Login                      string
	User                       string
	DeleteAccount              string
	Tracks                     string
	AddNewsletterSubscriber    string
	DeleteNewsletterSubscriber string
	AdminUploadTrack           string
	AdminDeleteTrack           string
}

type ApiConfig struct {
	ApiHost    string
	ApiPort    string
	ApiVersion string
}

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisUsername string
	RedisPassword string
}

type PostgresConfig struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUsername string
	PostgresPassword string
	PostgresDb       string
}

type Config struct {
	AllowedCORSOrigins []string
	Endpoints          *EndPoint
	ApiConfig          *ApiConfig
	RedisConfig        *RedisConfig
	PostgresConfig     *PostgresConfig
}
