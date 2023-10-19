package Branding

import (
	"io/ioutil"
	"strings"

	CNC "Rain/core/config/admin"
	viewsMapped "Rain/core/config/views"
)

func Load_Items() (int, error) {

	FlushStruct()

	WalkedUser, error := ioutil.ReadDir(CNC.BrandingFolderUser)
	if error != nil {
		return 0, error
	}

	WalkedAdmin, error := ioutil.ReadDir(CNC.BrandingFolderAttack)
	if error != nil {
		return 0, error
	}

	WalkedDefault, error := ioutil.ReadDir(CNC.BrandingFolderDefault)
	if error != nil {
		return 0, error
	}

	WalkedLogin, error := ioutil.ReadDir(CNC.BrandingFolderLogin)
	if error != nil {
		return 0, error
	}

	WalkedNewUser, error := ioutil.ReadDir(CNC.BrandingFolderNewUser)
	if error != nil {
		return 0, error
	}

	WalkedErrors, error := ioutil.ReadDir(CNC.BrandingFolderErrors)
	if error != nil {
		return 0, error
	}

	WalkedCaptcha, error := ioutil.ReadDir(CNC.BrandingFolderCaptcha)
	if error != nil {
		return 0, error
	}

	WalkedMFA, error := ioutil.ReadDir(CNC.BrandingFolderMFA)
	if error != nil {
		return 0, error
	}

	var Loaded_Items int = 0

	for _, File := range WalkedUser {
		Name := strings.Split(File.Name(), ".")

		if len(Name) < 1 {
			continue
		}

		FileContaining, error := ioutil.ReadFile(CNC.BrandingFolderUser + "/" + File.Name())
		if error != nil {
			continue
		}

		viewsMapped.NyxMux.Lock()
		viewsMapped.Branding[Name[0]] = string(FileContaining)
		viewsMapped.NyxMux.Unlock()

		Loaded_Items++
		continue
	}

	for _, File := range WalkedAdmin {
		Name := strings.Split(File.Name(), ".")

		if len(Name) < 1 {
			continue
		}

		FileContaining, error := ioutil.ReadFile(CNC.BrandingFolderAttack + "/" + File.Name())
		if error != nil {
			continue
		}

		viewsMapped.NyxMux.Lock()
		viewsMapped.Branding[Name[0]] = string(FileContaining)
		viewsMapped.NyxMux.Unlock()

		Loaded_Items++
		continue
	}

	for _, File := range WalkedDefault {
		Name := strings.Split(File.Name(), ".")

		if len(Name) < 1 {
			continue
		}

		FileContaining, error := ioutil.ReadFile(CNC.BrandingFolderDefault + "/" + File.Name())
		if error != nil {
			continue
		}

		viewsMapped.NyxMux.Lock()
		viewsMapped.Branding[Name[0]] = string(FileContaining)
		viewsMapped.NyxMux.Unlock()

		Loaded_Items++
		continue
	}

	for _, File := range WalkedLogin {
		Name := strings.Split(File.Name(), ".")

		if len(Name) < 1 {
			continue
		}

		FileContaining, error := ioutil.ReadFile(CNC.BrandingFolderLogin + "/" + File.Name())
		if error != nil {
			continue
		}

		viewsMapped.NyxMux.Lock()
		viewsMapped.Branding[Name[0]] = string(FileContaining)
		viewsMapped.NyxMux.Unlock()

		Loaded_Items++
		continue
	}

	for _, File := range WalkedNewUser {
		Name := strings.Split(File.Name(), ".")

		if len(Name) < 1 {
			continue
		}

		FileContaining, error := ioutil.ReadFile(CNC.BrandingFolderNewUser + "/" + File.Name())
		if error != nil {
			continue
		}

		viewsMapped.NyxMux.Lock()
		viewsMapped.Branding[Name[0]] = string(FileContaining)
		viewsMapped.NyxMux.Unlock()

		Loaded_Items++
		continue
	}

	for _, File := range WalkedErrors {
		Name := strings.Split(File.Name(), ".")

		if len(Name) < 1 {
			continue
		}

		FileContaining, error := ioutil.ReadFile(CNC.BrandingFolderErrors + "/" + File.Name())
		if error != nil {
			continue
		}

		viewsMapped.NyxMux.Lock()
		viewsMapped.Branding[Name[0]] = string(FileContaining)
		viewsMapped.NyxMux.Unlock()

		Loaded_Items++
		continue
	}

	for _, File := range WalkedCaptcha {
		Name := strings.Split(File.Name(), ".")

		if len(Name) < 1 {
			continue
		}

		FileContaining, error := ioutil.ReadFile(CNC.BrandingFolderCaptcha + "/" + File.Name())
		if error != nil {
			continue
		}

		viewsMapped.NyxMux.Lock()
		viewsMapped.Branding[Name[0]] = string(FileContaining)
		viewsMapped.NyxMux.Unlock()

		Loaded_Items++
		continue
	}

	for _, File := range WalkedMFA {
		Name := strings.Split(File.Name(), ".")

		if len(Name) < 1 {
			continue
		}

		FileContaining, error := ioutil.ReadFile(CNC.BrandingFolderMFA + "/" + File.Name())
		if error != nil {
			continue
		}

		viewsMapped.NyxMux.Lock()
		viewsMapped.Branding[Name[0]] = string(FileContaining)
		viewsMapped.NyxMux.Unlock()

		Loaded_Items++
		continue
	}


	return Loaded_Items, nil
}
