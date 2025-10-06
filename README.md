# ClusterNAS

**Easily combine and manage multiple drives into one smart local NAS.**

ClusterNAS is a fast, easy, and customizable local NAS solution for **Linux** and **Windows**. It lets you merge multiple drives into a single unified storage pool with just a few commands or through an intuitive Web UI. Perfect for home servers, small labs, or DIY NAS enthusiasts who want performance without complexity.

---

## 🚀 Features

* **Drive Pooling:** Combine multiple drives into one directory (via FUSE / mergerFS)
* **Cross-Platform:** Works on both **Linux** and **Windows**
* **Quick Setup:** Get started in minutes via CLI or Web UI
* **Web Dashboard:** Manage pools, drives, and shares directly from your browser
* **SMB File Sharing:** Built-in SMB support for local network access
* **REST API:** Automate and control ClusterNAS remotely
* **Customizable Storage:** Choose how drives are merged, replicated, or distributed

---

## ⚙️ Tech Stack

| Component      | Technology                      |
| -------------- | ------------------------------- |
| **Backend**    | Go (Golang)                     |
| **Frontend**   | React + Material UI             |
| **Filesystem** | mergerFS / FUSE                 |
| **Sharing**    | SMB / Embedded Web File Browser |
| **API**        | REST JSON API                   |

---

## 🧩 Installation

### Prerequisites

* Linux or Windows system
* Go 1.22+ installed
* (Optional) mergerFS for advanced pooling

### Clone the Repository

```bash
git clone https://github.com/<your-username>/ClusterNAS.git
cd ClusterNAS
```

### Build and Run

```bash
go build -o cln ./cmd/cln
./cln init
```

### Start Web UI

```bash
./cln web start
```

Then open **[http://localhost:8080](http://localhost:8080)** in your browser.

---

## 🧰 Example Usage

```bash
# Initialize ClusterNAS
cln init

# Detect and list available drives
cln drives detect

# Create a new pool and add drives
cln pool create main --add-all

# Enable SMB file sharing
cln share enable

# Launch Web UI
cln web start
```

---

## 🌐 API Example

```bash
curl http://localhost:8080/api/pools
```

Example response:

```json
[
  {
    "name": "main",
    "drives": ["/mnt/drive1", "/mnt/drive2"],
    "status": "mounted"
  }
]
```

---

## 🧭 Roadmap

* [x] Drive detection and pooling
* [x] Basic Web UI
* [x] SMB sharing support
* [ ] Drive health monitoring
* [ ] Pool redundancy modes (mirror, parity)
* [ ] Clustering support for multiple nodes

---

## 💡 Philosophy

> Simple setup, powerful features.
>
> ClusterNAS aims to make building a personal NAS as easy as running a few commands — without sacrificing flexibility, performance, or transparency.

---

## 🤝 Contributing

Contributions, pull requests, and feature suggestions are welcome!

1. Fork the repo
2. Create a new branch (`git checkout -b feature/awesome-feature`)
3. Commit your changes (`git commit -m 'Add awesome feature'`)
4. Push the branch (`git push origin feature/awesome-feature`)
5. Open a Pull Request

---

## 📜 License

This project is licensed under the **MIT License** — see the [LICENSE](LICENSE) file for details.

---

## ⭐ Acknowledgements

* Inspired by [mergerFS](https://github.com/trapexit/mergerfs)
* Built with Go, React, and open-source love 💙
