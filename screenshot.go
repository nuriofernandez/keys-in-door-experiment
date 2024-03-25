package main

import (
	"fmt"
	"github.com/bluenviron/gortsplib/v4"
	"github.com/bluenviron/gortsplib/v4/pkg/base"
	"github.com/bluenviron/gortsplib/v4/pkg/format"
	"github.com/bluenviron/gortsplib/v4/pkg/format/rtph264"
	"github.com/bluenviron/mediacommon/pkg/codecs/h264"
	"github.com/pion/rtp"
	"image"
	"log"
	"os"
)

var rtsp = fmt.Sprintf(
	"rtsp://%s:%s@%s:554/stream1",
	os.Getenv("RTSP_USER"),
	os.Getenv("RTSP_PASSWORD"),
	os.Getenv("RTSP_HOST"),
)

func screenShot() (*image.Image, error) {
	c := gortsplib.Client{}

	// parse URL
	u, err := base.ParseURL(rtsp)
	if err != nil {
		return nil, err
	}

	// connect to the server
	err = c.Start(u.Scheme, u.Host)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	// find available medias
	desc, _, err := c.Describe(u)
	if err != nil {
		return nil, err
	}

	// find the H264 media and format
	var forma *format.H264
	medi := desc.FindFormat(&forma)
	if medi == nil {
		return nil, fmt.Errorf("media not found")
	}

	// setup RTP/H264 -> H264 decoder
	rtpDec, err := forma.CreateDecoder()
	if err != nil {
		return nil, err
	}

	// setup H264 -> raw frames decoder
	frameDec := &h264Decoder{}
	err = frameDec.initialize()
	if err != nil {
		return nil, err
	}
	defer frameDec.close()

	// if SPS and PPS are present into the SDP, send them to the decoder
	if forma.SPS != nil {
		frameDec.decode(forma.SPS)
	}
	if forma.PPS != nil {
		frameDec.decode(forma.PPS)
	}

	// setup a single media
	_, err = c.Setup(desc.BaseURL, medi, 0, 0)
	if err != nil {
		return nil, err
	}

	iframeReceived := false

	// called when a RTP packet arrives
	var capturedImage *image.Image
	c.OnPacketRTP(medi, forma, func(pkt *rtp.Packet) {
		// extract access units from RTP packets
		au, err := rtpDec.Decode(pkt)
		if err != nil {
			if err != rtph264.ErrNonStartingPacketAndNoPrevious && err != rtph264.ErrMorePacketsNeeded {
				log.Printf("ERR: %v", err)
			}
			return
		}

		// wait for an I-frame
		if !iframeReceived {
			if !h264.IDRPresent(au) {
				log.Printf("waiting for an I-frame")
				return
			}
			iframeReceived = true
		}

		for _, nalu := range au {
			if capturedImage != nil {
				return
			}
			// convert NALUs into RGBA frames
			img, err := frameDec.decode(nalu)
			if err != nil {
				panic(err)
			}

			// wait for a frame
			if img == nil {
				continue
			}

			capturedImage = &img
			go c.Close()
			return
		}
	})

	// start playing
	_, err = c.Play(nil)
	if err != nil {
		return nil, err
	}

	// wait until a fatal error
	err = c.Wait()

	if capturedImage != nil {
		return capturedImage, nil
	}

	return capturedImage, err
}
