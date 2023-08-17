package utils

import (
	"bytes"
	"log"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func StreamFFmpeg(source string, endpoint string, sourceArgs map[string]interface{}, endpointArgs map[string]interface{}) error {
	cmd := ffmpeg_go.Input(
		source,
		sourceArgs,
	).Output(
		endpoint,
		endpointArgs,
	)
	var outb bytes.Buffer
	cmd.Compile().Stderr = &outb
	if err := cmd.Run(); err != nil {
		log.Printf("error: %s", string(outb.Bytes()))
		return err
	}
	return nil
}
