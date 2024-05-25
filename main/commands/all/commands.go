package all

import (
	"github.com/amnezia-vpn/amnezia-xray-core/main/commands/all/api"
	"github.com/amnezia-vpn/amnezia-xray-core/main/commands/all/tls"
	"github.com/amnezia-vpn/amnezia-xray-core/main/commands/base"
)

// go:generate go run github.com/amnezia-vpn/amnezia-xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
		cmdX25519,
		cmdWG,
	)
}
