package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	toUpperCuster := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + toUpperCuster
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return border + "\n" + welcomeMsg + "\n" + border
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	deleteStartString := strings.ReplaceAll(oldMsg, "*", " ")
	deleteSpaceString := strings.TrimSpace(deleteStartString)

	return deleteSpaceString
}
