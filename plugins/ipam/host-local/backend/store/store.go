package store

import (
	filemutex "github.com/alexflint/go-filemutex"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type Store struct {
	dataDir string
	mutex   *filemutex.FileMutex
}

const (
	defaultDataDir   = "/var/lib/cni/networks"
	LineBreak        = "\r\n"
	lastIpFilePrefix = "last_reserved_ip"
)

func New(subnet, dataDir string) (*Store, error) {
	if dataDir == "" {
		dataDir = defaultDataDir
	}
	dir := filepath.Join(dataDir, subnet)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}

	lk, err := filemutex.New(filepath.Join(dataDir, "lock"))
	if err != nil {
		return nil, err
	}

	return &Store{
		dataDir: dir,
		mutex:   lk,
	}, nil
}

// Reserve returns if success to create file
// create a file,default as `/var/lib/cni/networks/<subnet>/<ip>`
func (s *Store) Reserve(id string, ifname string, ip net.IP) (bool, error) {
	// Create a file
	f, err := os.OpenFile(filepath.Join(s.dataDir, ip.String()), os.O_RDWR|os.O_EXCL|os.O_CREATE, 0o600)
	if os.IsExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	defer f.Close()

	// write ifname in the file
	if _, err := f.WriteString(strings.TrimSpace(id) + LineBreak + ifname); err != nil {
		return false, err
	}

	// store it in latest reserve ip file
	lastFile, err := os.OpenFile(filepath.Join(s.dataDir, lastIpFilePrefix), os.O_RDWR|os.O_EXCL|os.O_CREATE, 0o600)

	return true, nil
}

func (s *Store) ReleaseByID(id string, ifname string) error {
	condition := strings.TrimSpace(id) + LineBreak + ifname
	var err error
	err = filepath.Walk(s.dataDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		if strings.TrimSpace(string(data)) == condition {
			if err := os.Remove(path); err != nil {
				return nil
			}
		}
		return nil
	})

	return err
}

func (s *Store) GetByID(id string, ifname string) []net.IP {
	condition := strings.TrimSpace(id) + LineBreak + ifname
	var ips []net.IP
	_ = filepath.Walk(s.dataDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		if strings.TrimSpace(string(data)) == condition {
			_, ipString := filepath.Split(path)
			if ip := net.ParseIP(ipString); ip != nil {
				ips = append(ips, ip)
			}
		}
		return nil
	})

	return ips
}

func (s *Store) Lock() error {
	return s.mutex.Lock()
}

func (s *Store) Unlock() error {
	return s.mutex.Unlock()
}

func (s *Store) Close() error {
	return s.mutex.Close()
}
