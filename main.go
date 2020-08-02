package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
  
  "github.com/google/uuid"
)

// Profile holds sender / recievers contact information
type Profile struct {
	name    string
	address string
	phone   string
	email   string
}

// Activity / action to be charged for
type Activity struct {
	quantity    float64
	description string
	unitPrice   float64
	amount      float64
}

// HeaderInformation information about the invoice for internal use
type HeaderInformation struct {
	invoiceID      string
	invoiceDate    string
	invoiceDueDate string
}

// FooterInformation holds calculation about the invoice {total, tax, subtotal}
type FooterInformation struct {
	subtotal float64
	tax      float64
	total    float64
}

func main() {
	var senderProfile Profile           // Sender Profile
	var recieverProfile Profile         // Reciever Profile
	var Activities []Activity           // Array of activites that the user will be charged for
	var InvoiceHeader HeaderInformation // Invoice Header
	var InvoiceFooter FooterInformation // Invoice Footer

	// Getting Profile Information

	// Sender Information
	fmt.Println("Sender Information")
	GetProfileInformation(&senderProfile)

	// Reciever Information
	fmt.Println("Reciever Information")
	GetProfileInformation(&recieverProfile)

	// Getting Invoice Header Information
	GetInvoiceHeader(&InvoiceHeader)

	// Getting Activities
	subTotal := GetActivities(&Activities)

	// Calculating Footer Details
	// Setting SubTotal for invoice footer
	InvoiceFooter.subtotal = subTotal
	GetInvoiceFooter(&InvoiceFooter)
}

// GetProfileInformation gets user to file profile information
func GetProfileInformation(profile *Profile) {
	// If the user wants to confirm if the entered input is correct or they want to update it
	for {
		fmt.Printf("Name: ")
		profile.name = ReadString()

		fmt.Printf("Address: ")
		profile.address = ReadString()

		fmt.Printf("Email: ")
		profile.email = ReadString()

		fmt.Printf("Phone: ")
		profile.phone = ReadString()

		// Printing Profile
		PrintProfile(profile)

		// Information Confirmation Prompt
		fmt.Printf("Is the Entered Input Correct? [Y - Correct / N - Update Information]\n:") // prompt
		if ReadString() == "Y" {
			break
		}
	}
}

// GetInvoiceHeader get the header information
func GetInvoiceHeader(Header *HeaderInformation) {
	for {
		// Invoice ID
		fmt.Printf("Invoice ID [Leave Blank for random ID]: ")
		Header.invoiceID = ReadString()
		if Header.invoiceID == "" {
			// Generating Random UUID
			invoiceUUID, err := uuid.NewRandom()
			if err != nil {
				panic(err)
			}
			// UUID to Text/String
			invoiceID := invoiceUUID.String()
			Header.invoiceID = invoiceID
		}

		// Invoice Date
		fmt.Printf("Invoice Date [Leave Blank for Today's Date]: ")
		Header.invoiceDate = ReadString()
		if Header.invoiceDate == "" {
			Header.invoiceDate = time.Now().Format("01-01-2000") // Formatting Date: MM-DD-YYYY
		}

		// Invoice Due Date
		fmt.Println("Invoice Due Date: ")
		Header.invoiceDueDate = ReadString()

		// Information Confirmation Prompt
		fmt.Printf("Is the Entered Input Correct? [Y - Correct / N - Update Information]\n:") // prompt
		if ReadString() == "Y" {
			break
		}
	}

}

// GetActivities gets the activities the reciever will be charged for
func GetActivities(Activities *[]Activity) float64 {
	fmt.Print("How many activities would you like to enter?\n:")
	numOfActivities := ReadInteger() // reading integer
	var subTotal float64 = 0

	for i := 0; i < numOfActivities; i++ {
		var activity Activity
		for {
			fmt.Printf("Activity #%d\n", i+1) // Activity Count

			// Quantity
			fmt.Printf("Quantity: ")
			activity.quantity = ReadFloat()

			// Description
			fmt.Printf("Description: ")
			activity.description = ReadString()

			// Unit Price
			fmt.Printf("Unit Price: ")
			activity.unitPrice = ReadFloat()

			// Amount: unitPrice * quantity
			activity.amount = float64(activity.unitPrice * activity.quantity)

			// Printing Activity Information
			PrintActivity(activity)
			// Information Confirmation Prompt
			fmt.Printf("Is the Entered Input Correct? [Y - Correct / N - Update Information]\n:") // prompt
			if ReadString() == "Y" {
				subTotal += activity.amount // addding activity ammount to subTotal
				break
			}
		}
	}

	return subTotal
}

