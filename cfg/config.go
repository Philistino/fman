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
	FmanConfigDir      = "/.config/fman/"
	FmanConfigFileName = "config.toml"

	// Config Defaults
	DefaultTheme            = "dracula"
	DefaultIcons            = "nerdfont"
	DefaultPreviewDelay     = 200
	DefaultPath             = "."
	DefaultDirsMixed        = false
	DefaultNoHidden         = false
	DefaultDoubleClickDelay = 500
)

// These pointers are a janky way to get Nonetype values so we can know
// if the user passed an argument or not either via the cli or config file, and then prioritise the config file or cli.
// This approach may get unruly if we have a lot of arguments because we will have to add args in a few places

// Cfg holds the configuration details for a session
type Cfg struct {
	Path             string `arg:"positional" help:"path to open. Defaults to current directory"`
	Icons            string `default:"" help:"icon set to use. Options are: nerdfont, emoji, none. Defaults to emoji"`
	Theme            string `default:"" help:"color theme to use. Defaults to dracula. Options are: brogrammer, catppuccin-frappe, catppuccin-latte, catppuccin-macchiato, catppuccin-mocha, dracula, everblush, gruvbox, nord"`
	DirsMixed        *bool  `arg:"--dirs-mixed" help:"do not sort files from directories. Defaults to false"`
	NoHidden         *bool  `arg:"--no-hidden" help:"do not show hidden files. Defaults to false"`
	PreviewDelay     *int   `arg:"--preview-delay" placeholder:"DELAY" help:"delay in milliseconds before opening a file for previewing. This is meant to reduce io. Defaults to 200"`
	DoubleClickDelay *int   `arg:"--double-click-delay" placeholder:"DELAY" help:"delay in milliseconds to register a second click as a double click. This is included for people with limited mobility. Defaults to 500"`
	// colorScheme theme.Theme // TODO: fetch colorscheme and icon map from theme and pin to config
}

// LoadConfig loads the configuration from the cli and config file (if present).
//
// It returns an error if the config file exists but could not be read or parsed
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

// mergeConfigs takes the values from the cli and the config file and
// prioritises the values from the cli, if they are not set.
func mergeConfigs(cmdCfg Cfg, fileCfg Cfg) Cfg {
	if cmdCfg.Path == "" {
		cmdCfg.Path = fileCfg.Path
	}
	if cmdCfg.Icons == "" {
		cmdCfg.Icons = fileCfg.Icons
	}
	if cmdCfg.Theme == "" {
		cmdCfg.Theme = fileCfg.Theme
	}
	if cmdCfg.DirsMixed == nil {
		cmdCfg.DirsMixed = fileCfg.DirsMixed
	}
	if cmdCfg.NoHidden == nil {
		cmdCfg.NoHidden = fileCfg.NoHidden
	}
	if cmdCfg.PreviewDelay == nil {
		cmdCfg.PreviewDelay = fileCfg.PreviewDelay
	}
	if cmdCfg.DoubleClickDelay == nil {
		cmdCfg.DoubleClickDelay = fileCfg.DoubleClickDelay
	}

	return cmdCfg
}

// for each config value, if neither config file or cli args provides a value
// then use default values to fill config object.
func setDefaults(cfg Cfg) Cfg {
	if cfg.Path == "" {
		cfg.Path = DefaultPath
	}
	if cfg.Icons == "" {
		cfg.Icons = DefaultIcons
	}
	if cfg.Theme == "" {
		cfg.Theme = DefaultTheme
	}
	if cfg.DirsMixed == nil {
		cfg.DirsMixed = new(bool)
		*cfg.DirsMixed = DefaultDirsMixed
	}
	if cfg.NoHidden == nil {
		cfg.NoHidden = new(bool)
		*cfg.NoHidden = DefaultNoHidden
	}
	if cfg.PreviewDelay == nil {
		cfg.PreviewDelay = new(int)
		*cfg.PreviewDelay = DefaultPreviewDelay
	}
	if cfg.DoubleClickDelay == nil {
		cfg.DoubleClickDelay = new(int)
		*cfg.DoubleClickDelay = DefaultDoubleClickDelay
	}
	return cfg
}
