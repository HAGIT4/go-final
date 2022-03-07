package api

type BonusServerInterface interface {
	ListenAndServe() (err error)
}
