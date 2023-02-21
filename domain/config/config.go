package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	App         AppObj         `required:"true" envconfig:"APP"`
	Db          DbObj          `required:"true" envconfig:"Db"`
	Firebase    FirebaseObj    `required:"true" envconfig:"FIREBASE"`
	Sendgrid    SendgridObj    `required:"true" envconfig:"SENDGRID"`
	Slack       SlackObj       `required:"true" envconfig:"SLACK"`
	Google      GoogleObj      `required:"true" envconfig:"GOOGLE"`
	ChatGpt     ChatGptObj     `required:"true" envconfig:"CHAT_GPT"`
	UberSuggest UberSuggestObj `required:"true" envconfig:"UBER_SUGGEST"`
}

var (
	App         AppObj
	Db          DbObj
	Firebase    FirebaseObj
	Sendgrid    SendgridObj
	Slack       SlackObj
	Google      GoogleObj
	ChatGpt     ChatGptObj
	UberSuggest UberSuggestObj
)

func New() error {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return err
	}

	// Set global variables
	App = NewAppObj(
		cfg.App.Env,
		cfg.App.Service,
		cfg.App.BatchType,
		cfg.App.Port,
		cfg.App.BasicUsers,
		cfg.App.BasicPasswords,
		cfg.App.BasicSecret,
		cfg.App.CorsDomains,
		cfg.App.ClientDomain,
		cfg.App.PythonDomain,
		cfg.App.LogFilePath,
	)

	Db = NewDbObj(
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name,
		cfg.Db.User,
		cfg.Db.Pass,
		cfg.Db.InstanceUnixSocket,
	)

	Firebase = NewFirebaseObj(
		cfg.Firebase.JsonFilePath,
		cfg.Firebase.WebApiKey,
	)

	Sendgrid = NewSendgridObj(
		cfg.Sendgrid.ApiKey)

	Slack = NewSlackObj(
		cfg.Slack.AccessToken,
		cfg.Slack.ChannelID,
	)

	Google = NewGoogleObj(
		cfg.Google.ApplicationCredentials,
	)

	ChatGpt = NewChatGptObj(
		cfg.ChatGpt.SecretKey,
	)

	UberSuggest = NewUberSuggestObj(
		cfg.UberSuggest.Email,
		cfg.UberSuggest.Password,
	)

	return nil
}

type AppObj struct {
	// AppEnv is the environment of the application.
	Env            string   `required:"true" split_words:"true"`
	// AppService is the service of the application.
	Service        string   `required:"true" split_words:"true"`
	// AppBatchType is the batch type of the application.
	BatchType      string   `required:"true" split_words:"true"`
	// AppPort is the port of the application.
	Port           int      `required:"true" split_words:"true"`
	// AppBasicUsers is the basic users of the application.
	BasicUsers     []string `required:"true" split_words:"true"`
	// AppBasicPasswords is the basic passwords of the application.
	BasicPasswords []string `required:"true" split_words:"true"`
	// AppBasicSecret is the basic secret of the application.
	BasicSecret    string   `required:"true" split_words:"true"`
	// AppCorsDomains is the cors domains of the application.
	CorsDomains    []string `required:"true" split_words:"true"`
	// AppClientDomain is the client domain of the application.
	ClientDomain   string   `required:"true" split_words:"true"`
	// AppPythonDomain is the python domain of the application.
	PythonDomain   string   `required:"true" split_words:"true"`
	// AppLogFilePath is the log file path of the application.
	LogFilePath    string   `required:"true" split_words:"true"`
}

func NewAppObj(
	env string,
	service string,
	batchType string,
	port int,
	basicUsers []string,
	basicPasswords []string,
	basicSecret string,
	corsDomains []string,
	clientDomain string,
	pythonDomain string,
	logFilePath string,
) AppObj {
	return AppObj{
		Env:            env,
		Service:        service,
		BatchType:      batchType,
		Port:           port,
		BasicUsers:     basicUsers,
		BasicPasswords: basicPasswords,
		CorsDomains:    corsDomains,
		ClientDomain:   clientDomain,
		PythonDomain:   pythonDomain,
		LogFilePath:    logFilePath,
	}
}

type DbObj struct {
	// DbHost is the host of the database.
	Host               string `required:"true" split_words:"true"`
	// DbPort is the port of the database.
	Port               int    `required:"true" split_words:"true"`
	// DbName is the name of the database.
	Name               string `required:"true" split_words:"true"`
	// DbUser is the user of the database.
	User               string `required:"true" split_words:"true"`
	// DbPass is the pass of the database.
	Pass               string `required:"true" split_words:"true"`
	// DbInstanceUnixSocket is the instance unix socket of the database.
	InstanceUnixSocket string `required:"true" split_words:"true"`
}

func NewDbObj(
	host string,
	port int,
	name string,
	user string,
	pass string,
	instanceUnixSocket string,
) DbObj {
	return DbObj{
		Host:               host,
		Port:               port,
		Name:               name,
		User:               user,
		Pass:               pass,
		InstanceUnixSocket: instanceUnixSocket,
	}
}

type FirebaseObj struct {
	// FirebaseJsonFilePath is the json file path of the firebase.
	JsonFilePath string `required:"true" split_words:"true"`
	// FirebaseWebApiKey is the web api key of the firebase.
	WebApiKey    string `required:"true" split_words:"true"`
}

func NewFirebaseObj(
	jsonFilePath string,
	webApiKey string,
) FirebaseObj {
	return FirebaseObj{
		JsonFilePath: jsonFilePath,
		WebApiKey:    webApiKey,
	}
}

type SendgridObj struct {
	// SendgridApiKey is the api key of the sendgrid.
	ApiKey string `required:"true" split_words:"true"`
}

func NewSendgridObj(
	apiKey string,
) SendgridObj {
	return SendgridObj{
		ApiKey: apiKey,
	}
}

type SlackObj struct {
	// SlackAccessToken is the access token of the slack.
	AccessToken string `required:"true" split_words:"true"`
	// SlackChannelID is the channel id of the slack.
	ChannelID   string `required:"true" split_words:"true"`
}

func NewSlackObj(
	accessToken string,
	channelID string,
) SlackObj {
	return SlackObj{
		AccessToken: accessToken,
		ChannelID:   channelID,
	}
}

type GoogleObj struct {
	// GoogleApplicationCredentials is the application credentials of the google.
	ApplicationCredentials string `required:"true" split_words:"true"`
}

func NewGoogleObj(
	applicationCredentials string,
) GoogleObj {
	return GoogleObj{
		ApplicationCredentials: applicationCredentials,
	}
}

type ChatGptObj struct {
	// ChatGptSecretKey is the secret key of the chat gpt.
	SecretKey string `required:"true" split_words:"true"`
}

func NewChatGptObj(
	secretKey string,
) ChatGptObj {
	return ChatGptObj{
		SecretKey: secretKey,
	}
}

type UberSuggestObj struct {
	// UberSuggestEmail is the email of the uber suggest.
	Email    string `required:"true" split_words:"true"`
	// UberSuggestPassword is the password of the uber suggest.
	Password string `required:"true" split_words:"true"`
}

func NewUberSuggestObj(
	email string,
	password string,
) UberSuggestObj {
	return UberSuggestObj{
		Email:    email,
		Password: password,
	}
}