package utils

import "github.com/32leaves/bel"

func ConvertToTSInterface(interfaceList []interface{}) {
	for _, _interface := range interfaceList {
		ts, err := bel.Extract(_interface)
		if err != nil {
			panic(err)
		}

		err = bel.Render(ts)
		if err != nil {
			panic(err)
		}
	}
}
