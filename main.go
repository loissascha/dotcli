/*
Copyright © 2024 Sascha Loishandl <sascha.loishandl@gmail.com>
*/
package main

import "github.com/loissascha/dotcli/cmd"

// TODO
// init command initializes a new 'project' in the current (or selected) directory (some file like .dotcli gets created where all the necessary information gets written to)
// add command let's you add new folders and files to be synced
// sync comand syncs all the currently added folders and files from the local file system into the 'project' folder
// rm command let's you remove folders or files from the sync

func main() {
	cmd.Execute()
}
