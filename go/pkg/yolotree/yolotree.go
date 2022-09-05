package yolotree

import (
	"berty.tech/yolo/v2/go/pkg/yolopb"
	"fmt"
	"sort"
)

type artifact struct {
	ID      string
	file    string
	entity  string
	project string
	build   string
	date    string
}

func DisplayTreeFormat(d *yolopb.Batch) {
	/*	for _, e := range d.Entities {
		fmt.Printf(" - entities: %s\n", e.Name)
		for _, p := range d.Projects {
			if p.GetHasOwnerID() == e.ID {
				fmt.Printf(" - > - projects: %s\n", p.Name)
				for _, b := range d.Builds {
					if b.GetHasProjectID() == p.ID {
						fmt.Printf(" - > > - builds: %s\n", b.ShortID)
						for _, a := range d.Artifacts {
							if a.HasBuildID == b.ID {
								fmt.Printf(" - > > > - artifacts: %s\n", a.LocalPath)
							}
						}
					}
				}
			}
		}
	}*/
	sort.Slice(d.Builds, func(i, j int) bool {
		return d.Builds[i].FinishedAt.Before(*d.Builds[j].FinishedAt)
	})
	for _, b := range d.Builds {
		fmt.Printf(" - builds: %s from %s\n", b.ShortID, b.GetHasProjectID())
		for _, a := range d.Artifacts {
			if a.HasBuildID == b.ID {
				fmt.Printf(" - > artifacts: %s\n", a.LocalPath)
			}
		}
		fmt.Println()
	}
}
