// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for RetentionExecution

// register flags to command
func registerModelRetentionExecutionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRetentionExecutionDryRun(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionExecutionEndTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionExecutionID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionExecutionPolicyID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionExecutionStartTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionExecutionStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionExecutionTrigger(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRetentionExecutionDryRun(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	dryRunDescription := ``

	var dryRunFlagName string
	if cmdPrefix == "" {
		dryRunFlagName = "dry_run"
	} else {
		dryRunFlagName = fmt.Sprintf("%v.dry_run", cmdPrefix)
	}

	var dryRunFlagDefault bool

	_ = cmd.PersistentFlags().Bool(dryRunFlagName, dryRunFlagDefault, dryRunDescription)

	return nil
}

func registerRetentionExecutionEndTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endTimeDescription := ``

	var endTimeFlagName string
	if cmdPrefix == "" {
		endTimeFlagName = "end_time"
	} else {
		endTimeFlagName = fmt.Sprintf("%v.end_time", cmdPrefix)
	}

	var endTimeFlagDefault string

	_ = cmd.PersistentFlags().String(endTimeFlagName, endTimeFlagDefault, endTimeDescription)

	return nil
}

func registerRetentionExecutionID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := ``

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

func registerRetentionExecutionPolicyID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	policyIdDescription := ``

	var policyIdFlagName string
	if cmdPrefix == "" {
		policyIdFlagName = "policy_id"
	} else {
		policyIdFlagName = fmt.Sprintf("%v.policy_id", cmdPrefix)
	}

	var policyIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(policyIdFlagName, policyIdFlagDefault, policyIdDescription)

	return nil
}

func registerRetentionExecutionStartTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startTimeDescription := ``

	var startTimeFlagName string
	if cmdPrefix == "" {
		startTimeFlagName = "start_time"
	} else {
		startTimeFlagName = fmt.Sprintf("%v.start_time", cmdPrefix)
	}

	var startTimeFlagDefault string

	_ = cmd.PersistentFlags().String(startTimeFlagName, startTimeFlagDefault, startTimeDescription)

	return nil
}

func registerRetentionExecutionStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := ``

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

func registerRetentionExecutionTrigger(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	triggerDescription := ``

	var triggerFlagName string
	if cmdPrefix == "" {
		triggerFlagName = "trigger"
	} else {
		triggerFlagName = fmt.Sprintf("%v.trigger", cmdPrefix)
	}

	var triggerFlagDefault string

	_ = cmd.PersistentFlags().String(triggerFlagName, triggerFlagDefault, triggerDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRetentionExecutionFlags(depth int, m *models.RetentionExecution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, dryRunAdded := retrieveRetentionExecutionDryRunFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || dryRunAdded

	err, endTimeAdded := retrieveRetentionExecutionEndTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endTimeAdded

	err, idAdded := retrieveRetentionExecutionIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, policyIdAdded := retrieveRetentionExecutionPolicyIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || policyIdAdded

	err, startTimeAdded := retrieveRetentionExecutionStartTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startTimeAdded

	err, statusAdded := retrieveRetentionExecutionStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, triggerAdded := retrieveRetentionExecutionTriggerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || triggerAdded

	return nil, retAdded
}

func retrieveRetentionExecutionDryRunFlags(depth int, m *models.RetentionExecution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	dryRunFlagName := fmt.Sprintf("%v.dry_run", cmdPrefix)
	if cmd.Flags().Changed(dryRunFlagName) {

		var dryRunFlagName string
		if cmdPrefix == "" {
			dryRunFlagName = "dry_run"
		} else {
			dryRunFlagName = fmt.Sprintf("%v.dry_run", cmdPrefix)
		}

		dryRunFlagValue, err := cmd.Flags().GetBool(dryRunFlagName)
		if err != nil {
			return err, false
		}
		m.DryRun = dryRunFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRetentionExecutionEndTimeFlags(depth int, m *models.RetentionExecution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	endTimeFlagName := fmt.Sprintf("%v.end_time", cmdPrefix)
	if cmd.Flags().Changed(endTimeFlagName) {

		var endTimeFlagName string
		if cmdPrefix == "" {
			endTimeFlagName = "end_time"
		} else {
			endTimeFlagName = fmt.Sprintf("%v.end_time", cmdPrefix)
		}

		endTimeFlagValue, err := cmd.Flags().GetString(endTimeFlagName)
		if err != nil {
			return err, false
		}
		m.EndTime = endTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRetentionExecutionIDFlags(depth int, m *models.RetentionExecution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRetentionExecutionPolicyIDFlags(depth int, m *models.RetentionExecution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	policyIdFlagName := fmt.Sprintf("%v.policy_id", cmdPrefix)
	if cmd.Flags().Changed(policyIdFlagName) {

		var policyIdFlagName string
		if cmdPrefix == "" {
			policyIdFlagName = "policy_id"
		} else {
			policyIdFlagName = fmt.Sprintf("%v.policy_id", cmdPrefix)
		}

		policyIdFlagValue, err := cmd.Flags().GetInt64(policyIdFlagName)
		if err != nil {
			return err, false
		}
		m.PolicyID = policyIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRetentionExecutionStartTimeFlags(depth int, m *models.RetentionExecution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	startTimeFlagName := fmt.Sprintf("%v.start_time", cmdPrefix)
	if cmd.Flags().Changed(startTimeFlagName) {

		var startTimeFlagName string
		if cmdPrefix == "" {
			startTimeFlagName = "start_time"
		} else {
			startTimeFlagName = fmt.Sprintf("%v.start_time", cmdPrefix)
		}

		startTimeFlagValue, err := cmd.Flags().GetString(startTimeFlagName)
		if err != nil {
			return err, false
		}
		m.StartTime = startTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveRetentionExecutionStatusFlags(depth int, m *models.RetentionExecution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveRetentionExecutionTriggerFlags(depth int, m *models.RetentionExecution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	triggerFlagName := fmt.Sprintf("%v.trigger", cmdPrefix)
	if cmd.Flags().Changed(triggerFlagName) {

		var triggerFlagName string
		if cmdPrefix == "" {
			triggerFlagName = "trigger"
		} else {
			triggerFlagName = fmt.Sprintf("%v.trigger", cmdPrefix)
		}

		triggerFlagValue, err := cmd.Flags().GetString(triggerFlagName)
		if err != nil {
			return err, false
		}
		m.Trigger = triggerFlagValue

		retAdded = true
	}

	return nil, retAdded
}
