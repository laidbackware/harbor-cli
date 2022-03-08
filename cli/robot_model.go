// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for Robot

// register flags to command
func registerModelRobotFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRobotCreationTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotDisable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotDuration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotEditable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotExpiresAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotLevel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotPermissions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotSecret(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotUpdateTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRobotCreationTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	creationTimeDescription := `The creation time of the robot.`

	var creationTimeFlagName string
	if cmdPrefix == "" {
		creationTimeFlagName = "creation_time"
	} else {
		creationTimeFlagName = fmt.Sprintf("%v.creation_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(creationTimeFlagName, "", creationTimeDescription)

	return nil
}

func registerRobotDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	descriptionDescription := `The description of the robot`

	var descriptionFlagName string
	if cmdPrefix == "" {
		descriptionFlagName = "description"
	} else {
		descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
	}

	var descriptionFlagDefault string

	_ = cmd.PersistentFlags().String(descriptionFlagName, descriptionFlagDefault, descriptionDescription)

	return nil
}

func registerRobotDisable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	disableDescription := `The disable status of the robot`

	var disableFlagName string
	if cmdPrefix == "" {
		disableFlagName = "disable"
	} else {
		disableFlagName = fmt.Sprintf("%v.disable", cmdPrefix)
	}

	var disableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(disableFlagName, disableFlagDefault, disableDescription)

	return nil
}

func registerRobotDuration(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	durationDescription := `The duration of the robot in days`

	var durationFlagName string
	if cmdPrefix == "" {
		durationFlagName = "duration"
	} else {
		durationFlagName = fmt.Sprintf("%v.duration", cmdPrefix)
	}

	var durationFlagDefault int64

	_ = cmd.PersistentFlags().Int64(durationFlagName, durationFlagDefault, durationDescription)

	return nil
}

func registerRobotEditable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	editableDescription := `The editable status of the robot`

	var editableFlagName string
	if cmdPrefix == "" {
		editableFlagName = "editable"
	} else {
		editableFlagName = fmt.Sprintf("%v.editable", cmdPrefix)
	}

	var editableFlagDefault bool

	_ = cmd.PersistentFlags().Bool(editableFlagName, editableFlagDefault, editableDescription)

	return nil
}

func registerRobotExpiresAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	expiresAtDescription := `The expiration data of the robot`

	var expiresAtFlagName string
	if cmdPrefix == "" {
		expiresAtFlagName = "expires_at"
	} else {
		expiresAtFlagName = fmt.Sprintf("%v.expires_at", cmdPrefix)
	}

	var expiresAtFlagDefault int64

	_ = cmd.PersistentFlags().Int64(expiresAtFlagName, expiresAtFlagDefault, expiresAtDescription)

	return nil
}

