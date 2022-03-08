// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/system_c_v_e_allowlist"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSystemcveAllowlistGetSystemCVEAllowlistCmd returns a cmd to handle operation getSystemCVEAllowlist
func makeOperationSystemcveAllowlistGetSystemCVEAllowlistCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getSystemCVEAllowlist",
		Short: `Get the system level allowlist of CVE.  This API can be called by all authenticated users.`,
		RunE:  runOperationSystemcveAllowlistGetSystemCVEAllowlist,
	}

	if err := registerOperationSystemcveAllowlistGetSystemCVEAllowlistParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSystemcveAllowlistGetSystemCVEAllowlist uses cmd flags to call endpoint api
func runOperationSystemcveAllowlistGetSystemCVEAllowlist(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := system_c_v_e_allowlist.NewGetSystemCVEAllowlistParams()
	if err, _ := retrieveOperationSystemcveAllowlistGetSystemCVEAllowlistXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSystemcveAllowlistGetSystemCVEAllowlistResult(appCli.SystemcveAllowlist.GetSystemCVEAllowlist(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSystemcveAllowlistGetSystemCVEAllowlistParamFlags registers all flags needed to fill params
func registerOperationSystemcveAllowlistGetSystemCVEAllowlistParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSystemcveAllowlistGetSystemCVEAllowlistXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSystemcveAllowlistGetSystemCVEAllowlistXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationSystemcveAllowlistGetSystemCVEAllowlistXRequestIDFlag(m *system_c_v_e_allowlist.GetSystemCVEAllowlistParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationSystemcveAllowlistGetSystemCVEAllowlistResult parses request result and return the string content
func parseOperationSystemcveAllowlistGetSystemCVEAllowlistResult(resp0 *system_c_v_e_allowlist.GetSystemCVEAllowlistOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*system_c_v_e_allowlist.GetSystemCVEAllowlistOK)
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
		resp1, ok := iResp1.(*system_c_v_e_allowlist.GetSystemCVEAllowlistUnauthorized)
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
		resp2, ok := iResp2.(*system_c_v_e_allowlist.GetSystemCVEAllowlistInternalServerError)
		if ok {
			if !swag.IsZero(resp2) && !swag.IsZero(resp2.Payload) {
				msgStr, err := json.Marshal(resp2.Payload)
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
