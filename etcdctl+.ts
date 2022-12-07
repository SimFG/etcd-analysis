const completionSpec: Fig.Spec = {
    name: "etcdctl+",
    description: "short tips for the command",
    subcommands: [
        { name: ["clear"], description: "Clear all etcd data" },
        {
            name: ["completion"],
            description: "Generate the autocompletion script for the specified shell",
            subcommands: [
                {
                    name: ["bash"],
                    description: "Generate the autocompletion script for bash",
                    options: [
                        {
                            name: ["--no-descriptions"],
                            description: "disable completion descriptions",
                        },
                    ],
                },
                {
                    name: ["fish"],
                    description: "Generate the autocompletion script for fish",
                    options: [
                        {
                            name: ["--no-descriptions"],
                            description: "disable completion descriptions",
                        },
                    ],
                },
                {
                    name: ["powershell"],
                    description: "Generate the autocompletion script for powershell",
                    options: [
                        {
                            name: ["--no-descriptions"],
                            description: "disable completion descriptions",
                        },
                    ],
                },
                {
                    name: ["zsh"],
                    description: "Generate the autocompletion script for zsh",
                    options: [
                        {
                            name: ["--no-descriptions"],
                            description: "disable completion descriptions",
                        },
                    ],
                },
            ],
        },
        {
            name: ["distribute"],
            description: "Show the data distribution of etcd",
            options: [
                {
                    name: ["--bucket"],
                    description: "Bucket Count",
                    args: [{ name: "bucket", default: "5", suggestions: ["5", "6", "7", "8", "9", "10"] }],
                },
                {
                    name: ["--type"],
                    description: "Distribution basis; key, value or kv",
                    args: [{ name: "type", default: "key", suggestions: ["key", "value", "kv"] }],
                },
            ],
        },
        {
            name: ["find"],
            description: "find the key from all etcd data",
            options: [
                {
                    name: ["--key"],
                    description: "Show the data like the key",
                    args: [{ name: "key" }],
                },
                {
                    name: ["--limit"],
                    description: "The limit of the show keys",
                    args: [{ name: "limit", default: "10", suggestions: ["5", "10", "15", "20"] }],
                },
                { name: ["--value"], description: "Show the value or not" },
            ],
        },
        { name: ["leader"], description: "Get the leader node info" },
        {
            name: ["look"],
            description: "Look all etcd data",
            options: [
                {
                    name: ["--filter"],
                    description: "The filter attribute",
                    args: [{ name: "filter", default: "none", suggestions: ["none", "key", "value", "kv"] }],
                },
                {
                    name: ["--filter-max"],
                    description: "The filter max value",
                    args: [{ name: "filter-max", default: "-1" }],
                },
                {
                    name: ["--filter-min"],
                    description: "The filter min value",
                    args: [{ name: "filter-min", default: "-1" }],
                },
                {
                    name: ["--hang"],
                    description:
                        "Get updates periodically, only '--write-out=file' takes effect",
                    args: [{ name: "hang", default: "false", suggestions: ["false", "true"] }],
                },
                {
                    name: ["--hang-interval"],
                    description: "Update interval, and the unit is 's'",
                    args: [{ name: "hang-interval", default: "2", suggestions: ["1", "2", "3"] }],
                },
                { name: ["--show-value"], description: "Show the value or not" },
                {
                    name: ["--write-out"],
                    description: "The looking type",
                    args: [{ name: "write-out", default: "stdout", suggestions: ["stdout", "file"] }],
                },
            ],
        },
        {
            name: ["help"],
            description: "Help about any command",
            subcommands: [
                { name: ["clear"], description: "Clear all etcd data" },
                {
                    name: ["completion"],
                    description:
                        "Generate the autocompletion script for the specified shell",
                    subcommands: [
                        {
                            name: ["bash"],
                            description: "Generate the autocompletion script for bash",
                        },
                        {
                            name: ["fish"],
                            description: "Generate the autocompletion script for fish",
                        },
                        {
                            name: ["powershell"],
                            description: "Generate the autocompletion script for powershell",
                        },
                        {
                            name: ["zsh"],
                            description: "Generate the autocompletion script for zsh",
                        },
                    ],
                },
                {
                    name: ["distribute"],
                    description: "Show the data distribution of etcd",
                },
                { name: ["find"], description: "find the key from all etcd data" },
                { name: ["leader"], description: "Get the leader node info" },
                { name: ["look"], description: "Look all etcd data" },
            ],
        },
    ],
    options: [
        {
            name: ["--Endpoints"],
            description: "etcd connect Endpoints",
            isPersistent: true,
            isRepeatable: true,
            args: [{ name: "Endpoints", default: "[127.0.0.1:2379]" }],
        },
        { name: ["--help", "-h"], description: "Display help", isPersistent: true },
    ],
};
export default completionSpec;
