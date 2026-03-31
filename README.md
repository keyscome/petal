# 🌱 Petal

Petal is a lightweight parallel task executor designed to run remote shell commands over SSH, driven by simple YAML configuration. It's ideal for infrastructure ops, deployment coordination, and one-off multi-host scripting.

## 📁 Project Structure

```bash
petal/
├── main.go                # Entry point
├── config/                # Task file loading and validation
│   └── config.go
├── executor/              # Executor interface and SSH implementation
│   └── ssh.go
├── task/                  # Task runner logic (concurrency + color output)
│   └── runner.go
├── color/                 # ANSI color printing utils
│   └── color.go
└── go.mod
```

⚙️ Installation & Build

```bash
git clone https://your.repo.url/keyscome.io/petal.git
cd petal
make build
```

## 🧪 Example Configuration (task.yml)

```yml
env:
  FAVORITE_COLOR: blue

tasks:
  - name: say-hello
    hosts: [node1, node2]
    cmd: echo "Hello from $(hostname)"

  - name: colored-message
    hosts: [node3]
    env:
      FAVORITE_COLOR: green
    cmd: echo "Color is $FAVORITE_COLOR"
```

## 🚀 Usage

Run all tasks:

`./petal`

Run selected task(s):

`./petal say-hello`

## 🌐 Web UI

Start the built-in web dashboard to browse and run tasks in a browser:

```bash
./petal -web
# or specify a custom address
./petal -web -addr :8080 -f task.yml
```

Then open **http://localhost:8080** in your browser.

The web UI shows all tasks defined in the task file. Click **▶ Run** on any task to execute it; real-time output streams directly into the page via Server-Sent Events.

## 🐳 Local Multi-Node Test Environment (Docker)

You can simulate a distributed SSH cluster on your PC using Docker. This spins up three containers (`node1`, `node2`, `node3`) — each running an SSH server — so you can test Petal's parallel task execution without real remote machines.

**Prerequisites:** Docker and Docker Compose installed.

> ⚠️ **Test-only keys** — `make docker-up` generates an unencrypted RSA key pair in `docker/ssh/` (git-ignored). These keys are for local Docker testing only and must never be used against production machines.

### Quick start

```bash
# 1. Start the test nodes (generates an SSH key pair and updates ~/.ssh/config)
make docker-up

# 2. Build petal
make build

# 3. Run the bundled test tasks
./build/petal-linux-amd64 -f docker/task.yml
# or run a single task
./build/petal-linux-amd64 -f docker/task.yml ping-all
```

### What it sets up

| Container | SSH port on localhost | Hostname |
|-----------|----------------------|----------|
| node1     | 2221                 | node1    |
| node2     | 2222                 | node2    |
| node3     | 2223                 | node3    |

`setup.sh` (called by `make docker-up`) will:
1. Generate a disposable RSA key pair in `docker/ssh/` (excluded from git).
2. Build and start the three SSH containers.
3. Append `node1` / `node2` / `node3` entries to `~/.ssh/config`.

### Tear down

```bash
make docker-down
```



- ✅ YAML-based configuration
- ✅ Global + task-level environment support
- ✅ Parallel SSH command execution
- ✅ Color-coded terminal output
- ✅ Browser-based web UI with real-time output streaming
- ⏳ Planned: Logically Orchestrate Tasks
- ⏳ Planned: Sensitive Environment Variables

## 🤝 Contributing

  1. Fork the repo and create a feature branch
  2. Follow modular structure and idiomatic Go
  3. Run go fmt ./... and go vet ./... before submitting

## 📄 License

This project is dual-licensed under Apache 2.0 and commercial terms. See `COMMERCIAL.md` for details.
