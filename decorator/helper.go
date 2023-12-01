package decorator

import "os"

func ensureFolder(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0644)
		}
	}

	if !f.IsDir() {
		return os.Mkdir(path, 0644)
	}

	return nil
}
