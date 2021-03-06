package main

import "github.com/codegangsta/cli"

var tasks = []string{"cook", "clean", "laundry", "eat", "sleep", "code"}

var commandHelps = map[string]string{
	"init":      "[-b BASE_PATH]",
	"clone":     "[--] [<git clone option>...]",
	"available": "",
	"install":   "[-n NAME] [-d] [-p] <tag|branch|commit> <configure option>...]",
	"list":      "[-f pretty|plain|json] [-d]",
	"uninstall": "<version>",
}

var clusterSubcommandHelps = map[string]string{
	"create": "[-n NAME] <version> [<initdb option>...]",
	"list":   "[-f pretty|plain|json] [-d]",
	"start":  "[-p PORT|--find-free-port] <cluster>",
	"stop":   "<cluster>",
	"psql":   "<cluster> [<psql options>...]",
	"env":    "<cluster>",
	"remove": "<cluster>",
	"edit":   "<cluster> <file>",
}

var commands = []cli.Command{
	initCommand,
	cloneCommand,
	updateCommand,
	availableCommand,
	installCommand,
	listCommand,
	uninstallCommand,
	clusterCommand,
}

var initCommand = cli.Command{
	Name:        "init",
	Usage:       "Initialize pgenv environment",
	Description: `During initialization process, a config file will be created at ~/.pgenv/config.json. Although this file will be created under the base directory as a default, the path of the config file is independent from a pgenv base directory, and not configurable.`,
	Action:      DoInit,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "base-path,b",
			Usage: "Path of pgenv base directory. default: ~/.pgenv",
		},
		cli.StringFlag{
			Name:  "repository-path,r",
			Usage: "Path of PostgreSQL git repository. If this parameter is not set, `init` command will clone it from the official remote repository.",
		},
	},
}

var cloneCommand = cli.Command{
	Name:        "clone",
	Usage:       "Clone postgresql git repository",
	Description: `Execute git clone to retrive a clone of the official postgresql git repository. You can set any git option by passing options after '--'`,
	Action:      DoClone,
}

var updateCommand = cli.Command{
	Name:   "update",
	Usage:  "Update a local git repository",
	Action: DoUpdate,
}

var availableCommand = cli.Command{
	Name:   "available",
	Usage:  "List available versions",
	Action: DoAvailable,
}

var installCommand = cli.Command{
	Name:  "install",
	Usage: "Build and install a specified version",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name,n",
			Usage: "A name for this installation. a default value is a version number in x.y.z format.",
		},
		cli.BoolFlag{
			Name:  "debug,d",
			Usage: "Enable debug build swith (i.e. --enable-debug --enable-cassert)",
		},
		cli.BoolFlag{
			Name:  "parallel,p",
			Usage: "Allow multiple jobs of make command",
		},
	},
	Action:       DoInstall,
	BashComplete: InstallCompletion,
}

var listCommand = cli.Command{
	Name:  "list",
	Usage: "List installed versions",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format,f",
			Usage: "Output format. available option: pretty plain json. default: pretty.",
		},
		cli.BoolFlag{
			Name:  "detail,d",
			Usage: "Extend output with detailed information",
		},
	},
	Action: DoList,
}

var uninstallCommand = cli.Command{
	Name:         "uninstall",
	Usage:        "Uninstall a specified version",
	Action:       DoUninstall,
	BashComplete: UninstallCompletion,
}

var clusterCommands = []cli.Command{
	clusterCreateCommand,
	clusterRemoveCommand,
	clusterListCommand,
	clusterEnvCommand,
	clusterStartCommand,
	clusterStopCommand,
	clusterPsqlCommand,
	clusterEditCommand,
}

var clusterCommand = cli.Command{
	Name:        "cluster",
	Usage:       "Manage PostgreSQL clusters",
	Subcommands: clusterCommands,
	Before: func(c *cli.Context) error {
		args := c.Args()
		if len(args) > 0 {
			updateCommandHelp(args[0], clusterSubcommandHelps)
		}
		return nil
	},
}

var clusterCreateCommand = cli.Command{
	Name:    "create",
	Aliases: []string{"initdb"},
	Usage:   "Execute initdb to create a new cluster",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "name,n",
			Usage: "A name for this cluster. A specified <version> is used as a default value.",
		},
	},
	Action:       DoClusterCreate,
	BashComplete: ClusterCreateCompletion,
}

var clusterRemoveCommand = cli.Command{
	Name:         "remove",
	Usage:        "Remove a specified cluster",
	Action:       DoClusterRemove,
	BashComplete: ClusterRemoveCompletion,
}

var clusterListCommand = cli.Command{
	Name:  "list",
	Usage: "List clusters",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "format,f",
			Usage: "Output format. available option: pretty plain json. default: pretty.",
		},
		cli.BoolFlag{
			Name:  "detail,d",
			Usage: "Extend output with detailed information",
		},
	},
	Action: DoClusterList,
}

var clusterEnvCommand = cli.Command{
	Name:         "env",
	Usage:        "Show shell scripts to enable a specified cluster",
	Action:       DoClusterEnv,
	BashComplete: ClusterEnvCompletion,
}

var clusterStartCommand = cli.Command{
	Name:  "start",
	Usage: "Start a postgresql process with a specified cluster",
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "port,p",
			Usage: "A port that a postgresql process will listen.",
		},
		cli.BoolFlag{
			Name:  "find-free-port,f",
			Usage: "Find a random free port and use it as a listening port",
		},
	},
	Action:       DoClusterStart,
	BashComplete: ClusterStartCompletion,
}

var clusterStopCommand = cli.Command{
	Name:         "stop",
	Usage:        "Stop a postgresql process with a specified cluster",
	Action:       DoClusterStop,
	BashComplete: ClusterStopCompletion,
}

var clusterPsqlCommand = cli.Command{
	Name:            "psql",
	Usage:           "Run psql for a specified cluster",
	Action:          DoClusterPsql,
	BashComplete:    ClusterPsqlCompletion,
	SkipFlagParsing: true,
}

var clusterEditCommand = cli.Command{
	Name:         "edit",
	Usage:        "Edit files in a cluster directory",
	Action:       DoClusterEdit,
	BashComplete: ClusterEditCompletion,
}
