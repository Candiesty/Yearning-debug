package i18n

var en_US = map[int]string{
	ER_MYSQL_CONNECTION_FAILED:     "MySQL connection failed! Please check the configuration information",
	ER_QUERY_NO_DATA_SOURCE:        "No data source displayed for query",
	INFO_QUERY:                     "Query",
	INFO_TRANSFERRED_TO_NEXT_AGENT: "Transferred to next agent",
	INFO_SUBMITTED:                 "Submitted",
	ER_KEY_DECRYPTION_FAILED:       "Key decryption failed, unable to decrypt current password!",
	INFO_INITIALIZATION_SUCCESS_USERNAME_PASSWORD_RUN_COMMAND: "Initialization successful! Username: admin Password: Yearning_admin Please run ./Yearning run, default address: http://<host>:8000",
	INFO_ALREADY_INITIALIZED:                                  "Already initialized, please do not execute again",
	INFO_CHECKING_UPDATE:                                      "Checking for updates...",
	INFO_DATA_UPDATED:                                         "Data updated!",
	INFO_FIX_DESTRUCTIVE_CHANGE:                               "Fixing destructive change",
	INFO_FIX_SUCCESS:                                          "Fix successful!",
	INFO_YEARNING_MYSQL_AUDIT_PLATFORM:                        "Yearning MySQL Audit Platform",
	INFO_CONFIG_FILE_PATH:                                     "Configuration file path, default is conf.toml. If there is no need to move the configuration file, no configuration is required!",
	ER_DESTRUCTIVE_VERSION_UPGRADE_FIX:                        "Destructive version upgrade fix",
	INFO_RESET_SUPER_ADMIN_PASSWORD:                           "Reset super admin password",
	INFO_ADMIN_PASSWORD_RESET:                                 "Admin password has been reset to Yearning_admin",
	INFO_YEARNING_INSTALLATION_AND_DATA_INITIALIZATION:        "Yearning installation and data initialization",
	INFO_YEARNING_STARTUP:                                     "Start Yearning",
	INFO_YEARNING_STARTUP_PORT:                                "Yearning startup port",
	INFO_PLATFORM_ADDRESS:                                     "Platform address displayed in DingTalk/Email notifications",
	INFO_SPONSOR_KEY:                                          "Sponsor key",
	ER_LOGIN:                                                  "Account/password incorrect, please enter the correct account password!",
	ER_REGISTER:                                               "Registration channel not open!",
	ER_REQ_BIND:                                               "Parameter error!",
	ER_REQ_FAKE:                                               "Illegal parameter passing!",
	ER_REQ_PASSWORD_FAKE:                                      "Password change failed!",
	INFO_ORDER_IS_CREATE:                                      "Order has been created!",
	INFO_ORDER_IS_DUP:                                         "Do not submit duplicate orders!",
	INFO_ORDER_IS_EDIT:                                        "Order has been edited!",
	INFO_ORDER_IS_DELETE:                                      "Order has been deleted!",
	INFO_ORDER_IS_CLEAR:                                       "Order has been cleared",
	INFO_ORDER_IS_AGREE:                                       "Order has been approved",
	INFO_ORDER_IS_UNDO:                                        "Order has been revoked",
	INFO_ORDER_IS_REJECT:                                      "Order has been rejected",
	INFO_ORDER_IS_ALL_END:                                     "All orders have been terminated",
	INFO_ORDER_IS_END:                                         "Order has been terminated",
	INFO_ORDER_IS_ALL_CANCEL:                                  "All orders have been canceled",
	INFO_DATA_IS_DELETE:                                       "Data has been deleted!",
	INFO_DATA_IS_EDIT:                                         "Data has been edited!",
	INFO_DATA_IS_UPDATED:                                      "Data has been updated",
	INFO_ORDER_DELAY_SUCCESS:                                  "Order delay successful",
	INFO_RULE_IS_ADD:                                          "Rule has been added",
	INFO_RULE_IS_UPDATED:                                      "Rule has been edited",
	RULE_IS_DELETE:                                            "Rule has been deleted",
	INFO_LIBRARY_NAME_TABLE_NAME:                              "Please select the database name and table name before clicking to get the table structure information",
	ER_ILLEGAL_GET_INFO:                                       "Illegal information retrieval",
	UNDO_MESSAGE_ERROR:                                        "Order status has changed! Unable to undo",
	UNDO_MESSAGE_SUCCESS:                                      "Order has been undone!",
	COMMENT_IS_POST:                                           "Comment posted successfully",
	INFO_REGISTRATION_SUCCESS:                                 "Registration successful!",
	ER_USER_ALREADY_EXISTS:                                    "User already exists, please register again!",
	INFO_OIDC_LOGIN_DISABLED:                                  "OIDC login not enabled",
	CREATE_MESSAGE_SUCCESS:                                    "autoTask task added successfully!",
	CREATE_MESSAGE_ERROR:                                      "Do not add tasks with the same name repeatedly!",
	EDIT_MESSAGE_SUCCESS:                                      "AutoTask information updated!",
	EDIT_MESSAGE_ACTIVE:                                       "AutoTask order status updated!",
	CONN_TEST_SUCCESS:                                         "Database instance connected successfully!",
	DB_SAVE_SUCCESS:                                           "Connection name added successfully!",
	ERR_DB_SAVE:                                               "The SecretKey value in the config.toml file must be 16 characters!",
	DB_EDIT_SUCCESS:                                           "Data source information updated!",
	GROUP_DELETE_SUCCESS:                                      "Permission group ID: %s has been deleted",
	GROUP_CREATE_SUCCESS:                                      "%s permission group created/edited!",
	GROUP_EDIT_SUCCESS:                                        "%s permissions have been updated!",
	WEBHOOK_TEST:                                              "Test message sent! Please check your inbox!",
	MAIL_TEST:                                                 "Test email sent! Please check your inbox!",
	ERR_LDAP_TEST:                                             "LDAP connection failed, please check the configuration/test user password!",
	SUCCESS_LDAP_TEST:                                         "LDAP connection successful!",
	ER_MISSING_DATA_SOURCE:                                    "No data source added for workflow! Unable to submit order",
	ER_USER_REGUSTER:                                          "User already exists, please register again!",
	USER_REGUSTER_SUCCESS:                                     "Registration successful!",
	USER_DELETE_SUCCESS:                                       "User %s has been deleted",
	USER_EDIT_SUCCESS:                                         "User information modified successfully!",
	USER_EDIT_PASSWORD_SUCCESS:                                "Password changed successfully!",
	ADMIN_NOT_DELETE:                                          "admin user cannot be deleted!",
	ADMIN_HAVE_DELETE_OTHER:                                   "Non-admin users cannot delete other users",
	USER_PROLICY_EDIT_SUCCESS:                                 "%s permissions have been updated!",
	USER_CANNOT_DELETE:                                        "User %s is currently an auditor for workflow %s. Please delete the user auditor in the relevant node before deleting",
	BOARD_MESSAGE_SAVE:                                        "Announcement saved",
	ORDER_POST_SUCCESS:                                        "Order submitted, please wait for the auditor to review it!",
	CUSTOM_PASSWORD_SUCCESS:                                   "Personal information modified successfully!",
	ORDER_AGREE_MESSAGE:                                       "Approved and transferred to %s",
	ORDER_REJECT_MESSAGE:                                      "Rejected",
	ORDER_AGREE_STATE:                                         "Order has been transferred!",
	ORDER_REJECT_STATE:                                        "Order has been rejected!",
	ORDER_KILL_STATE:                                          "Delayed order has been terminated!",
	ORDER_EXECUTE_STATE:                                       "Approved and executed order!",
	ORDER_DELAY_KILL_DETAIL:                                   "Kill command has been sent! It will automatically cancel when the execution time is reached, and the status has been changed to execution failure!",
	ORDER_NOT_SEARCH:                                          "Someone has already approved this stage/You are not the auditor of this stage! Operation does not comply with idempotence",
	ER_MISSING_SQL_STATEMENT:                                  "No SQL statement is currently filled in",
	ER_BLOB_FIELD_NOT_DISPLAYABLE:                             "Blob field cannot be displayed",
	INFO_SENSITIVE_FIELD:                                      "****Sensitive field",
	INFO_QUERY_AUDIT_DISABLED:                                 "Query audit is currently disabled, users can freely query",
	ER_USER_NO_PERMISSION:                                     "User %s does not have permission for data source %s, unable to perform this operation",
	ER_DATABASE_CONNECTION_FAILED:                             "Database connection failed",
}
