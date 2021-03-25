package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	//"sort"

	"github.com/spf13/cobra"

	"github.com/dwburke/caterpillar/hash"
	"github.com/dwburke/caterpillar/util"
)

func init() {
	hashCmd.Flags().String("output", "", "file to save the json to (defaults to '<dir>.json')")
	hashCmd.Flags().Bool("write", false, "write json file (just generates hashes and compares with previous version by default)")

	rootCmd.AddCommand(hashCmd)
}

var hashCmd = &cobra.Command{
	Use:                   "hash <dir>",
	Short:                 "hash",
	Long:                  `hash`,
	Args:                  cobra.ExactArgs(1),
	Example:               "hash foo",
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		var differences bool
		var exit_code int

		dir := util.TrimSuffix(args[0], "/")
		dir, err := filepath.Abs(dir)
		if err != nil {
			return err
		}

		save_file := dir + ".json"

		output_file, _ := cmd.Flags().GetString("output")
		if output_file != "" {
			save_file = filepath.Clean(output_file)
		}

		old_files := make(map[string]*hash.FileData)

		_, err = os.Stat(save_file)
		if err != nil {
			if !os.IsNotExist(err) {
				return err
			}
		} else {
			old_files, err = hash.ReadJson(save_file)
			if err != nil {
				return err
			}
		}

		files, err := hash.HashTree(dir)
		if err != nil {
			return err
		}

		for k, v := range old_files {
			nfile, ok := files[k]
			if !ok {
				differences = true
				fmt.Printf("%s : removed\n", k)
				exit_code |= (1 << 3)
				continue
			}
			if nfile.Hash != v.Hash {
				differences = true
				fmt.Printf("%s : hash changed (%s)\n", k, nfile.Hash)
				exit_code |= (1 << 1)
			}
			if nfile.FileMode != v.FileMode {
				differences = true
				fmt.Printf("%s : filemode changed [%v] -> [%v]\n", k, v.FileMode, nfile.FileMode)
				exit_code |= (1 << 4)
			}
			if nfile.Permissions != v.Permissions {
				differences = true
				fmt.Printf("%s : file permissions changed [%v] -> [%v]\n", k, v.Permissions, nfile.Permissions)
				exit_code |= (1 << 5)
			}
		}

		//var names []string

		for k, v := range files {
			//names = append(names, k)
			if _, ok := old_files[k]; !ok {
				differences = true
				fmt.Printf("%s : added (%s)\n", v.Name, v.Hash)
				exit_code |= (1 << 2)
			}
		}

		//sort.Strings(names)
		//for _, n := range names {
		//v := files[n]
		//if v.Hash == "" {
		//continue
		//}
		//fmt.Printf("%32s %s\n", v.Hash, v.Name)
		//}

		if differences {
			if write_file, _ := cmd.Flags().GetBool("write"); write_file {
				if err = hash.SaveJson(save_file, files); err != nil {
					return err
				}
			}
		}

		os.Exit(exit_code)

		return nil
	},
}
