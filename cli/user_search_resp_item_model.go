// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for UserSearchRespItem

// register flags to command
func registerModelUserSearchRespItemFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserSearchRespItemUserID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserSearchRespItemUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserSearchRespItemUserID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	userIdDescription := `The ID of the user.`

	var userIdFlagName string
	if cmdPrefix == "" {
		userIdFlagName = "user_id"
	} else {
		userIdFlagName = fmt.Sprintf("%v.user_id", cmdPrefix)
	}

	var userIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(userIdFlagName, userIdFlagDefault, userIdDescription)

	return nil
}

func registerUserSearchRespItemUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	usernameDescription := ``

	var usernameFlagName string
	if cmdPrefix == "" {
		usernameFlagName = "username"
	} else {
		usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
	}

	var usernameFlagDefault string

	_ = cmd.PersistentFlags().String(usernameFlagName, usernameFlagDefault, usernameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserSearchRespItemFlags(depth int, m *models.UserSearchRespItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, userIdAdded := retrieveUserSearchRespItemUserIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || userIdAdded

	err, usernameAdded := retrieveUserSearchRespItemUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usernameAdded

	return nil, retAdded
}

func retrieveUserSearchRespItemUserIDFlags(depth int, m *models.UserSearchRespItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	userIdFlagName := fmt.Sprintf("%v.user_id", cmdPrefix)
	if cmd.Flags().Changed(userIdFlagName) {

		var userIdFlagName string
		if cmdPrefix == "" {
			userIdFlagName = "user_id"
		} else {
			userIdFlagName = fmt.Sprintf("%v.user_id", cmdPrefix)
		}

		userIdFlagValue, err := cmd.Flags().GetInt64(userIdFlagName)
		if err != nil {
			return err, false
		}
		m.UserID = userIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserSearchRespItemUsernameFlags(depth int, m *models.UserSearchRespItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usernameFlagName := fmt.Sprintf("%v.username", cmdPrefix)
	if cmd.Flags().Changed(usernameFlagName) {

		var usernameFlagName string
		if cmdPrefix == "" {
			usernameFlagName = "username"
		} else {
			usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
		}

		usernameFlagValue, err := cmd.Flags().GetString(usernameFlagName)
		if err != nil {
			return err, false
		}
		m.Username = usernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}