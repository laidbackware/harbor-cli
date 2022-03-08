// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for Metrics

// register flags to command
func registerModelMetricsFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerMetricsErrorTaskCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricsPendingTaskCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricsRunningTaskCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricsScheduledTaskCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricsStoppedTaskCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricsSuccessTaskCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerMetricsTaskCount(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerMetricsErrorTaskCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	errorTaskCountDescription := `The count of error task`

	var errorTaskCountFlagName string
	if cmdPrefix == "" {
		errorTaskCountFlagName = "error_task_count"
	} else {
		errorTaskCountFlagName = fmt.Sprintf("%v.error_task_count", cmdPrefix)
	}

	var errorTaskCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(errorTaskCountFlagName, errorTaskCountFlagDefault, errorTaskCountDescription)

	return nil
}

func registerMetricsPendingTaskCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	pendingTaskCountDescription := `The count of pending task`

	var pendingTaskCountFlagName string
	if cmdPrefix == "" {
		pendingTaskCountFlagName = "pending_task_count"
	} else {
		pendingTaskCountFlagName = fmt.Sprintf("%v.pending_task_count", cmdPrefix)
	}

	var pendingTaskCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(pendingTaskCountFlagName, pendingTaskCountFlagDefault, pendingTaskCountDescription)

	return nil
}

func registerMetricsRunningTaskCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	runningTaskCountDescription := `The count of running task`

	var runningTaskCountFlagName string
	if cmdPrefix == "" {
		runningTaskCountFlagName = "running_task_count"
	} else {
		runningTaskCountFlagName = fmt.Sprintf("%v.running_task_count", cmdPrefix)
	}

	var runningTaskCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(runningTaskCountFlagName, runningTaskCountFlagDefault, runningTaskCountDescription)

	return nil
}

func registerMetricsScheduledTaskCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	scheduledTaskCountDescription := `The count of scheduled task`

	var scheduledTaskCountFlagName string
	if cmdPrefix == "" {
		scheduledTaskCountFlagName = "scheduled_task_count"
	} else {
		scheduledTaskCountFlagName = fmt.Sprintf("%v.scheduled_task_count", cmdPrefix)
	}

	var scheduledTaskCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(scheduledTaskCountFlagName, scheduledTaskCountFlagDefault, scheduledTaskCountDescription)

	return nil
}

func registerMetricsStoppedTaskCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	stoppedTaskCountDescription := `The count of stopped task`

	var stoppedTaskCountFlagName string
	if cmdPrefix == "" {
		stoppedTaskCountFlagName = "stopped_task_count"
	} else {
		stoppedTaskCountFlagName = fmt.Sprintf("%v.stopped_task_count", cmdPrefix)
	}

	var stoppedTaskCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(stoppedTaskCountFlagName, stoppedTaskCountFlagDefault, stoppedTaskCountDescription)

	return nil
}

func registerMetricsSuccessTaskCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	successTaskCountDescription := `The count of success task`

	var successTaskCountFlagName string
	if cmdPrefix == "" {
		successTaskCountFlagName = "success_task_count"
	} else {
		successTaskCountFlagName = fmt.Sprintf("%v.success_task_count", cmdPrefix)
	}

	var successTaskCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(successTaskCountFlagName, successTaskCountFlagDefault, successTaskCountDescription)

	return nil
}

