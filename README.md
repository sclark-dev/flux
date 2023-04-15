# Flux
Flux is designed as a lightweight replacement for Technic's Solder software. Built from the ground up for high performance, and low memory usage. This software is designed to run anywhere, from a high-performance server to a Raspberry Pi. Fully backwards compatible with Technic's Solder API, this is designed as a self-hosted alternative to the various SaaS websites like Technic Pack, Modrinth, Curseforge, etc.

### Usage
To use this server, just download the [Latest Release](https://github.com/sclark-dev/flux/releases/latest) and run it via terminal. All features are documented via the `-h` argument, however the most common uses are listed below. <!--The examples given are for Linux and MacOS. Windows is done largly the same, just with a different executable name and using backslashes.-->

### Building
To build this from source, you'll need Go 1.20 installed.

### FAQs
* If this is written in Go, how can I upload it to a hosting service?

You can't use it on a traditional web hosting service. This is designed to run on home servers, or via PaaS providers such as Digital Ocean or Heroku.

* Is there a launcher for this?

No, not yet. I will be looking into getting support added to public builds of Prism Launcher. This software is very much in development, and the main feature is to be able to connect this through the Technic platform.