// GetInvoiceFooter calulate the requried fields for the footer
func GetInvoiceFooter(footer *FooterInformation) {
	for {
		// Getting Tax
		fmt.Printf("Enter Tax [No Percent Symbols]: ")
		footer.tax = ReadFloat() // getting tax value
		footer.total = footer.subtotal + ((footer.subtotal / 100) * footer.tax)

		// Print Footer Information
		PrintFooter(footer)

		// Information Confirmation Prompt
		fmt.Printf("Is the Entered Input Correct? [Y - Correct / N - Update Information]\n:") // prompt
		if ReadString() == "Y" {
			break
		}
	}
}

/*
	======= PRINT HELEPR =======
*/

// PrintProfile prints all the fields in the Profile struct
func PrintProfile(profile *Profile) {
	// Printing Name, Address, Email, Phone
	fmt.Printf("Name: %s\n", profile.name)
	fmt.Printf("Address: %s\n", profile.address)
	fmt.Printf("Email: %s\n", profile.email)
	fmt.Printf("Phone: %s\n", profile.phone)
}

// PrintHeader prints all the fields in the activity struct
func PrintHeader(header HeaderInformation) {
	fmt.Printf("Invoice ID: %s\n", header.invoiceID)
	fmt.Printf("Invoice Date: %s\n", header.invoiceDate)
	fmt.Printf("Invoice Due Date: %s\n1", header.invoiceDueDate)
}

// PrintActivity prints all the fields in the activity struct
func PrintActivity(activity Activity) {
	// Printing Quantity, Description, Unit Price, Amount
	fmt.Printf("Quantity: %.2f\n", activity.quantity)
	fmt.Printf("Description: %s\n", activity.description)
	fmt.Printf("Unit Price: %.2f\n", activity.unitPrice)
	fmt.Printf("Amount: %.2f\n", activity.amount)
}

// PrintFooter prints all the feilds in the activity struct
func PrintFooter(footer *FooterInformation) {
	fmt.Printf("Tax: %.2f\n", footer.tax)
	fmt.Printf("Tax: %.2f\n", footer.subtotal)
	fmt.Printf("Tax: %.2f\n", footer.total)
}

/*
	======= INPUT HELPERS =======
*/

// ReadString reads string / input from standard input
func ReadString() string {
	// Creating a reader to read lines
	reader := bufio.NewReader(os.Stdin)       // reading from standard input
	userInput, err := reader.ReadString('\n') // '\n' delimiter
	// If there was an error reading from standard input
	if err != nil {
		panic(err)
	}

	// Strip the String of '\r' & '\n'
	userInput = userInput[:len(userInput)-2]
	return userInput // returning userInput
}

// ReadInteger returns a integer read from standard input
func ReadInteger() int {
	// Creating a reader to read lines
	reader := bufio.NewReader(os.Stdin)       // stdin reader
	userInput, err := reader.ReadString('\n') // '\n' delimiter
	// If there was an error reading from userInput
	if err != nil {
		panic(err)
	}

	// Removing \n and checking if the entered input is an integer
	userInput = userInput[:len(userInput)-2] // removing '\r\n'
	num, err := strconv.Atoi(userInput)      // parsing string to integer
	if err != nil {
		panic(err)
	}
	return num // returning num
}

// ReadFloat returns a float64 from userInput
func ReadFloat() float64 {
	// Creating a reader to read lines
	reader := bufio.NewReader(os.Stdin)       // stdin reader
	userInput, err := reader.ReadString('\n') // '\n' delimiter
	// If there was an error reading from userInput
	if err != nil {
		panic(err)
	}

	// Removing \n and checking if the entered input is an integer
	userInput = userInput[:len(userInput)-2]        // removing '\r\n'
	float, err := strconv.ParseFloat(userInput, 64) // parsing string to float64
	if err != nil {
		panic(err)
	}

	return float // returning float

}
