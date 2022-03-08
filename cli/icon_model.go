// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for Icon

// register flags to command
func registerModelIconFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerIconContent(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerIconContentType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerIconContent(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	contentDescription := `The base64 encoded content of the icon`

	var contentFlagName string
	if cmdPrefix == "" {
		contentFlagName = "content"
	} else {
		contentFlagName = fmt.Sprintf("%v.content", cmdPrefix)
	}

	var contentFlagDefault string

	_ = cmd.PersistentFlags().String(contentFlagName, contentFlagDefault, contentDescription)

	return nil
}

func registerIconContentType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	contentTypeDescription := `The content type of the icon`

	var contentTypeFlagName string
	if cmdPrefix == "" {
		contentTypeFlagName = "content-type"
	} else {
		contentTypeFlagName = fmt.Sprintf("%v.content-type", cmdPrefix)
	}

	var contentTypeFlagDefault string

	_ = cmd.PersistentFlags().String(contentTypeFlagName, contentTypeFlagDefault, contentTypeDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelIconFlags(depth int, m *models.Icon, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, contentAdded := retrieveIconContentFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contentAdded

	err, contentTypeAdded := retrieveIconContentTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || contentTypeAdded

	return nil, retAdded
}

func retrieveIconContentFlags(depth int, m *models.Icon, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contentFlagName := fmt.Sprintf("%v.content", cmdPrefix)
	if cmd.Flags().Changed(contentFlagName) {

		var contentFlagName string
		if cmdPrefix == "" {
			contentFlagName = "content"
		} else {
			contentFlagName = fmt.Sprintf("%v.content", cmdPrefix)
		}

		contentFlagValue, err := cmd.Flags().GetString(contentFlagName)
		if err != nil {
			return err, false
		}
		m.Content = contentFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveIconContentTypeFlags(depth int, m *models.Icon, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	contentTypeFlagName := fmt.Sprintf("%v.content-type", cmdPrefix)
	if cmd.Flags().Changed(contentTypeFlagName) {

		var contentTypeFlagName string
		if cmdPrefix == "" {
			contentTypeFlagName = "content-type"
		} else {
			contentTypeFlagName = fmt.Sprintf("%v.content-type", cmdPrefix)
		}

		contentTypeFlagValue, err := cmd.Flags().GetString(contentTypeFlagName)
		if err != nil {
			return err, false
		}
		m.ContentType = contentTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}