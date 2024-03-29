package server

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func ProcessURL(url string) ([]*YtdlpInfo, error) {
	list, err := ytdlp_get_list(url)
	if err != nil {
		return nil, fmt.Errorf("process_url: %s", err)
	}

	if list.Type == "video" {
		info, err := ytdlp_get_info(url)
		if err != nil {
			return nil, fmt.Errorf("process_url: info: %s", err)
		}
		return []*YtdlpInfo{info}, nil
	}

	infos := make([]*YtdlpInfo, 0, len(list.Entries))

	for _, e := range list.Entries {
		info, err := ytdlp_get_info(e.URL)
		if err != nil {
			return nil, fmt.Errorf("process_url: info: %s", err)
		}
		infos = append(infos, info)
	}

	return infos, nil
}

func ytdlp_get_list(url string) (*YtdlpList, error) {
	args := []string{"--skip-download", "--no-warning", "--flat-playlist", "--dump-single-json", url}

	cmd := exec.Command("yt-dlp", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("ytdlp-list: %s", err)
	}

	list := &YtdlpList{}
	if err = json.Unmarshal(out, list); err != nil {
		return nil, fmt.Errorf("ytdlp-list: %s", err)
	}

	return list, nil
}

func ytdlp_get_info(url string) (*YtdlpInfo, error) {
	args := []string{"--skip-download", "--no-warning", "--dump-single-json", url}

	cmd := exec.Command("yt-dlp", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("ytdlp-info: %s", err)
	}

	info := &YtdlpInfo{}
	if err = json.Unmarshal(out, info); err != nil {
		return nil, fmt.Errorf("ytdlp-info: %s", err)
	}

	return info, nil
}
