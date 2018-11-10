package model

// db describes interface of database needed by API
// to communicate with it
type db interface {
	GetAd(adID int64) (*AdItem, error)
	GetAds(limit int, offset int) ([]*AdItem, error)
	GetUserWithID(userID int64) (*User, error)
	GetUserWithEmail(email string) (*User, error)
	NewUser(user *User) (int64, error)
	NewAd(ad *AdItem) (int64, error)
	EditUser(user *User) (int64, error)
	EditAd(ad *AdItem) (int64, error)
	RemoveUser(userID int64) (int64, error)
	RemoveAd(adID int64) (int64, error)
}