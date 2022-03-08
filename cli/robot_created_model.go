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

// Schema cli for RobotCreated

// register flags to command
func registerModelRobotCreatedFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRobotCreatedCreationTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreatedExpiresAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreatedID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreatedName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRobotCreatedSecret(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRobotCreatedCreationTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreatedExpiresAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreatedID(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreatedName(depth int, cmdPrefix string, cmd *cobra.Command) error {
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

func registerRobotCreatedSecret(depth int, cmdPrefix string, cmd *cobra.Command) error {
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
func retrieveModelRobotCreatedFlags(depth int, m *models.RobotCreated, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, creationTimeAdded := retrieveRobotCreatedCreationTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || creationTimeAdded

	err, expiresAtAdded := retrieveRobotCreatedExpiresAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || expiresAtAdded

	err, idAdded := retrieveRobotCreatedIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, nameAdded := retrieveRobotCreatedNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || nameAdded

	err, secretAdded := retrieveRobotCreatedSecretFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || secretAdded

	return nil, retAdded
}

func retrieveRobotCreatedCreationTimeFlags(depth int, m *models.RobotCreated, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreatedExpiresAtFlags(depth int, m *models.RobotCreated, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreatedIDFlags(depth int, m *models.RobotCreated, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreatedNameFlags(depth int, m *models.RobotCreated, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRobotCreatedSecretFlags(depth int, m *models.RobotCreated, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
