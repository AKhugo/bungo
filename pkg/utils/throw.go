package utils

import "fmt"


func ThrowError(err error, message string) (error) {


	if err != nil {
		return fmt.Errorf("error occured while %s: %s", message, err.Error());
	}

	return nil;
}