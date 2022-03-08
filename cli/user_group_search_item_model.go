// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for UserGroupSearchItem

// register flags to command
func registerModelUserGroupSearchItemFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerUserGroupSearchItemGroupName(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserGroupSearchItemGroupType(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerUserGroupSearchItemID(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerUserGroupSearchItemGroupName(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	groupNameDescription := `The name of the user group`

	var groupNameFlagName string
	if cmdPrefix == "" {
		groupNameFlagName = "group_name"
	} else {
		groupNameFlagName = fmt.Sprintf("%v.group_name", cmdPrefix)
	}

	var groupNameFlagDefault string

	_ = cmd.PersistentFlags().String(groupNameFlagName, groupNameFlagDefault, groupNameDescription)

	return nil
}

func registerUserGroupSearchItemGroupType(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	groupTypeDescription := `The group type, 1 for LDAP group, 2 for HTTP group.`

	var groupTypeFlagName string
	if cmdPrefix == "" {
		groupTypeFlagName = "group_type"
	} else {
		groupTypeFlagName = fmt.Sprintf("%v.group_type", cmdPrefix)
	}

	var groupTypeFlagDefault int64

	_ = cmd.PersistentFlags().Int64(groupTypeFlagName, groupTypeFlagDefault, groupTypeDescription)

	return nil
}

func registerUserGroupSearchItemID(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	idDescription := `The ID of the user group`

	var idFlagName string
	if cmdPrefix == "" {
		idFlagName = "id"
	} else {
		idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
	}

	var idFlagDefault int64

	_ = cmd.PersistentFlags().Int64(idFlagName, idFlagDefault, idDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelUserGroupSearchItemFlags(depth int, m *models.UserGroupSearchItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, groupNameAdded := retrieveUserGroupSearchItemGroupNameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || groupNameAdded

	err, groupTypeAdded := retrieveUserGroupSearchItemGroupTypeFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || groupTypeAdded

	err, idAdded := retrieveUserGroupSearchItemIDFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || idAdded

	return nil, retAdded
}

func retrieveUserGroupSearchItemGroupNameFlags(depth int, m *models.UserGroupSearchItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	groupNameFlagName := fmt.Sprintf("%v.group_name", cmdPrefix)
	if cmd.Flags().Changed(groupNameFlagName) {

		var groupNameFlagName string
		if cmdPrefix == "" {
			groupNameFlagName = "group_name"
		} else {
			groupNameFlagName = fmt.Sprintf("%v.group_name", cmdPrefix)
		}

		groupNameFlagValue, err := cmd.Flags().GetString(groupNameFlagName)
		if err != nil {
			return err, false
		}
		m.GroupName = groupNameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserGroupSearchItemGroupTypeFlags(depth int, m *models.UserGroupSearchItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	groupTypeFlagName := fmt.Sprintf("%v.group_type", cmdPrefix)
	if cmd.Flags().Changed(groupTypeFlagName) {

		var groupTypeFlagName string
		if cmdPrefix == "" {
			groupTypeFlagName = "group_type"
		} else {
			groupTypeFlagName = fmt.Sprintf("%v.group_type", cmdPrefix)
		}

		groupTypeFlagValue, err := cmd.Flags().GetInt64(groupTypeFlagName)
		if err != nil {
			return err, false
		}
		m.GroupType = groupTypeFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveUserGroupSearchItemIDFlags(depth int, m *models.UserGroupSearchItem, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	idFlagName := fmt.Sprintf("%v.id", cmdPrefix)
	if cmd.Flags().Changed(idFlagName) {

		var idFlagName string
		if cmdPrefix == "" {
			idFlagName = "id"
		} else {
			idFlagName = fmt.Sprintf("%v.id", cmdPrefix)
		}

		idFlagValue, err := cmd.Flags().GetInt64(idFlagName)
		if err != nil {
			return err, false
		}
		m.ID = idFlagValue

		retAdded = true
	}

	return nil, retAdded
}