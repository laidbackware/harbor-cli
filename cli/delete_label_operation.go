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

// makeOperationLabelDeleteLabelCmd returns a cmd to handle operation deleteLabel
func makeOperationLabelDeleteLabelCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "DeleteLabel",
		Short: `Delete the label specified by ID.
`,
		RunE: runOperationLabelDeleteLabel,
	}

	if err := registerOperationLabelDeleteLabelParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLabelDeleteLabel uses cmd flags to call endpoint api
func runOperationLabelDeleteLabel(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := label.NewDeleteLabelParams()
	if err, _ := retrieveOperationLabelDeleteLabelXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLabelDeleteLabelLabelIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLabelDeleteLabelResult(appCli.Label.DeleteLabel(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLabelDeleteLabelParamFlags registers all flags needed to fill params
func registerOperationLabelDeleteLabelParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLabelDeleteLabelXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLabelDeleteLabelLabelIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLabelDeleteLabelXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationLabelDeleteLabelLabelIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationLabelDeleteLabelXRequestIDFlag(m *label.DeleteLabelParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationLabelDeleteLabelLabelIDFlag(m *label.DeleteLabelParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationLabelDeleteLabelResult parses request result and return the string content
func parseOperationLabelDeleteLabelResult(resp0 *label.DeleteLabelOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning deleteLabelOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*label.DeleteLabelBadRequest)
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
		resp2, ok := iResp2.(*label.DeleteLabelUnauthorized)
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
		resp3, ok := iResp3.(*label.DeleteLabelNotFound)
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
		resp4, ok := iResp4.(*label.DeleteLabelInternalServerError)
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

	// warning: non schema response deleteLabelOK is not supported by go-swagger cli yet.

	return "", nil
}
