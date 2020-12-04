package transfer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	md "github.com/sonr-io/core/internal/models"

	msgio "github.com/libp2p/go-msgio"
	"google.golang.org/protobuf/proto"
)

// ** Constants for Chunking Data ** //
const B64ChunkSize = 31998 // Adjusted for Base64 -- has to be divisible by 3
const BufferChunkSize = 32000

// ^ Chunks string based on B64ChunkSize ^ //
func chunkBase64(s string, B64ChunkSize int) []string {
	chunkSize := B64ChunkSize
	ss := make([]string, 0, len(s)/chunkSize+1)
	for len(s) > 0 {
		if len(s) < chunkSize {
			chunkSize = len(s)
		}
		// Create Current Chunk String
		ss, s = append(ss, s[:chunkSize]), s[chunkSize:]
	}
	return ss
}

// ^ write file as Base64 in Msgio to Stream ^ //
func writeBase64ToStream(writer msgio.WriteCloser, meta *md.Metadata) {
	// Initialize Buffer
	imgBuffer := new(bytes.Buffer)

	// @ Check Image type
	if meta.Mime.Subtype == "jpeg" {
		// Get JPEG Encoded Buffer
		err := EncodeJpegBuffer(imgBuffer, meta)
		if err != nil {
			log.Fatalln(err)
		}
	} else if meta.Mime.Subtype == "png" {
		// Get PNG Encoded Buffer
		err := EncodePngBuffer(imgBuffer, meta)
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Encode Buffer to base 64
	imgBytes := imgBuffer.Bytes()
	data := base64.StdEncoding.EncodeToString(imgBytes)
	total := int32(len(data))

	// Iterate for Entire file as String
	for i, chunk := range chunkBase64(data, B64ChunkSize) {
		log.Println("Chunk Number: ", i)
		// Create Block Protobuf from Chunk
		chunk := md.Chunk{
			Size:    int32(len(chunk)),
			B64:     chunk,
			Current: int32(i),
			Total:   total,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(&chunk)
		if err != nil {
			log.Fatalln(err)
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

// ^ write file as Bytes in Msgio to Stream ^ //
func writeBytesToStream(writer msgio.WriteCloser, meta *md.Metadata, total int32) {
	// Open File
	file, err := os.Open(meta.Path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Set Chunk Variables
	ps := make([]byte, BufferChunkSize)

	// Iterate file
	for i := 0; ; i++ {
		// Read Bytes
		bytesread, err := file.Read(ps)

		// Check for Error
		if err != nil {
			// Non EOF Error
			if err != io.EOF {
				fmt.Println(err)
			}
			// File Complete
			break
		}

		// Create Block Protobuf from Chunk
		chunk := md.Chunk{
			Size:    int32(len(ps[:bytesread])),
			Buffer:  ps[:bytesread],
			Current: int32(i),
			Total:   total,
		}

		// Convert to bytes
		bytes, err := proto.Marshal(&chunk)
		if err != nil {
			log.Fatalln(err)
		}

		// Write Message Bytes to Stream
		err = writer.WriteMsg(bytes)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
