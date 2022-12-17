package aws

func NewAwsOssStore() *AwsOssStore {
	return &AwsOssStore{}
}

type AwsOssStore struct {
}

func (t *AwsOssStore) Upload(bucketName, objectKey, fileName string) error {
	return nil
}
