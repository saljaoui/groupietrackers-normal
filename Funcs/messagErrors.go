package Groupie_tracker

func ErrorsMessage() AllMessageErrors {
	return AllMessageErrors{
		NotFound:                 "Page Not Found",
		BadRequest:               "Bad Request",
		InternalError:            "Internal Server Error",
		DescriptionNotFound:      "Sorry, the page you are looking for does not exist. It might have been moved or deleted. Please check the URL or return to the homepage.",
		DescriptionBadRequest:    "The request could not be understood by the server due to malformed syntax. Please verify your input and try again.",
		DescriptionInternalError: "An unexpected error occurred on the server. We are working to resolve this issue. Please try again later.",
	}
}