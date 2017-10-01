package main

import (
	"flag"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/soniakeys/quant/median"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

const VERSION = "1.0.3"

func Resize(ifile string, ofile string, width uint, height uint, oformat string) int {
	//
	// input
	//
	var in io.Reader
	var err error
	if ifile == "-" {
		in = os.Stdin
	} else {
		in, err = os.Open(ifile)
		if err != nil {
			log.Fatal(err)
			return 1
		}
	}

	//
	// decode
	//
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
	img, fmt, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
		return 1
	}

	//
	// resize
	//
	m := resize.Resize(width, height, img, resize.Bicubic)

	//
	// output
	//
	var out io.Writer
	if ofile == "-" {
		out = os.Stdout
	} else {
		out, err = os.Create(ofile)
		if err != nil {
			log.Fatal(err)
			return 1
		}
	}

	if oformat == "-" {
		oformat = fmt
	}
	if oformat == "png" {
		err = png.Encode(out, m)
	} else if oformat == "gif" {
		opt := gif.Options{}
		opt.NumColors = 256
		opt.Quantizer = median.Quantizer(256)
		err = gif.Encode(out, m, &opt)
	} else {
		opt := jpeg.Options{}
		opt.Quality = 100
		err = jpeg.Encode(out, m, &opt)
	}
	if err != nil {
		log.Fatal(err)
		return 1
	}

	return 0
}

func main() {
	//
	// option
	//
	printUsage := flag.Bool("help", false, "print this help message.")
	ifile := flag.String("ifile", "-", "input-file name. ('-' is stdin) ")
	ofile := flag.String("ofile", "-", "output-file name. ('-' is stdout) ")
	oformat := flag.String("oformat", "-", "output-file format. [jpg|png|gif] ('-' is same as input)")
	width := flag.Uint("width", 0, "resize width. (0 is keep aspect ratio)")
	height := flag.Uint("height", 0, "resize height. (0 is keep aspect ratio)")
	flag.Parse()

	if *printUsage {
		fmt.Fprintf(os.Stderr, "samuneiru: The image thumbnailer. (Ver.%s)\n", VERSION)
		fmt.Fprintf(os.Stderr, "usage: %s [-width | -height] [other options]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *width == 0 && *height == 0 {
		log.Println("You must set -width or -height")
		os.Exit(2)
	}

	if *width != 0 && *height != 0 {
		log.Println("You can't set both -width and -height")
		os.Exit(2)
	}

	os.Exit(Resize(*ifile, *ofile, *width, *height, *oformat))
}
