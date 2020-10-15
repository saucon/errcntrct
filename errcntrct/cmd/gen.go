/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/Saucon/errcntrct/errcntrct/utils"
	"github.com/spf13/cobra"
)

var sourcepath string
var packagename string
var outputpath string

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Your JSON contract to const in GoLang",
	Long: `A tool to convert your JSON contract into golang. 
For example:

	errcntrct gen -i errorContract.json -o example/const.go

With this you dont need to declare const in go, just copy 
the generated code to your project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generated error JSON contract to const in go language")
		source, err := cmd.Flags().GetString("source")
		if err != nil {
			panic(err)
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			panic(err)
		}

		pkg, err := cmd.Flags().GetString("package")
		if err != nil {
			panic(err)
		}

		if err := utils.GenerateCodeFile(source, output, pkg, utils.TemplateJSONtoGolangConst); err != nil {
			panic(err)
		}

		fmt.Println("SUCCESS GENERATE FILE ..........ok \n")
		fmt.Println("JSON file resource path :\n",source)
		fmt.Println("Package in golang :\n",pkg)
		fmt.Println("File generated path :\n",output)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")
	genCmd.Flags().StringVarP(&sourcepath, "source", "s", "", "Source .json file ")
	rootCmd.MarkFlagRequired("source")
	genCmd.Flags().StringVarP(&outputpath, "output", "o", "", "Output const file ")
	rootCmd.MarkFlagRequired("output")
	genCmd.Flags().StringVarP(&packagename, "package", "p", "", "Package const file ")
	rootCmd.MarkFlagRequired("package")


	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

