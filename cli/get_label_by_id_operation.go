// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/label"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationLabelGetLabelByIDCmd returns a cmd to handle operation getLabelById
func makeOperationLabelGetLabelByIDCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "GetLabelByID",
		Short: `This endpoint let user get the label by specific ID.
`,
		RunE: runOperationLabelGetLabelByID,
	}

	if err := registerOperationLabelGetLabelByIDParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLabelGetLabelByID uses cmd flags to call endpoint api
func runOperationLabelGetLabelByID(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := label.NewGetLabelByIDParams()
	if err, _ := retrieveOperationLabelGetLabelByIDXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLabelGetLabelByIDLabelIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLabelGetLabelByIDResult(appCli.Label.GetLabelByID(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLabelGetLabelByIDParamFlags registers all flags needed to fill params
func registerOperationLabelGetLabelByIDParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLabelGetLabelByIDXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLabelGetLabelByIDLabelIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLabelGetLabelByIDXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationLabelGetLabelByIDLabelIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	labelIdDescription := `Required. Label ID`

	var labelIdFlagName string
	if cmdPrefix == "" {
		labelIdFlagName = "label_id"
	} else {
		labelIdFlagName = fmt.Sprintf("%v.label_id", cmdPrefix)
	}

	var labelIdFlagDefault int64

	_ = cmd.PersistentFlags().Int64(labelIdFlagName, labelIdFlagDefault, labelIdDescription)

	return nil
}

func retrieveOperationLabelGetLabelByIDXRequestIDFlag(m *label.GetLabelByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationLabelGetLabelByIDLabelIDFlag(m *label.GetLabelByIDParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("label_id") {

		var labelIdFlagName string
		if cmdPrefix == "" {
			labelIdFlagName = "label_id"
		} else {
			labelIdFlagName = fmt.Sprintf("%v.label_id", cmdPrefix)
		}

		labelIdFlagValue, err := cmd.Flags().GetInt64(labelIdFlagName)
		if err != nil {
			return err, false
		}
		m.LabelID = labelIdFlagValue

	}
	return nil, retAdded
}

// parseOperationLabelGetLabelByIDResult parses request result and return the string content
func parseOperationLabelGetLabelByIDResult(resp0 *label.GetLabelByIDOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*label.GetLabelByIDOK)
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
		resp1, ok := iResp1.(*label.GetLabelByIDUnauthorized)
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
		resp2, ok := iResp2.(*label.GetLabelByIDNotFound)
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
		resp3, ok := iResp3.(*label.GetLabelByIDInternalServerError)
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

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}