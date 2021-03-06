// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/ldap"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationLdapPingLdapCmd returns a cmd to handle operation pingLdap
func makeOperationLdapPingLdapCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use: "pingLdap",
		Short: `This endpoint ping the available ldap service for test related configuration parameters.
`,
		RunE: runOperationLdapPingLdap,
	}

	if err := registerOperationLdapPingLdapParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationLdapPingLdap uses cmd flags to call endpoint api
func runOperationLdapPingLdap(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := ldap.NewPingLdapParams()
	if err, _ := retrieveOperationLdapPingLdapXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationLdapPingLdapLdapconfFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationLdapPingLdapResult(appCli.Ldap.PingLdap(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationLdapPingLdapParamFlags registers all flags needed to fill params
func registerOperationLdapPingLdapParamFlags(cmd *cobra.Command) error {
	if err := registerOperationLdapPingLdapXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationLdapPingLdapLdapconfParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationLdapPingLdapXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationLdapPingLdapLdapconfParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var ldapconfFlagName string
	if cmdPrefix == "" {
		ldapconfFlagName = "ldapconf"
	} else {
		ldapconfFlagName = fmt.Sprintf("%v.ldapconf", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(ldapconfFlagName, "", "Optional json string for [ldapconf]. ldap configuration. support input ldap service configuration. If it is a empty request, will load current configuration from the system.")

	// add flags for body
	if err := registerModelLdapConfFlags(0, "ldapConf", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationLdapPingLdapXRequestIDFlag(m *ldap.PingLdapParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationLdapPingLdapLdapconfFlag(m *ldap.PingLdapParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("ldapconf") {
		// Read ldapconf string from cmd and unmarshal
		ldapconfValueStr, err := cmd.Flags().GetString("ldapconf")
		if err != nil {
			return err, false
		}

		ldapconfValue := models.LdapConf{}
		if err := json.Unmarshal([]byte(ldapconfValueStr), &ldapconfValue); err != nil {
			return fmt.Errorf("cannot unmarshal ldapconf string in models.LdapConf: %v", err), false
		}
		m.Ldapconf = &ldapconfValue
	}
	ldapconfValueModel := m.Ldapconf
	if swag.IsZero(ldapconfValueModel) {
		ldapconfValueModel = &models.LdapConf{}
	}
	err, added := retrieveModelLdapConfFlags(0, ldapconfValueModel, "ldapConf", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Ldapconf = ldapconfValueModel
	}
	if dryRun && debug {

		ldapconfValueDebugBytes, err := json.Marshal(m.Ldapconf)
		if err != nil {
			return err, false
		}
		logDebugf("Ldapconf dry-run payload: %v", string(ldapconfValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationLdapPingLdapResult parses request result and return the string content
func parseOperationLdapPingLdapResult(resp0 *ldap.PingLdapOK, respErr error) (string, error) {
	if respErr != nil {

		var iResp0 interface{} = respErr
		resp0, ok := iResp0.(*ldap.PingLdapOK)
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
		resp1, ok := iResp1.(*ldap.PingLdapBadRequest)
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
		resp2, ok := iResp2.(*ldap.PingLdapUnauthorized)
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
		resp3, ok := iResp3.(*ldap.PingLdapForbidden)
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
		resp4, ok := iResp4.(*ldap.PingLdapInternalServerError)
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
