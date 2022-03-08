// Code generated by go-swagger; DO NOT EDIT.

package cli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/laidbackware/harbor-cli/models"
	"github.com/spf13/cobra"
)

// Schema cli for LdapUser

// register flags to command
func registerModelLdapUserFlags(depth int, cmdPrefix string, cmd *cobra.Command) error {

	if err := registerLdapUserEmail(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLdapUserRealname(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	if err := registerLdapUserUsername(depth, cmdPrefix, cmd); err != nil {
		return err
	}

	return nil
}

func registerLdapUserEmail(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	emailDescription := `The user email address from 'mail' or 'email' attribute.`

	var emailFlagName string
	if cmdPrefix == "" {
		emailFlagName = "email"
	} else {
		emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
	}

	var emailFlagDefault string

	_ = cmd.PersistentFlags().String(emailFlagName, emailFlagDefault, emailDescription)

	return nil
}

func registerLdapUserRealname(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	realnameDescription := `The user realname from 'uid' or 'cn' attribute.`

	var realnameFlagName string
	if cmdPrefix == "" {
		realnameFlagName = "realname"
	} else {
		realnameFlagName = fmt.Sprintf("%v.realname", cmdPrefix)
	}

	var realnameFlagDefault string

	_ = cmd.PersistentFlags().String(realnameFlagName, realnameFlagDefault, realnameDescription)

	return nil
}

func registerLdapUserUsername(depth int, cmdPrefix string, cmd *cobra.Command) error {
	if depth > maxDepth {
		return nil
	}

	usernameDescription := `ldap username.`

	var usernameFlagName string
	if cmdPrefix == "" {
		usernameFlagName = "username"
	} else {
		usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
	}

	var usernameFlagDefault string

	_ = cmd.PersistentFlags().String(usernameFlagName, usernameFlagDefault, usernameDescription)

	return nil
}

// retrieve flags from commands, and set value in model. Return true if any flag is passed by user to fill model field.
func retrieveModelLdapUserFlags(depth int, m *models.LdapUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	retAdded := false

	err, emailAdded := retrieveLdapUserEmailFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || emailAdded

	err, realnameAdded := retrieveLdapUserRealnameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || realnameAdded

	err, usernameAdded := retrieveLdapUserUsernameFlags(depth, m, cmdPrefix, cmd)
	if err != nil {
		return err, false
	}
	retAdded = retAdded || usernameAdded

	return nil, retAdded
}

func retrieveLdapUserEmailFlags(depth int, m *models.LdapUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	emailFlagName := fmt.Sprintf("%v.email", cmdPrefix)
	if cmd.Flags().Changed(emailFlagName) {

		var emailFlagName string
		if cmdPrefix == "" {
			emailFlagName = "email"
		} else {
			emailFlagName = fmt.Sprintf("%v.email", cmdPrefix)
		}

		emailFlagValue, err := cmd.Flags().GetString(emailFlagName)
		if err != nil {
			return err, false
		}
		m.Email = emailFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLdapUserRealnameFlags(depth int, m *models.LdapUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	realnameFlagName := fmt.Sprintf("%v.realname", cmdPrefix)
	if cmd.Flags().Changed(realnameFlagName) {

		var realnameFlagName string
		if cmdPrefix == "" {
			realnameFlagName = "realname"
		} else {
			realnameFlagName = fmt.Sprintf("%v.realname", cmdPrefix)
		}

		realnameFlagValue, err := cmd.Flags().GetString(realnameFlagName)
		if err != nil {
			return err, false
		}
		m.Realname = realnameFlagValue

		retAdded = true
	}

	return nil, retAdded
}

func retrieveLdapUserUsernameFlags(depth int, m *models.LdapUser, cmdPrefix string, cmd *cobra.Command) (error, bool) {
	if depth > maxDepth {
		return nil, false
	}
	retAdded := false

	usernameFlagName := fmt.Sprintf("%v.username", cmdPrefix)
	if cmd.Flags().Changed(usernameFlagName) {

		var usernameFlagName string
		if cmdPrefix == "" {
			usernameFlagName = "username"
		} else {
			usernameFlagName = fmt.Sprintf("%v.username", cmdPrefix)
		}

		usernameFlagValue, err := cmd.Flags().GetString(usernameFlagName)
		if err != nil {
			return err, false
		}
		m.Username = usernameFlagValue

		retAdded = true
	}

	return nil, retAdded
}
