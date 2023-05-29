package repository

import "minder/src/server/model"

type UserRepo interface {
	Create(user *model.User) error
	Explore() (*[]model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(id string, user *model.User) error
	Delete(id string) error
	CountUsers() (*int, error)
}

type UserPhotoRepo interface {
	Create(user *model.UserPhoto) error
	FindById(id string) (*model.UserPhoto, error)
	FindByUserId(userId string) (*[]model.UserPhoto, error)
	Delete(id string) error
}

type UserMembershipRepo interface {
	Create(user *model.UserMembership) error
	FindById(id string) (*model.UserMembership, error)
	FindByUserId(userId string) (*model.UserMembership, error)
}

type UserSwipeRepo interface {
	Create(user *model.UserSwipe) error
	FindById(id string) (*model.UserSwipe, error)
	FindSwipes(userId string) (*[]model.UserSwipe, error)
	FindSwipesToday(userId string) (*[]model.UserSwipe, error)
	FindLike(userId string) (*[]model.UserSwipe, error)
	FindLikeToday(userId string) (*[]model.UserSwipe, error)
	FindPass(userId string) (*[]model.UserSwipe, error)
	FindPassToday(userId string) (*[]model.UserSwipe, error)
	FindFavourite(userId string) (*[]model.UserSwipe, error)
	FindFavouriteToday(userId string) (*[]model.UserSwipe, error)
}

type InterestRepo interface {
	GetInterests() (*[]model.Interest, error)
	CreateInterest(m *model.Interest) error
	UpdateInterest(id string, m *model.Interest) error
	DeleteInterest(id string) error
	GetInterestById(id string) (*model.Interest, error)
}

type LocationRepo interface {
	GetLocations() (*[]model.Location, error)
	CreateLocation(m *model.Location) error
	UpdateLocation(id string, m *model.Location) error
	DeleteLocation(id string, m *model.Location) error
	GetLocationById(id string) (*model.Location, error)
}

type MembershipRepo interface {
	GetMemberships() (*[]model.Membership, error)
	GetMembershipById(id string) (*model.Membership, error)
	GetMembershipByName(name string) (*model.Membership, error)
	CreateMembership(m *model.Membership) error
	UpdateMembership(id string, m *model.Membership) error
	DeleteMembership(id string) error
}

type PrivilegeRepo interface {
	GetPrivileges() (*[]model.Privilege, error)
	CreatePrivilege(m *model.Privilege) error
	UpdatePrivilege(id string, m *model.Privilege) error
	DeletePrivilege(id string) error
	GetPrivilegeById(id string) (*model.Privilege, error)
	GetPrivilegeByBulkId(ids []string) (*[]model.Privilege, error)
}

type MembershipPrivilegeRepo interface {
	GetMembershipPrivileges() (*[]model.MembershipPrivilege, error)
	CreateMembershipPrivilege(m *model.MembershipPrivilege) error
	UpdateMembershipPrivilege(id string, m *model.MembershipPrivilege) error
	DeleteMembershipPrivilege(id string) error
	DeleteMembershipPrivilegeByMemberId(id string) error
	DeleteMembershipPrivilegeByPrivilegeId(id string) error
	GetMembershipPrivilegeById(id string) (*model.MembershipPrivilege, error)
	GetMembershipPrivilegeByMemberId(memberId string) (*[]model.MembershipPrivilege, error)
	GetMembershipPrivilegeByBulkMemberId(bulkId []string) (*[]model.MembershipPrivilege, error)
}

type PurchaseRepo interface {
	GetAllPurchases() (*[]model.Purchase, error)
	CreatePurchase(m *model.Purchase) error
	GetPendingPurchases() (*[]model.Purchase, error)
	GetLastPurchaseByDate(date string) (*model.Purchase, error)
	GetFailedPurchases() (*[]model.Purchase, error)
	GetSuccessPurchases() (*[]model.Purchase, error)
	GetPurchaseById(id string) (*model.Purchase, error)
	UpdatePurchase(id string, m *model.Purchase) error
}
