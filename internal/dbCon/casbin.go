package dbCon

import (
	"fmt"

	"github.com/Lavina-Tech-LLC/lavina-utils/llog"
)

type (
	rule struct {
		policyName string
		rules      [][]string
	}
)

var (
	rules = []rule{
		{
			"p",
			[][]string{
				{"noAuth", "netFree", "GET"},
				{"someUserGroup", "some API or ObjGroup", "GET"},
			},
		},
		{
			"g2",
			[][]string{
				{"/network/ip", "netFree"},
			},
		},
		{
			"g",
			[][]string{
				{"LavinaTest", "someUserGroup"},
			},
		},
	}
)

func PopulateCasbinDefaults() {
	e := GetCasbin

	for _ , rule:= range rules[0].rules{
		if ! e.HasPolicy(rule){
			_, err := e.AddPolicy(rule)
			if err != nil {
				llog.Error(fmt.Sprintf("Error g2"))
			}
		}
	}

	for _ , rule:= range rules[1].rules{
		if ! e.HasNamedGroupingPolicy("g2", rule){
			_, err := e.AddNamedGroupingPolicy("g2", rule)
			if err != nil {
				llog.Error(fmt.Sprintf("Error g2"))
			}
		}
	}

	for _ , rule:= range rules[2].rules{
		if ! e.HasNamedGroupingPolicy("g", rule){
			_, err := e.AddNamedGroupingPolicy("g", rule)
			if err != nil {
				llog.Error(fmt.Sprintf("Error g"))
			}
		}
	}

	e.SavePolicy()

	llog.Info("Succesfully updated policies..")

}