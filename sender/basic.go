package sender

// BaseSenderHandler the abstract sender, indicate the interface methods
// all sender handler should realize
type BaseSenderHandler interface {
	message()
	handler()
}