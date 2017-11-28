package main

import "github.com/heroku/herald"
// todo: name that herlad
import "time"
import "log"
import "fmt"



func main() {

	for {

		// Get a list of the buildpacks (as types).
		buildpacks := herald.GetBuildpacks()

		// Iterate over them.
		for _, bp := range(buildpacks) {

			// Download and extract each Buildpack.
			bp, path := bp.BPDownload()
			log.Printf("Buildpack '%s' downloaded to '%s'!", bp, path)

			// Find all version executables for the given buildpack. 
			executables := bp.FindVersionScripts()
			for _, exe := range(executables) {

				// TODO: Ensure chmod for the executable.
				exe.EnsureExecutable()

				// Execute the executable, print the results. 
				fmt.Print(exe.Execute())

				// 
				log.Printf("Executing '%s:%s' script…", bp, exe)
			}
		}

		log.Print("Sleeping for 10 minutes…")

		// Sleep for ten minutes. 
		time.Sleep(10*time.Minute)

		}

}