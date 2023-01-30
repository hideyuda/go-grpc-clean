package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	App      App      `required:"true" envconfig:"APP"`
	Db       Db       `required:"true" envconfig:"Db"`
	Firebase Firebase `required:"true" envconfig:"FIREBASE"`
	Sendgrid Sendgrid `required:"true" envconfig:"SENDGRID"`
	Slack    Slack    `required:"true" envconfig:"SLACK"`
	Google   Google   `required:"true" envconfig:"GOOGLE"`
}

func New() (Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return Config{}, err
	}

	// Set global variables
	AppEnv = c.App.Env
	AppService = c.App.Service
	AppBatchType = c.App.BatchType
	AppPort = c.App.Port
	AppBasicUsers = c.App.BasicUsers
	AppBasicPasswords = c.App.BasicPasswords
	AppCorsDomains = c.App.CorsDomains
	DbHost = c.Db.Host
	DbPort = c.Db.Port
	DbName = c.Db.Name
	DbUser = c.Db.User
	DbPass = c.Db.Pass
	FirebaseJsonFilePath = c.Firebase.JsonFilePath
	FirebaseWebApiKey = c.Firebase.WebApiKey
	SendgridApiKey = c.Sendgrid.ApiKey
	// SlackWebhookURL = c.Slack.WebhookURL
	// GoogleClientID = c.Google.ClientID
	// GoogleClientSecret = c.Google.ClientSecret

	return c, nil
}

type App struct {
	Env            string   `required:"true" split_words:"true"`
	Service        string   `required:"true" split_words:"true"`
	BatchType      string   `required:"true" split_words:"true"`
	Port           int      `required:"true" split_words:"true"`
	BasicUsers     []string `required:"true" split_words:"true"`
	BasicPasswords []string `required:"true" split_words:"true"`
	CorsDomains    []string `required:"true" split_words:"true"`
}

type Db struct {
	Host string `required:"true" split_words:"true"`
	Port int    `required:"true" split_words:"true"`
	Name string `required:"true" split_words:"true"`
	User string `required:"true" split_words:"true"`
	Pass string `required:"true" split_words:"true"`
}

type Firebase struct {
	JsonFilePath string `required:"true" split_words:"true"`
	WebApiKey    string `required:"true" split_words:"true"`
}

type Sendgrid struct {
	ApiKey string `required:"true" split_words:"true"`
}

type Slack struct {
	AccessToken string `required:"true" split_words:"true"`
	ChanelID    string `required:"true" split_words:"true"`
}

type Google struct {
	ApplicationCredentials string `required:"true" split_words:"true"`
}

var (
	// AppEnv is the environment of the application.
	AppEnv string
	// AppService is the service name of the application.
	AppService string
	// AppBatchType is the batch type of the application.
	AppBatchType string
	// AppPort is the port of the application.
	AppPort int
	// AppBasicUsers is the basic users of the application.
	AppBasicUsers []string
	// AppBasicPasswords is the basic passwords of the application.
	AppBasicPasswords []string
	// AppCorsDomains is the cors domains of the application.
	AppCorsDomains []string
	// DbHost is the host of the database.
	DbHost string
	// DbPort is the port of the database.
	DbPort int
	// DbName is the name of the database.
	DbName string
	// DbUser is the user of the database.
	DbUser string
	// DbPass is the password of the database.
	DbPass string
	// FirebaseJsonFilePath is the Json file path of the firebase.
	FirebaseJsonFilePath string
	// FirebaseWebApiKey is the web Api key of the firebase.
	FirebaseWebApiKey string
	// SendgridApiKey is the Api key of the sendgrid.
	SendgridApiKey string
	// SlackWebhookUrl is the webhook url of the slack.
	SlackWebhookUrl string
	// GoogleClientID is the client id of the google.
	GoogleClientID string
	// GoogleClientSecret is the client secret of the google.
	GoogleClientSecret string
)
