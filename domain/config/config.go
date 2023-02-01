package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	App      AppObj      `required:"true" envconfig:"APP"`
	Db       DbObj       `required:"true" envconfig:"Db"`
	Firebase FirebaseObj `required:"true" envconfig:"FIREBASE"`
	Sendgrid SendgridObj `required:"true" envconfig:"SENDGRID"`
	Slack    SlackObj    `required:"true" envconfig:"SLACK"`
	Google   GoogleObj   `required:"true" envconfig:"GOOGLE"`
}

func New() (Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return Config{}, err
	}

	// Set global variables
	App = AppObj{
		Env:            c.App.Env,
		Service:        c.App.Service,
		BatchType:      c.App.BatchType,
		Port:           c.App.Port,
		BasicUsers:     c.App.BasicUsers,
		BasicPasswords: c.App.BasicPasswords,
		CorsDomains:    c.App.CorsDomains,
		LogFilePath:    c.App.LogFilePath,
	}

	Db = DbObj{
		Host:               c.Db.Host,
		Port:               c.Db.Port,
		Name:               c.Db.Name,
		User:               c.Db.User,
		Pass:               c.Db.Pass,
		InstanceUnixSocket: c.Db.InstanceUnixSocket,
	}

	Firebase = FirebaseObj{
		JsonFilePath: c.Firebase.JsonFilePath,
		WebApiKey:    c.Firebase.WebApiKey,
	}

	Sendgrid = SendgridObj{
		ApiKey: c.Sendgrid.ApiKey,
	}

	// SlackVar = Slack{
	// 	WebhookURL: c.Slack.WebhookURL,
	// }

	// GoogleVar = Google{
	// 	ClientID:     c.Google.ClientID,
	// 	ClientSecret: c.Google.ClientSecret,
	// }
	// AppEnv = c.App.Env
	// AppService = c.App.Service
	// AppBatchType = c.App.BatchType
	// AppPort = c.App.Port
	// AppBasicUsers = c.App.BasicUsers
	// AppBasicPasswords = c.App.BasicPasswords
	// AppCorsDomains = c.App.CorsDomains
	// DbHost = c.Db.Host
	// DbPort = c.Db.Port
	// DbName = c.Db.Name
	// DbUser = c.Db.User
	// DbPass = c.Db.Pass
	// DbInstanceUnixSocket = c.Db.InstanceUnixSocket
	// FirebaseJsonFilePath = c.Firebase.JsonFilePath
	// FirebaseWebApiKey = c.Firebase.WebApiKey
	// SendgridApiKey = c.Sendgrid.ApiKey
	// SlackWebhookURL = c.Slack.WebhookURL
	// GoogleClientID = c.Google.ClientID
	// GoogleClientSecret = c.Google.ClientSecret

	return c, nil
}

type AppObj struct {
	Env            string   `required:"true" split_words:"true"`
	Service        string   `required:"true" split_words:"true"`
	BatchType      string   `required:"true" split_words:"true"`
	Port           int      `required:"true" split_words:"true"`
	BasicUsers     []string `required:"true" split_words:"true"`
	BasicPasswords []string `required:"true" split_words:"true"`
	CorsDomains    []string `required:"true" split_words:"true"`
	LogFilePath    string   `required:"true" split_words:"true"`
}

type DbObj struct {
	Host               string `required:"true" split_words:"true"`
	Port               int    `required:"true" split_words:"true"`
	Name               string `required:"true" split_words:"true"`
	User               string `required:"true" split_words:"true"`
	Pass               string `required:"true" split_words:"true"`
	InstanceUnixSocket string `required:"true" split_words:"true"`
}

type FirebaseObj struct {
	JsonFilePath string `required:"true" split_words:"true"`
	WebApiKey    string `required:"true" split_words:"true"`
}

type SendgridObj struct {
	ApiKey string `required:"true" split_words:"true"`
}

type SlackObj struct {
	AccessToken string `required:"true" split_words:"true"`
	ChanelID    string `required:"true" split_words:"true"`
}

type GoogleObj struct {
	ApplicationCredentials string `required:"true" split_words:"true"`
}

var (
	App      AppObj
	Db       DbObj
	Firebase FirebaseObj
	Sendgrid SendgridObj
	Slack    SlackObj
	Google   GoogleObj
)

// var (
// 	// AppEnv is the environment of the application.
// 	AppEnv string
// 	// AppService is the service name of the application.
// 	AppService string
// 	// AppBatchType is the batch type of the application.
// 	AppBatchType string
// 	// AppPort is the port of the application.
// 	AppPort int
// 	// AppBasicUsers is the basic users of the application.
// 	AppBasicUsers []string
// 	// AppBasicPasswords is the basic passwords of the application.
// 	AppBasicPasswords []string
// 	// AppCorsDomains is the cors domains of the application.
// 	AppCorsDomains []string
// 	// AppLogFilePath is the log file path of the application.
// 	AppLogFilePath string
// 	// DbHost is the host of the database.
// 	DbHost string
// 	// DbPort is the port of the database.
// 	DbPort int
// 	// DbName is the name of the database.
// 	DbName string
// 	// DbUser is the user of the database.
// 	DbUser string
// 	// DbPass is the password of the database.
// 	DbPass string
// 	// DbInstanceUnixSocket is the instance unix socket of the database.
// 	DbInstanceUnixSocket string
// 	// FirebaseJsonFilePath is the Json file path of the firebase.
// 	FirebaseJsonFilePath string
// 	// FirebaseWebApiKey is the web Api key of the firebase.
// 	FirebaseWebApiKey string
// 	// SendgridApiKey is the Api key of the sendgrid.
// 	SendgridApiKey string
// 	// SlackWebhookUrl is the webhook url of the slack.
// 	SlackWebhookUrl string
// 	// GoogleClientID is the client id of the google.
// 	GoogleClientID string
// 	// GoogleClientSecret is the client secret of the google.
// 	GoogleClientSecret string
// )
