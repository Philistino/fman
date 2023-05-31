package cfg

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/alexflint/go-arg"
)

const (
	// Config Metadata
	XdgConfigDir       = "/.config/"
	FmanConfigDir      = "/.config/fman/"
	FmanConfigFileName = "config.toml"

	// Config Defaults
	DefaultTheme = "default"
	DefaultIcons = "nerdfont"
)

type Cfg struct {
	Path      string `arg:"positional" default:"."`
	Icons     string `default:"" help:"Icon set to use. Options are: nerdfont, material"`
	DirsMixed bool   `arg:"--dirs-mixed" default:"false" help:"Do not sort files from directories"`
	NoHidden  bool   `arg:"--no-hidden" default:"false" help:"Do not show hidden files"`
	Theme     string `default:"" help:"Color theme to use. Options are: brogrammer, catppuccin-frappe, catppuccin-latte, catppuccin-macchiato, catppuccin-mocha, default, dracula, everblush, gruvbox, nord"`
	// colorScheme theme.Theme // TODO: fetch colorscheme and icon map from theme and pin to config
}

func LoadConfig() (Cfg, error) {
	cli := loadCliArgs()
	fileCfg, err := loadConfigFile()
	cfg := mergeConfigs(cli, fileCfg)
	cfg = setDefaults(cfg)
	return cfg, err
}

func loadCliArgs() Cfg {
	var cli Cfg
	arg.MustParse(&cli)
	return cli
}

func loadConfigFile() (Cfg, error) {
	var fileCfg Cfg
	home, err := os.UserHomeDir()
	if err != nil {
		return fileCfg, err
	}

	fileContents, err := os.ReadFile(filepath.Join(home, FmanConfigDir, FmanConfigFileName))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fileCfg, nil
		} else {
			return fileCfg, err
		}
	}

	_, err = toml.Decode(string(fileContents), &fileCfg)
	if err != nil {
		return fileCfg, errors.New("could not decode config file")
	}
	return fileCfg, nil
}

func mergeConfigs(cmdCfg Cfg, fileCfg Cfg) Cfg {
	if cmdCfg.Theme == "" {
		cmdCfg.Theme = fileCfg.Theme
	}
	if cmdCfg.Icons == "" {
		cmdCfg.Icons = fileCfg.Icons
	}
	// because DirsMixed defaults to false, we will use the value from the config file
	// if the value is not provided by the user at the command line.
	if !cmdCfg.DirsMixed {
		cmdCfg.DirsMixed = fileCfg.DirsMixed
	}
	// same as above
	if !cmdCfg.NoHidden {
		cmdCfg.NoHidden = fileCfg.NoHidden
	}

	return cmdCfg
}

func setDefaults(cfg Cfg) Cfg {
	// For each config value if neither config file or cli args provides a value
	// then use default values to fill config object.
	if cfg.Theme == "" {
		cfg.Theme = DefaultTheme
	}
	if cfg.Icons == "" {
		cfg.Icons = DefaultIcons
	}
	return cfg
}
