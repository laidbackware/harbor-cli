// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/retention"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationRetentionGetRentenitionMetadataCmd returns a cmd to handle operation getRentenitionMetadata
func makeOperationRetentionGetRentenitionMetadataCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "getRentenitionMetadata",
		Short: `Get Retention Metadatas.`,
		RunE:  runOperationRetentionGetRentenitionMetadata,
	}

	if err := registerOperationRetentionGetRentenitionMetadataParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationRetentionGetRentenitionMetadata uses cmd flags to call endpoint api
func runOperationRetentionGetRentenitionMetadata(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := retention.NewGetRentenitionMetadataParams()
	if err, _ := retrieveOperationRetentionGetRentenitionMetadataXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationRetentionGetRentenitionMetadataResult(appCli.Retention.GetRentenitionMetadata(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationRetentionGetRentenitionMetadataParamFlags registers all flags needed to fill params
func registerOperationRetentionGetRentenitionMetadataParamFlags(cmd *cobra.Command) error {
	if err := registerOperationRetentionGetRentenitionMetadataXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationRetentionGetRentenitionMetadataXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationRetentionGetRentenitionMetadataXRequestIDFlag(m *retention.GetRentenitionMetadataParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationRetentionGetRentenitionMetadataResult parses request result and return the string content
func parseOperationRetentionGetRentenitionMetadataResult(resp0 *retention.GetRentenitionMetadataOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*retention.GetRentenitionMetadataOK)
		if ok {
			if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
				msgStr, err := json.Marshal(resp0.Payload)
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
