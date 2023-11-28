// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Code generated by schema-generate. DO NOT EDIT.

package protocol

// AdditionalRepositoriesItems
type AdditionalRepositoriesItems struct {

	// Path to where the repository should be checked out relative to `/workspace`. Defaults to the simple repository name.
	CheckoutLocation string `yaml:"checkoutLocation,omitempty" json:"checkoutLocation,omitempty"`

	// The url of the git repository to clone. Supports any context URLs.
	Url string `yaml:"url" json:"url"`
}

// CoreDump Configure the default action of certain signals is to cause a process to terminate and produce a core dump file, a file containing an image of the process's memory at the time of termination. Disabled by default.
type CoreDump struct {
	Enabled bool `yaml:"enabled,omitempty" json:"enabled,omitempty"`

	// the hard limit acts as a ceiling for the soft limit. For more details please check https://man7.org/linux/man-pages/man2/getrlimit.2.html
	HardLimit float64 `yaml:"hardLimit,omitempty" json:"hardLimit,omitempty"`

	// upper limit on the size of the core dump file that will be produced if it receives a core dump signal
	SoftLimit float64 `yaml:"softLimit,omitempty" json:"softLimit,omitempty"`
}

// Env Environment variables to set.
type Env struct {
}

// Github Configures Gitpod's GitHub app (deprecated)
type Github struct {

	// Set to true to enable workspace prebuilds, false to disable them. Defaults to true. (deprecated)
	Prebuilds interface{} `yaml:"prebuilds,omitempty" json:"prebuilds,omitempty"`
}

// GitpodConfig
type GitpodConfig struct {

	// List of additional repositories that are part of this project.
	AdditionalRepositories []*AdditionalRepositoriesItems `yaml:"additionalRepositories,omitempty" json:"additionalRepositories,omitempty"`

	// Path to where the repository should be checked out relative to `/workspace`. Defaults to the simple repository name.
	CheckoutLocation string `yaml:"checkoutLocation,omitempty" json:"checkoutLocation,omitempty"`

	// Configure the default action of certain signals is to cause a process to terminate and produce a core dump file, a file containing an image of the process's memory at the time of termination. Disabled by default.
	CoreDump *CoreDump `yaml:"coreDump,omitempty" json:"coreDump,omitempty"`

	// Experimental network configuration in workspaces (deprecated). Enabled by default
	ExperimentalNetwork bool `yaml:"experimentalNetwork,omitempty" json:"experimentalNetwork,omitempty"`

	// Git config values should be provided in pairs. E.g. `core.autocrlf: input`. See https://git-scm.com/docs/git-config#_values.
	GitConfig map[string]string `yaml:"gitConfig,omitempty" json:"gitConfig,omitempty"`

	// Configures Gitpod's GitHub app (deprecated)
	Github *Github `yaml:"github,omitempty" json:"github,omitempty"`

	// The Docker image to run your workspace in.
	Image interface{} `yaml:"image,omitempty" json:"image,omitempty"`

	// Configure JetBrains integration
	Jetbrains *Jetbrains `yaml:"jetbrains,omitempty" json:"jetbrains,omitempty"`

	// The main repository, containing the dev environment configuration.
	MainConfiguration string `yaml:"mainConfiguration,omitempty" json:"mainConfiguration,omitempty"`

	// List of exposed ports.
	Ports []*PortsItems `yaml:"ports,omitempty" json:"ports,omitempty"`

	// List of tasks to run on start. Each task will open a terminal in the IDE.
	Tasks []*TasksItems `yaml:"tasks,omitempty" json:"tasks,omitempty"`

	// Configure VS Code integration
	Vscode *Vscode `yaml:"vscode,omitempty" json:"vscode,omitempty"`

	// Path to where the IDE's workspace should be opened. Supports vscode's `*.code-workspace` files.
	WorkspaceLocation string `yaml:"workspaceLocation,omitempty" json:"workspaceLocation,omitempty"`
}

// Image_object The Docker image to run your workspace in.
type Image_object struct {

	// Relative path to the context path (optional). Should only be set if you need to copy files into the image.
	Context string `yaml:"context,omitempty" json:"context,omitempty"`

	// Relative path to a docker file.
	File string `yaml:"file" json:"file"`
}

// Jetbrains Configure JetBrains integration
type Jetbrains struct {

	// Configure CLion integration
	Clion *JetbrainsProduct `yaml:"clion,omitempty" json:"clion,omitempty"`

	// Configure GoLand integration
	Goland *JetbrainsProduct `yaml:"goland,omitempty" json:"goland,omitempty"`

	// Configure IntelliJ integration
	Intellij *JetbrainsProduct `yaml:"intellij,omitempty" json:"intellij,omitempty"`

	// Configure PhpStorm integration
	Phpstorm *JetbrainsProduct `yaml:"phpstorm,omitempty" json:"phpstorm,omitempty"`

	// List of plugins which should be installed for all JetBrains product for users of this workspace. From the JetBrains Marketplace page, find a page of the required plugin, select 'Versions' tab, click any version to copy pluginId (short name such as org.rust.lang) of the plugin you want to install.
	Plugins []string `yaml:"plugins,omitempty" json:"plugins,omitempty"`

	// Configure PyCharm integration
	Pycharm *JetbrainsProduct `yaml:"pycharm,omitempty" json:"pycharm,omitempty"`

	// Configure Rider integration
	Rider *JetbrainsProduct `yaml:"rider,omitempty" json:"rider,omitempty"`

	// Configure RubyMine integration
	Rubymine *JetbrainsProduct `yaml:"rubymine,omitempty" json:"rubymine,omitempty"`

	// Configure WebStorm integration
	Webstorm *JetbrainsProduct `yaml:"webstorm,omitempty" json:"webstorm,omitempty"`
}

