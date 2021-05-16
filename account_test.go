package helium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountList(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	accounts, err := account.List(NoQuery)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 100, len(accounts.Data))
}

func TestAccountRichest(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	accounts, err := account.Richest(5)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 5, len(accounts.Data))
}

func TestAccountGet(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	accounts, err := account.Get("13WRNw4fmssJBvMqMnREwe1eCvUVXfnWXSXGcWXyVvAnQUF3D9R")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "13WRNw4fmssJBvMqMnREwe1eCvUVXfnWXSXGcWXyVvAnQUF3D9R", accounts.Data.Address)
}

func TestAccountHotspots(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	hotspots, err := account.Hotspots("13WRNw4fmssJBvMqMnREwe1eCvUVXfnWXSXGcWXyVvAnQUF3D9R")
	if err != nil {
		t.Error(err)
	}
	assert.Greater(t, len(hotspots.Data), 0)
}

func TestAccountOuis(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	ouis, err := account.Ouis("13tyMLKRFYURNBQqLSqNJg9k41maP1A7Bh8QYxR13oWv7EnFooc")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(ouis.Data))
}

func TestAccountActivity(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	activity, err := account.Activity("13GCcF7oGb6waFBzYDMmydmXx4vNDUZGX4LE3QUh8eSBG53s5bx")
	if err != nil {
		t.Error(err)
	}
	assert.Greater(t, len(activity.Data), 0)
}

func TestAccountActivityCount(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	activityCount, err := account.ActivityCount("13GCcF7oGb6waFBzYDMmydmXx4vNDUZGX4LE3QUh8eSBG53s5bx")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, activityCount)
}

func TestAccountElections(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	elections, err := account.Elections("146MwmL9eJJCdrykbgdL3dobdChP4Ut34mCZMR3Hv9HXTeBJQzC")
	if err != nil {
		t.Error(err)
	}
	t.Log(elections.Data)
	assert.Greater(t, len(elections.Data), 0)
}

func TestAccountChallenges(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	activityCount, err := account.Challenges("146MwmL9eJJCdrykbgdL3dobdChP4Ut34mCZMR3Hv9HXTeBJQzC")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, activityCount)
}

func TestAccountPendingTransactions(t *testing.T) {
	client := DefaultClient()
	account := client.Account()
	activityCount, err := account.ActivityCount("13GCcF7oGb6waFBzYDMmydmXx4vNDUZGX4LE3QUh8eSBG53s5bx")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, activityCount)
}
