// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for OverallHealthStatus

// register flags to command
func registerModelOverallHealthStatusFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerOverallHealthStatusComponents(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerOverallHealthStatusStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerOverallHealthStatusComponents(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: components []*ComponentHealthStatus array type is not supported by go-swagger cli yet

	return nil
}

func registerOverallHealthStatusStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := `The overall health status. It is 'healthy' only when all the components status are 'healthy'`

	var statusFlagName string
	if cmdPrefix == "" {
		statusFlagName = "status"
	} else {
		statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
	}

	var statusFlagDefault string

	_ = cmd.PersistentFlags().String(statusFlagName, statusFlagDefault, statusDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelOverallHealthStatusFlags(depth int, m *models.OverallHealthStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, componentsAdded := retrieveOverallHealthStatusComponentsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || componentsAdded

	err, statusAdded := retrieveOverallHealthStatusStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	return nil, retAdded
}

func retrieveOverallHealthStatusComponentsFlags(depth int, m *models.OverallHealthStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	componentsFlagName := fmt.Sprintf("%v.components", cmdPrefix)
	if cmd.Flags().Changed(componentsFlagName) {
		// warning: components array type []*ComponentHealthStatus is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveOverallHealthStatusStatusFlags(depth int, m *models.OverallHealthStatus, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusFlagName := fmt.Sprintf("%v.status", cmdPrefix)
	if cmd.Flags().Changed(statusFlagName) {

		var statusFlagName string
		if cmdPrefix == "" {
			statusFlagName = "status"
		} else {
			statusFlagName = fmt.Sprintf("%v.status", cmdPrefix)
		}

		statusFlagValue, err := cmd.Flags().GetString(statusFlagName)
		if err != nil {
			return err, false
		}
		m.Status = statusFlagValue

		retAdded = true
	}

	return nil, retAdded
}