# 🖥️ Brejax System Stats API

## My First Project Released with Golang
This was a great exercise for me as I wanted to build a high-performance API server. I transitioned to Go, and I find the language to be very versatile and comfortable to work with. 
I will continue to develop and update this API. Enjoy!

A lightweight and customizable **system monitoring API** written in Go, providing real-time stats on CPU, memory, disk usage, network activity, and system load. 🚀

## 🌟 Features

- 🔍 **Comprehensive System Metrics**: Monitor CPU, RAM, disk usage, network activity, and system load.
- ⚙️ **Customizable Endpoints**: Modify API endpoint paths and server port via `config.json`.
- 🛠️ **Lightweight & Fast**: Perfect for Linux (Debian/Ubuntu) environments.
- 📡 **API-First Design**: Ready for integration with dashboards or external monitoring tools.

---

## 🚀 Installation & Setup

### 1. **Clone this repository**
```bash
git clone https://github.com/mambuzrrr/brejax-system-stats-api.git
cd brejax-system-stats-api
```

### 2. Install dependencies
Make sure you have Go installed (`go version` should return a valid version).
```bash
go mod tidy
```

### 3. Configure the API
Edit the `config.json` file to customize your port and endpoint paths:
```bash
{
  "port": 8123,
  "endpoints": {
    "stats": "/system-stats",
    "cpu": "/cpu-info",
    "memory": "/memory-info",
    "disk": "/disk-info",
    "network": "/network-info",
    "load": "/load-info"
  }
}
```

### 4. Run the server
```bash
go run main.go
```

## 📡 API Endpoints
|Endpoint|Description|
| --- | --- | 
| `/system-stats` | Full system stats: CPU, memory, disk, etc. | 
| `/cpu-info` | Current CPU usage percentage. | 
| `/memory-info` | RAM usage (total, used, percentage). |
| `/disk-info` | Disk usage (total, used, percentage). |
| `/network-info` | Network I/O statistics. |
| `/load-info` | System load averages (1, 5, 15 mins). |

### Example: Full Stats Output (`/system-stats`)
```bash
{
  "cpu_usage": [5.1],
  "memory_usage": {
    "total": 16777216000,
    "used": 8208394240,
    "used_percent": 48.92
  },
  "disk_usage": {
    "total": 500107862016,
    "used": 200123435008,
    "used_percent": 40.02
  },
  "network_stats": [
    {
      "name": "eth0",
      "bytes_sent": 123456789,
      "bytes_recv": 987654321
    }
  ],
  "system_load": {
    "load1": 0.12,
    "load5": 0.15,
    "load15": 0.18
  }
}
```
