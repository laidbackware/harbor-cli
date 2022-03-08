// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for Access

// register flags to command
func registerModelAccessFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerAccessAction(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAccessEffect(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerAccessResource(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerAccessAction(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	actionDescription := `The action of the access`

	var actionFlagName string
	if cmdPrefix == "" {
		actionFlagName = "action"
	} else {
		actionFlagName = fmt.Sprintf("%v.action", cmdPrefix)
	}

	var actionFlagDefault string

	_ = cmd.PersistentFlags().String(actionFlagName, actionFlagDefault, actionDescription)

	return nil
}

func registerAccessEffect(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	effectDescription := `The effect of the access`

	var effectFlagName string
	if cmdPrefix == "" {
		effectFlagName = "effect"
	} else {
		effectFlagName = fmt.Sprintf("%v.effect", cmdPrefix)
	}

	var effectFlagDefault string

	_ = cmd.PersistentFlags().String(effectFlagName, effectFlagDefault, effectDescription)

	return nil
}

func registerAccessResource(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	resourceDescription := `The resource of the access`

	var resourceFlagName string
	if cmdPrefix == "" {
		resourceFlagName = "resource"
	} else {
		resourceFlagName = fmt.Sprintf("%v.resource", cmdPrefix)
	}

	var resourceFlagDefault string

	_ = cmd.PersistentFlags().String(resourceFlagName, resourceFlagDefault, resourceDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelAccessFlags(depth int, m *models.Access, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, actionAdded := retrieveAccessActionFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || actionAdded

	err, effectAdded := retrieveAccessEffectFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || effectAdded

	err, resourceAdded := retrieveAccessResourceFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || resourceAdded

	return nil, retAdded
}

func retrieveAccessActionFlags(depth int, m *models.Access, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	actionFlagName := fmt.Sprintf("%v.action", cmdPrefix)
	if cmd.Flags().Changed(actionFlagName) {

		var actionFlagName string
		if cmdPrefix == "" {
			actionFlagName = "action"
		} else {
			actionFlagName = fmt.Sprintf("%v.action", cmdPrefix)
		}

		actionFlagValue, err := cmd.Flags().GetString(actionFlagName)
		if err != nil {
			return err, false
		}
		m.Action = actionFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAccessEffectFlags(depth int, m *models.Access, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	effectFlagName := fmt.Sprintf("%v.effect", cmdPrefix)
	if cmd.Flags().Changed(effectFlagName) {

		var effectFlagName string
		if cmdPrefix == "" {
			effectFlagName = "effect"
		} else {
			effectFlagName = fmt.Sprintf("%v.effect", cmdPrefix)
		}

		effectFlagValue, err := cmd.Flags().GetString(effectFlagName)
		if err != nil {
			return err, false
		}
		m.Effect = effectFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveAccessResourceFlags(depth int, m *models.Access, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	resourceFlagName := fmt.Sprintf("%v.resource", cmdPrefix)
	if cmd.Flags().Changed(resourceFlagName) {

		var resourceFlagName string
		if cmdPrefix == "" {
			resourceFlagName = "resource"
		} else {
			resourceFlagName = fmt.Sprintf("%v.resource", cmdPrefix)
		}

		resourceFlagValue, err := cmd.Flags().GetString(resourceFlagName)
		if err != nil {
			return err, false
		}
		m.Resource = resourceFlagValue

		retAdded = true
	}

	return nil, retAdded
}
