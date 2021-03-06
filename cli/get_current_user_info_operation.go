// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/user"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationUserGetCurrentUserInfoCmd returns a cmd to handle operation getCurrentUserInfo
func makeOperationUserGetCurrentUserInfoCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getCurrentUserInfo",
		Short: ``,
		RunE:  runOperationUserGetCurrentUserInfo,
	}

	if err := registerOperationUserGetCurrentUserInfoParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationUserGetCurrentUserInfo uses cmd flags to call endpoint api
func runOperationUserGetCurrentUserInfo(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := user.NewGetCurrentUserInfoParams()
	if err, _ := retrieveOperationUserGetCurrentUserInfoXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationUserGetCurrentUserInfoResult(appCli.User.GetCurrentUserInfo(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationUserGetCurrentUserInfoParamFlags registers all flags needed to fill params
func registerOperationUserGetCurrentUserInfoParamFlags(cmd *cobra.Command) error {
	if err := registerOperationUserGetCurrentUserInfoXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationUserGetCurrentUserInfoXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationUserGetCurrentUserInfoXRequestIDFlag(m *user.GetCurrentUserInfoParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationUserGetCurrentUserInfoResult parses request result and return the string content
func parseOperationUserGetCurrentUserInfoResult(resp0 *user.GetCurrentUserInfoOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*user.GetCurrentUserInfoOK)
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
		resp1, ok := iResp1.(*user.GetCurrentUserInfoUnauthorized)
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
		resp2, ok := iResp2.(*user.GetCurrentUserInfoInternalServerError)
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
