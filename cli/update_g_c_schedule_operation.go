// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/gc"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationGcUpdateGCScheduleCmd returns a cmd to handle operation updateGCSchedule
func makeOperationGcUpdateGCScheduleCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "updateGCSchedule",
		Short: `This endpoint is for update gc schedule.
`,
		RunE: runOperationGcUpdateGCSchedule,
	}

	if err := registerOperationGcUpdateGCScheduleParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationGcUpdateGCSchedule uses cmd flags to call endpoint api
func runOperationGcUpdateGCSchedule(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := gc.NewUpdateGCScheduleParams()
	if err, _ := retrieveOperationGcUpdateGCScheduleXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationGcUpdateGCScheduleScheduleFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationGcUpdateGCScheduleResult(appCli.Gc.UpdateGCSchedule(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationGcUpdateGCScheduleParamFlags registers all flags needed to fill params
func registerOperationGcUpdateGCScheduleParamFlags(cmd *cobra.Command) error {
	if err := registerOperationGcUpdateGCScheduleXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationGcUpdateGCScheduleScheduleParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationGcUpdateGCScheduleXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	xRequestIdDescription := `An unique ID for the request`

	var xRequestIdFlagName string
	if cmdPrefix == "" {
		xRequestIdFlagName = "X-Request-Id"
	} else {
		xRequestIdFlagName = fmt.Sprintf("%v.X-Request-Id", cmdPrefix)
	}

	var xRequestIdFlagDefault string

	_ = cmd.PersistentFlags().String(xRequestIdFlagName, xRequestIdFlagDefault, xRequestIdDescription)

	return nil
}
func registerOperationGcUpdateGCScheduleScheduleParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var scheduleFlagName string
	if cmdPrefix == "" {
		scheduleFlagName = "schedule"
	} else {
		scheduleFlagName = fmt.Sprintf("%v.schedule", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(scheduleFlagName, "", "Optional json string for [schedule]. Updates of gcs schedule.")

	// add flags for body
	if err := registerModelScheduleFlags(0, "schedule", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationGcUpdateGCScheduleXRequestIDFlag(m *gc.UpdateGCScheduleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("X-Request-Id") {

		var xRequestIdFlagName string
		if cmdPrefix == "" {
			xRequestIdFlagName = "X-Request-Id"
		} else {
			xRequestIdFlagName = fmt.Sprintf("%v.X-Request-Id", cmdPrefix)
		}

		xRequestIdFlagValue, err := cmd.Flags().GetString(xRequestIdFlagName)
		if err != nil {
			return err, false
		}
		m.XRequestID = &xRequestIdFlagValue

	}
	return nil, retAdded
}
func retrieveOperationGcUpdateGCScheduleScheduleFlag(m *gc.UpdateGCScheduleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("schedule") {
		// Read schedule string from cmd and unmarshal
		scheduleValueStr, err := cmd.Flags().GetString("schedule")
		if err != nil {
			return err, false
		}

		scheduleValue := models.Schedule{}
		if err := json.Unmarshal([]byte(scheduleValueStr), &scheduleValue); err != nil {
			return fmt.Errorf("cannot unmarshal schedule string in models.Schedule: %v", err), false
		}
		m.Schedule = &scheduleValue
	}
	scheduleValueModel := m.Schedule
	if swag.IsZero(scheduleValueModel) {
		scheduleValueModel = &models.Schedule{}
	}
	err, added := retrieveModelScheduleFlags(0, scheduleValueModel, "schedule", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Schedule = scheduleValueModel
	}
	if dryRun && debug {

		scheduleValueDebugBytes, err := json.Marshal(m.Schedule)
		if err != nil {
			return err, false
		}
		logDebugf("Schedule dry-run payload: %v", string(scheduleValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationGcUpdateGCScheduleResult parses request result and return the string content
func parseOperationGcUpdateGCScheduleResult(resp0 *gc.UpdateGCScheduleOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateGCScheduleOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*gc.UpdateGCScheduleBadRequest)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp2 interface{} = respErr
		resp2, ok := iResp2.(*gc.UpdateGCScheduleUnauthorized)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp3 interface{} = respErr
		resp3, ok := iResp3.(*gc.UpdateGCScheduleForbidden)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*gc.UpdateGCScheduleInternalServerError)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response updateGCScheduleOK is not supported by go-swagger cli yet.

	return "", nil
}
