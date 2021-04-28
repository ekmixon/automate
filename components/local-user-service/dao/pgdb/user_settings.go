package pgdb

import (
	"database/sql"
	"encoding/json"

	usersettings "github.com/chef/automate/api/external/user_settings"

	"github.com/chef/automate/api/external/lib/errorutils"

	"github.com/pkg/errors"
)

const uniqueViolation = "23505"

type UserSettings struct {
	ID            string          `db:"id"`
	Settings      json.RawMessage `db:"-"`
	ConnectorType string          `db:"conn_type"`
	Enabled       bool            `db:"enabled"`
}

func (db *DB) DeleteUserSettings(settings *usersettings.DeleteUserSettingsRequest) error {
	var count int64 = 0
	var err error
	err = Transact(db, func(tx *DBTrans) error {
		count, err = tx.Delete(&UserSettings{ID: settings.Id})
		if count == 0 {
			return errorutils.ProcessSQLNotFound(sql.ErrNoRows, settings.Id, "DeleteUserSettings")
		}
		if err != nil {
			return errorutils.ProcessSQLNotFound(err, settings.Id, "DeleteUserSettings")
		}
		return nil
	})
	return err
}

func (db *DB) UpdateUserSettings(settings *usersettings.UpdateUserSettingsRequest) error {
	dbUserSettings := updateToUserSetting(settings)
	var err error
	var count int64 = 0
	err = Transact(db, func(tx *DBTrans) error {
		if count, err = tx.Update(dbUserSettings); err != nil {
			return errors.Wrap(err, "UpdateUserSetting: unable to update")
		}
		if count == 0 {
			return errorutils.ProcessSQLNotFound(sql.ErrNoRows, settings.Id, "UpdateUserSettings")
		}
		return nil
	})
	return err
}

func updateToUserSetting(settingsReq *usersettings.UpdateUserSettingsRequest) *UserSettings {
	userSettings := UserSettings{}
	userSettings.ID = settingsReq.Id

	return &userSettings
}

func dbToGetUserSettingsResponse(settings *UserSettings) *usersettings.GetUserSettingsResponse {
	userSettingsResp := usersettings.GetUserSettingsResponse{}
	userSettingsResp.Id = settings.ID

	return &userSettingsResp
}

func (db *DB) GetUserSettings(settings *usersettings.GetUserSettingsRequest) (*usersettings.GetUserSettingsResponse, error) {
	var err error
	var obj interface{}
	var userSettings *UserSettings
	err = Transact(db, func(tx *DBTrans) error {
		if obj, err = tx.Get(UserSettings{}, settings.Id); err != nil {
			return errors.Wrap(err, "GetDestination: unable to get destination")
		}
		if obj == nil {
			userSettings = &UserSettings{}
			err = errorutils.ProcessSQLNotFound(sql.ErrNoRows, settings.Id, "GetUserSettings")
		} else {
			userSettings = obj.(*UserSettings)
		}
		return err
	})
	result := dbToGetUserSettingsResponse(userSettings)
	if err != nil {
		return result, err
	}
	return result, err
}
