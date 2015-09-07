package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func DoClusterRemove(c *cli.Context) {
	args := c.Args()
	if len(args) != 1 {
		showHelpAndExit(c, fmt.Sprint("<cluster name> must be specified"))
	}

	name := args[0]
	cluster, err := NewCluster(name)
	if err != nil {
		log.WithField("err", err).Fatal("failed to get a cluster")
	}

	log.WithField("name", name).Info("remove a cluster")
	if err := os.RemoveAll(cluster.Path()); err != nil {
		log.WithField("err", err).Fatal("failed to remove a cluster")
	}
}

func ClusterRemoveCompletion(c *cli.Context) {

	if len(c.Args()) == 0 {
		for _, c := range AllClusters() {
			fmt.Println(c.Name)
		}
	}
}
