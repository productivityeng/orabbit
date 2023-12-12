// Package entities provides the test cases for the UserEntity struct in the user package.

package entities

import (
	"testing"
	"time"

	entities2 "github.com/productivityeng/orabbit/locker/entities"
	"github.com/stretchr/testify/assert"
)

// TestIsLockedShouldReturnTrueWhenAtLeastOneLockerHasDisabledAtNil tests the IsLocked method of the UserEntity struct.
// It verifies that the IsLocked method returns true when at least one locker has DisabledAt set to nil.
func TestIsLockedShouldReturnTrueWhenAtLeastOneLockerHasDisabledAtNil(t *testing.T) {
	lockerArray := [5]entities2.LockerEntity{{
		DisabledAt: &time.Time{},
	}, {DisabledAt: &time.Time{}}, {DisabledAt: nil}}

	user := UserEntity{
		Lockers: lockerArray[:],
	}

	isLocked := user.IsLocked()
	assert.True(t, isLocked)
}

// TestIsLockedShouldReturnFalseWhenAllLockersHaveDisabledAtNotNil tests the IsLocked method of the UserEntity struct.
// It verifies that the IsLocked method returns false when all lockers have DisabledAt set to a non-nil value.
func TestIsLockedShouldReturnFalseWhenAllLockersHaveDisabledAtNotNil(t *testing.T) {
	lockerArray := [5]entities2.LockerEntity{{
		DisabledAt: &time.Time{},
	}, {DisabledAt: &time.Time{}}, {DisabledAt: &time.Time{}}, {DisabledAt: &time.Time{}}, {DisabledAt: &time.Time{}}}

	user := UserEntity{
		Lockers: lockerArray[:],
	}

	isLocked := user.IsLocked()
	assert.False(t, isLocked)
}


// ...

// TestUserInListByNameShouldReturnUserEntityWhenUsernameExists tests the UserInListByName function of the UserEntityList type.
// It verifies that the function returns the UserEntity when the username exists in the list.
func TestUserInListByNameShouldReturnUserEntityWhenUsernameExists(t *testing.T) {
	list := UserEntityList{
		{Username: "user1"},
		{Username: "user2"},
		{Username: "user3"},
	}

	username := "user2"
	user := list.UserInListByName(username)

	assert.NotNil(t, user)
	assert.Equal(t, username, user.Username)
}

// TestUserInListByNameShouldReturnNilWhenUsernameDoesNotExist tests the UserInListByName function of the UserEntityList type.
// It verifies that the function returns nil when the username does not exist in the list.
func TestUserInListByNameShouldReturnNilWhenUsernameDoesNotExist(t *testing.T) {
	list := UserEntityList{
		{Username: "user1"},
		{Username: "user2"},
		{Username: "user3"},
	}

	username := "user4"
	user := list.UserInListByName(username)

	assert.Nil(t, user)
}

// ...

