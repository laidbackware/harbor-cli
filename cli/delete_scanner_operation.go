// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/scanner"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationScannerDeleteScannerCmd returns a cmd to handle operation deleteScanner
func makeOperationScannerDeleteScannerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "deleteScanner",
		Short: `Deletes the specified scanner registration.
`,
		RunE: runOperationScannerDeleteScanner,
	}

	if err := registerOperationScannerDeleteScannerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationScannerDeleteScanner uses cmd flags to call endpoint api
func runOperationScannerDeleteScanner(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := scanner.NewDeleteScannerParams()
	if err, _ := retrieveOperationScannerDeleteScannerXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationScannerDeleteScannerRegistrationIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationScannerDeleteScannerResult(appCli.Scanner.DeleteScanner(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationScannerDeleteScannerParamFlags registers all flags needed to fill params
func registerOperationScannerDeleteScannerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationScannerDeleteScannerXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationScannerDeleteScannerRegistrationIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationScannerDeleteScannerXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationScannerDeleteScannerRegistrationIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	registrationIdDescription := `Required. The scanner registration identifier.`

	var registrationIdFlagName string
	if cmdPrefix == "" {
		registrationIdFlagName = "registration_id"
	} else {
		registrationIdFlagName = fmt.Sprintf("%v.registration_id", cmdPrefix)
	}

	var registrationIdFlagDefault string

	_ = cmd.PersistentFlags().String(registrationIdFlagName, registrationIdFlagDefault, registrationIdDescription)

	return nil
}

func retrieveOperationScannerDeleteScannerXRequestIDFlag(m *scanner.DeleteScannerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationScannerDeleteScannerRegistrationIDFlag(m *scanner.DeleteScannerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("registration_id") {

		var registrationIdFlagName string
		if cmdPrefix == "" {
			registrationIdFlagName = "registration_id"
		} else {
			registrationIdFlagName = fmt.Sprintf("%v.registration_id", cmdPrefix)
		}

		registrationIdFlagValue, err := cmd.Flags().GetString(registrationIdFlagName)
		if err != nil {
			return err, false
		}
		m.RegistrationID = registrationIdFlagValue

	}
	return nil, retAdded
}

// parseOperationScannerDeleteScannerResult parses request result and return the string content
func parseOperationScannerDeleteScannerResult(resp0 *scanner.DeleteScannerOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*scanner.DeleteScannerOK)
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
		resp1, ok := iResp1.(*scanner.DeleteScannerUnauthorized)
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
		resp2, ok := iResp2.(*scanner.DeleteScannerForbidden)
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
		resp3, ok := iResp3.(*scanner.DeleteScannerNotFound)
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
		resp4, ok := iResp4.(*scanner.DeleteScannerInternalServerError)
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

	if !swag.IsZero(resp0) && !swag.IsZero(resp0.Payload) {
		msgStr, err := json.Marshal(resp0.Payload)
		if err != nil {
			return "", err
		}
		return string(msgStr), nil
	}

	return "", nil
}
