package tencent

func NewTencentOssStore() *TencentOssStore {
	return &TencentOssStore{}
}

type TencentOssStore struct {
}

func (t *TencentOssStore) Upload(bucketName, objectKey, fileName string) error {
	return nil
}
