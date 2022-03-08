// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/scanner"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationScannerUpdateScannerCmd returns a cmd to handle operation updateScanner
func makeOperationScannerUpdateScannerCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "updateScanner",
		Short: `Updates the specified scanner registration.
`,
		RunE: runOperationScannerUpdateScanner,
	}

	if err := registerOperationScannerUpdateScannerParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationScannerUpdateScanner uses cmd flags to call endpoint api
func runOperationScannerUpdateScanner(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := scanner.NewUpdateScannerParams()
	if err, _ := retrieveOperationScannerUpdateScannerXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationScannerUpdateScannerRegistrationFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationScannerUpdateScannerRegistrationIDFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationScannerUpdateScannerResult(appCli.Scanner.UpdateScanner(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationScannerUpdateScannerParamFlags registers all flags needed to fill params
func registerOperationScannerUpdateScannerParamFlags(cmd *cobra.Command) error {
	if err := registerOperationScannerUpdateScannerXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationScannerUpdateScannerRegistrationParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationScannerUpdateScannerRegistrationIDParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationScannerUpdateScannerXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationScannerUpdateScannerRegistrationParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var registrationFlagName string
	if cmdPrefix == "" {
		registrationFlagName = "registration"
	} else {
		registrationFlagName = fmt.Sprintf("%v.registration", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(registrationFlagName, "", "Optional json string for [registration]. A scanner registraiton to be updated.")

	// add flags for body
	if err := registerModelScannerRegistrationReqFlags(0, "scannerRegistrationReq", cmd); err != nil {
		return err
	}

	return nil
}
func registerOperationScannerUpdateScannerRegistrationIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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

func retrieveOperationScannerUpdateScannerXRequestIDFlag(m *scanner.UpdateScannerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationScannerUpdateScannerRegistrationFlag(m *scanner.UpdateScannerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("registration") {
		// Read registration string from cmd and unmarshal
		registrationValueStr, err := cmd.Flags().GetString("registration")
		if err != nil {
			return err, false
		}

		registrationValue := models.ScannerRegistrationReq{}
		if err := json.Unmarshal([]byte(registrationValueStr), &registrationValue); err != nil {
			return fmt.Errorf("cannot unmarshal registration string in models.ScannerRegistrationReq: %v", err), false
		}
		m.Registration = &registrationValue
	}
	registrationValueModel := m.Registration
	if swag.IsZero(registrationValueModel) {
		registrationValueModel = &models.ScannerRegistrationReq{}
	}
	err, added := retrieveModelScannerRegistrationReqFlags(0, registrationValueModel, "scannerRegistrationReq", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Registration = registrationValueModel
	}
	if dryRun && debug {

		registrationValueDebugBytes, err := json.Marshal(m.Registration)
		if err != nil {
			return err, false
		}
		logDebugf("Registration dry-run payload: %v", string(registrationValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}
func retrieveOperationScannerUpdateScannerRegistrationIDFlag(m *scanner.UpdateScannerParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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

// parseOperationScannerUpdateScannerResult parses request result and return the string content
func parseOperationScannerUpdateScannerResult(resp0 *scanner.UpdateScannerOK, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning updateScannerOK is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*scanner.UpdateScannerUnauthorized)
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
		resp2, ok := iResp2.(*scanner.UpdateScannerForbidden)
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
		resp3, ok := iResp3.(*scanner.UpdateScannerNotFound)
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
		resp4, ok := iResp4.(*scanner.UpdateScannerInternalServerError)
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

	// warning: non schema response updateScannerOK is not supported by go-swagger cli yet.

	return "", nil
}