func registerMetricsTaskCount(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	taskCountDescription := `The count of task`

	var taskCountFlagName string
	if cmdPrefix == "" {
		taskCountFlagName = "task_count"
	} else {
		taskCountFlagName = fmt.Sprintf("%v.task_count", cmdPrefix)
	}

	var taskCountFlagDefault int64

	_ = cmd.PersistentFlags().Int64(taskCountFlagName, taskCountFlagDefault, taskCountDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelMetricsFlags(depth int, m *models.Metrics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, errorTaskCountAdded := retrieveMetricsErrorTaskCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || errorTaskCountAdded

	err, pendingTaskCountAdded := retrieveMetricsPendingTaskCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || pendingTaskCountAdded

	err, runningTaskCountAdded := retrieveMetricsRunningTaskCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || runningTaskCountAdded

	err, scheduledTaskCountAdded := retrieveMetricsScheduledTaskCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scheduledTaskCountAdded

	err, stoppedTaskCountAdded := retrieveMetricsStoppedTaskCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || stoppedTaskCountAdded

	err, successTaskCountAdded := retrieveMetricsSuccessTaskCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || successTaskCountAdded

	err, taskCountAdded := retrieveMetricsTaskCountFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || taskCountAdded

	return nil, retAdded
}

func retrieveMetricsErrorTaskCountFlags(depth int, m *models.Metrics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	errorTaskCountFlagName := fmt.Sprintf("%v.error_task_count", cmdPrefix)
	if cmd.Flags().Changed(errorTaskCountFlagName) {

		var errorTaskCountFlagName string
		if cmdPrefix == "" {
			errorTaskCountFlagName = "error_task_count"
		} else {
			errorTaskCountFlagName = fmt.Sprintf("%v.error_task_count", cmdPrefix)
		}

		errorTaskCountFlagValue, err := cmd.Flags().GetInt64(errorTaskCountFlagName)
		if err != nil {
			return err, false
		}
		m.ErrorTaskCount = errorTaskCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricsPendingTaskCountFlags(depth int, m *models.Metrics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	pendingTaskCountFlagName := fmt.Sprintf("%v.pending_task_count", cmdPrefix)
	if cmd.Flags().Changed(pendingTaskCountFlagName) {

		var pendingTaskCountFlagName string
		if cmdPrefix == "" {
			pendingTaskCountFlagName = "pending_task_count"
		} else {
			pendingTaskCountFlagName = fmt.Sprintf("%v.pending_task_count", cmdPrefix)
		}

		pendingTaskCountFlagValue, err := cmd.Flags().GetInt64(pendingTaskCountFlagName)
		if err != nil {
			return err, false
		}
		m.PendingTaskCount = pendingTaskCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricsRunningTaskCountFlags(depth int, m *models.Metrics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	runningTaskCountFlagName := fmt.Sprintf("%v.running_task_count", cmdPrefix)
	if cmd.Flags().Changed(runningTaskCountFlagName) {

		var runningTaskCountFlagName string
		if cmdPrefix == "" {
			runningTaskCountFlagName = "running_task_count"
		} else {
			runningTaskCountFlagName = fmt.Sprintf("%v.running_task_count", cmdPrefix)
		}

		runningTaskCountFlagValue, err := cmd.Flags().GetInt64(runningTaskCountFlagName)
		if err != nil {
			return err, false
		}
		m.RunningTaskCount = runningTaskCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricsScheduledTaskCountFlags(depth int, m *models.Metrics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scheduledTaskCountFlagName := fmt.Sprintf("%v.scheduled_task_count", cmdPrefix)
	if cmd.Flags().Changed(scheduledTaskCountFlagName) {

		var scheduledTaskCountFlagName string
		if cmdPrefix == "" {
			scheduledTaskCountFlagName = "scheduled_task_count"
		} else {
			scheduledTaskCountFlagName = fmt.Sprintf("%v.scheduled_task_count", cmdPrefix)
		}

		scheduledTaskCountFlagValue, err := cmd.Flags().GetInt64(scheduledTaskCountFlagName)
		if err != nil {
			return err, false
		}
		m.ScheduledTaskCount = scheduledTaskCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricsStoppedTaskCountFlags(depth int, m *models.Metrics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	stoppedTaskCountFlagName := fmt.Sprintf("%v.stopped_task_count", cmdPrefix)
	if cmd.Flags().Changed(stoppedTaskCountFlagName) {

		var stoppedTaskCountFlagName string
		if cmdPrefix == "" {
			stoppedTaskCountFlagName = "stopped_task_count"
		} else {
			stoppedTaskCountFlagName = fmt.Sprintf("%v.stopped_task_count", cmdPrefix)
		}

		stoppedTaskCountFlagValue, err := cmd.Flags().GetInt64(stoppedTaskCountFlagName)
		if err != nil {
			return err, false
		}
		m.StoppedTaskCount = stoppedTaskCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricsSuccessTaskCountFlags(depth int, m *models.Metrics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	successTaskCountFlagName := fmt.Sprintf("%v.success_task_count", cmdPrefix)
	if cmd.Flags().Changed(successTaskCountFlagName) {

		var successTaskCountFlagName string
		if cmdPrefix == "" {
			successTaskCountFlagName = "success_task_count"
		} else {
			successTaskCountFlagName = fmt.Sprintf("%v.success_task_count", cmdPrefix)
		}

		successTaskCountFlagValue, err := cmd.Flags().GetInt64(successTaskCountFlagName)
		if err != nil {
			return err, false
		}
		m.SuccessTaskCount = successTaskCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveMetricsTaskCountFlags(depth int, m *models.Metrics, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	taskCountFlagName := fmt.Sprintf("%v.task_count", cmdPrefix)
	if cmd.Flags().Changed(taskCountFlagName) {

		var taskCountFlagName string
		if cmdPrefix == "" {
			taskCountFlagName = "task_count"
		} else {
			taskCountFlagName = fmt.Sprintf("%v.task_count", cmdPrefix)
		}

		taskCountFlagValue, err := cmd.Flags().GetInt64(taskCountFlagName)
		if err != nil {
			return err, false
		}
		m.TaskCount = taskCountFlagValue

		retAdded = true
	}

	return nil, retAdded
}
