// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/preheat"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationPreheatPingInstancesCmd returns a cmd to handle operation pingInstances
func makeOperationPreheatPingInstancesCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "PingInstances",
		Short: `This endpoint checks status of a instance, the instance can be given by ID or Endpoint URL (together with credential)
`,
		RunE: runOperationPreheatPingInstances,
	}

	if err := registerOperationPreheatPingInstancesParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationPreheatPingInstances uses cmd flags to call endpoint api
func runOperationPreheatPingInstances(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := preheat.NewPingInstancesParams()
	if err, _ := retrieveOperationPreheatPingInstancesXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationPreheatPingInstancesInstanceFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationPreheatPingInstancesResult(appCli.Preheat.PingInstances(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationPreheatPingInstancesParamFlags registers all flags needed to fill params
func registerOperationPreheatPingInstancesParamFlags(cmd *cobra.Command) error {
	if err := registerOperationPreheatPingInstancesXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationPreheatPingInstancesInstanceParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationPreheatPingInstancesXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationPreheatPingInstancesInstanceParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var instanceFlagName string
	if cmdPrefix == "" {
		instanceFlagName = "instance"
	} else {
		instanceFlagName = fmt.Sprintf("%v.instance", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(instanceFlagName, "", "Optional json string for [instance]. The JSON object of instance.")

	// add flags for body
	if err := registerModelInstanceFlags(0, "instance", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationPreheatPingInstancesXRequestIDFlag(m *preheat.PingInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationPreheatPingInstancesInstanceFlag(m *preheat.PingInstancesParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("instance") {
		// Read instance string from cmd and unmarshal
		instanceValueStr, err := cmd.Flags().GetString("instance")
		if err != nil {
			return err, false
		}

		instanceValue := models.Instance{}
		if err := json.Unmarshal([]byte(instanceValueStr), &instanceValue); err != nil {
			return fmt.Errorf("cannot unmarshal instance string in models.Instance: %v", err), false
		}
		m.Instance = &instanceValue
	}
	instanceValueModel := m.Instance
	if swag.IsZero(instanceValueModel) {
		instanceValueModel = &models.Instance{}
	}
	err, added := retrieveModelInstanceFlags(0, instanceValueModel, "instance", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Instance = instanceValueModel
	}
	if dryRun && debug {

		instanceValueDebugBytes, err := json.Marshal(m.Instance)
		if err != nil {
			return err, false
		}
		logDebugf("Instance dry-run payload: %v", string(instanceValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationPreheatPingInstancesResult parses request result and return the string content
func parseOperationPreheatPingInstancesResult(resp0 *preheat.PingInstancesOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning pingInstancesOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*preheat.PingInstancesBadRequest)
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
		resp2, ok := iResp2.(*preheat.PingInstancesUnauthorized)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		// Non schema case: warning pingInstancesNotFound is not supported

		var iResp4 interface{} = respErr
		resp4, ok := iResp4.(*preheat.PingInstancesInternalServerError)
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

	// warning: non schema response pingInstancesOK is not supported by go-swagger cli yet.

	return "", nil
}
