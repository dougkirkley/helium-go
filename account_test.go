package helium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountList(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	accounts, err := account.List(&AccountListInput{})
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 100, len(accounts.Data))
}

func TestAccountRichest(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountRichestInput{
		Limit: 5,
	}
	accounts, err := account.Richest(input)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 5, len(accounts.Data))
}

func TestAccountGet(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountInput{
		ID: "13WRNw4fmssJBvMqMnREwe1eCvUVXfnWXSXGcWXyVvAnQUF3D9R",
	}
	accounts, err := account.Get(input)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "13WRNw4fmssJBvMqMnREwe1eCvUVXfnWXSXGcWXyVvAnQUF3D9R", accounts.Data.Address)
}

func TestAccountHotspots(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountInput{
		ID: "13WRNw4fmssJBvMqMnREwe1eCvUVXfnWXSXGcWXyVvAnQUF3D9R",
	}
	hotspots, err := account.Hotspots(input)
	if err != nil {
		t.Error(err)
	}
	assert.Greater(t, len(hotspots.Data), 0)
}

func TestAccountOuis(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountInput{
		ID: "13tyMLKRFYURNBQqLSqNJg9k41maP1A7Bh8QYxR13oWv7EnFooc",
	}
	ouis, err := account.Ouis(input)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(ouis.Data))
}

func TestAccountActivity(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountInput{
		ID: "13GCcF7oGb6waFBzYDMmydmXx4vNDUZGX4LE3QUh8eSBG53s5bx",
	}
	activity, err := account.Activity(input)
	if err != nil {
		t.Error(err)
	}
	assert.Greater(t, len(activity.Data), 0)
}

func TestAccountActivityCount(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountInput{
		ID: "13GCcF7oGb6waFBzYDMmydmXx4vNDUZGX4LE3QUh8eSBG53s5bx",
	}
	activityCount, err := account.ActivityCount(input)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, activityCount)
}

func TestAccountElections(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountInput{
		ID: "146MwmL9eJJCdrykbgdL3dobdChP4Ut34mCZMR3Hv9HXTeBJQzC",
	}
	elections, err := account.Elections(input)
	if err != nil {
		t.Error(err)
	}
	t.Log(elections.Data)
	assert.Equal(t, len(elections.Data), 0)
}

func TestAccountChallenges(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountInput{
		ID: "146MwmL9eJJCdrykbgdL3dobdChP4Ut34mCZMR3Hv9HXTeBJQzC",
	}
	activityCount, err := account.Challenges(input)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, activityCount)
}

func TestAccountPendingTransactions(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	input := &AccountInput{
		ID: "13GCcF7oGb6waFBzYDMmydmXx4vNDUZGX4LE3QUh8eSBG53s5bx",
	}
	activityCount, err := account.ActivityCount(input)
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, activityCount)
}
