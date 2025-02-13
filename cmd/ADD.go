/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// ADDCmd represents the ADD command
var ADDCmd = &cobra.Command{
	Use:   "ADD",
	Short: "Garvit-77 ADD <description>",
	Long:  `-`,
	Run: func(cmd *cobra.Command, args []string) {
		task := args[0]
		add(task)
	},
}

func add(task string) {
	writefile, err := os.OpenFile("tasks.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error Opening file")
	}
	defer writefile.Close()

	readfile, err := os.OpenFile("tasks.csv", os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error Opening Read file")
	}
	defer readfile.Close()

	reader := csv.NewReader(readfile)
	taskslice, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Your Csv is Unreadable: ", err)
		return
	}

	writer := csv.NewWriter(writefile)
	defer writer.Flush()

	if len(taskslice) == 0 {
		writer.Write([]string{
			"Task ID", "Desciprtion", "Done",
		})
	}

	id := strconv.FormatInt(time.Now().Unix(), 10)
	err = writer.Write([]string{
		id, task, "false",
	})
	if err != nil {
		fmt.Println("Your task wasn't added")
		return
	} else {
		fmt.Println("Successfully Added your task")
	}
}

func init() {
	rootCmd.AddCommand(ADDCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ADDCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ADDCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
