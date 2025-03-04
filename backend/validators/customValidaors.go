package validators

import (
	"fmt"
	"libraryManagement/internal/dto"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Global validator instance
var Validate *validator.Validate

func init() {
	fmt.Println("initialized validators")
	Validate = validator.New()
	_ = Validate.RegisterValidation("alpha_space", ValidateName)
	_ = Validate.RegisterValidation("password", ValidatePassword)
	_ = Validate.RegisterValidation("isbn", IsValidISBN)
	_ = Validate.RegisterValidation("phone", ValidatePhoneNumber)
}

func ValidateName(fl validator.FieldLevel) bool {
	nameRegex := regexp.MustCompile(`^[A-Za-z\s]+$`)
	fmt.Println(fl.Field().String())
	return nameRegex.MatchString(fl.Field().String())
}
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpper {
		return false
	}

	// Check for at least one lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !hasLower {
		return false
	}

	// Check for at least one digit
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	if !hasDigit {
		return false
	}

	// Check for at least one special character
	hasSpecial := regexp.MustCompile(`[@$!%*?&]`).MatchString(password)

	if !hasSpecial {
		return false
	}

	// If all checks pass, the password is valid
	return true
}

func ValidatePhoneNumber(fl validator.FieldLevel) bool {
	// Regular expression to match 10-digit Indian numbers starting with 6-9
	re := regexp.MustCompile(`^[6789]\d{9}$`)
	return re.MatchString(fl.Field().String())
}

// Custom ISBN Validator
func IsValidISBN(fl validator.FieldLevel) bool {
	var isbnRegex = regexp.MustCompile(`^(97(8|9))?\d{9}(\d|X)$`)

	return isbnRegex.MatchString(fl.Field().String())
}

func IsValidateBook(book dto.RequestUpdateBook) bool {

	nameRegex := regexp.MustCompile(`^[A-Za-z\s]+$`)
	validateTitle := nameRegex.MatchString(book.Title)
	valdiatePublisher := nameRegex.MatchString(book.Title)
	validateAuthor := nameRegex.MatchString(book.Authors)

	if book.Publisher == "" {
		valdiatePublisher = true
	}
	if book.Authors == "" {
		validateAuthor = true
	}
	if book.Title == "" {
		validateTitle = true
	}
	fmt.Println(validateTitle, valdiatePublisher, validateAuthor)
	return valdiatePublisher && validateAuthor && validateTitle
}

// init() runs automatically when the package is imported
