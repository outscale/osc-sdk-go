[![Project Graduated](https://docs.outscale.com/fr/userguide/_images/Project-Graduated-green.svg)](https://docs.outscale.com/en/userguide/Open-Source-Projects.html) [![](https://dcbadge.limes.pink/api/server/HUVtY5gT6s?style=flat\&theme=default-inverted)](https://discord.gg/HUVtY5gT6s)

# Outscale SDK for Golang

<p align="center">
  <img alt="Outscale SDK Go" src="https://img.icons8.com/?size=100&id=2866&format=png&color=000000" width="100px">
</p>

---

## 🌐 Links

* Documentation: [https://docs.outscale.com/en/](https://docs.outscale.com/en/)
* Go package reference: [https://pkg.go.dev/github.com/outscale/osc-sdk-go/v2](https://pkg.go.dev/github.com/outscale/osc-sdk-go/v2)
* Project repository: [https://github.com/outscale/osc-sdk-go](https://github.com/outscale/osc-sdk-go)
* OUTSCALE website: [https://outscale.com/](https://outscale.com/)
* Join our community on [Discord](https://discord.gg/HUVtY5gT6s)

---

## 📄 Table of Contents

* [Overview](#-overview)
* [Requirements](#-requirements)
* [Installation](#-installation)
* [Configuration](#-configuration)
* [Usage](#-usage)
* [Examples](#-examples)
* [License](#-license)
* [Contributing](#-contributing) 

---

## 🧭 Overview

**Outscale SDK for Golang** (`osc-sdk-go`) is the official Go SDK to interact with the OUTSCALE APIs.

It provides:

* A typed Go client for OUTSCALE cloud APIs.
* Helpers to build and send requests using Go modules.
* Integration that fits naturally into existing Go applications and tooling.

You will need an OUTSCALE account and API credentials to use this SDK.

---

## ✅ Requirements

* Go 1.18+ with modules enabled (`GO111MODULE=on` recommended).
* Git (if you plan to clone the repository).
* Access to the OUTSCALE API (valid access key / secret key or basic auth).
* Network access to OUTSCALE API endpoints.

---

## ⚙ Installation

### Option 1: Use Go modules (recommended)

Add the module to your project:

```bash
go get github.com/outscale/osc-sdk-go/v2
```

This will add `github.com/outscale/osc-sdk-go/v2` to your `go.mod` file.

### Option 2: Install from source

```bash
git clone https://github.com/outscale/osc-sdk-go.git
cd osc-sdk-go
# use as a local module, or reference it via your go.mod 'replace' directive
```

---

## 🛠 Configuration

You typically configure the SDK in Go code when creating the client (for example: region, credentials, and optional HTTP settings).

A typical flow is:

1. Build a configuration (region, credentials, etc.).
2. Create a new API client from this configuration.
3. Use the generated services to call the OUTSCALE APIs.

Refer to the GoDoc / pkg.go.dev reference and the [`examples/`](examples/) directory for concrete configuration patterns and authentication approaches.

---

## 🚀 Usage

Once the module is installed and configured in your code, you can start calling the API.

A high-level example flow:

1. Import the SDK in your Go file.
2. Initialize the configuration and API client.
3. Call the desired operation (for example, reading VMs, volumes, or networking resources).
4. Handle the response and errors as in any standard Go program.

See the [Examples](#-examples) section below for more details.

---

## 💡 Examples

The repository includes an [`examples/`](examples/) directory to help you get started quickly.

You will find examples showing how to:

* Initialize a client.
* Authenticate against the OUTSCALE API.
* Perform common operations on compute, network, and storage resources.
* Handle responses and errors in idiomatic Go.

We recommend starting from these examples and adapting them to your own use case.

---

## 📜 License

**Outscale SDK for Golang** is released under the **BSD-3-Clause** license.

© Outscale SAS

See [LICENSE](./LICENSE) for full details.

This project is compliant with [REUSE](https://reuse.software/).

---

## 🤝 Contributing

We welcome contributions!

Please read our [Contributing Guidelines](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) before submitting a pull request.
