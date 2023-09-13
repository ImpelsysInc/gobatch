package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

func init() {
	LoadEnv()
}

const (
	STATUS_COMPLETED   = "COMPLETED"
	STATUS_PENDING     = "PENDING"
	STATUS_IN_PROGRESS = "IN-PROGRESS"
	STATUS_FAILED      = "FAILED"
)

type Conf struct {
	Debug      bool   `env:"DEBUG,required"`
	DebugLevel string `env:"DEBUG_LEVEL"`
	Server     serverConf
	Db         dbConf
	Log        logConf
	App        appConf
	Gateway    gatewayConf
	Aws        awsConf
}

type serverConf struct {
	Port         string        `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

type logConf struct {
	LogFilePath string `env:"Log_FILE_PATH"`
	LogFileName string `env:"LOG_FILE_NAME"`
}

type dbConf struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DbName   string `env:"DB_NAME,required"`
}

type appConf struct {
	BaseURL                      string `env:"APP_BASE_URL"`
	Lang                         string `env:"APP_LANG"`
	Name                         string `env:"SERVICE_NAME"`
	Step1                        string `env:"STEP1_NAME"`
	Step2                        string `env:"STEP2_NAME"`
	Step3                        string `env:"STEP3_NAME"`
	Step4                        string `env:"STEP4_NAME"`
	CEJobName                    string `env:"CE_JOB_NAME"`
	IntelligoSubscriptionJobName string `env:"INTELLIGO_SUBSCRIPTION_JOB_NAME"`
	IntelligoSurveyJobName       string `env:"INTELLIGO_SURVEY_JOB_NAME"`
	CEJobEnabled                 bool   `env:"CE_JOB_ENABLED"`
	SurveyJobEnabled             bool   `env:"SURVEY_JOB_ENABLED"`
	SubscriptionJobEnabled       bool   `env:"USER_SUBSCRIPTION_JOB_ENABLED"`
	QuestionJobEnabled           bool   `env:"QUESTION_JOB_ENABLED"`
	IntelligoQuestionJobName     string `env:"INTELLIGO_QUESTION_JOB_NAME"`
	BatchSize                    uint   `env:"BATCH_SIZE"`
	CEBatchSize                  uint   `env:"CE_BATCH_SIZE"`
	MaxRunningJobs               int    `env:"MAX_RUNNING_JOBS"`
	MaxRunningSteps              int    `env:"MAX_RUNNING_STEPS"`
	RootDir                      string
}

type gatewayConf struct {
	URL    string `env:"API_GATEWAY_URL"`
	Prefix string `env:"API_GATEWAY_PREFIX"`
}

type awsConf struct {
	AwsRegion             string `env:"AWS_REGION"`
	AwsAccessKeyID        string `env:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey    string `env:"AWS_SECRET_ACCESS_KEY"`
	AwsTempBucketName     string `env:"AWS_S3_TEMP_BUCKET_NAME"`
	AwsFolderPrefix       string `env:"AWS_S3_FOLDER_PREFIX"`
	AwsUrlExpiry          int    `env:"AWS_S3_URL_EXPIRY"`
	ContentDeliveryURL    string `env:"CONTENT_DELIVERY_URL"`
	AwsDeliveryBucketName string `env:"AWS_S3_DELIVERY_BUCKET_NAME"`
}

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)

	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func AppConfig() *Conf {
	var c Conf

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	if len(c.App.RootDir) <= 0 {
		c.App.RootDir = GetRootDir()
	}

	if len(c.App.Lang) <= 0 {
		c.App.Lang = "en-US"
	}

	if len(c.App.BaseURL) <= 0 {
		c.App.BaseURL = "api/v1"
	}

	if len(c.Log.LogFilePath) <= 0 {
		c.Log.LogFilePath = c.App.RootDir + "/log"
	}

	if len(c.Log.LogFileName) <= 0 {
		c.Log.LogFileName = "micro.log"
	}

	if len(c.App.Name) <= 0 {
		c.App.Name = "MicroService"
	}

	return &c
}

func LoadEnv() {

	dir, err := os.Getwd()
	path := dir + "/.env"

	if err != nil {
		log.Println("Not able to get current working director")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		dir = filepath.Dir(dir)
		path = dir + "/.env"
	}

	if err := godotenv.Load(path); err != nil {
		log.Println("No .env file found in path " + path)
	}
}