// JetbrainsProduct
type JetbrainsProduct struct {

	// List of plugins which should be installed for users of this workspace. From the JetBrains Marketplace page, find a page of the required plugin, select 'Versions' tab, click any version to copy pluginId (short name such as org.rust.lang) of the plugin you want to install.
	Plugins []string `yaml:"plugins,omitempty" json:"plugins,omitempty"`

	// Enable warming up of JetBrains backend in prebuilds.
	Prebuilds *Prebuilds `yaml:"prebuilds,omitempty" json:"prebuilds,omitempty"`

	// Configure JVM options, for instance '-Xmx=4096m'.
	Vmoptions string `yaml:"vmoptions,omitempty" json:"vmoptions,omitempty"`
}

// PortsItems
type PortsItems struct {

	// A description to identify what is this port used for.
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// Port name.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// What to do when a service on this port was detected. 'notify' (default) will show a notification asking the user what to do. 'open-browser' will open a new browser tab. 'open-preview' will open in the preview on the right of the IDE. 'ignore' will do nothing.
	OnOpen string `yaml:"onOpen,omitempty" json:"onOpen,omitempty"`

	// The port number (e.g. 1337) or range (e.g. 3000-3999) to expose.
	Port interface{} `yaml:"port" json:"port"`

	// The protocol of workspace port.
	Protocol string `yaml:"protocol,omitempty" json:"protocol,omitempty"`

	// Whether the port visibility should be private or public. 'private' (default) will only allow users with workspace access to access the port. 'public' will allow everyone with the port URL to access the port.
	Visibility string `yaml:"visibility,omitempty" json:"visibility,omitempty"`
}

// Prebuilds Enable warming up of JetBrains backend in prebuilds.
type Prebuilds struct {

	// Whether only stable, latest or both versions should be warmed up. Default is stable only.
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

// Prebuilds_object Set to true to enable workspace prebuilds, false to disable them. Defaults to true. (deprecated)
type Prebuilds_object struct {

	// Add a Review in Gitpod badge to pull requests. Defaults to true.
	AddBadge bool `yaml:"addBadge,omitempty" json:"addBadge,omitempty"`

	// Add a commit check to pull requests. Set to 'fail-on-error' if you want broken prebuilds to block merging. Defaults to true.
	AddCheck interface{} `yaml:"addCheck,omitempty" json:"addCheck,omitempty"`

	// Add a label to a PR when it's prebuilt. Set to true to use the default label (prebuilt-in-gitpod) or set to a string to use a different label name. This is a beta feature and may be unreliable. Defaults to false.
	AddLabel interface{} `yaml:"addLabel,omitempty" json:"addLabel,omitempty"`

	// Enable prebuilds for all branches. Defaults to false.
	Branches bool `yaml:"branches,omitempty" json:"branches,omitempty"`

	// Enable prebuilds for the default branch (typically master). Defaults to true.
	Master bool `yaml:"master,omitempty" json:"master,omitempty"`

	// Enable prebuilds for pull-requests from the original repo. Defaults to true.
	PullRequests bool `yaml:"pullRequests,omitempty" json:"pullRequests,omitempty"`

	// Enable prebuilds for pull-requests from any repo (e.g. from forks). Defaults to false.
	PullRequestsFromForks bool `yaml:"pullRequestsFromForks,omitempty" json:"pullRequestsFromForks,omitempty"`
}

// TasksItems
type TasksItems struct {

	// A shell command to run before `init` and the main `command`. This command is executed on every start and is expected to terminate. If it fails, the following commands will not be executed.
	Before string `yaml:"before,omitempty" json:"before,omitempty"`

	// The main shell command to run after `before` and `init`. This command is executed last on every start and doesn't have to terminate.
	Command string `yaml:"command,omitempty" json:"command,omitempty"`

	// Environment variables to set.
	Env *Env `yaml:"env,omitempty" json:"env,omitempty"`

	// A shell command to run between `before` and the main `command`. This command is executed only on after initializing a workspace with a fresh clone, but not on restarts and snapshots. This command is expected to terminate. If it fails, the `command` property will not be executed.
	Init string `yaml:"init,omitempty" json:"init,omitempty"`

	// Name of the task. Shown on the tab of the opened terminal.
	Name string `yaml:"name,omitempty" json:"name,omitempty"`

	// The panel/area where to open the terminal. Default is 'bottom' panel.
	OpenIn string `yaml:"openIn,omitempty" json:"openIn,omitempty"`

	// The opening mode. Default is 'tab-after'.
	OpenMode string `yaml:"openMode,omitempty" json:"openMode,omitempty"`

	// A shell command to run after `before`. This command is executed only on during workspace prebuilds. This command is expected to terminate. If it fails, the workspace build fails.
	Prebuild string `yaml:"prebuild,omitempty" json:"prebuild,omitempty"`
}

// Vscode Configure VS Code integration
type Vscode struct {

	// List of extensions which should be installed for users of this workspace. The identifier of an extension is always '${publisher}.${name}'. For example: 'vscode.csharp'.
	Extensions []string `yaml:"extensions,omitempty" json:"extensions,omitempty"`
}
