### Skyn 🧐

Skyn Provides TCP Port Scanning Through A Clean, Modular CLI — Built For Speed And Extensibility.

#### 🚀 Features

- 🧐 Scans Popular TCP Ports With Millisecond Precision
- 🧭 Intuitive CLI Interface Powered By Go's `flag` Package
- 🧱 Modular Architecture For Easy Expansion
- ⚡ Lightning-Fast Execution With Zero Dependencies
- 🌍 Fully Written In English For Global Accessibility

#### 📦 Installation

##### 📥 Clone The Repository

```bash
git clone https://github.com/shermsql/Skyn.git
```

##### 📁 Go To Project Directory

```bash
cd Skyn
```

##### ⚙️ Compile The Project

```bash
go build .\cmd\skyn
```

#### 🧱 Platform Builds

Skyn Can Be Built For Multiple Platforms Using Go's Cross-Compilation.

##### 🪟 Windows (64-Bit)

```bash
GOOS=windows GOARCH=amd64 go build -o Skyn.exe .\cmd\skyn
```

##### 🍎 macOS (Intel & Apple Silicon)

```bash
GOOS=darwin GOARCH=amd64 go build -o Skyn .\cmd\skyn
```

##### 🐧 Linux (64-Bit)

```bash
GOOS=linux GOARCH=amd64 go build -o Skyn .\cmd\skyn
```

##### 🐧 Linux (ARM)

```bash
GOOS=linux GOARCH=arm64 go build -o Skyn .\cmd\skyn
```

##### ℹ️ All Builds Produce A Standalone Binary With No External Dependencies.

##### 🧪 Tested On Windows, macOS, And Linux — Skyn Runs Smoothly Across Platforms.

#### 🛠️ Usage

```bash
Skyn -t <Target>
```

#### 📑 Parameters

| Flag     | Type   | Description       | Example                   |
|----------|--------|-------------------|---------------------------|
| `-t`     | String | Target Host Or IP | `localhost` / `127.0.0.1` |
| `-h`     | Bool   | Show Help Message | `-h`                      |

#### 📖 Examples

##### ⚡ Basic Scan

Scans Popular TCP Ports On The Target Host.

```bash
go run .\cmd\skyn -t localhost
```

##### 🆘 Help Output

Displays Skyn's ASCII Logo And Help Message When Parameters Are Missing Or `-h` Is Used.

```bash
go run .\cmd\skyn -h
```

#### 🗂️ Project Structure

Skyn Follows A Clean, Layered Architecture — Separating CLI Logic, Argument Parsing, Scanning, And UI Output.

```bash
Skyn/
├── cmd/skyn          # Entry Point
├── internal/args/    # CLI Argument Parser
├── internal/scanner/ # TCP Port Scanner
├── internal/ui/      # CLI Output
│   ├── Help.txt      # Help Text
│   ├── print.go      # Stylized Terminal Output
│   └── ui.go         # Help Printer Logic
├── go.mod            # Go Module Definition
├── README.md         # Project Documentation
├── LICENSE.txt       # MIT License
```

#### 📄 License

Licensed Under The MIT License — Feel Free To Use, Modify, And Share With Proper Attribution. Scanning Belongs To Everyone.

#### 💡 Philosophy & Motivation

Skyn Was Born From A Desire To Combine Network Utility With Architectural Elegance.
It Aims To Be More Than Just A Scanner — Skyn Is A Statement That CLI Tools Can Be Both Functional And Beautiful.
No External Libraries, No Clutter — Just Pure Go, Clean Design, And A Focus On Clarity.
The Project Reflects A Belief That Technical Precision And Emotional Resonance Can Coexist.

#### 🛣️ Roadmap

- 🌐 Add UDP Scanning Support
- 📁 Enable Output Logging To File
- 🧪 Implement Unit Tests For Scanner Logic
- 📚 Expand Documentation With Visual Examples
- 🧩 Introduce Plugin System For Custom Scan Profiles
- 🌍 Add Localization Support For CLI Messages

#### 🤝 Contributing

##### 💬 Contributions Are Welcome — Whether It's Code, Documentation, Or Ideas.

##### 🛠️ To Get Started

- 🍴 Fork The Repository
- 🌿 Create A Feature Branch
- 📤 Submit A Pull Request With Clear Description
- 🧘 Respect The Project's Minimalist And Modular Philosophy

##### 🧪 No External Inspirations Were Used — Skyn Is Built Entirely From Scratch, With Original Architecture And Design.