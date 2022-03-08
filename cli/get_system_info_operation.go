// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/systeminfo"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSysteminfoGetSystemInfoCmd returns a cmd to handle operation getSystemInfo
func makeOperationSysteminfoGetSystemInfoCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "getSystemInfo",
		Short: `This API is for retrieving general system info, this can be called by anonymous request.  Some attributes will be omitted in the response when this API is called by anonymous request.
`,
		RunE: runOperationSysteminfoGetSystemInfo,
	}

	if err := registerOperationSysteminfoGetSystemInfoParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSysteminfoGetSystemInfo uses cmd flags to call endpoint api
func runOperationSysteminfoGetSystemInfo(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := systeminfo.NewGetSystemInfoParams()
	if err, _ := retrieveOperationSysteminfoGetSystemInfoXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSysteminfoGetSystemInfoResult(appCli.Systeminfo.GetSystemInfo(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSysteminfoGetSystemInfoParamFlags registers all flags needed to fill params
func registerOperationSysteminfoGetSystemInfoParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSysteminfoGetSystemInfoXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSysteminfoGetSystemInfoXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSysteminfoGetSystemInfoXRequestIDFlag(m *systeminfo.GetSystemInfoParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSysteminfoGetSystemInfoResult parses request result and return the string content
func parseOperationSysteminfoGetSystemInfoResult(resp0 *systeminfo.GetSystemInfoOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*systeminfo.GetSystemInfoOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*systeminfo.GetSystemInfoInternalServerError)
		if ok {
			if !swag.IsZero(resp1) && !swag.IsZero(resp1.Payload) {
				msgStr, err := json.Marshal(resp1.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
