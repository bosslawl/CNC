package ParseINI

import (
	"os"
	"fmt"

	"Rain/core/config/admin"

	"github.com/alexeyco/simpletable"
	ini "gopkg.in/ini.v1"
)

func CheckINI() error {
	
	if _, err := os.Stat(CNC.BrandingFolderTables + "users.ini"); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(CNC.BrandingFolderTables + "sessions.ini"); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(CNC.BrandingFolderTables + "servers.ini"); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(CNC.BrandingFolderTables + "plans.ini"); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(CNC.BrandingFolderTables + "myrunning.ini"); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(CNC.BrandingFolderTables + "globalrunning.ini"); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(CNC.BrandingFolderTables + "commands.ini"); os.IsNotExist(err) {
		return err
	}

	if _, err := os.Stat(CNC.BrandingFolderTables + "clients.ini"); os.IsNotExist(err) {
		return err
	}

	return nil
}

func ParseTableUsers(table *simpletable.Table) {
	cfg, err := ini.Load(CNC.BrandingFolderTables + "users.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	switch(cfg.Section("").Key("style").MustInt()) {
		case 1:
			table.SetStyle(simpletable.StyleCompact)
		case 2:
			table.SetStyle(simpletable.StyleCompactLite)
		case 3:
			table.SetStyle(simpletable.StyleCompactClassic)
		case 4:
			table.SetStyle(simpletable.StyleDefault)
		case 5:
			table.SetStyle(simpletable.StyleMarkdown)
		case 6:
			table.SetStyle(simpletable.StyleRounded)
		case 7:
			table.SetStyle(simpletable.StyleUnicode)
		default:
			table.SetStyle(simpletable.StyleUnicode)
	}
}

func ParseTableSessions(table *simpletable.Table) {
	cfg, err := ini.Load(CNC.BrandingFolderTables + "sessions.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	switch(cfg.Section("").Key("style").MustInt()) {
		case 1:
			table.SetStyle(simpletable.StyleCompact)
		case 2:
			table.SetStyle(simpletable.StyleCompactLite)
		case 3:
			table.SetStyle(simpletable.StyleCompactClassic)
		case 4:
			table.SetStyle(simpletable.StyleDefault)
		case 5:
			table.SetStyle(simpletable.StyleMarkdown)
		case 6:
			table.SetStyle(simpletable.StyleRounded)
		case 7:
			table.SetStyle(simpletable.StyleUnicode)
		default:
			table.SetStyle(simpletable.StyleUnicode)
	}
}

func ParseTableServers(table *simpletable.Table) {
	cfg, err := ini.Load(CNC.BrandingFolderTables + "servers.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	switch(cfg.Section("").Key("style").MustInt()) {
		case 1:
			table.SetStyle(simpletable.StyleCompact)
		case 2:
			table.SetStyle(simpletable.StyleCompactLite)
		case 3:
			table.SetStyle(simpletable.StyleCompactClassic)
		case 4:
			table.SetStyle(simpletable.StyleDefault)
		case 5:
			table.SetStyle(simpletable.StyleMarkdown)
		case 6:
			table.SetStyle(simpletable.StyleRounded)
		case 7:
			table.SetStyle(simpletable.StyleUnicode)
		default:
			table.SetStyle(simpletable.StyleUnicode)
	}
}

func ParseTablePlans(table *simpletable.Table) {
	cfg, err := ini.Load(CNC.BrandingFolderTables + "plans.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	switch(cfg.Section("").Key("style").MustInt()) {
		case 1:
			table.SetStyle(simpletable.StyleCompact)
		case 2:
			table.SetStyle(simpletable.StyleCompactLite)
		case 3:
			table.SetStyle(simpletable.StyleCompactClassic)
		case 4:
			table.SetStyle(simpletable.StyleDefault)
		case 5:
			table.SetStyle(simpletable.StyleMarkdown)
		case 6:
			table.SetStyle(simpletable.StyleRounded)
		case 7:
			table.SetStyle(simpletable.StyleUnicode)
		default:
			table.SetStyle(simpletable.StyleUnicode)
	}
}

func ParseTableMyRunning(table *simpletable.Table) {
	cfg, err := ini.Load(CNC.BrandingFolderTables + "myrunning.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	switch(cfg.Section("").Key("style").MustInt()) {
		case 1:
			table.SetStyle(simpletable.StyleCompact)
		case 2:
			table.SetStyle(simpletable.StyleCompactLite)
		case 3:
			table.SetStyle(simpletable.StyleCompactClassic)
		case 4:
			table.SetStyle(simpletable.StyleDefault)
		case 5:
			table.SetStyle(simpletable.StyleMarkdown)
		case 6:
			table.SetStyle(simpletable.StyleRounded)
		case 7:
			table.SetStyle(simpletable.StyleUnicode)
		default:
			table.SetStyle(simpletable.StyleUnicode)
	}
}

func ParseTableGlobalRunning(table *simpletable.Table) {
	cfg, err := ini.Load(CNC.BrandingFolderTables + "globalrunning.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	switch(cfg.Section("").Key("style").MustInt()) {
		case 1:
			table.SetStyle(simpletable.StyleCompact)
		case 2:
			table.SetStyle(simpletable.StyleCompactLite)
		case 3:
			table.SetStyle(simpletable.StyleCompactClassic)
		case 4:
			table.SetStyle(simpletable.StyleDefault)
		case 5:
			table.SetStyle(simpletable.StyleMarkdown)
		case 6:
			table.SetStyle(simpletable.StyleRounded)
		case 7:
			table.SetStyle(simpletable.StyleUnicode)
		default:
			table.SetStyle(simpletable.StyleUnicode)
	}
}

func ParseTableCommands(table *simpletable.Table) {
	cfg, err := ini.Load(CNC.BrandingFolderTables + "commands.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	switch(cfg.Section("").Key("style").MustInt()) {
		case 1:
			table.SetStyle(simpletable.StyleCompact)
		case 2:
			table.SetStyle(simpletable.StyleCompactLite)
		case 3:
			table.SetStyle(simpletable.StyleCompactClassic)
		case 4:
			table.SetStyle(simpletable.StyleDefault)
		case 5:
			table.SetStyle(simpletable.StyleMarkdown)
		case 6:
			table.SetStyle(simpletable.StyleRounded)
		case 7:
			table.SetStyle(simpletable.StyleUnicode)
		default:
			table.SetStyle(simpletable.StyleUnicode)
	}
}

func ParseTableClients(table *simpletable.Table) {
	cfg, err := ini.Load(CNC.BrandingFolderTables + "clients.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	switch(cfg.Section("").Key("style").MustInt()) {
		case 1:
			table.SetStyle(simpletable.StyleCompact)
		case 2:
			table.SetStyle(simpletable.StyleCompactLite)
		case 3:
			table.SetStyle(simpletable.StyleCompactClassic)
		case 4:
			table.SetStyle(simpletable.StyleDefault)
		case 5:
			table.SetStyle(simpletable.StyleMarkdown)
		case 6:
			table.SetStyle(simpletable.StyleRounded)
		case 7:
			table.SetStyle(simpletable.StyleUnicode)
		default:
			table.SetStyle(simpletable.StyleUnicode)
	}
}