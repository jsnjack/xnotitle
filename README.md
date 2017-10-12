xnotitle
====

### What is it?
xnotitle hides title bar for every maximized window that matches specific pattern. It was developed as a replacement for [HTitle](https://addons.mozilla.org/en-US/firefox/addon/htitle/) Firefox addon as it is no longer works since Firefox 57

### How to use?

#### Dependencies
xnotitle works in Gnome 3 and has the following dependency:
```bash
sudo dnf install wmctrl
```

Start `xnotitle` to monitor and hide title bar of the windows.
```
Usage of ./xnotitle:
  -debug
    	Display debug information
  -name string
    	Hide title bar for windows which title contains <name> (default "Mozilla Firefox")
  -period int
    	Check for new windows every <period> ms (default 1000)
```

#### Automatically start `xnotitle`
Create a file `~/.config/autostart/xnotitle.desktop` with the following content (example):
```
[Desktop Entry]
Type=Application
Name=xnotitle
Comment=xnotitle hides title bar of the specified windows
Exec=/home/jsn/go/src/github.com/jsnjack/xnotitle/xnotitle
Terminal=false
```

#### Close Firefox window
As `xnotitle` hides the title bar and all window controls it is no longer comfortable to close the window. You can add close window button to the toolbar with this extension [Close Window Button](https://addons.mozilla.org/en-US/firefox/addon/close-window-button/)
