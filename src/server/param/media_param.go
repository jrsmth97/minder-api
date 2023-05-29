package param

type UploadFile struct {
	File string `validate:"required"`
}
