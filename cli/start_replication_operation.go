// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"

	"github.com/laidbackware/harbor-cli/client/replication"
	"github.com/laidbackware/harbor-cli/models"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

// makeOperationReplicationStartReplicationCmd returns a cmd to handle operation startReplication
func makeOperationReplicationStartReplicationCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "startReplication",
		Short: `Start one replication execution according to the policy`,
		RunE:  runOperationReplicationStartReplication,
	}

	if err := registerOperationReplicationStartReplicationParamFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

// runOperationReplicationStartReplication uses cmd flags to call endpoint api
func runOperationReplicationStartReplication(cmd *cobra.Command, args []string) error {
	appCli, err := makeClient(cmd, args)
	if err != nil {
		return err
	}
	// retrieve flag values from cmd and fill params
	params := replication.NewStartReplicationParams()
	if err, _ := retrieveOperationReplicationStartReplicationXRequestIDFlag(params, "", cmd); err != nil {
		return err
	}
	if err, _ := retrieveOperationReplicationStartReplicationExecutionFlag(params, "", cmd); err != nil {
		return err
	}
	if dryRun {

		logDebugf("dry-run flag specified. Skip sending request.")
		return nil
	}
	// make request and then print result
	msgStr, err := parseOperationReplicationStartReplicationResult(appCli.Replication.StartReplication(params, nil))
	if err != nil {
		return err
	}
	if !debug {

		fmt.Println(msgStr)
	}
	return nil
}

// registerOperationReplicationStartReplicationParamFlags registers all flags needed to fill params
func registerOperationReplicationStartReplicationParamFlags(cmd *cobra.Command) error {
	if err := registerOperationReplicationStartReplicationXRequestIDParamFlags("", cmd); err != nil {
		return err
	}
	if err := registerOperationReplicationStartReplicationExecutionParamFlags("", cmd); err != nil {
		return err
	}
	return nil
}

func registerOperationReplicationStartReplicationXRequestIDParamFlags(cmdPrefix string, cmd *cobra.Command) error {

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
func registerOperationReplicationStartReplicationExecutionParamFlags(cmdPrefix string, cmd *cobra.Command) error {

	var executionFlagName string
	if cmdPrefix == "" {
		executionFlagName = "execution"
	} else {
		executionFlagName = fmt.Sprintf("%v.execution", cmdPrefix)
	}

	_ = cmd.PersistentFlags().String(executionFlagName, "", "Optional json string for [execution]. The ID of policy that the execution belongs to")

	// add flags for body
	if err := registerModelStartReplicationExecutionFlags(0, "startReplicationExecution", cmd); err != nil {
		return err
	}

	return nil
}

func retrieveOperationReplicationStartReplicationXRequestIDFlag(m *replication.StartReplicationParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
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
func retrieveOperationReplicationStartReplicationExecutionFlag(m *replication.StartReplicationParams, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false
	if cmd.Flags().Changed("execution") {
		// Read execution string from cmd and unmarshal
		executionValueStr, err := cmd.Flags().GetString("execution")
		if err != nil {
			return err, false
		}

		executionValue := models.StartReplicationExecution{}
		if err := json.Unmarshal([]byte(executionValueStr), &executionValue); err != nil {
			return fmt.Errorf("cannot unmarshal execution string in models.StartReplicationExecution: %v", err), false
		}
		m.Execution = &executionValue
	}
	executionValueModel := m.Execution
	if swag.IsZero(executionValueModel) {
		executionValueModel = &models.StartReplicationExecution{}
	}
	err, added := retrieveModelStartReplicationExecutionFlags(0, executionValueModel, "startReplicationExecution", cmd)
	if err != nil {
		return err, false
	}
	if added {
		m.Execution = executionValueModel
	}
	if dryRun && debug {

		executionValueDebugBytes, err := json.Marshal(m.Execution)
		if err != nil {
			return err, false
		}
		logDebugf("Execution dry-run payload: %v", string(executionValueDebugBytes))
	}
	retAdded = retAdded || added

	return nil, retAdded
}

// parseOperationReplicationStartReplicationResult parses request result and return the string content
func parseOperationReplicationStartReplicationResult(resp0 *replication.StartReplicationCreated, respErr error) (string, error) {
	if respErr != nil {

		// Non schema case: warning startReplicationCreated is not supported

		var iResp1 interface{} = respErr
		resp1, ok := iResp1.(*replication.StartReplicationBadRequest)
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
		resp2, ok := iResp2.(*replication.StartReplicationUnauthorized)
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
		resp3, ok := iResp3.(*replication.StartReplicationForbidden)
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
		resp4, ok := iResp4.(*replication.StartReplicationInternalServerError)
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

	// warning: non schema response startReplicationCreated is not supported by go-swagger cli yet.

	return "", nil
}
