package credentials

type Wukong struct {
	Token string
}

func NewWukong(token string) *Credentials {
	return New(&Wukong{Token: token})
}

func (w *Wukong) Retrieve() (Value, error) {
	return Value{
		AccessKeyID: w.Token,
		SignerType:  SignatureWukong,
	}, nil
}

func (w *Wukong) IsExpired() bool {
	return false
}
