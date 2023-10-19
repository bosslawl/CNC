package Branding

import viewsMapped "Rain/core/config/views"

func FlushStruct() error {
	for Key, _ := range viewsMapped.Branding {
		delete(viewsMapped.Branding, Key)
	}
	return nil
}
