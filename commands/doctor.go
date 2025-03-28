package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/JscorpTech/jst-go/models"
	"github.com/spf13/cobra"
)

func Doctor() *cobra.Command {
	return &cobra.Command{
		Use: "create-doctor",
		Run: func(cmd *cobra.Command, args []string) {
			db := bootstrap.NewDatabase(bootstrap.NewEnv())
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Ism: ")
			firstName, _ := reader.ReadString('\n')
			fmt.Print("lastName: ")
			lastName, _ := reader.ReadString('\n')
			fmt.Print("phone: ")
			phone, _ := reader.ReadString('\n')
			fmt.Print("specialty: ")
			specialty, _ := reader.ReadString('\n')
			fmt.Print("workStart: ")
			workStartString, _ := reader.ReadString('\n')
			fmt.Print("workEnd: ")
			workEndString, _ := reader.ReadString('\n')
			loc, _ := time.LoadLocation("Asia/Tashkent")
			workStart, _ := time.ParseInLocation("15:04", strings.TrimSpace(workStartString), loc)
			workEnd, _ := time.ParseInLocation("15:04", strings.TrimSpace(workEndString), loc)
			db.Create(
				&models.Doctor{
					FirstName: strings.TrimSpace(firstName),
					LastName:  strings.TrimSpace(lastName),
					Phone:     strings.TrimSpace(phone),
					Specialty: strings.TrimSpace(specialty),
					WorkStart: workStart,
					WorkEnd:   workEnd,
				},
			)

		},
	}
}