func registerRobotID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of the robot`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

func registerRobotLevel(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	levelDescription := `The level of the robot, project or system`

	var levelFlagName string
	if cmdPrefix == "" {
		levelFlagName = "level"
	} else {
		levelFlagName = fmt.Sprintf("%v.level", cmdPrefix)
	}

	var levelFlagDefault string

	_ = cmd.PersistentFlags().String(levelFlagName, levelFlagDefault, levelDescription)

	return nil
}

func registerRobotName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	nameDescription := `The name of the tag`

	var nameFlagName string
	if cmdPrefix == "" {
		nameFlagName = "name"
	} else {
		nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
	}

	var nameFlagDefault string

	_ = cmd.PersistentFlags().String(nameFlagName, nameFlagDefault, nameDescription)

	return nil
}

func registerRobotPermissions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: permissions []*RobotPermission array type is not supported by go-swagger cli yet

	return nil
}

func registerRobotSecret(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	secretDescription := `The secret of the robot`

	var secretFlagName string
	if cmdPrefix == "" {
		secretFlagName = "secret"
	} else {
		secretFlagName = fmt.Sprintf("%v.secret", cmdPrefix)
	}

	var secretFlagDefault string

	_ = cmd.PersistentFlags().String(secretFlagName, secretFlagDefault, secretDescription)

	return nil
}

func registerRobotUpdateTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updateTimeDescription := `The update time of the robot.`

	var updateTimeFlagName string
	if cmdPrefix == "" {
		updateTimeFlagName = "update_time"
	} else {
		updateTimeFlagName = fmt.Sprintf("%v.update_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(updateTimeFlagName, "", updateTimeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRobotFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, creationTimeAdded := retrieveRobotCreationTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || creationTimeAdded

	err, descriptionAdded := retrieveRobotDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, disableAdded := retrieveRobotDisableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || disableAdded

	err, durationAdded := retrieveRobotDurationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || durationAdded

	err, editableAdded := retrieveRobotEditableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || editableAdded

	err, expiresAtAdded := retrieveRobotExpiresAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || expiresAtAdded

	err, idAdded := retrieveRobotIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, levelAdded := retrieveRobotLevelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || levelAdded

	err, nameAdded := retrieveRobotNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, permissionsAdded := retrieveRobotPermissionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || permissionsAdded

	err, secretAdded := retrieveRobotSecretFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secretAdded

	err, updateTimeAdded := retrieveRobotUpdateTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updateTimeAdded

	return nil, retAdded
}

func retrieveRobotCreationTimeFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	creationTimeFlagName := fmt.Sprintf("%v.creation_time", cmdPrefix)
	if cmd.Flags().Changed(creationTimeFlagName) {

		var creationTimeFlagName string
		if cmdPrefix == "" {
			creationTimeFlagName = "creation_time"
		} else {
			creationTimeFlagName = fmt.Sprintf("%v.creation_time", cmdPrefix)
		}

		creationTimeFlagValueStr, err := cmd.Flags().GetString(creationTimeFlagName)
		if err != nil {
			return err, false
		}
		var creationTimeFlagValue strfmt.DateTime
		if err := creationTimeFlagValue.UnmarshalText([]byte(creationTimeFlagValueStr)); err != nil {
			return err, false
		}
		m.CreationTime = creationTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotDescriptionFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	descriptionFlagName := fmt.Sprintf("%v.description", cmdPrefix)
	if cmd.Flags().Changed(descriptionFlagName) {

		var descriptionFlagName string
		if cmdPrefix == "" {
			descriptionFlagName = "description"
		} else {
			descriptionFlagName = fmt.Sprintf("%v.description", cmdPrefix)
		}

		descriptionFlagValue, err := cmd.Flags().GetString(descriptionFlagName)
		if err != nil {
			return err, false
		}
		m.Description = descriptionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotDisableFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	disableFlagName := fmt.Sprintf("%v.disable", cmdPrefix)
	if cmd.Flags().Changed(disableFlagName) {

		var disableFlagName string
		if cmdPrefix == "" {
			disableFlagName = "disable"
		} else {
			disableFlagName = fmt.Sprintf("%v.disable", cmdPrefix)
		}

		disableFlagValue, err := cmd.Flags().GetBool(disableFlagName)
		if err != nil {
			return err, false
		}
		m.Disable = disableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotDurationFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	durationFlagName := fmt.Sprintf("%v.duration", cmdPrefix)
	if cmd.Flags().Changed(durationFlagName) {

		var durationFlagName string
		if cmdPrefix == "" {
			durationFlagName = "duration"
		} else {
			durationFlagName = fmt.Sprintf("%v.duration", cmdPrefix)
		}

		durationFlagValue, err := cmd.Flags().GetInt64(durationFlagName)
		if err != nil {
			return err, false
		}
		m.Duration = durationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotEditableFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	editableFlagName := fmt.Sprintf("%v.editable", cmdPrefix)
	if cmd.Flags().Changed(editableFlagName) {

		var editableFlagName string
		if cmdPrefix == "" {
			editableFlagName = "editable"
		} else {
			editableFlagName = fmt.Sprintf("%v.editable", cmdPrefix)
		}

		editableFlagValue, err := cmd.Flags().GetBool(editableFlagName)
		if err != nil {
			return err, false
		}
		m.Editable = editableFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotExpiresAtFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	expiresAtFlagName := fmt.Sprintf("%v.expires_at", cmdPrefix)
	if cmd.Flags().Changed(expiresAtFlagName) {

		var expiresAtFlagName string
		if cmdPrefix == "" {
			expiresAtFlagName = "expires_at"
		} else {
			expiresAtFlagName = fmt.Sprintf("%v.expires_at", cmdPrefix)
		}

		expiresAtFlagValue, err := cmd.Flags().GetInt64(expiresAtFlagName)
		if err != nil {
			return err, false
		}
		m.ExpiresAt = expiresAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotIDFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotLevelFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	levelFlagName := fmt.Sprintf("%v.level", cmdPrefix)
	if cmd.Flags().Changed(levelFlagName) {

		var levelFlagName string
		if cmdPrefix == "" {
			levelFlagName = "level"
		} else {
			levelFlagName = fmt.Sprintf("%v.level", cmdPrefix)
		}

		levelFlagValue, err := cmd.Flags().GetString(levelFlagName)
		if err != nil {
			return err, false
		}
		m.Level = levelFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotNameFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	nameFlagName := fmt.Sprintf("%v.name", cmdPrefix)
	if cmd.Flags().Changed(nameFlagName) {

		var nameFlagName string
		if cmdPrefix == "" {
			nameFlagName = "name"
		} else {
			nameFlagName = fmt.Sprintf("%v.name", cmdPrefix)
		}

		nameFlagValue, err := cmd.Flags().GetString(nameFlagName)
		if err != nil {
			return err, false
		}
		m.Name = nameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotPermissionsFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	permissionsFlagName := fmt.Sprintf("%v.permissions", cmdPrefix)
	if cmd.Flags().Changed(permissionsFlagName) {
		// warning: permissions array type []*RobotPermission is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRobotSecretFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	secretFlagName := fmt.Sprintf("%v.secret", cmdPrefix)
	if cmd.Flags().Changed(secretFlagName) {

		var secretFlagName string
		if cmdPrefix == "" {
			secretFlagName = "secret"
		} else {
			secretFlagName = fmt.Sprintf("%v.secret", cmdPrefix)
		}

		secretFlagValue, err := cmd.Flags().GetString(secretFlagName)
		if err != nil {
			return err, false
		}
		m.Secret = secretFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRobotUpdateTimeFlags(depth int, m *models.Robot, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updateTimeFlagName := fmt.Sprintf("%v.update_time", cmdPrefix)
	if cmd.Flags().Changed(updateTimeFlagName) {

		var updateTimeFlagName string
		if cmdPrefix == "" {
			updateTimeFlagName = "update_time"
		} else {
			updateTimeFlagName = fmt.Sprintf("%v.update_time", cmdPrefix)
		}

		updateTimeFlagValueStr, err := cmd.Flags().GetString(updateTimeFlagName)
		if err != nil {
			return err, false
		}
		var updateTimeFlagValue strfmt.DateTime
		if err := updateTimeFlagValue.UnmarshalText([]byte(updateTimeFlagValueStr)); err != nil {
			return err, false
		}
		m.UpdateTime = updateTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}