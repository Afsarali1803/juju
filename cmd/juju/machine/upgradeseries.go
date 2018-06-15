// Copyright 2018 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package machine

import (
	"fmt"
	"strings"

	"github.com/juju/cmd"
	"github.com/juju/errors"
	"github.com/juju/gnuflag"
	"github.com/juju/os/series"
	"gopkg.in/juju/names.v2"

	jujucmd "github.com/juju/juju/cmd"
	"github.com/juju/juju/cmd/modelcmd"
)

// Actions
const (
	PrepareCommand  = "prepare"
	CompleteCommand = "complete"
)

// NewUpgradeSeriesCommand returns a command which upgrades the series of
// an application or machine.
func NewUpgradeSeriesCommand() cmd.Command {
	return modelcmd.Wrap(&upgradeSeriesCommand{})
}

type upgradeMachineSeriesAPI interface {
}

// upgradeSeriesCommand is responsible for updating the series of an application or machine.
type upgradeSeriesCommand struct {
	modelcmd.ModelCommandBase
	// modelcmd.IAASOnlyCommand

	// upgradeMachineSeriesClient upgradeMachineSeriesAPI

	prepCommand   string
	force         bool
	machineNumber string
	series        string
}

var upgradeSeriesDoc = `
Upgrade a machine's operating system series.

upgrade-series allows users to perform a managed upgrade of the operating system
series of a machine. This command is performed in two steps; prepare and complete.

The "prepare" step notifies Juju that a series upgrade is taking place for a given
machine and as such Juju guards that machine against operations that would
interfere with the upgrade process.

The "complete" step notifies juju that the managed upgrade has been successfully completed.

It should be noted that once the prepare command is issued there is no way to
cancel or abort the process. Once you commit to prepare you must complete the
process or you will end up with an unusable machine!

The requested series must be explicitly supported by all charms deployed to
the specified machine. To override this constraint the --force option may be used.

The --force option should be used with caution since using a charm on a machine
running an unsupported series may cause unexpected behavior. Alternately, if the
requested series is supported in later revisions of the charm, upgrade-charm can
run beforehand.

Examples:
	juju upgrade-series prepare <machine> <series>
        juju upgrade-series prepare <machine> <series> --force
	juju upgrade-series complete <machine>

See also:
    machines
    status
    upgrade-charm
    set-series
`

func (c *upgradeSeriesCommand) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "upgrade-series",
		Args:    "<action> [args]",
		Purpose: "Upgrade a machine's series.",
		Doc:     upgradeSeriesDoc,
	}
}

func (c *upgradeSeriesCommand) SetFlags(f *gnuflag.FlagSet) {
	c.ModelCommandBase.SetFlags(f)
	f.BoolVar(&c.force, "force", false, "Upgrade even if the series is not supported by the charm and/or related subordinate charms.")
}

// Init implements cmd.Command.
func (c *upgradeSeriesCommand) Init(args []string) error {
	numArguments := 3

	if len(args) < 1 {
		return errors.Errorf("wrong number of arguments")
	}

	prepCommandStrings := []string{PrepareCommand, CompleteCommand}
	prepCommand, err := checkPrepCommands(prepCommandStrings, args[0])
	if err != nil {
		return errors.Annotate(err, "invalid argument")
	}
	c.prepCommand = prepCommand

	if c.prepCommand == CompleteCommand {
		numArguments = 2
	}

	if len(args) != numArguments {
		return errors.Errorf("wrong number of arguments")
	}

	if names.IsValidMachine(args[1]) {
		c.machineNumber = args[1]
	} else {
		return errors.Errorf("%q is an invalid machine name", args[1])
	}

	if c.prepCommand == PrepareCommand {
		series, err := checkSeries(series.SupportedSeries(), args[2])
		if err != nil {
			return err
		}
		c.series = series
	}

	return nil
}

// Run implements cmd.Run.
func (c *upgradeSeriesCommand) Run(ctx *cmd.Context) error {
	if c.prepCommand == PrepareCommand {
		err := c.promptConfirmation(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *upgradeSeriesCommand) promptConfirmation(ctx *cmd.Context) error {
	var confirmationMsg = `
WARNING This command will mark machine %q as being upgraded to series %q
This operation cannot be reverted or canceled once started.

Continue [y/N]? `[1:]
	fmt.Fprintf(ctx.Stdout, confirmationMsg, c.machineNumber, c.series)

	if err := jujucmd.UserConfirmYes(ctx); err != nil {
		return errors.Annotate(err, "upgrade series")
	}

	return nil
}

func checkPrepCommands(prepCommands []string, argCommand string) (string, error) {
	for _, prepCommand := range prepCommands {
		if prepCommand == argCommand {
			return prepCommand, nil
		}
	}

	return "", errors.Errorf("%q is an invalid upgrade-series command", argCommand)
}

func checkSeries(supportedSeries []string, seriesArgument string) (string, error) {
	for _, series := range supportedSeries {
		if series == strings.ToLower(seriesArgument) {
			return series, nil
		}
	}

	return "", errors.Errorf("%q is an unsupported series", seriesArgument)
}