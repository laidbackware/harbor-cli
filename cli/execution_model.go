// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for Execution

// register flags to command
func registerModelExecutionFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerExecutionEndTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionExtraAttrs(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionMetrics(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionStartTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionStatusMessage(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionTrigger(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionVendorID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerExecutionVendorType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerExecutionEndTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	endTimeDescription := `The end time of execution`

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

func registerExecutionExtraAttrs(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: extra_attrs ExtraAttrs map type is not supported by go-swagger cli yet

	return nil
}

func registerExecutionID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of execution`

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

func registerExecutionMetrics(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var metricsFlagName string
	if cmdPrefix == "" {
		metricsFlagName = "metrics"
	} else {
		metricsFlagName = fmt.Sprintf("%v.metrics", cmdPrefix)
	}

	if err := registerModelMetricsFlags(depth+1, metricsFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerExecutionStartTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	startTimeDescription := `The start time of execution`

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

func registerExecutionStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusDescription := `The status of execution`

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

func registerExecutionStatusMessage(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	statusMessageDescription := `The status message of execution`

	var statusMessageFlagName string
	if cmdPrefix == "" {
		statusMessageFlagName = "status_message"
	} else {
		statusMessageFlagName = fmt.Sprintf("%v.status_message", cmdPrefix)
	}

	var statusMessageFlagDefault string

	_ = cmd.PersistentFlags().String(statusMessageFlagName, statusMessageFlagDefault, statusMessageDescription)

	return nil
}

func registerExecutionTrigger(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	triggerDescription := `The trigger of execution`

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

func registerExecutionVendorID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vendorIdDescription := `The vendor id of execution`

	var vendorIdFlagName string
	if cmdPrefix == "" {
		vendorIdFlagName = "vendor_id"
	} else {
		vendorIdFlagName = fmt.Sprintf("%v.vendor_id", cmdPrefix)
	}

	var vendorIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(vendorIdFlagName, vendorIdFlagDefault, vendorIdDescription)

	return nil
}

func registerExecutionVendorType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	vendorTypeDescription := `The vendor type of execution`

	var vendorTypeFlagName string
	if cmdPrefix == "" {
		vendorTypeFlagName = "vendor_type"
	} else {
		vendorTypeFlagName = fmt.Sprintf("%v.vendor_type", cmdPrefix)
	}

	var vendorTypeFlagDefault string

	_ = cmd.PersistentFlags().String(vendorTypeFlagName, vendorTypeFlagDefault, vendorTypeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelExecutionFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, endTimeAdded := retrieveExecutionEndTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || endTimeAdded

	err, extraAttrsAdded := retrieveExecutionExtraAttrsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || extraAttrsAdded

	err, idAdded := retrieveExecutionIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, metricsAdded := retrieveExecutionMetricsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metricsAdded

	err, startTimeAdded := retrieveExecutionStartTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || startTimeAdded

	err, statusAdded := retrieveExecutionStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusAdded

	err, statusMessageAdded := retrieveExecutionStatusMessageFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || statusMessageAdded

	err, triggerAdded := retrieveExecutionTriggerFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || triggerAdded

	err, vendorIdAdded := retrieveExecutionVendorIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vendorIdAdded

	err, vendorTypeAdded := retrieveExecutionVendorTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || vendorTypeAdded

	return nil, retAdded
}

func retrieveExecutionEndTimeFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveExecutionExtraAttrsFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	extraAttrsFlagName := fmt.Sprintf("%v.extra_attrs", cmdPrefix)
	if cmd.Flags().Changed(extraAttrsFlagName) {
		// warning: extra_attrs map type ExtraAttrs is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveExecutionIDFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveExecutionMetricsFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	metricsFlagName := fmt.Sprintf("%v.metrics", cmdPrefix)
	if cmd.Flags().Changed(metricsFlagName) {
		// info: complex object metrics Metrics is retrieved outside this Changed() block
	}
	metricsFlagValue := m.Metrics
	if swag.IsZero(metricsFlagValue) {
		metricsFlagValue = &models.Metrics{}
	}

	err, metricsAdded := retrieveModelMetricsFlags(depth+1, metricsFlagValue, metricsFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || metricsAdded
	if metricsAdded {
		m.Metrics = metricsFlagValue
	}

	return nil, retAdded
}

func retrieveExecutionStartTimeFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveExecutionStatusFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveExecutionStatusMessageFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	statusMessageFlagName := fmt.Sprintf("%v.status_message", cmdPrefix)
	if cmd.Flags().Changed(statusMessageFlagName) {

		var statusMessageFlagName string
		if cmdPrefix == "" {
			statusMessageFlagName = "status_message"
		} else {
			statusMessageFlagName = fmt.Sprintf("%v.status_message", cmdPrefix)
		}

		statusMessageFlagValue, err := cmd.Flags().GetString(statusMessageFlagName)
		if err != nil {
			return err, false
		}
		m.StatusMessage = statusMessageFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecutionTriggerFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveExecutionVendorIDFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vendorIdFlagName := fmt.Sprintf("%v.vendor_id", cmdPrefix)
	if cmd.Flags().Changed(vendorIdFlagName) {

		var vendorIdFlagName string
		if cmdPrefix == "" {
			vendorIdFlagName = "vendor_id"
		} else {
			vendorIdFlagName = fmt.Sprintf("%v.vendor_id", cmdPrefix)
		}

		vendorIdFlagValue, err := cmd.Flags().GetInt64(vendorIdFlagName)
		if err != nil {
			return err, false
		}
		m.VendorID = vendorIdFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveExecutionVendorTypeFlags(depth int, m *models.Execution, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	vendorTypeFlagName := fmt.Sprintf("%v.vendor_type", cmdPrefix)
	if cmd.Flags().Changed(vendorTypeFlagName) {

		var vendorTypeFlagName string
		if cmdPrefix == "" {
			vendorTypeFlagName = "vendor_type"
		} else {
			vendorTypeFlagName = fmt.Sprintf("%v.vendor_type", cmdPrefix)
		}

		vendorTypeFlagValue, err := cmd.Flags().GetString(vendorTypeFlagName)
		if err != nil {
			return err, false
		}
		m.VendorType = vendorTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
