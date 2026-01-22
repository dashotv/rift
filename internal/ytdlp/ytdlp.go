package ytdlp

import (
	"encoding/json"
	"os/exec"

	"github.com/dashotv/fae"
)

var YTDLP_OPTIONS = `--skip-download --no-warning --dump-single-json --cookies cookies.txt --add-header 'sec-ch-ua:"Not(A:Brand";v="8", "Chromium";v="144", "Brave";v="144"' --add-header 'sec-ch-ua-mobile:?0' --add-header 'sec-ch-ua-platform:macOS' --add-header 'User-Agent:Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36'`

func ProcessURL(url string) ([]*Info, error) {
	list, err := ytdlp_get_list(url)
	if err != nil {
		return nil, fae.Wrap(err, "getting list")
	}

	if list.Type == "video" {
		info, err := ytdlp_get_info(url)
		if err != nil {
			return nil, fae.Wrap(err, "getting info")
		}
		return []*Info{info}, nil
	}

	infos := make([]*Info, 0)
	errors := make([]error, 0)
	for _, e := range list.Entries {
		info, err := ytdlp_get_info(e.URL)
		if err != nil {
			// return nil, fae.Wrap(err, "process_url: info")
			// fmt.Printf("process_url: info: %s\n", err)
			errors = append(errors, fae.Wrapf(err, "getting info: %s", e.URL))
			continue
		}
		infos = append(infos, info)
	}

	if len(errors) > 0 {
		return nil, fae.Wrapf(errors[0], "%d errors, showing first", len(errors))
	}

	return infos, nil
}

func ytdlp_get_list(url string) (*List, error) {
	// args := []string{"--skip-download", "--no-warning", "--flat-playlist", "--dump-single-json", url}
	// cmd := exec.Command("yt-dlp", args...)

	sh := "yt-dlp " + YTDLP_OPTIONS + " --flat-playlist  " + url
	cmd := exec.Command("sh", "-c", sh)

	out, err := cmd.Output()
	if err != nil {
		return nil, fae.Wrap(err, "running command")
	}

	list := &List{}
	if err = json.Unmarshal(out, list); err != nil {
		return nil, fae.Wrap(err, "unmarshalling json")
	}

	return list, nil
}

func ytdlp_get_info(url string) (*Info, error) {
	// args := []string{"--skip-download", "--no-warning", "--dump-single-json", url}
	// cmd := exec.Command("yt-dlp", args...)

	sh := "yt-dlp " + YTDLP_OPTIONS + " " + url
	cmd := exec.Command("sh", "-c", sh)

	out, err := cmd.Output()
	if err != nil {
		// fmt.Printf("ytdlp_get_info: %s\n", out)
		return nil, fae.Wrap(err, "running command")
	}

	info := &Info{}
	if err = json.Unmarshal(out, info); err != nil {
		return nil, fae.Wrap(err, "unmarshalling json")
	}

	return info, nil
}
