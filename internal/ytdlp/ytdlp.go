package ytdlp

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/dashotv/fae"
)

func ProcessURL(url string) ([]*Info, error) {
	list, err := ytdlp_get_list(url)
	if err != nil {
		return nil, fae.Errorf("process_url: %s", err)
	}

	if list.Type == "video" {
		info, err := ytdlp_get_info(url)
		if err != nil {
			return nil, fae.Errorf("process_url: info: %s", err)
		}
		return []*Info{info}, nil
	}

	infos := make([]*Info, 0)

	for _, e := range list.Entries {
		info, err := ytdlp_get_info(e.URL)
		if err != nil {
			// return nil, fae.Errorf("process_url: info: %s", err)
			fmt.Printf("process_url: info: %s\n", err)
			continue
		}
		infos = append(infos, info)
	}

	return infos, nil
}

func ytdlp_get_list(url string) (*List, error) {
	args := []string{"--skip-download", "--no-warning", "--flat-playlist", "--dump-single-json", url}

	cmd := exec.Command("yt-dlp", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, fae.Errorf("ytdlp-list: %s", err)
	}

	list := &List{}
	if err = json.Unmarshal(out, list); err != nil {
		return nil, fae.Errorf("ytdlp-list: %s", err)
	}

	return list, nil
}

func ytdlp_get_info(url string) (*Info, error) {
	args := []string{"--skip-download", "--no-warning", "--dump-single-json", url}

	cmd := exec.Command("yt-dlp", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, fae.Errorf("ytdlp-info: %s", err)
	}

	info := &Info{}
	if err = json.Unmarshal(out, info); err != nil {
		return nil, fae.Errorf("ytdlp-info: %s", err)
	}

	return info, nil
}
