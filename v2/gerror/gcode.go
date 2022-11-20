package gerror

// ================================================================================================================
// Common error code definition.
// There are reserved internal error code by framework: code < 1000.
// ================================================================================================================

var (
	CodeNil                      = LocalCode{"", "", 0}                                                               // No error code specified.
	CodeOK                       = LocalCode{code: "OK", message: "OK", httpStatus: 200}                              // It is OK.
	CodeInternalError            = LocalCode{"error.InternalError", "Internal Error", 500}                            // An error occurred internally.
	CodeValidationFailed         = LocalCode{"error.ValidationFailed", "Validation Failed", 400}                      // Data validation failed.
	CodeDbOperationError         = LocalCode{"error.CodeDbOperationError", "Database Operation Error", 500}           // Database operation error.
	CodeInvalidParameter         = LocalCode{"error.CodeInvalidParameter", "Invalid Parameter", 400}                  // The given parameter for current operation is invalid.
	CodeMissingParameter         = LocalCode{"error.CodeMissingParameter", "Missing Parameter", 400}                  // Parameter for current operation is missing.
	CodeInvalidOperation         = LocalCode{"error.CodeInvalidOperation", "Invalid Operation", 400}                  // The function cannot be used like this.
	CodeInvalidConfiguration     = LocalCode{"error.CodeInvalidConfiguration", "Invalid Configuration", 400}          // The configuration is invalid for current operation.
	CodeMissingConfiguration     = LocalCode{"error.CodeMissingConfiguration", "Missing Configuration", 400}          // The configuration is missing for current operation.
	CodeNotImplemented           = LocalCode{"error.CodeNotImplemented", "Not Implemented", 500}                      // The operation is not implemented yet.
	CodeNotSupported             = LocalCode{"error.CodeNotSupported", "Not Supported", 405}                          // The operation is not supported yet.
	CodeOperationFailed          = LocalCode{"error.CodeOperationFailed", "Operation Failed", 500}                    // I tried, but I cannot give you what you want.
	CodeNotAuthorized            = LocalCode{"error.CodeNotAuthorized", "Not Authorized", 401}                        // Not Authorized.
	CodeSecurityReason           = LocalCode{"error.CodeSecurityReason", "Security Reason", 403}                      // Security Reason.
	CodeServerBusy               = LocalCode{"error.CodeServerBusy", "Server Is Busy", 502}                           // Server is busy, please try again later.
	CodeUnknown                  = LocalCode{"error.CodeUnknown", "Unknown Error", 500}                               // Unknown error.
	CodeNotFound                 = LocalCode{"error.NotFound", "Not Found", 404}                                      // Resource does not exist.
	CodeInvalidRequest           = LocalCode{"error.CodeInvalidRequest", "Invalid Request", 400}                      // Invalid request.
	CodeBusinessValidationFailed = LocalCode{"error.CodeBusinessValidationFailed", "Business Validation Failed", 400} // Business validation failed.
)
