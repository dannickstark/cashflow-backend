package hooks

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
)

// A function to bind all the hooks for the users
func BindUsersHooks(app *pocketbase.PocketBase) {
	app.OnRecordAfterCreateRequest("users").Add(func(e *core.RecordCreateEvent) error {
		// Get the user ID from the event
		userID := e.Record.Id

		// Call the CreateUserSetting function to create a new Setting record for the user
		_, err := CreateUserSetting(app, userID)
		if err != nil {
			return err
		}

		return nil
	})
}

// Create a new Setting record for the user
func CreateUserSetting(app *pocketbase.PocketBase, userID string) (*models.Record, error) {
	collection, err := app.Dao().FindCollectionByNameOrId("settings")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	record := models.NewRecord(collection)
	form := forms.NewRecordUpsert(app, record)

	settingData := map[string]any{
		"defaultCurrency": "USD",
		"user":            userID,
	}

	form.LoadData(settingData)

	if err := form.Submit(); err != nil {
		return nil, err
	}
	return record, nil
}
