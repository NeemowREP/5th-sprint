package actioninfo

import "log"

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, item := range dataset {
		if err := dp.Parse(item); err != nil {
			log.Printf("problems with parsing '%s' : %v", item, err)
			continue
		}
		text, err := dp.ActionInfo()
		if err != nil {
			log.Printf("problems with taking info for '%s' : %v", item, err)
			continue
		}
		log.Print(text)
	}
}
