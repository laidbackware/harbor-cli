// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/scan_all"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationScanAllUpdateScanAllScheduleCmd returns a cmd to handle operation updateScanAllSchedule
func makeOperationScanAllUpdateScanAllScheduleCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "updateScanAllSchedule",
		Short: `This endpoint is for updating the schedule of scan all job, which scans all of images in Harbor.`,
		RunE:  runOperationScanAllUpdateScanAllSchedule,
	}

	if err := registerOperationScanAllUpdateScanAllScheduleParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationScanAllUpdateScanAllSchedule uses cmd flags to call endpoint api
func runOperationScanAllUpdateScanAllSchedule(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := scan_all.NewUpdateScanAllScheduleParams()
	if err, _ := retrieveOperationScanAllUpdateScanAllScheduleXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationScanAllUpdateScanAllScheduleScheduleFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationScanAllUpdateScanAllScheduleResult(appCli.ScanAll.UpdateScanAllSchedule(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationScanAllUpdateScanAllScheduleParamFlags registers all flags needed to fill params
func registerOperationScanAllUpdateScanAllScheduleParamFlags(cmd *cobra.Command) error {
	if err := registerOperationScanAllUpdateScanAllScheduleXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationScanAllUpdateScanAllScheduleScheduleParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationScanAllUpdateScanAllScheduleXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationScanAllUpdateScanAllScheduleScheduleParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var scheduleFlagName string
	if cmdPrefix == "" {
		scheduleFlagName = "schedule"
	} else {
		scheduleFlagName = fmt.Sprintf("%v.schedule", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(scheduleFlagName, "", "Optional json string for [schedule]. Updates the schedule of scan all job, which scans all of images in Harbor.")

	// add flags for body
	if err := registerModelScheduleFlags(0, "schedule", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationScanAllUpdateScanAllScheduleXRequestIDFlag(m *scan_all.UpdateScanAllScheduleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationScanAllUpdateScanAllScheduleScheduleFlag(m *scan_all.UpdateScanAllScheduleParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationScanAllUpdateScanAllScheduleResult parses request result and return the string content
func parseOperationScanAllUpdateScanAllScheduleResult(resp0 *scan_all.UpdateScanAllScheduleOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateScanAllScheduleOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*scan_all.UpdateScanAllScheduleBadRequest)
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
		resp2, ok := iResp2.(*scan_all.UpdateScanAllScheduleUnauthorized)
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
		resp3, ok := iResp3.(*scan_all.UpdateScanAllScheduleForbidden)
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
		resp4, ok := iResp4.(*scan_all.UpdateScanAllSchedulePreconditionFailed)
		if ok {
			if !swag.IsZero(resp4) && !swag.IsZero(resp4.Payload) {
				msgStr, err := json.Marshal(resp4.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp5 interface{} = respErr
		resp5, ok := iResp5.(*scan_all.UpdateScanAllScheduleInternalServerError)
		if ok {
			if !swag.IsZero(resp5) && !swag.IsZero(resp5.Payload) {
				msgStr, err := json.Marshal(resp5.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response updateScanAllScheduleOK is not supported by go-swagger cli yet.

	return "", nil
}
