import pulumi

import infra

pulumi.export('subnet_network', infra.subnet_network)
