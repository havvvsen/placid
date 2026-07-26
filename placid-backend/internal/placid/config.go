package placid

type Secret struct {
	JwtSecretKey string
}

type EndPoint struct {
	Register              string
	Login                 string
	User                  string
	DeleteAccount         string
	Tracks                string
	SubscribeNewsletter   string
	UnsubscribeNewsletter string
	AdminUploadTrack      string
	AdminDeleteTrack      string
}

type ApiConfig struct {
	ApiHost    string
	ApiPort    string
	ApiVersion string
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
	Secrets            *Secret
	PostgresConfig     *PostgresConfig
}
