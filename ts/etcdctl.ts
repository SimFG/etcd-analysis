const completionSpec: Fig.Spec = {
    name: "etcdctl",
    description: "A simple command line client for etcd3.",
    subcommands: [
      {
        name: ["alarm"],
        description: "Alarm related commands",
        subcommands: [
          { name: ["disarm"], description: "Disarms all alarms" },
          { name: ["list"], description: "Lists all alarms" },
        ],
      },
      {
        name: ["auth"],
        description: "Enable or disable authentication",
        subcommands: [
          { name: ["disable"], description: "Disables authentication" },
          { name: ["enable"], description: "Enables authentication" },
          { name: ["status"], description: "Returns authentication status" },
        ],
      },
      {
        name: ["check"],
        description: "commands for checking properties of the etcd cluster",
        subcommands: [
          {
            name: ["datascale"],
            description:
              "Check the memory usage of holding data for different workloads on a given server endpoint.",
            options: [
              {
                name: ["--auto-compact"],
                description:
                  "Compact storage with last revision after test is finished.",
              },
              {
                name: ["--auto-defrag"],
                description: "Defragment storage after test is finished.",
              },
              {
                name: ["--load"],
                description:
                  "The datascale check's workload model. Accepted workloads: s(small), m(medium), l(large), xl(xLarge)",
                args: [{ name: "load", default: "s" }],
              },
              {
                name: ["--prefix"],
                description: "The prefix for writing the datascale check's keys.",
                args: [{ name: "prefix", default: "/etcdctl-check-datascale/" }],
              },
            ],
          },
          {
            name: ["perf"],
            description: "Check the performance of the etcd cluster",
            options: [
              {
                name: ["--auto-compact"],
                description:
                  "Compact storage with last revision after test is finished.",
              },
              {
                name: ["--auto-defrag"],
                description: "Defragment storage after test is finished.",
              },
              {
                name: ["--load"],
                description:
                  "The performance check's workload model. Accepted workloads: s(small), m(medium), l(large), xl(xLarge). Different workload models use different configurations in terms of number of clients and expected throughtput.",
                args: [{ name: "load", default: "s" }],
              },
              {
                name: ["--prefix"],
                description:
                  "The prefix for writing the performance check's keys.",
                args: [{ name: "prefix", default: "/etcdctl-check-perf/" }],
              },
            ],
          },
        ],
      },
      {
        name: ["compaction"],
        description: "Compacts the event history in etcd",
        options: [
          {
            name: ["--physical"],
            description:
              "'true' to wait for compaction to physically remove all old revisions",
          },
        ],
      },
      { name: ["completion"], description: "Generate completion script" },
      {
        name: ["defrag"],
        description:
          "Defragments the storage of the etcd members with given endpoints",
        options: [
          {
            name: ["--cluster"],
            description: "use all endpoints from the cluster member list",
            isPersistent: true,
          },
        ],
      },
      {
        name: ["del"],
        description:
          "Removes the specified key or range of keys [key, range_end)",
        options: [
          {
            name: ["--from-key"],
            description:
              "delete keys that are greater than or equal to the given key using byte compare",
          },
          { name: ["--prefix"], description: "delete keys with matching prefix" },
          { name: ["--prev-kv"], description: "return deleted key-value pairs" },
          { name: ["--range"], description: "delete range of keys" },
        ],
      },
      {
        name: ["downgrade"],
        description: "Downgrade related commands",
        subcommands: [
          {
            name: ["cancel"],
            description: "Cancel the ongoing downgrade action to cluster",
          },
          {
            name: ["enable"],
            description: "Start a downgrade action to cluster",
          },
          {
            name: ["validate"],
            description:
              "Validate downgrade capability before starting downgrade",
          },
        ],
      },
      {
        name: ["elect"],
        description: "Observes and participates in leader election",
        options: [{ name: ["--listen", "-l"], description: "observation mode" }],
      },
      {
        name: ["endpoint"],
        description: "Endpoint related commands",
        subcommands: [
          {
            name: ["hashkv"],
            description:
              "Prints the KV history hash for each endpoint in --endpoints",
            options: [
              {
                name: ["--rev"],
                description: "maximum revision to hash (default: all revisions)",
                isPersistent: true,
                args: [{ name: "rev", default: "0" }],
              },
            ],
          },
          {
            name: ["health"],
            description:
              "Checks the healthiness of endpoints specified in `--endpoints` flag",
          },
          {
            name: ["status"],
            description:
              "Prints out the status of endpoints specified in `--endpoints` flag",
          },
        ],
        options: [
          {
            name: ["--cluster"],
            description: "use all endpoints from the cluster member list",
            isPersistent: true,
          },
        ],
      },
      {
        name: ["get"],
        description: "Gets the key or a range of keys",
        options: [
          {
            name: ["--consistency"],
            description: "Linearizable(l) or Serializable(s)",
            args: [{ name: "consistency", default: "l" }],
          },
          { name: ["--count-only"], description: "Get only the count" },
          {
            name: ["--from-key"],
            description:
              "Get keys that are greater than or equal to the given key using byte compare",
          },
          { name: ["--keys-only"], description: "Get only the keys" },
          {
            name: ["--limit"],
            description: "Maximum number of results",
            args: [{ name: "limit", default: "0" }],
          },
          {
            name: ["--order"],
            description:
              "Order of results; ASCEND or DESCEND (ASCEND by default)",
            args: [{ name: "order" }],
          },
          { name: ["--prefix"], description: "Get keys with matching prefix" },
          {
            name: ["--print-value-only"],
            description:
              'Only write values when using the "simple" output format',
          },
          {
            name: ["--rev"],
            description: "Specify the kv revision",
            args: [{ name: "rev", default: "0" }],
          },
          {
            name: ["--sort-by"],
            description: "Sort target; CREATE, KEY, MODIFY, VALUE, or VERSION",
            args: [{ name: "sort-by" }],
          },
        ],
      },
      {
        name: ["lease"],
        description: "Lease related commands",
        subcommands: [
          { name: ["grant"], description: "Creates leases" },
          {
            name: ["keep-alive"],
            description: "Keeps leases alive (renew)",
            options: [
              {
                name: ["--once"],
                description:
                  "Resets the keep-alive time to its original value and cobrautl.Exits immediately",
              },
            ],
          },
          { name: ["list"], description: "List all active leases" },
          { name: ["revoke"], description: "Revokes leases" },
          {
            name: ["timetolive"],
            description: "Get lease information",
            options: [
              {
                name: ["--keys"],
                description: "Get keys attached to this lease",
              },
            ],
          },
        ],
      },
      {
        name: ["lock"],
        description: "Acquires a named lock",
        options: [
          {
            name: ["--ttl"],
            description: "timeout for session",
            args: [{ name: "ttl", default: "10" }],
          },
        ],
      },
      {
        name: ["make-mirror"],
        description: "Makes a mirror at the destination etcd cluster",
        options: [
          {
            name: ["--dest-cacert"],
            description:
              "Verify certificates of TLS enabled secure servers using this CA bundle",
            args: [{ name: "dest-cacert" }],
          },
          {
            name: ["--dest-cert"],
            description:
              "Identify secure client using this TLS certificate file for the destination cluster",
            args: [{ name: "dest-cert" }],
          },
          {
            name: ["--dest-insecure-transport"],
            description: "Disable transport security for client connections",
          },
          {
            name: ["--dest-key"],
            description: "Identify secure client using this TLS key file",
            args: [{ name: "dest-key" }],
          },
          {
            name: ["--dest-password"],
            description:
              "Destination password for authentication (if this option is used, --user option shouldn't include password)",
            args: [{ name: "dest-password" }],
          },
          {
            name: ["--dest-prefix"],
            description:
              "destination prefix to mirror a prefix to a different prefix in the destination cluster",
            args: [{ name: "dest-prefix" }],
          },
          {
            name: ["--dest-user"],
            description:
              "Destination username[:password] for authentication (prompt if password is not supplied)",
            args: [{ name: "dest-user" }],
          },
          {
            name: ["--max-txn-ops"],
            description:
              "Maximum number of operations permitted in a transaction during syncing updates.",
            args: [{ name: "max-txn-ops", default: "128" }],
          },
          {
            name: ["--no-dest-prefix"],
            description:
              "mirror key-values to the root of the destination cluster",
          },
          {
            name: ["--prefix"],
            description: "Key-value prefix to mirror",
            args: [{ name: "prefix" }],
          },
          {
            name: ["--rev"],
            description: "Specify the kv revision to start to mirror",
            args: [{ name: "rev", default: "0" }],
          },
        ],
      },
      {
        name: ["member"],
        description: "Membership related commands",
        subcommands: [
          {
            name: ["add"],
            description: "Adds a member into the cluster",
            options: [
              {
                name: ["--learner"],
                description: "indicates if the new member is raft learner",
              },
              {
                name: ["--peer-urls"],
                description: "comma separated peer URLs for the new member.",
                args: [{ name: "peer-urls" }],
              },
            ],
          },
          { name: ["list"], description: "Lists all members in the cluster" },
          {
            name: ["promote"],
            description: "Promotes a non-voting member in the cluster",
          },
          { name: ["remove"], description: "Removes a member from the cluster" },
          {
            name: ["update"],
            description: "Updates a member in the cluster",
            options: [
              {
                name: ["--peer-urls"],
                description: "comma separated peer URLs for the updated member.",
                args: [{ name: "peer-urls" }],
              },
            ],
          },
        ],
      },
      {
        name: ["move-leader"],
        description: "Transfers leadership to another etcd cluster member.",
      },
      {
        name: ["put"],
        description: "Puts the given key into the store",
        options: [
          {
            name: ["--ignore-lease"],
            description: "updates the key using its current lease",
          },
          {
            name: ["--ignore-value"],
            description: "updates the key using its current value",
          },
          {
            name: ["--lease"],
            description: "lease ID (in hexadecimal) to attach to the key",
            args: [{ name: "lease", default: "0" }],
          },
          {
            name: ["--prev-kv"],
            description: "return the previous key-value pair before modification",
          },
        ],
      },
      {
        name: ["role"],
        description: "Role related commands",
        subcommands: [
          { name: ["add"], description: "Adds a new role" },
          { name: ["delete"], description: "Deletes a role" },
          { name: ["get"], description: "Gets detailed information of a role" },
          {
            name: ["grant-permission"],
            description: "Grants a key to a role",
            options: [
              {
                name: ["--from-key"],
                description:
                  "grant a permission of keys that are greater than or equal to the given key using byte compare",
              },
              { name: ["--prefix"], description: "grant a prefix permission" },
            ],
          },
          { name: ["list"], description: "Lists all roles" },
          {
            name: ["revoke-permission"],
            description: "Revokes a key from a role",
            options: [
              {
                name: ["--from-key"],
                description:
                  "revoke a permission of keys that are greater than or equal to the given key using byte compare",
              },
              { name: ["--prefix"], description: "revoke a prefix permission" },
            ],
          },
        ],
      },
      {
        name: ["snapshot"],
        description: "Manages etcd node snapshots",
        subcommands: [
          {
            name: ["save"],
            description: "Stores an etcd node backend snapshot to a given file",
          },
        ],
      },
      {
        name: ["txn"],
        description: "Txn processes all the requests in one transaction",
        options: [
          {
            name: ["--interactive", "-i"],
            description: "Input transaction in interactive mode",
          },
        ],
      },
      {
        name: ["user"],
        description: "User related commands",
        subcommands: [
          {
            name: ["add"],
            description: "Adds a new user",
            options: [
              {
                name: ["--interactive"],
                description:
                  "Read password from stdin instead of interactive terminal",
              },
              {
                name: ["--new-user-password"],
                description: "Supply password from the command line flag",
                args: [{ name: "new-user-password" }],
              },
              {
                name: ["--no-password"],
                description:
                  "Create a user without password (CN based auth only)",
              },
            ],
          },
          { name: ["delete"], description: "Deletes a user" },
          {
            name: ["get"],
            description: "Gets detailed information of a user",
            options: [
              {
                name: ["--detail"],
                description: "Show permissions of roles granted to the user",
              },
            ],
          },
          { name: ["grant-role"], description: "Grants a role to a user" },
          { name: ["list"], description: "Lists all users" },
          {
            name: ["passwd"],
            description: "Changes password of user",
            options: [
              {
                name: ["--interactive"],
                description:
                  "If true, read password from stdin instead of interactive terminal",
              },
            ],
          },
          { name: ["revoke-role"], description: "Revokes a role from a user" },
        ],
      },
      { name: ["version"], description: "Prints the version of etcdctl" },
      {
        name: ["watch"],
        description: "Watches events stream on keys or prefixes",
        options: [
          { name: ["--interactive", "-i"], description: "Interactive mode" },
          {
            name: ["--prefix"],
            description: "Watch on a prefix if prefix is set",
          },
          {
            name: ["--prev-kv"],
            description:
              "get the previous key-value pair before the event happens",
          },
          {
            name: ["--progress-notify"],
            description: "get periodic watch progress notification from server",
          },
          {
            name: ["--rev"],
            description: "Revision to start watching",
            args: [{ name: "rev", default: "0" }],
          },
        ],
      },
      {
        name: ["help"],
        description: "Help about any command",
        subcommands: [
          {
            name: ["alarm"],
            description: "Alarm related commands",
            subcommands: [
              { name: ["disarm"], description: "Disarms all alarms" },
              { name: ["list"], description: "Lists all alarms" },
            ],
          },
          {
            name: ["auth"],
            description: "Enable or disable authentication",
            subcommands: [
              { name: ["disable"], description: "Disables authentication" },
              { name: ["enable"], description: "Enables authentication" },
              { name: ["status"], description: "Returns authentication status" },
            ],
          },
          {
            name: ["check"],
            description: "commands for checking properties of the etcd cluster",
            subcommands: [
              {
                name: ["datascale"],
                description:
                  "Check the memory usage of holding data for different workloads on a given server endpoint.",
              },
              {
                name: ["perf"],
                description: "Check the performance of the etcd cluster",
              },
            ],
          },
          {
            name: ["compaction"],
            description: "Compacts the event history in etcd",
          },
          { name: ["completion"], description: "Generate completion script" },
          {
            name: ["defrag"],
            description:
              "Defragments the storage of the etcd members with given endpoints",
          },
          {
            name: ["del"],
            description:
              "Removes the specified key or range of keys [key, range_end)",
          },
          {
            name: ["downgrade"],
            description: "Downgrade related commands",
            subcommands: [
              {
                name: ["cancel"],
                description: "Cancel the ongoing downgrade action to cluster",
              },
              {
                name: ["enable"],
                description: "Start a downgrade action to cluster",
              },
              {
                name: ["validate"],
                description:
                  "Validate downgrade capability before starting downgrade",
              },
            ],
          },
          {
            name: ["elect"],
            description: "Observes and participates in leader election",
          },
          {
            name: ["endpoint"],
            description: "Endpoint related commands",
            subcommands: [
              {
                name: ["hashkv"],
                description:
                  "Prints the KV history hash for each endpoint in --endpoints",
              },
              {
                name: ["health"],
                description:
                  "Checks the healthiness of endpoints specified in `--endpoints` flag",
              },
              {
                name: ["status"],
                description:
                  "Prints out the status of endpoints specified in `--endpoints` flag",
              },
            ],
          },
          { name: ["get"], description: "Gets the key or a range of keys" },
          {
            name: ["lease"],
            description: "Lease related commands",
            subcommands: [
              { name: ["grant"], description: "Creates leases" },
              { name: ["keep-alive"], description: "Keeps leases alive (renew)" },
              { name: ["list"], description: "List all active leases" },
              { name: ["revoke"], description: "Revokes leases" },
              { name: ["timetolive"], description: "Get lease information" },
            ],
          },
          { name: ["lock"], description: "Acquires a named lock" },
          {
            name: ["make-mirror"],
            description: "Makes a mirror at the destination etcd cluster",
          },
          {
            name: ["member"],
            description: "Membership related commands",
            subcommands: [
              { name: ["add"], description: "Adds a member into the cluster" },
              { name: ["list"], description: "Lists all members in the cluster" },
              {
                name: ["promote"],
                description: "Promotes a non-voting member in the cluster",
              },
              {
                name: ["remove"],
                description: "Removes a member from the cluster",
              },
              {
                name: ["update"],
                description: "Updates a member in the cluster",
              },
            ],
          },
          {
            name: ["move-leader"],
            description: "Transfers leadership to another etcd cluster member.",
          },
          { name: ["put"], description: "Puts the given key into the store" },
          {
            name: ["role"],
            description: "Role related commands",
            subcommands: [
              { name: ["add"], description: "Adds a new role" },
              { name: ["delete"], description: "Deletes a role" },
              {
                name: ["get"],
                description: "Gets detailed information of a role",
              },
              {
                name: ["grant-permission"],
                description: "Grants a key to a role",
              },
              { name: ["list"], description: "Lists all roles" },
              {
                name: ["revoke-permission"],
                description: "Revokes a key from a role",
              },
            ],
          },
          {
            name: ["snapshot"],
            description: "Manages etcd node snapshots",
            subcommands: [
              {
                name: ["save"],
                description:
                  "Stores an etcd node backend snapshot to a given file",
              },
            ],
          },
          {
            name: ["txn"],
            description: "Txn processes all the requests in one transaction",
          },
          {
            name: ["user"],
            description: "User related commands",
            subcommands: [
              { name: ["add"], description: "Adds a new user" },
              { name: ["delete"], description: "Deletes a user" },
              {
                name: ["get"],
                description: "Gets detailed information of a user",
              },
              { name: ["grant-role"], description: "Grants a role to a user" },
              { name: ["list"], description: "Lists all users" },
              { name: ["passwd"], description: "Changes password of user" },
              {
                name: ["revoke-role"],
                description: "Revokes a role from a user",
              },
            ],
          },
          { name: ["version"], description: "Prints the version of etcdctl" },
          {
            name: ["watch"],
            description: "Watches events stream on keys or prefixes",
          },
        ],
      },
    ],
    options: [
      {
        name: ["--cacert"],
        description:
          "verify certificates of TLS-enabled secure servers using this CA bundle",
        isPersistent: true,
        args: [{ name: "cacert" }],
      },
      {
        name: ["--cert"],
        description: "identify secure client using this TLS certificate file",
        isPersistent: true,
        args: [{ name: "cert" }],
      },
      {
        name: ["--command-timeout"],
        description: "timeout for short running command (excluding dial timeout)",
        isPersistent: true,
        args: [{ name: "command-timeout", default: "5s" }],
      },
      {
        name: ["--debug"],
        description: "enable client-side debug logging",
        isPersistent: true,
      },
      {
        name: ["--dial-timeout"],
        description: "dial timeout for client connections",
        isPersistent: true,
        args: [{ name: "dial-timeout", default: "2s" }],
      },
      {
        name: ["--discovery-srv", "-d"],
        description:
          "domain name to query for SRV records describing cluster endpoints",
        isPersistent: true,
        args: [{ name: "discovery-srv" }],
      },
      {
        name: ["--discovery-srv-name"],
        description: "service name to query when using DNS discovery",
        isPersistent: true,
        args: [{ name: "discovery-srv-name" }],
      },
      {
        name: ["--endpoints"],
        description: "gRPC endpoints",
        isPersistent: true,
        isRepeatable: true,
        args: [{ name: "endpoints", default: "[127.0.0.1:2379]" }],
      },
      {
        name: ["--hex"],
        description: "print byte strings as hex encoded strings",
        isPersistent: true,
      },
      {
        name: ["--insecure-discovery"],
        description: "accept insecure SRV records describing cluster endpoints",
        isPersistent: true,
      },
      {
        name: ["--insecure-skip-tls-verify"],
        description:
          "skip server certificate verification (CAUTION: this option should be enabled only for testing purposes)",
        isPersistent: true,
      },
      {
        name: ["--insecure-transport"],
        description: "disable transport security for client connections",
        isPersistent: true,
      },
      {
        name: ["--keepalive-time"],
        description: "keepalive time for client connections",
        isPersistent: true,
        args: [{ name: "keepalive-time", default: "2s" }],
      },
      {
        name: ["--keepalive-timeout"],
        description: "keepalive timeout for client connections",
        isPersistent: true,
        args: [{ name: "keepalive-timeout", default: "6s" }],
      },
      {
        name: ["--key"],
        description: "identify secure client using this TLS key file",
        isPersistent: true,
        args: [{ name: "key" }],
      },
      {
        name: ["--password"],
        description:
          "password for authentication (if this option is used, --user option shouldn't include password)",
        isPersistent: true,
        args: [{ name: "password" }],
      },
      {
        name: ["--user"],
        description:
          "username[:password] for authentication (prompt if password is not supplied)",
        isPersistent: true,
        args: [{ name: "user" }],
      },
      {
        name: ["--write-out", "-w"],
        description:
          "set the output format (fields, json, protobuf, simple, table)",
        isPersistent: true,
        args: [{ name: "write-out", default: "simple" }],
      },
      { name: ["--help", "-h"], description: "Display help", isPersistent: true },
    ],
  };
  export default completionSpec;
  
