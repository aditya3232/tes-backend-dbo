package constant

// All constant for pagination and order
const (
	DefaultLimit = 10

	DefaultPage = 1

	DefaultSort = "id"

	DefaultOrder = "asc"

	DefaultCode = 200
)

// All constant message for response
const (
	DefaultMessage = "Success."

	SuccessMessage = "The request was processed successfully."

	FailedMessage = "The request failed to process."

	DataFound = "Data found."

	DataNotFound = "Data not found."

	CannotProcessRequest = "Cannot process request."

	InvalidRequest = "Invalid request."

	SuccessCreateData = "Successfully created new data."

	FailedCreateData = "Failed to create new data."

	SuccessGetData = "Successfully retrieved data."

	SuccessUpdateData = "Successfully updated data."

	FailedUpdateData = "Failed to update data."

	SuccessDeleteData = "Successfully deleted data."

	FailedDeleteData = "Failed to delete data."

	FailedUnauthorized = "Unauthorized."

	LoginSuccess = "Login success."

	LoginFailed = "Login failed."

	LogoutSuccess = "Logout success."

	LogoutFailed = "Logout failed."

	TokenNotValid = "Token not valid."
)
