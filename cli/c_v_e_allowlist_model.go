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

// Schema cli for CVEAllowlist

// register flags to command
func registerModelCVEAllowlistFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerCVEAllowlistCreationTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCVEAllowlistExpiresAt(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCVEAllowlistID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCVEAllowlistItems(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCVEAllowlistProjectID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerCVEAllowlistUpdateTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerCVEAllowlistCreationTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	creationTimeDescription := `The creation time of the allowlist.`

	var creationTimeFlagName string
	if cmdPrefix == "" {
		creationTimeFlagName = "creation_time"
	} else {
		creationTimeFlagName = fmt.Sprintf("%v.creation_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(creationTimeFlagName, "", creationTimeDescription)

	return nil
}

func registerCVEAllowlistExpiresAt(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	expiresAtDescription := `the time for expiration of the allowlist, in the form of seconds since epoch.  This is an optional attribute, if its not set the CVE allowlist does not expire.`

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

func registerCVEAllowlistID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `ID of the allowlist`

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

func registerCVEAllowlistItems(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: items []*CVEAllowlistItem array type is not supported by go-swagger cli yet

	return nil
}

func registerCVEAllowlistProjectID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	projectIdDescription := `ID of the project which the allowlist belongs to.  For system level allowlist this attribute is zero.`

	var projectIdFlagName string
	if cmdPrefix == "" {
		projectIdFlagName = "project_id"
	} else {
		projectIdFlagName = fmt.Sprintf("%v.project_id", cmdPrefix)
	}

	var projectIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(projectIdFlagName, projectIdFlagDefault, projectIdDescription)

	return nil
}

func registerCVEAllowlistUpdateTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updateTimeDescription := `The update time of the allowlist.`

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
func retrieveModelCVEAllowlistFlags(depth int, m *models.CVEAllowlist, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, creationTimeAdded := retrieveCVEAllowlistCreationTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || creationTimeAdded

	err, expiresAtAdded := retrieveCVEAllowlistExpiresAtFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || expiresAtAdded

	err, idAdded := retrieveCVEAllowlistIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, itemsAdded := retrieveCVEAllowlistItemsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || itemsAdded

	err, projectIdAdded := retrieveCVEAllowlistProjectIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || projectIdAdded

	err, updateTimeAdded := retrieveCVEAllowlistUpdateTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updateTimeAdded

	return nil, retAdded
}

func retrieveCVEAllowlistCreationTimeFlags(depth int, m *models.CVEAllowlist, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveCVEAllowlistExpiresAtFlags(depth int, m *models.CVEAllowlist, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
		m.ExpiresAt = &expiresAtFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCVEAllowlistIDFlags(depth int, m *models.CVEAllowlist, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveCVEAllowlistItemsFlags(depth int, m *models.CVEAllowlist, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	itemsFlagName := fmt.Sprintf("%v.items", cmdPrefix)
	if cmd.Flags().Changed(itemsFlagName) {
		// warning: items array type []*CVEAllowlistItem is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveCVEAllowlistProjectIDFlags(depth int, m *models.CVEAllowlist, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	projectIdFlagName := fmt.Sprintf("%v.project_id", cmdPrefix)
	if cmd.Flags().Changed(projectIdFlagName) {

		var projectIdFlagName string
		if cmdPrefix == "" {
			projectIdFlagName = "project_id"
		} else {
			projectIdFlagName = fmt.Sprintf("%v.project_id", cmdPrefix)
		}

		projectIdFlagValue, err := cmd.Flags().GetInt64(projectIdFlagName)
		if err != nil {
			return err, false
		}
		m.ProjectID = projectIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveCVEAllowlistUpdateTimeFlags(depth int, m *models.CVEAllowlist, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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