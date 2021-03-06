// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for RetentionRuleParamMetadata

// register flags to command
func registerModelRetentionRuleParamMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRetentionRuleParamMetadataRequired(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionRuleParamMetadataType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionRuleParamMetadataUnit(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRetentionRuleParamMetadataRequired(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	requiredDescription := ``

	var requiredFlagName string
	if cmdPrefix == "" {
		requiredFlagName = "required"
	} else {
		requiredFlagName = fmt.Sprintf("%v.required", cmdPrefix)
	}

	var requiredFlagDefault bool

	_ = cmd.PersistentFlags().Bool(requiredFlagName, requiredFlagDefault, requiredDescription)

	return nil
}

func registerRetentionRuleParamMetadataType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	typeDescription := ``

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

func registerRetentionRuleParamMetadataUnit(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	unitDescription := ``

	var unitFlagName string
	if cmdPrefix == "" {
		unitFlagName = "unit"
	} else {
		unitFlagName = fmt.Sprintf("%v.unit", cmdPrefix)
	}

	var unitFlagDefault string

	_ = cmd.PersistentFlags().String(unitFlagName, unitFlagDefault, unitDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRetentionRuleParamMetadataFlags(depth int, m *models.RetentionRuleParamMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, requiredAdded := retrieveRetentionRuleParamMetadataRequiredFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || requiredAdded

	err, typeAdded := retrieveRetentionRuleParamMetadataTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || typeAdded

	err, unitAdded := retrieveRetentionRuleParamMetadataUnitFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || unitAdded

	return nil, retAdded
}

func retrieveRetentionRuleParamMetadataRequiredFlags(depth int, m *models.RetentionRuleParamMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	requiredFlagName := fmt.Sprintf("%v.required", cmdPrefix)
	if cmd.Flags().Changed(requiredFlagName) {

		var requiredFlagName string
		if cmdPrefix == "" {
			requiredFlagName = "required"
		} else {
			requiredFlagName = fmt.Sprintf("%v.required", cmdPrefix)
		}

		requiredFlagValue, err := cmd.Flags().GetBool(requiredFlagName)
		if err != nil {
			return err, false
		}
		m.Required = requiredFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRetentionRuleParamMetadataTypeFlags(depth int, m *models.RetentionRuleParamMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRetentionRuleParamMetadataUnitFlags(depth int, m *models.RetentionRuleParamMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	unitFlagName := fmt.Sprintf("%v.unit", cmdPrefix)
	if cmd.Flags().Changed(unitFlagName) {

		var unitFlagName string
		if cmdPrefix == "" {
			unitFlagName = "unit"
		} else {
			unitFlagName = fmt.Sprintf("%v.unit", cmdPrefix)
		}

		unitFlagValue, err := cmd.Flags().GetString(unitFlagName)
		if err != nil {
			return err, false
		}
		m.Unit = unitFlagValue

		retAdded = true
	}

	return nil, retAdded
}
