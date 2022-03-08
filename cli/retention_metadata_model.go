// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for RetentionMetadata

// register flags to command
func registerModelRetentionMetadataFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerRetentionMetadataScopeSelectors(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionMetadataTagSelectors(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerRetentionMetadataTemplates(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerRetentionMetadataScopeSelectors(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: scope_selectors []*RetentionSelectorMetadata array type is not supported by go-swagger cli yet

	return nil
}

func registerRetentionMetadataTagSelectors(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: tag_selectors []*RetentionSelectorMetadata array type is not supported by go-swagger cli yet

	return nil
}

func registerRetentionMetadataTemplates(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	// warning: templates []*RetentionRuleMetadata array type is not supported by go-swagger cli yet

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelRetentionMetadataFlags(depth int, m *models.RetentionMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, scopeSelectorsAdded := retrieveRetentionMetadataScopeSelectorsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || scopeSelectorsAdded

	err, tagSelectorsAdded := retrieveRetentionMetadataTagSelectorsFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || tagSelectorsAdded

	err, templatesAdded := retrieveRetentionMetadataTemplatesFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || templatesAdded

	return nil, retAdded
}

func retrieveRetentionMetadataScopeSelectorsFlags(depth int, m *models.RetentionMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	scopeSelectorsFlagName := fmt.Sprintf("%v.scope_selectors", cmdPrefix)
	if cmd.Flags().Changed(scopeSelectorsFlagName) {
		// warning: scope_selectors array type []*RetentionSelectorMetadata is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRetentionMetadataTagSelectorsFlags(depth int, m *models.RetentionMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	tagSelectorsFlagName := fmt.Sprintf("%v.tag_selectors", cmdPrefix)
	if cmd.Flags().Changed(tagSelectorsFlagName) {
		// warning: tag_selectors array type []*RetentionSelectorMetadata is not supported by go-swagger cli yet
	}

	return nil, retAdded
}

func retrieveRetentionMetadataTemplatesFlags(depth int, m *models.RetentionMetadata, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	templatesFlagName := fmt.Sprintf("%v.templates", cmdPrefix)
	if cmd.Flags().Changed(templatesFlagName) {
		// warning: templates array type []*RetentionRuleMetadata is not supported by go-swagger cli yet
	}

	return nil, retAdded
}
