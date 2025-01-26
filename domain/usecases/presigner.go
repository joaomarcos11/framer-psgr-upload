package usecases

type Presigner interface {
	GeneratePreSignedUrl(bucket, key string) (string, error)
}
