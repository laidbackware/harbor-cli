// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for InternalConfigurationValue

// register flags to command
func registerModelInternalConfigurationValueFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerInternalConfigurationValueEditable(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerInternalConfigurationValueValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerInternalConfigurationValueEditable(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	editableDescription := `The configure item can be updated or not`

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

func registerInternalConfigurationValueValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: value interface{} map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelInternalConfigurationValueFlags(depth int, m *models.InternalConfigurationValue, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, editableAdded := retrieveInternalConfigurationValueEditableFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || editableAdded

	err, valueAdded := retrieveInternalConfigurationValueValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveInternalConfigurationValueEditableFlags(depth int, m *models.InternalConfigurationValue, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveInternalConfigurationValueValueFlags(depth int, m *models.InternalConfigurationValue, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valueFlagName := fmt.Sprintf("%v.value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {
		// warning: value map type interface{} is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
