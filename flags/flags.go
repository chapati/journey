package flags

import (
	"flag"
	"log"
)

var (
	Log     = ""
	LogFlag = "log"

	CustomPath     = ""
	CustomPathFlag = "custom-path"

	IsInDevMode     = false
	IsInDevModeFlag = "dev"

	HttpPort     = ""
	HttpPortFlag = "http-port"

	HttpsPort     = ""
	HttpsPortFlag = "https-port"

	CustomBuiltInPath     = ""
	CustomBuiltInPathFlag = "custom-built-in-path"
)

func init() {
	// Parse all flags
	parseFlags()
	if IsInDevMode {
		log.Println("Starting Journey in developer mode...")
	}
}

func parseFlags() {
	// Check if the log should be output to a file
	flag.StringVar(&Log, LogFlag, "", "Use this option to save to log output to a file. Note: Journey needs create, read, and write access to that file. Example: -log=path/to/log.txt")

	// Check if a custom content path has been provided by the user
	flag.StringVar(&CustomPath, CustomPathFlag, "", "Specify a custom path to store content files. Note: Journey needs read and write access to that path. A theme folder needs to be located in the custon path under content/themes. Example: -custom-path=/absolute/path/to/custom/folder")

	// Check if the dvelopment mode flag was provided by the user
	flag.BoolVar(&IsInDevMode, IsInDevModeFlag, false, "Use this flag flag to put Journey in developer mode. Features of developer mode: Themes and plugins will be recompiled immediately after changes to the files. Example: -dev")

	// Check if the http port that was set in the config was overridden by the user
	flag.StringVar(&HttpPort, HttpPortFlag, "", "Use this option to override the HTTP port that was set in the config.json. Example: -http-port=8080")

	// Check if the http port that was set in the config was overridden by the user
	flag.StringVar(&HttpsPort, HttpsPortFlag, "", "Use this option to override the HTTPS port that was set in the config.json. Example: -https-port=8081")

	// Check if custom built-in path has been provided by user
	flag.StringVar(&CustomBuiltInPath, CustomBuiltInPathFlag, "", "Specify a custom path to store builtin files. Read-only access is needed.")

	flag.Parse()
}
