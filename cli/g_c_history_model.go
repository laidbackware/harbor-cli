// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/spf13/cobra"
)

// Schema cli for GCHistory

// register flags to command
func registerModelGCHistoryFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerGCHistoryCreationTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGCHistoryDeleted(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGCHistoryID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGCHistoryJobKind(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGCHistoryJobName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGCHistoryJobParameters(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGCHistoryJobStatus(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGCHistorySchedule(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerGCHistoryUpdateTime(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerGCHistoryCreationTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	creationTimeDescription := `the creation time of gc job.`

	var creationTimeFlagName string
	if cmdPrefix == "" {
		creationTimeFlagName = "creation_time"
	} else {
		creationTimeFlagName = fmt.Sprintf("%v.creation_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(creationTimeFlagName, "", creationTimeDescription)

	return nil
}

func registerGCHistoryDeleted(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	deletedDescription := `if gc job was deleted.`

	var deletedFlagName string
	if cmdPrefix == "" {
		deletedFlagName = "deleted"
	} else {
		deletedFlagName = fmt.Sprintf("%v.deleted", cmdPrefix)
	}

	var deletedFlagDefault bool

	_ = cmd.PersistentFlags().Bool(deletedFlagName, deletedFlagDefault, deletedDescription)

	return nil
}

func registerGCHistoryID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `the id of gc job.`

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

func registerGCHistoryJobKind(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	jobKindDescription := `the job kind of gc job.`

	var jobKindFlagName string
	if cmdPrefix == "" {
		jobKindFlagName = "job_kind"
	} else {
		jobKindFlagName = fmt.Sprintf("%v.job_kind", cmdPrefix)
	}

	var jobKindFlagDefault string

	_ = cmd.PersistentFlags().String(jobKindFlagName, jobKindFlagDefault, jobKindDescription)

	return nil
}

func registerGCHistoryJobName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	jobNameDescription := `the job name of gc job.`

	var jobNameFlagName string
	if cmdPrefix == "" {
		jobNameFlagName = "job_name"
	} else {
		jobNameFlagName = fmt.Sprintf("%v.job_name", cmdPrefix)
	}

	var jobNameFlagDefault string

	_ = cmd.PersistentFlags().String(jobNameFlagName, jobNameFlagDefault, jobNameDescription)

	return nil
}

func registerGCHistoryJobParameters(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	jobParametersDescription := `the job parameters of gc job.`

	var jobParametersFlagName string
	if cmdPrefix == "" {
		jobParametersFlagName = "job_parameters"
	} else {
		jobParametersFlagName = fmt.Sprintf("%v.job_parameters", cmdPrefix)
	}

	var jobParametersFlagDefault string

	_ = cmd.PersistentFlags().String(jobParametersFlagName, jobParametersFlagDefault, jobParametersDescription)

	return nil
}

func registerGCHistoryJobStatus(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	jobStatusDescription := `the status of gc job.`

	var jobStatusFlagName string
	if cmdPrefix == "" {
		jobStatusFlagName = "job_status"
	} else {
		jobStatusFlagName = fmt.Sprintf("%v.job_status", cmdPrefix)
	}

	var jobStatusFlagDefault string

	_ = cmd.PersistentFlags().String(jobStatusFlagName, jobStatusFlagDefault, jobStatusDescription)

	return nil
}

func registerGCHistorySchedule(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	var scheduleFlagName string
	if cmdPrefix == "" {
		scheduleFlagName = "schedule"
	} else {
		scheduleFlagName = fmt.Sprintf("%v.schedule", cmdPrefix)
	}

	if err := registerModelScheduleObjFlags(depth+1, scheduleFlagName, cmd); err != nil {
		return err
	}

	return nil
}

func registerGCHistoryUpdateTime(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	updateTimeDescription := `the update time of gc job.`

	var updateTimeFlagName string
	if cmdPrefix == "" {
		updateTimeFlagName = "update_time"
	} else {
		updateTimeFlagName = fmt.Sprintf("%v.update_time", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(updateTimeFlagName, "", updateTimeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelGCHistoryFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, creationTimeAdded := retrieveGCHistoryCreationTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || creationTimeAdded

	err, deletedAdded := retrieveGCHistoryDeletedFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || deletedAdded

	err, idAdded := retrieveGCHistoryIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	err, jobKindAdded := retrieveGCHistoryJobKindFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || jobKindAdded

	err, jobNameAdded := retrieveGCHistoryJobNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || jobNameAdded

	err, jobParametersAdded := retrieveGCHistoryJobParametersFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || jobParametersAdded

	err, jobStatusAdded := retrieveGCHistoryJobStatusFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || jobStatusAdded

	err, scheduleAdded := retrieveGCHistoryScheduleFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scheduleAdded

	err, updateTimeAdded := retrieveGCHistoryUpdateTimeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || updateTimeAdded

	return nil, retAdded
}

func retrieveGCHistoryCreationTimeFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	creationTimeFlagName := fmt.Sprintf("%v.creation_time", cmdPrefix)
	if cmd.Flags().Changed(creationTimeFlagName) {

		var creationTimeFlagName string
		if cmdPrefix == "" {
			creationTimeFlagName = "creation_time"
		} else {
			creationTimeFlagName = fmt.Sprintf("%v.creation_time", cmdPrefix)
		}

		creationTimeFlagValueStr, err := cmd.Flags().GetString(creationTimeFlagName)
		if err != nil {
			return err, false
		}
		var creationTimeFlagValue strfmt.DateTime
		if err := creationTimeFlagValue.UnmarshalText([]byte(creationTimeFlagValueStr)); err != nil {
			return err, false
		}
		m.CreationTime = creationTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGCHistoryDeletedFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	deletedFlagName := fmt.Sprintf("%v.deleted", cmdPrefix)
	if cmd.Flags().Changed(deletedFlagName) {

		var deletedFlagName string
		if cmdPrefix == "" {
			deletedFlagName = "deleted"
		} else {
			deletedFlagName = fmt.Sprintf("%v.deleted", cmdPrefix)
		}

		deletedFlagValue, err := cmd.Flags().GetBool(deletedFlagName)
		if err != nil {
			return err, false
		}
		m.Deleted = deletedFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGCHistoryIDFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

func retrieveGCHistoryJobKindFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	jobKindFlagName := fmt.Sprintf("%v.job_kind", cmdPrefix)
	if cmd.Flags().Changed(jobKindFlagName) {

		var jobKindFlagName string
		if cmdPrefix == "" {
			jobKindFlagName = "job_kind"
		} else {
			jobKindFlagName = fmt.Sprintf("%v.job_kind", cmdPrefix)
		}

		jobKindFlagValue, err := cmd.Flags().GetString(jobKindFlagName)
		if err != nil {
			return err, false
		}
		m.JobKind = jobKindFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGCHistoryJobNameFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	jobNameFlagName := fmt.Sprintf("%v.job_name", cmdPrefix)
	if cmd.Flags().Changed(jobNameFlagName) {

		var jobNameFlagName string
		if cmdPrefix == "" {
			jobNameFlagName = "job_name"
		} else {
			jobNameFlagName = fmt.Sprintf("%v.job_name", cmdPrefix)
		}

		jobNameFlagValue, err := cmd.Flags().GetString(jobNameFlagName)
		if err != nil {
			return err, false
		}
		m.JobName = jobNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGCHistoryJobParametersFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	jobParametersFlagName := fmt.Sprintf("%v.job_parameters", cmdPrefix)
	if cmd.Flags().Changed(jobParametersFlagName) {

		var jobParametersFlagName string
		if cmdPrefix == "" {
			jobParametersFlagName = "job_parameters"
		} else {
			jobParametersFlagName = fmt.Sprintf("%v.job_parameters", cmdPrefix)
		}

		jobParametersFlagValue, err := cmd.Flags().GetString(jobParametersFlagName)
		if err != nil {
			return err, false
		}
		m.JobParameters = jobParametersFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGCHistoryJobStatusFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	jobStatusFlagName := fmt.Sprintf("%v.job_status", cmdPrefix)
	if cmd.Flags().Changed(jobStatusFlagName) {

		var jobStatusFlagName string
		if cmdPrefix == "" {
			jobStatusFlagName = "job_status"
		} else {
			jobStatusFlagName = fmt.Sprintf("%v.job_status", cmdPrefix)
		}

		jobStatusFlagValue, err := cmd.Flags().GetString(jobStatusFlagName)
		if err != nil {
			return err, false
		}
		m.JobStatus = jobStatusFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveGCHistoryScheduleFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scheduleFlagName := fmt.Sprintf("%v.schedule", cmdPrefix)
	if cmd.Flags().Changed(scheduleFlagName) {
		// info: complex object schedule ScheduleObj is retrieved outside this Changed() block
	}
	scheduleFlagValue := m.Schedule
	if swag.IsZero(scheduleFlagValue) {
		scheduleFlagValue = &models.ScheduleObj{}
	}

	err, scheduleAdded := retrieveModelScheduleObjFlags(depth+1, scheduleFlagValue, scheduleFlagName, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scheduleAdded
	if scheduleAdded {
		m.Schedule = scheduleFlagValue
	}

	return nil, retAdded
}

func retrieveGCHistoryUpdateTimeFlags(depth int, m *models.GCHistory, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	updateTimeFlagName := fmt.Sprintf("%v.update_time", cmdPrefix)
	if cmd.Flags().Changed(updateTimeFlagName) {

		var updateTimeFlagName string
		if cmdPrefix == "" {
			updateTimeFlagName = "update_time"
		} else {
			updateTimeFlagName = fmt.Sprintf("%v.update_time", cmdPrefix)
		}

		updateTimeFlagValueStr, err := cmd.Flags().GetString(updateTimeFlagName)
		if err != nil {
			return err, false
		}
		var updateTimeFlagValue strfmt.DateTime
		if err := updateTimeFlagValue.UnmarshalText([]byte(updateTimeFlagValueStr)); err != nil {
			return err, false
		}
		m.UpdateTime = updateTimeFlagValue

		retAdded = true
	}

	return nil, retAdded
}
