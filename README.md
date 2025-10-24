# GoPulse

**GoPulse** is a Bluetooth audio profile manager for Linux, built in Go.  


# Installation 

```bash
curl -fsSL https://raw.githubusercontent.com/anirbandotdev/gopulse/main/install.sh | bash
```

# Waybar config

```json
"custom/gopulse": {
    "format": "⚛ {}",
    "exec": "gopulse current",
    "on-click": "gopulse toggle",
    "interval": 3
},
```