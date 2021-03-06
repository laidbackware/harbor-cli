// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for RegistryEndpoint

// register flags to command
func registerModelRegistryEndpointFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRegistryEndpointKey(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryEndpointValue(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegistryEndpointKey(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	keyDescription := `The endpoint key`

	var keyFlagName string
	if cmdPrefix == "" {
		keyFlagName = "key"
	} else {
		keyFlagName = fmt.Sprintf("%v.key", cmdPrefix)
	}

	var keyFlagDefault string

	_ = cmd.PersistentFlags().String(keyFlagName, keyFlagDefault, keyDescription)

	return nil
}

func registerRegistryEndpointValue(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	valueDescription := `The endpoint value`

	var valueFlagName string
	if cmdPrefix == "" {
		valueFlagName = "value"
	} else {
		valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
	}

	var valueFlagDefault string

	_ = cmd.PersistentFlags().String(valueFlagName, valueFlagDefault, valueDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRegistryEndpointFlags(depth int, m *models.RegistryEndpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, keyAdded := retrieveRegistryEndpointKeyFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || keyAdded

	err, valueAdded := retrieveRegistryEndpointValueFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || valueAdded

	return nil, retAdded
}

func retrieveRegistryEndpointKeyFlags(depth int, m *models.RegistryEndpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	keyFlagName := fmt.Sprintf("%v.key", cmdPrefix)
	if cmd.Flags().Changed(keyFlagName) {

		var keyFlagName string
		if cmdPrefix == "" {
			keyFlagName = "key"
		} else {
			keyFlagName = fmt.Sprintf("%v.key", cmdPrefix)
		}

		keyFlagValue, err := cmd.Flags().GetString(keyFlagName)
		if err != nil {
			return err, false
		}
		m.Key = keyFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRegistryEndpointValueFlags(depth int, m *models.RegistryEndpoint, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	valueFlagName := fmt.Sprintf("%v.value", cmdPrefix)
	if cmd.Flags().Changed(valueFlagName) {

		var valueFlagName string
		if cmdPrefix == "" {
			valueFlagName = "value"
		} else {
			valueFlagName = fmt.Sprintf("%v.value", cmdPrefix)
		}

		valueFlagValue, err := cmd.Flags().GetString(valueFlagName)
		if err != nil {
			return err, false
		}
		m.Value = valueFlagValue

		retAdded = true
	}

	return nil, retAdded
}
