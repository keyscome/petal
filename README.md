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

## 🔧 Features

- ✅ YAML-based configuration
- ✅ Global + task-level environment support
- ✅ Parallel SSH command execution
- ✅ Color-coded terminal output
- ⏳ Planned: Logically Orchestrate Tasks
- ⏳ Planned: Senstive Environment Variables

## 🤝 Contributing

  1. Fork the repo and create a feature branch
  2. Follow modular structure and idiomatic Go
  3. Run go fmt ./... and go vet ./... before submitting

## 📄 License

This project is dual-licensed under Apache 2.0 and commercial terms. See `COMMERCIAL.md` for details.
