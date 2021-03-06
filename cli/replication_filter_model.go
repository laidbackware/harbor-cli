// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for ReplicationFilter

// register flags to command
func registerModelReplicationFilterFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerReplicationFilterDecoration(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReplicationFilterType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerReplicationFilterValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerReplicationFilterDecoration(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	decorationDescription := `matches or excludes the result`

	var decorationFlagName string
	if cmdPrefix == "" {
		decorationFlagName = "decoration"
	} else {
		decorationFlagName = fmt.Sprintf("%v.decoration", cmdPrefix)
	}

	var decorationFlagDefault string

	_ = cmd.PersistentFlags().String(decorationFlagName, decorationFlagDefault, decorationDescription)

	return nil
}

func registerReplicationFilterType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := `The replication policy filter type.`

	var typeFlagName string
	if cmdPrefix == "" {
		typeFlagName = "type"
	} else {
		typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
	}

	var typeFlagDefault string

	_ = cmd.PersistentFlags().String(typeFlagName, typeFlagDefault, typeDescription)

	return nil
}

func registerReplicationFilterValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: value interface{} map type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelReplicationFilterFlags(depth int, m *models.ReplicationFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, decorationAdded := retrieveReplicationFilterDecorationFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || decorationAdded

	err, typeAdded := retrieveReplicationFilterTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, valueAdded := retrieveReplicationFilterValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveReplicationFilterDecorationFlags(depth int, m *models.ReplicationFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	decorationFlagName := fmt.Sprintf("%v.decoration", cmdPrefix)
	if cmd.Flags().Changed(decorationFlagName) {

		var decorationFlagName string
		if cmdPrefix == "" {
			decorationFlagName = "decoration"
		} else {
			decorationFlagName = fmt.Sprintf("%v.decoration", cmdPrefix)
		}

		decorationFlagValue, err := cmd.Flags().GetString(decorationFlagName)
		if err != nil {
			return err, false
		}
		m.Decoration = decorationFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReplicationFilterTypeFlags(depth int, m *models.ReplicationFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	typeFlagName := fmt.Sprintf("%v.type", cmdPrefix)
	if cmd.Flags().Changed(typeFlagName) {

		var typeFlagName string
		if cmdPrefix == "" {
			typeFlagName = "type"
		} else {
			typeFlagName = fmt.Sprintf("%v.type", cmdPrefix)
		}

		typeFlagValue, err := cmd.Flags().GetString(typeFlagName)
		if err != nil {
			return err, false
		}
		m.Type = typeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveReplicationFilterValueFlags(depth int, m *models.ReplicationFilter, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
