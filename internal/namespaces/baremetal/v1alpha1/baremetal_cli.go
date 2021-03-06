// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package baremetal

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/baremetal/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		baremetalRoot(),
		baremetalServer(),
		baremetalOs(),
		baremetalIP(),
		baremetalBmc(),
		baremetalOffer(),
		baremetalServerList(),
		baremetalServerGet(),
		baremetalServerCreate(),
		baremetalServerUpdate(),
		baremetalServerInstall(),
		baremetalServerDelete(),
		baremetalServerReboot(),
		baremetalServerStart(),
		baremetalServerStop(),
	)
}
func baremetalRoot() *core.Command {
	return &core.Command{
		Short:     `Baremetal API`,
		Long:      ``,
		Namespace: "baremetal",
	}
}

func baremetalServer() *core.Command {
	return &core.Command{
		Short:     `A server is a denomination of a type of instances provided by Scaleway`,
		Long:      `A server is a denomination of a type of instances provided by Scaleway.`,
		Namespace: "baremetal",
		Resource:  "server",
	}
}

func baremetalOs() *core.Command {
	return &core.Command{
		Short:     `An Operating System (OS) is the underlying software installed on your server`,
		Long:      `An Operating System (OS) is the underlying software installed on your server.`,
		Namespace: "baremetal",
		Resource:  "os",
	}
}

func baremetalIP() *core.Command {
	return &core.Command{
		Short:     `IP fail-over management`,
		Long:      `IP fail-over management.`,
		Namespace: "baremetal",
		Resource:  "ip",
	}
}

func baremetalBmc() *core.Command {
	return &core.Command{
		Short:     `Baseboard Management Controller (BMC) offers a low-level access to your baremetal instance`,
		Long:      `Baseboard Management Controller (BMC) offers a low-level access to your baremetal instance.`,
		Namespace: "baremetal",
		Resource:  "bmc",
	}
}

func baremetalOffer() *core.Command {
	return &core.Command{
		Short:     `Commercial offers`,
		Long:      `Commercial offers.`,
		Namespace: "baremetal",
		Resource:  "offer",
	}
}

func baremetalServerList() *core.Command {
	return &core.Command{
		Short:     `List servers`,
		Long:      `List all created servers.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "list",
		ArgsType:  reflect.TypeOf(baremetal.ListServersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Short:      `Order of the servers`,
				Required:   false,
				EnumValues: []string{"created_at_asc", "created_at_desc"},
			},
			{
				Name:     "tags.{index}",
				Short:    `Filter servers by tags`,
				Required: false,
			},
			{
				Name:     "status.{index}",
				Short:    `Filter servers by status`,
				Required: false,
			},
			{
				Name:     "name",
				Short:    `Filter servers by name`,
				Required: false,
			},
			{
				Name:     "organization-id",
				Short:    `Filter servers by organization ID`,
				Required: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.ListServersRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			resp, err := api.ListServers(request, scw.WithAllPages())
			if err != nil {
				return nil, err
			}
			return resp.Servers, nil

		},
		Examples: []*core.Example{
			{
				Short:   "List all servers on your default zone",
				Request: `null`,
			},
		},
	}
}

func baremetalServerGet() *core.Command {
	return &core.Command{
		Short:     `Get server`,
		Long:      `Get the server associated with the given ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "get",
		ArgsType:  reflect.TypeOf(baremetal.GetServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "server-id",
				Short:    `ID of the server`,
				Required: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.GetServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.GetServer(request)

		},
	}
}

func baremetalServerCreate() *core.Command {
	return &core.Command{
		Short:     `Create server`,
		Long:      `Create a new server. Once the server is created, you probably want to install an OS.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "create",
		ArgsType:  reflect.TypeOf(baremetal.CreateServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "offer-id",
				Short:    `Offer ID of the new server`,
				Required: true,
			},
			{
				Name:     "name",
				Short:    `Name of the server (≠hostname)`,
				Required: true,
			},
			{
				Name:     "description",
				Short:    `Description associated to the server, max 255 characters`,
				Required: true,
			},
			{
				Name:     "tags.{index}",
				Short:    `Tags to associate to the server`,
				Required: false,
			},
			core.OrganizationIDArgSpec(),
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.CreateServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.CreateServer(request)

		},
	}
}

func baremetalServerUpdate() *core.Command {
	return &core.Command{
		Short:     `Update server`,
		Long:      `Update the server associated with the given ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "update",
		ArgsType:  reflect.TypeOf(baremetal.UpdateServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "server-id",
				Short:    `ID of the server to update`,
				Required: true,
			},
			{
				Name:     "name",
				Short:    `Name of the server (≠hostname), not updated if null`,
				Required: false,
			},
			{
				Name:     "description",
				Short:    `Description associated to the server, max 255 characters, not updated if null`,
				Required: false,
			},
			{
				Name:     "tags",
				Short:    `Tags associated to the server, not updated if null`,
				Required: false,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.UpdateServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.UpdateServer(request)

		},
	}
}

func baremetalServerInstall() *core.Command {
	return &core.Command{
		Short:     `Install server`,
		Long:      `Install an OS on the server associated with the given ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "install",
		ArgsType:  reflect.TypeOf(baremetal.InstallServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "server-id",
				Short:    `Server ID to install`,
				Required: true,
			},
			{
				Name:     "os-id",
				Short:    `ID of the OS to install on the server`,
				Required: true,
			},
			{
				Name:     "hostname",
				Short:    `Hostname of the server`,
				Required: true,
			},
			{
				Name:     "ssh-key-ids.{index}",
				Short:    `SSH key IDs authorized on the server`,
				Required: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.InstallServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.InstallServer(request)

		},
	}
}

func baremetalServerDelete() *core.Command {
	return &core.Command{
		Short:     `Delete server`,
		Long:      `Delete the server associated with the given ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "delete",
		ArgsType:  reflect.TypeOf(baremetal.DeleteServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "server-id",
				Short:    `ID of the server to delete`,
				Required: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.DeleteServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.DeleteServer(request)

		},
	}
}

func baremetalServerReboot() *core.Command {
	return &core.Command{
		Short:     `Reboot server`,
		Long:      `Reboot the server associated with the given ID, use boot param to reboot in rescue.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "reboot",
		ArgsType:  reflect.TypeOf(baremetal.RebootServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "server-id",
				Short:    `ID of the server to reboot`,
				Required: true,
			},
			{
				Name:       "boot-type",
				Short:      `The type of boot`,
				Required:   false,
				EnumValues: []string{"normal", "rescue"},
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.RebootServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.RebootServer(request)

		},
	}
}

func baremetalServerStart() *core.Command {
	return &core.Command{
		Short:     `Start server`,
		Long:      `Start the server associated with the given ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "start",
		ArgsType:  reflect.TypeOf(baremetal.StartServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "server-id",
				Short:    `ID of the server to start`,
				Required: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.StartServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.StartServer(request)

		},
	}
}

func baremetalServerStop() *core.Command {
	return &core.Command{
		Short:     `Stop server`,
		Long:      `Stop the server associated with the given ID.`,
		Namespace: "baremetal",
		Resource:  "server",
		Verb:      "stop",
		ArgsType:  reflect.TypeOf(baremetal.StopServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:     "server-id",
				Short:    `ID of the server to stop`,
				Required: true,
			},
			core.ZoneArgSpec(scw.ZoneFrPar2),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*baremetal.StopServerRequest)

			client := core.ExtractClient(ctx)
			api := baremetal.NewAPI(client)
			return api.StopServer(request)

		},
	}
}
