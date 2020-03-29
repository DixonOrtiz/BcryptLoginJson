package functions

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
)

var users []User

// User type which is a struct with two field
// User.Name string
// User.Password string
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// SelectOption function
// This functions is called in main and show a menu with options to the user
func SelectOption() {
	fmt.Printf("\n%s\n", color.YellowString("[selectOption() function CALLED]"))

	option := 10

	for option != 6 {
		ShowMenu()
		fmt.Printf("\t%s", color.WhiteString("Enter an option: "))
		_, _ = fmt.Scanf("%d", &option)

		switch option {
		case 1:
			u, p := RegisterUser()
			AppendUser(&users, u, p)

		case 2:
			userEntered, passwordEntered := GetUser()
			LoginUser(users, userEntered, passwordEntered)

		case 3:
			fmt.Println(users)

		case 5:
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()

		case 6:
			break

		}

	}

	fmt.Printf("%s\n\n", color.YellowString("[selectOption() function EXITED SUCCESSFULLY]"))
}

// ShowMenu function
// This function displays the options of the application's menu.
func ShowMenu() {
	fmt.Printf("\n%s\n", color.YellowString("[ShowMenu() function CALLED]"))

	fmt.Printf("\t%s\n", color.WhiteString("STRUCTS, JSON AND BCRYPT GO PROGRAM"))
	fmt.Printf("\t\t%s%s\n", color.GreenString("1)"), color.WhiteString("Register an user"))
	fmt.Printf("\t\t%s%s\n", color.GreenString("2)"), color.WhiteString("Login an user"))
	fmt.Printf("\t\t%s%s\n", color.GreenString("3)"), color.WhiteString("List all users"))
	fmt.Printf("\t\t%s%s\n", color.GreenString("4)"), color.WhiteString("Convert []User to Json an list"))
	fmt.Printf("\t\t%s%s\n", color.GreenString("5)"), color.WhiteString("Clean Screen"))
	fmt.Printf("\t\t%s%s\n", color.GreenString("6)"), color.WhiteString("Enter 6 to exit"))

	fmt.Printf("%s\n\n", color.YellowString("[ShowMenu() function EXITED SUCCESSFULLY]"))
}

// GetUser function
// This function get an user from de database.
func GetUser() (name, password string) {
	fmt.Printf("\n%s", color.YellowString("[GetUser() function CALLED]"))

	fmt.Printf("%s", color.GreenString("\n\tEnter your user:\n"))

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\t%s", color.WhiteString("\t\tEnter Name: "))
	name, _ = reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Printf("\t%s", color.WhiteString("\t\tEnter Password: "))
	password, _ = reader.ReadString('\n')
	password = strings.TrimSuffix(password, "\n")

	fmt.Printf("%s\n\n", color.YellowString("[GetUser() function EXITED SUCCESSFULLY]"))
	return name, password
}

// RegisterUser function.
// This function allows to register an user in the system.
func RegisterUser() (name, hashedPassword string) {
	fmt.Printf("\n%s", color.YellowString("[RegisterUser() function CALLED]"))

	fmt.Printf("%s", color.GreenString("\n\tRegister an user:\n"))

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\t%s", color.WhiteString("\t\tEnter Name: "))
	name, _ = reader.ReadString('\n')
	name = strings.TrimSuffix(name, "\n")

	fmt.Printf("\t%s", color.WhiteString("\t\tEnter Password: "))
	password, _ := reader.ReadString('\n')
	password = strings.TrimSuffix(password, "\n")

	hashedPassword = string(HashPassword(password))

	fmt.Printf("\t%s%s, %s%s\n", color.GreenString("User: {"), name, hashedPassword, color.GreenString("} added successfully!"))

	fmt.Printf("%s\n\n", color.YellowString("[RegisterUser() function EXITED SUCCESSFULLY]"))
	return name, hashedPassword

}

// HashPassword function.
// This function hash a password from string to encrypt data.
func HashPassword(password string) (hashedPassword []byte) {
	fmt.Printf("\n%s", color.YellowString("[HashPassword() function CALLED]"))
	cost := 4

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n\n", color.YellowString("\n[HashPassword() function EXITED SUCCESSFULLY]"))
	return hashedPassword
}

// AppendUser function.
// This function allows to append an user into de database.
func AppendUser(users *[]User, name, password string) {
	fmt.Printf("%s\n", color.YellowString("[AppendUser() function CALLED]"))
	user := User{
		Name:     name,
		Password: password,
	}

	*users = append(*users, user)
	fmt.Printf("%s\n\n", color.YellowString("[AppendUser() function EXITED SUCCESSFULLY]"))
}

// DidUserExists function.
// This function check if the introduce by the user exists in the database
func DidUserExists(users []User, name string) (b bool, i int) {
	fmt.Printf("\n%s", color.YellowString("[DidUserExists() function CALLED]\n"))
	b = false

	for i, user := range users {
		if user.Name == name {
			b = true

			fmt.Printf("%s\n\n", color.YellowString("[DidUserExists() function EXITED SUCCESSFULLY]"))
			return b, i
		}
	}

	fmt.Printf("%s\n\n", color.YellowString("[DidUserExists() function EXITED SUCCESSFULLY]"))

	i = -1
	return b, i
}

// LoginUser function.
//	This function allow to login in the application with valid credentials.
func LoginUser(users []User, name, password string) {
	fmt.Printf("\n%s", color.YellowString("[LoginUser() function CALLED]\n"))

	userExist, indexUser := DidUserExists(users, name)

	if userExist {

		err := bcrypt.CompareHashAndPassword([]byte(users[indexUser].Password), []byte(password))
		if err != nil {
			fmt.Println("You can't access :(")
			return
		}

		fmt.Printf("\t\t%s%s%s\n",
			color.WhiteString("Welcome back "),
			color.GreenString(name),
			color.WhiteString("!\n"),
		)

	} else {
		fmt.Println("The user you entered does not exist :(")
	}

	fmt.Printf("%s\n", color.YellowString("[LoginUser() function EXITED SUCCESSFULLY]"))

}
