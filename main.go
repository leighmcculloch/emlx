package main

import (
	"flag"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strconv"

	"github.com/DusanKasan/parsemail"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	flag.Parse()
	for _, arg := range flag.Args() {
		fmt.Fprintln(os.Stderr, "File:", arg)
		f, err := os.Open(arg)
		if err != nil {
			return err
		}
		defer f.Close()
		err = extract(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func extract(f *os.File) error {
	mail, err := parsemail.Parse(f)
	if err != nil {
		return err
	}

	dir := mail.Date.Format("2006-01-02") + " - " + mail.Subject
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}

	i := 0
	for _, a := range mail.Attachments {
		i++
		filename := strconv.Itoa(i) + " - " + a.Filename
		path := filepath.Join(dir, filename)
		fmt.Fprintln(os.Stderr, "Attachment:", path)
		w, err := os.Create(path)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, a.Data)
		if err != nil {
			return err
		}
	}

	for _, e := range mail.EmbeddedFiles {
		i++
		filename := e.CID
		_, params, err := mime.ParseMediaType(e.ContentType)
		if err != nil {
			return err
		}
		if fn, ok := params["name"]; ok {
			filename = fn
		}
		filename = strconv.Itoa(i) + " - " + filename
		path := filepath.Join(dir, filename)
		fmt.Fprintln(os.Stderr, "Attachment:", path)
		w, err := os.Create(path)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, e.Data)
		if err != nil {
			return err
		}
	}

	return nil
}
