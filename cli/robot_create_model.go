// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for RobotCreate

// register flags to command
func registerModelRobotCreateFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRobotCreateDescription(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreateDisable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreateDuration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreateLevel(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreateName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreatePermissions(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreateSecret(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRobotCreateDescription(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreateDisable(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreateDuration(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreateLevel(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreateName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreatePermissions(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: permissions []*RobotPermission array type is not supported by go-swagger cli yet

	return nil
}

func registerRobotCreateSecret(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRobotCreateFlags(depth int, m *models.RobotCreate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, descriptionAdded := retrieveRobotCreateDescriptionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || descriptionAdded

	err, disableAdded := retrieveRobotCreateDisableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || disableAdded

	err, durationAdded := retrieveRobotCreateDurationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || durationAdded

	err, levelAdded := retrieveRobotCreateLevelFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || levelAdded

	err, nameAdded := retrieveRobotCreateNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, permissionsAdded := retrieveRobotCreatePermissionsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || permissionsAdded

	err, secretAdded := retrieveRobotCreateSecretFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secretAdded

	return nil, retAdded
}

func retrieveRobotCreateDescriptionFlags(depth int, m *models.RobotCreate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreateDisableFlags(depth int, m *models.RobotCreate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreateDurationFlags(depth int, m *models.RobotCreate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreateLevelFlags(depth int, m *models.RobotCreate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreateNameFlags(depth int, m *models.RobotCreate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreatePermissionsFlags(depth int, m *models.RobotCreate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreateSecretFlags(depth int, m *models.RobotCreate, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
