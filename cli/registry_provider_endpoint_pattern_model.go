// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for RegistryProviderEndpointPattern

// register flags to command
func registerModelRegistryProviderEndpointPatternFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRegistryProviderEndpointPatternEndpointType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRegistryProviderEndpointPatternEndpoints(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRegistryProviderEndpointPatternEndpointType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endpointTypeDescription := `The endpoint type`

	var endpointTypeFlagName string
	if cmdPrefix == "" {
		endpointTypeFlagName = "endpoint_type"
	} else {
		endpointTypeFlagName = fmt.Sprintf("%v.endpoint_type", cmdPrefix)
	}

	var endpointTypeFlagDefault string

	_ = cmd.PersistentFlags().String(endpointTypeFlagName, endpointTypeFlagDefault, endpointTypeDescription)

	return nil
}

func registerRegistryProviderEndpointPatternEndpoints(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: endpoints []*RegistryEndpoint array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRegistryProviderEndpointPatternFlags(depth int, m *models.RegistryProviderEndpointPattern, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endpointTypeAdded := retrieveRegistryProviderEndpointPatternEndpointTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointTypeAdded

	err, endpointsAdded := retrieveRegistryProviderEndpointPatternEndpointsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endpointsAdded

	return nil, retAdded
}

func retrieveRegistryProviderEndpointPatternEndpointTypeFlags(depth int, m *models.RegistryProviderEndpointPattern, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointTypeFlagName := fmt.Sprintf("%v.endpoint_type", cmdPrefix)
	if cmd.Flags().Changed(endpointTypeFlagName) {

		var endpointTypeFlagName string
		if cmdPrefix == "" {
			endpointTypeFlagName = "endpoint_type"
		} else {
			endpointTypeFlagName = fmt.Sprintf("%v.endpoint_type", cmdPrefix)
		}

		endpointTypeFlagValue, err := cmd.Flags().GetString(endpointTypeFlagName)
		if err != nil {
			return err, false
		}
		m.EndpointType = endpointTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRegistryProviderEndpointPatternEndpointsFlags(depth int, m *models.RegistryProviderEndpointPattern, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endpointsFlagName := fmt.Sprintf("%v.endpoints", cmdPrefix)
	if cmd.Flags().Changed(endpointsFlagName) {
		// warning: endpoints array type []*RegistryEndpoint is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
