type DynamoDBConfig struct {
	Region          string
	EndpointURL     string
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

func NewDynamoDBClient(config DynamoDBConfig) *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	if config.Region != "" {
		cfg.Region = config.Region
	}

	if config.EndpointURL != "" {
		cfg.EndpointResolver = aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: config.EndpointURL,
			}, nil
		})
	}

	if config.AccessKeyID != "" && config.SecretAccessKey != "" {
		cfg.Credentials = aws.NewStaticCredentialsProvider(config.AccessKeyID, config.SecretAccessKey, config.SessionToken)
	}

	return dynamodb.NewFromConfig(cfg)
}
