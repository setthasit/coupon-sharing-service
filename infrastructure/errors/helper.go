package errors

/*
	CompareError - compare the given error to set of errors.
	the function will check if the given error can be match with api error according to the list,
	and will return boolean which will return `true` if the error match with api error
	and will return error that match with the given error in the from of *APIError
*/
func CompareError(from error, to ...*APIError) (bool, *APIError) {
	apiErr, ok := from.(*APIError)
	if !ok {
		return false, nil
	}

	for _, t := range to {
		if apiErr.Error() == t.Error() {
			return true, apiErr
		}
	}

	return false, nil
}
