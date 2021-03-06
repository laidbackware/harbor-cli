// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/system_c_v_e_allowlist"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationSystemcveAllowlistPutSystemCVEAllowlistCmd returns a cmd to handle operation putSystemCVEAllowlist
func makeOperationSystemcveAllowlistPutSystemCVEAllowlistCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "putSystemCVEAllowlist",
		Short: `This API overwrites the system level allowlist of CVE with the list in request body.  Only system Admin has permission to call this API.`,
		RunE:  runOperationSystemcveAllowlistPutSystemCVEAllowlist,
	}

	if err := registerOperationSystemcveAllowlistPutSystemCVEAllowlistParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationSystemcveAllowlistPutSystemCVEAllowlist uses cmd flags to call endpoint api
func runOperationSystemcveAllowlistPutSystemCVEAllowlist(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := system_c_v_e_allowlist.NewPutSystemCVEAllowlistParams()
	if err, _ := retrieveOperationSystemcveAllowlistPutSystemCVEAllowlistXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationSystemcveAllowlistPutSystemCVEAllowlistAllowlistFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationSystemcveAllowlistPutSystemCVEAllowlistResult(appCli.SystemcveAllowlist.PutSystemCVEAllowlist(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationSystemcveAllowlistPutSystemCVEAllowlistParamFlags registers all flags needed to fill params
func registerOperationSystemcveAllowlistPutSystemCVEAllowlistParamFlags(cmd *cobra.Command) error {
	if err := registerOperationSystemcveAllowlistPutSystemCVEAllowlistXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationSystemcveAllowlistPutSystemCVEAllowlistAllowlistParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationSystemcveAllowlistPutSystemCVEAllowlistXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationSystemcveAllowlistPutSystemCVEAllowlistAllowlistParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var allowlistFlagName string
	if cmdPrefix == "" {
		allowlistFlagName = "allowlist"
	} else {
		allowlistFlagName = fmt.Sprintf("%v.allowlist", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(allowlistFlagName, "", "Optional json string for [allowlist]. The allowlist with new content")

	// add flags for body
	if err := registerModelCVEAllowlistFlags(0, "cVEAllowlist", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationSystemcveAllowlistPutSystemCVEAllowlistXRequestIDFlag(m *system_c_v_e_allowlist.PutSystemCVEAllowlistParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationSystemcveAllowlistPutSystemCVEAllowlistAllowlistFlag(m *system_c_v_e_allowlist.PutSystemCVEAllowlistParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("allowlist") {
		// Read allowlist string from cmd and unmarshal
		allowlistValueStr, err := cmd.Flags().GetString("allowlist")
		if err != nil {
			return err, false
		}

		allowlistValue := models.CVEAllowlist{}
		if err := json.Unmarshal([]byte(allowlistValueStr), &allowlistValue); err != nil {
			return fmt.Errorf("cannot unmarshal allowlist string in models.CVEAllowlist: %v", err), false
		}
		m.Allowlist = &allowlistValue
	}
	allowlistValueModel := m.Allowlist
	if swag.IsZero(allowlistValueModel) {
		allowlistValueModel = &models.CVEAllowlist{}
	}
	err, added := retrieveModelCVEAllowlistFlags(0, allowlistValueModel, "cVEAllowlist", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Allowlist = allowlistValueModel
	}
	if dryRun && debug {

		allowlistValueDebugBytes, err := json.Marshal(m.Allowlist)
		if err != nil {
			return err, false
		}
		logDebugf("Allowlist dry-run payload: %v", string(allowlistValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationSystemcveAllowlistPutSystemCVEAllowlistResult parses request result and return the string content
func parseOperationSystemcveAllowlistPutSystemCVEAllowlistResult(resp0 *system_c_v_e_allowlist.PutSystemCVEAllowlistOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning putSystemCVEAllowlistOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*system_c_v_e_allowlist.PutSystemCVEAllowlistUnauthorized)
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
		resp2, ok := iResp2.(*system_c_v_e_allowlist.PutSystemCVEAllowlistForbidden)
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
		resp3, ok := iResp3.(*system_c_v_e_allowlist.PutSystemCVEAllowlistInternalServerError)
		if ok {
			if !swag.IsZero(resp3) && !swag.IsZero(resp3.Payload) {
				msgStr, err := json.Marshal(resp3.Payload)
				if err != nil {
					return "", err
				}
				return string(msgStr), nil
			}
		}

		return "", respErr
	}

	// warning: non schema response putSystemCVEAllowlistOK is not supported by go-swagger cli yet.

	return "", nil
}
