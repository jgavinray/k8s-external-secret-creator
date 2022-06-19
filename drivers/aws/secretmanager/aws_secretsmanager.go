package secretmanager

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/jgavinrary/k8s-external-secrets-creator/drivers"
)

type SecretsManagerService struct {
	// Name of the key to be stored
	Name      string `json:"name"`
	AccessKey string
}

// init function to self register against the provider interface
func init() {
	drivers.RegisterDriver("aws_secretsmanager", func() drivers.Provider {
		p, err := New()
		if err != nil {
			return nil
		}
		return p
	})
}

func New() (*SecretsManagerService, error) {
	return &SecretsManagerService{}, nil
}

func (s SecretsManagerService) LoadConfiguration() string {
	value, ok := os.LookupEnv("accessKey")
	if ok {
		lastFive := value[len(value)-5:]
		log.Log.Info(fmt.Sprintf("Using Access Key that ends in '***************%s'.", lastFive))
		return value
	}
	log.Log.Info("Unable to locate an access key, returning empty string")
	return ""

}

// Put Parameter takes in a key/value pair of strings and creates a new secret in AWS Secret Manager
func (s SecretsManagerService) PutParameter(key string, value string) error {
	awsRegion := "us-east-1"

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(awsRegion)},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		return err
	}

	secretManagerConnection := secretsmanager.New(sess, aws.NewConfig().WithRegion(awsRegion))

	secretInput := &secretsmanager.CreateSecretInput{
		Name:         aws.String(key),
		SecretString: aws.String(value),
		Description:  aws.String("Created by k8s-external-secrets-creator"),
	}

	result, err := secretManagerConnection.CreateSecret(secretInput)

	if err != nil {
		// The AWS SDK will return aws specific errors, but if something just outright fails
		// because of this, we will attempt to cast the error type.  If the error is of type
		// awserr, then we create a log message with the details of the error and bubble up.
		// If the type casting fails, then we return the error as is.
		// https://stackoverflow.com/questions/18416042/golang-type-conversion-not-working-as-i-expected
		awsErr, ok := err.(awserr.Error)
		if ok {
			s := fmt.Sprintf("AWS specific error %s", awsErr.Code())
			return errors.New(s)
		}
		return err
	}

	response := fmt.Sprintf("Secret %s version %s was created with ARN value of %s", *result.Name, *result.VersionId, *result.ARN)
	log.Log.Info(response)
	return nil

}
