package provisioner

import (
	"testing"
)

func TestNewGuestCommands(t *testing.T) {
	_, err := NewGuestCommands("Amiga")
	if err == nil {
		t.Fatalf("Should have returned an err for unsupported OS type")
	}
}

func TestCreateDir(t *testing.T) {
	guestCmd, err := NewGuestCommands(UnixOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", UnixOSType)
	}
	cmd := guestCmd.CreateDir("/tmp/tempdir")
	if cmd != "mkdir -p '/tmp/tempdir'" {
		t.Fatalf("Unexpected Unix create dir cmd: %s", cmd)
	}

	guestCmd, err = NewGuestCommands(WindowsOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", WindowsOSType)
	}
	cmd = guestCmd.CreateDir("C:\\Windows\\Temp\\tempdir")
	if cmd != "New-Item -ItemType directory -Force -ErrorAction SilentlyContinue -Path \"C:\\Windows\\Temp\\tempdir\"" {
		t.Fatalf("Unexpected Windows create dir cmd: %s", cmd)
	}
}

func TestChmodExecutable(t *testing.T) {
	guestCmd, err := NewGuestCommands(UnixOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", UnixOSType)
	}
	cmd := guestCmd.ChmodExecutable("/usr/local/bin/script.sh")
	if cmd != "chmod +x '/usr/local/bin/script.sh'" {
		t.Fatalf("Unexpected Unix chmod +x cmd: %s", cmd)
	}

	guestCmd, err = NewGuestCommands(WindowsOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", WindowsOSType)
	}
	cmd = guestCmd.ChmodExecutable("C:\\Program Files\\SomeApp\\someapp.exe")
	if cmd != "echo 'skipping chmod C:\\Program Files\\SomeApp\\someapp.exe'" {
		t.Fatalf("Unexpected Windows chmod +x cmd: %s", cmd)
	}
}

func TestRemoveDir(t *testing.T) {
	guestCmd, err := NewGuestCommands(UnixOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", UnixOSType)
	}
	cmd := guestCmd.RemoveDir("/tmp/somedir")
	if cmd != "rm -rf '/tmp/somedir'" {
		t.Fatalf("Unexpected Unix remove dir cmd: %s", cmd)
	}

	guestCmd, err = NewGuestCommands(WindowsOSType)
	if err != nil {
		t.Fatalf("Failed to create new GuestCommands for OS: %s", WindowsOSType)
	}
	cmd = guestCmd.RemoveDir("C:\\Temp\\SomeDir")
	if cmd != "rm \"C:\\Temp\\SomeDir\" -recurse -force" {
		t.Fatalf("Unexpected Windows remove dir cmd: %s", cmd)
	}
}
