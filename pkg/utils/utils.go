package utils

/**
 * CheckError is a helper function to check if an error is not nil, if it's not nil, it panics
 * So much better than writing if err != nil { panic(err) } every time (╯°□°）╯︵ ┻━┻
 */
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
